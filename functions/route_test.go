package functions

import (
	"testing"

	ironClient "github.com/iron-io/functions_go"

	"github.com/mutemaniac/steel/models"
)

func TestCreateRoute(t *testing.T) {
	type args struct {
		route models.ExRouteWrapper
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
				route: models.ExRouteWrapper{
					AppName: "tang",
					Runtime: "golang",
					Route: ironClient.Route{
						Path:  "/golangdemo",
						Image: "mutemaniac/golangdemo:0.1",
					},
					Code: `package main

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
				},
			},
			wantErr: false,
		},
		{
			name: "python",
			args: args{
				route: models.ExRouteWrapper{
					AppName: "tang",
					Runtime: "python",
					Route: ironClient.Route{
						Path:  "/pythondemo",
						Image: "mutemaniac/pythondemo:0.1",
					},
					Code: `import sys
sys.path.append("packages")
import os
import json

name = "World"
if not os.isatty(sys.stdin.fileno()):
	obj = json.loads(sys.stdin.read())
	if obj["name"] != "":
		name = obj["name"]

print "Hello", name, "!!!"`,
				},
			},
			wantErr: false,
		},
		{
			name: "node",
			args: args{
				route: models.ExRouteWrapper{
					AppName: "tang",
					Runtime: "node",
					Route: ironClient.Route{
						Path:  "/nodedemo",
						Image: "mutemaniac/nodedemo:0.1",
					},
					Code: `name = "World";
fs = require('fs');
try {
	obj = JSON.parse(fs.readFileSync('/dev/stdin').toString())
	if (obj.name != "") {
		name = obj.name
	}
} catch(e) {}
console.log("Hello", name, "from Node!");`,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateRoute(tt.args.route); (err != nil) != tt.wantErr {
				t.Errorf("CreateRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
