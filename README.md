# email
CLI to send emails via smtp relay Simple Auth.
Attach Files via files flags, pass file path delimited by `,`


## install
```
go install github.com/joematpal/email
```

## usage
It will error if one uses smtp.gmail.com the auth will panic if not enabled
https://myaccount.google.com/lesssecureapps

## example
```
EMAIL_PASSWORD=$SUPER_SECRET_PASS \
EMAIL_HOST=smtp.gmail.com \
EMAIL_FROM=joseph@test.com \
EMAIL_TEMPLATE=/path/to/template \
email --to recepient@test.com \
--subject "test subject" \
--file /path/to/file1,/path/to/file2
```
## help
```
NAME:
   email - A new cli application

USAGE:
   email [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, --files value  
   --password value              [$EMAIL_PASSWORD]
   --host value                  [$EMAIL_HOST]
   --from value                  [$EMAIL_FROM]
   --port value                 (default: "587") [$EMAIL_PORT]
   --template value              [$EMAIL_TEMPLATE]
   --to value                   
   --subject value              
   --data value                 provide path to json file
   --help, -h                   show help (default: false)
```