package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func Login(server string, username string, password string) error {
	cmd := exec.Command("docker", "login", "-u", username, "-p", password, server)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(server + username + password)
		return err
	}
	return nil
}
