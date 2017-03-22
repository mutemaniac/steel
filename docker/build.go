package docker

import (
	"io/ioutil"
	"os"

	"github.com/mutemaniac/steel/config"
	"github.com/mutemaniac/steel/docker/langs"
)

// Build Build docker image. Paramater may be changed sometimes.
func Build(code string, lang string, image string, appname string) error {
	//Generate a temporary directory
	dir, err := ioutil.TempDir(config.CodeFileTmpDir, appname)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	langHelper, err := langs.New(lang)
	if err != nil {
		return nil
	}
	//Save source code file.
	fullpath, err := SaveCode(dir, langHelper, code)
	if err != nil {
		return err
	}
	err = GenerateDockerfile(langHelper, fullpath)
	if err != nil {
		return err
	}
	//Docker build
	err = dockerbuild(langHelper, image)
	if err != nil {
		return err
	}
	// Docker push
	err = dockerpush(image)
	if err != nil {
		return err
	}

	return nil
}

func dockerbuild(lang langs.LangHelper, image string) error {
	return nil
}

func dockerpush(image string) error {
	return nil
}
