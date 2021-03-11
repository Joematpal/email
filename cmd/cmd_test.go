package cmd

import "testing"

func Test_parseTemplate(t *testing.T) {
	type args struct {
		templateFileName string
		data             map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				templateFileName: "email.tmpl",
				data: map[string]interface{}{
					"test":   "test_string",
					"sender": "jmp",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTemplate(tt.args.templateFileName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
