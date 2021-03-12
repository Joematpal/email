package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"

	"github.com/joematpal/email/internal/flags"
	"github.com/urfave/cli/v2"
)

var delimeter = "dll!769042**cd1"

type Sender interface {
	Send(from string, to []string, data interface{}, template string, files []string) (bool, error)
}

func NewApp() *cli.App {
	return &cli.App{
		Name:  "email",
		Flags: flags.Email,
		Action: func(c *cli.Context) error {

			data := map[string]interface{}{}
			dataPath := c.String("data")
			if dataPath != "" {
				dataFile, err := os.ReadFile(dataPath)
				if err != nil {
					return fmt.Errorf("data file: %v", err)
				}
				if err := json.Unmarshal(dataFile, &data); err != nil {
					return err
				}
			}

			from := c.String("from")
			if from == "" {
				return errors.New("sender is empty")
			}
			files := strings.Split(c.String("files"), ",")

			subject := c.String("subject")
			if subject == "" {
				return errors.New("subject is empty")
			}
			to := strings.Split(c.String("to"), ",")

			e := &Email{
				opts: Options{
					host:     c.String("host"),
					password: c.String("password"),
					port:     c.String("port"),
				},
			}
			sent, err := e.Send(subject, from, to, data, c.String("template"), files)
			if err != nil {
				return err
			}

			if !sent {
				return errors.New("not sent")
			}
			return nil
		},
	}
}

func parseTemplate(templateFileName string, data map[string]interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", fmt.Errorf("%s: %v", templateFileName, err)
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", fmt.Errorf("execute tmplt: %v", err)
	}

	return buf.String(), nil
}

type Options struct {
	host     string
	password string
	port     string
}

type Email struct {
	opts Options
}

func (e *Email) Send(subject, from string, to []string, data map[string]interface{}, template string, files []string) (bool, error) {
	// Create the auth
	emailAuth := smtp.PlainAuth("", from, e.opts.password, e.opts.host)

	emailBody, err := parseTemplate(template, data)
	if err != nil {
		return false, err
	}
	sb := new(bytes.Buffer)

	sb.WriteString("MIME-Version: 1.0\r\n")
	fs := []string{}
	for _, file := range files {
		if file == "" {
			continue
		}
		fs = append(fs, file)
	}
	sb.WriteString(fmt.Sprintf("Subject: %s \n", subject))
	sb.WriteString(fmt.Sprintf("From: %s\r\n", from))
	sb.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ";")))

	if len(fs) > 0 {
		sb.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter))
	} else {
		sb.WriteString(fmt.Sprintf("MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"))
	}

	if len(fs) > 0 {
		sb.WriteString(fmt.Sprintf("\r\n--%s\r\n", delimeter))
	}

	sb.WriteString(emailBody)

	// read files
	for _, file := range fs {
		filename := filepath.Base(file)

		sb.WriteString(fmt.Sprintf("\r\n--%s\r\n", delimeter))
		sb.WriteString("Content-Type: applications/zip; charset=\"utf-8\"\r\n")
		sb.WriteString("Content-Transfer-Encoding: base64\r\n")
		sb.WriteString("Content-Disposition: attachment;filename=\"" + filename + "\"\r\n")
		rawFile, err := os.ReadFile(file)
		if err != nil {
			return false, fmt.Errorf("open attachment: %s: %v", file, err)
		}

		sb.WriteString(base64.StdEncoding.EncodeToString(rawFile))
		sb.WriteString("\r\n")
	}

	// Create Connection
	addr := fmt.Sprintf("%s:%v", e.opts.host, e.opts.port)

	c, err := smtp.Dial(addr)
	if err != nil {
		return false, fmt.Errorf("dial: %s: %v", addr, err)
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: e.opts.host}

		if err = c.StartTLS(config); err != nil {
			return false, err
		}
	}
	if err = c.Auth(emailAuth); err != nil {
		return false, err
	}

	if err = c.Mail(from); err != nil {
		return false, err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return false, err
		}
	}

	wr, err := c.Data()
	if err != nil {
		return false, err
	}

	if _, err := wr.Write(sb.Bytes()); err != nil {
		return false, err
	}
	if err := wr.Close(); err != nil {
		return false, err
	}

	return true, c.Quit()
}
