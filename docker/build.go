package docker

import (
	"io/ioutil"
	"os"

	"github.com/mutemaniac/steel/config"
)

// Build Build docker image. Paramater may be changed sometimes.
func Build(code string, lang string, image string, appname string) error {
	//Generate a temp directory
	dir, err := ioutil.TempDir(config.CodeFileTmpDir, appname)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	// TODO save file
	fullpath, err := SaveCode(dir, code, "")
	if err != nil {
		return err
	}
	// TODO dockerfile
	err = GenerateDockerfile(lang, fullpath)
	if err != nil {
		//TODO
		return err
	}
	// TODO docker build
	dockerbuild(image)
	// TODO docker push
	dockerpush(image)

	return nil
}

func dockerbuild(image string) error {
	return nil
}

func dockerpush(image string) error {
	return nil
}
