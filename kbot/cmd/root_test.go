package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	for _, tt := range []struct {
		name    string
		args    []string
		token   string
		want    string
		wantErr bool
	}{
		{"help", []string{"--help"}, "", "start", false},
		{"version", []string{"--version"}, "", version, false},
		{"missing token", []string{"start"}, "", "TELE_TOKEN is required", true},
		{"blank token", []string{"start"}, "  ", "TELE_TOKEN is required", true},
		{"invalid timeout", []string{"start", "--poll-timeout=0s"}, "test", "poll-timeout", true},
		{"excessive timeout", []string{"start", "--poll-timeout=2m"}, "test", "poll-timeout", true},
		{"unexpected argument", []string{"start", "extra"}, "", "unknown command", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("TELE_TOKEN", tt.token)
			root := NewRootCommand()
			var out bytes.Buffer
			root.SetOut(&out)
			root.SetErr(&out)
			root.SetArgs(tt.args)
			err := root.Execute()
			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr = %v", err, tt.wantErr)
			}
			if err != nil {
				out.WriteString(err.Error())
			}
			if !strings.Contains(out.String(), tt.want) {
				t.Errorf("output = %q, want substring %q", out.String(), tt.want)
			}
		})
	}
}
