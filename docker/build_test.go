package docker

import (
	"testing"

	"github.com/mutemaniac/steel/docker/langs"
)

func TestBuild(t *testing.T) {
	type args struct {
		code    string
		lang    string
		image   string
		appname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "go",
			args: args{
				code: `package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
}

func main() {
	p := &Person{Name: "World"}
	json.NewDecoder(os.Stdin).Decode(p)
	fmt.Printf("Hello %v!", p.Name)
}`,
				lang:    "golang",
				image:   "mutemaniac/iron-go:0.1",
				appname: "tang",
			},
			wantErr: false,
		},
		{
			name: "python",
			args: args{
				code: `import sys
sys.path.append("packages")
import os
import json

name = "World"
if not os.isatty(sys.stdin.fileno()):
	obj = json.loads(sys.stdin.read())
	if obj["name"] != "":
		name = obj["name"]

print "Hello", name, "!!!"`,
				lang:    "python",
				image:   "mutemaniac/iron-python:0.1",
				appname: "tang",
			},
			wantErr: false,
		},
		{
			name: "node",
			args: args{
				code: `name = "World";
fs = require('fs');
try {
	obj = JSON.parse(fs.readFileSync('/dev/stdin').toString())
	if (obj.name != "") {
		name = obj.name
	}
} catch(e) {}
console.log("Hello", name, "from Node!");`,
				image:   "mutemaniac/iron-node:0.1",
				lang:    "node",
				appname: "tang",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Build(tt.args.code, tt.args.lang, tt.args.image, tt.args.appname); (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Must(i langs.LangHelper, err error) langs.LangHelper {
	if err != nil {
		panic(err)
	}
	return i
}
func Test_dockerbuild(t *testing.T) {

	type args struct {
		dir    string
		helper langs.LangHelper
		image  string
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
				dir:    "/Users/tang/Documents/test/golang",
				helper: Must(langs.New("golang", "/Users/tang/Documents/test/golang")),
				image:  "mutemaniac/iron-go:0.1",
			},
			wantErr: false,
		},
		{
			name: "python",
			args: args{
				dir:    "/Users/tang/Documents/test/python",
				helper: Must(langs.New("python", "/Users/tang/Documents/test/python")),
				image:  "mutemaniac/iron-python:0.1",
			},
			wantErr: false,
		},
		{
			name: "node",
			args: args{
				dir:    "/Users/tang/Documents/test/node",
				image:  "mutemaniac/iron-node:0.1",
				helper: Must(langs.New("node", "/Users/tang/Documents/test/node")),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := dockerbuild(tt.args.dir, tt.args.helper, tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("%s: dockerbuild() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func Test_dockerpush(t *testing.T) {
	type args struct {
		image string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// Add test cases.
	// {
	// 	name: "go",
	// 	args: args{
	// 		image: "mutemaniac/service-cloud:0.1",
	// 	},
	// 	wantErr: false,
	// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := dockerpush(tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("dockerpush() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_exists(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "go",
			args: args{
				name: "/Users/tang/Documents/test/golang",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exists(tt.args.name); got != tt.want {
				t.Errorf("exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
