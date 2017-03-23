package docker

import (
	"testing"

	"os"

	"path/filepath"

	"github.com/mutemaniac/steel/docker/langs"
	"github.com/mutemaniac/steel/docker/langs/golang"
	"github.com/mutemaniac/steel/docker/langs/node"
	"github.com/mutemaniac/steel/docker/langs/python"
)

func TestGenerateDockerfile(t *testing.T) {
	type args struct {
		lang langs.LangHelper
		dir  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Add test cases.
		{
			name: "go",
			args: args{
				dir:  "/Users/tang/Documents/test/golang",
				lang: &golang.GoLangHelper{},
			},
			wantErr: false,
		},
		{
			name: "python",
			args: args{
				dir:  "/Users/tang/Documents/test/python",
				lang: &python.PythonHelper{},
			},
			wantErr: false,
		},
		{
			name: "node",
			args: args{
				dir:  "/Users/tang/Documents/test/node",
				lang: &node.NodeLangHelper{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateDockerfile(tt.args.lang, tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("GenerateDockerfile() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				os.Remove(filepath.Join(tt.args.dir, "Dockerfile"))
			}
		})
	}
}
