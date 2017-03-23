package docker

import (
	"testing"

	"github.com/mutemaniac/steel/docker/langs"
	"github.com/mutemaniac/steel/docker/langs/golang"
	"github.com/mutemaniac/steel/docker/langs/node"
	"github.com/mutemaniac/steel/docker/langs/python"
)

func TestSaveCode(t *testing.T) {
	type args struct {
		dir    string
		lang   langs.LangHelper
		source string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "go",
			args: args{
				dir:  "/Users/tang/Documents/test/golang",
				lang: &golang.GoLangHelper{},
				source: `package main

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
}`},

			want:    "/Users/tang/Documents/test/golang/func.go",
			wantErr: false,
		},
		{
			name: "python",
			args: args{
				dir:  "/Users/tang/Documents/test/python",
				lang: &python.PythonHelper{},
				source: `import sys
sys.path.append("packages")
import os
import json

name = "World"
if not os.isatty(sys.stdin.fileno()):
	obj = json.loads(sys.stdin.read())
	if obj["name"] != "":
		name = obj["name"]

print "Hello", name, "!!!"`},

			want:    "/Users/tang/Documents/test/python/func.py",
			wantErr: false,
		},
		{
			name: "node",
			args: args{
				dir:  "/Users/tang/Documents/test/node",
				lang: &node.NodeLangHelper{},
				source: `name = "World";
fs = require('fs');
try {
	obj = JSON.parse(fs.readFileSync('/dev/stdin').toString())
	if (obj.name != "") {
		name = obj.name
	}
} catch(e) {}
console.log("Hello", name, "from Node!");`},

			want:    "/Users/tang/Documents/test/node/func.js",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SaveCode(tt.args.dir, tt.args.lang, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SaveCode() = %v, want %v", got, tt.want)
			}
			//os.Remove(got)
		})
	}
}
