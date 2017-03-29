package docker

import "testing"

func TestLogin(t *testing.T) {
	type args struct {
		server   string
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// Add test cases.
	// {
	// 	name: "dcloudos",
	// 	args: args{
	// 		server: "dcatalog.hnaresearch.com",
	// 	},
	// 	wantErr: false,
	// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Login(tt.args.server, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
