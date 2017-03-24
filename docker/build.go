package docker

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mutemaniac/steel/config"
	"github.com/mutemaniac/steel/docker/langs"
)

// Build Build docker image. Paramater may be changed sometimes.
func Build(code string, lang string, image string, appname string) error {
	fmt.Println("Enter Build.")
	//Generate a temporary directory
	dir, err := ioutil.TempDir(config.CodeFileTmpDir, appname)
	if err != nil {
		fmt.Printf("TempDir error %v.", err)
		return err
	}
	defer os.RemoveAll(dir)
	fmt.Println("New langHelper.", lang)
	langHelper, err := NewLangHelper(lang, dir)
	if err != nil {
		fmt.Printf("New langHelper error %v.\n", err)
		return err
	}
	//Save source code file.
	fmt.Println("SaveCode.")
	_, err = SaveCode(dir, langHelper, code)
	if err != nil {
		return err
	}
	//Docker build
	fmt.Println("dockerbuild.")
	err = dockerbuild(dir, langHelper, image)
	if err != nil {
		return err
	}
	// Docker push
	fmt.Println("dockerbuild.")
	err = dockerpush(image)
	if err != nil {
		return err
	}

	return nil
}

func dockerbuild(dir string, helper langs.LangHelper, image string) error {
	fmt.Println("Enter dockerbuild.")
	dockerfile := filepath.Join(dir, "Dockerfile")
	if !exists(dockerfile) {
		err := GenerateDockerfile(helper, dir)
		if err != nil {
			return err
		}
	}
	if helper.HasPreBuild() {
		err := helper.PreBuild()
		if err != nil {
			fmt.Printf("PreBuild: %v \n", err)
			return err
		}
	}
	cmd := exec.Command("docker", "build", "-t", image, ".")
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running docker build: %v", err)
	}
	if helper != nil {
		err := helper.AfterBuild()
		if err != nil {
			fmt.Printf("AfterBuild: %v \n", err)
			return err
		}
	}
	return nil
}

func dockerpush(image string) error {
	fmt.Println("Enter dockerpush.")
	cmd := exec.Command("docker", "push", image)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running docker push: %v", err)
	}
	return nil
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
