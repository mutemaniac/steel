package docker

import (
	"github.com/mutemaniac/steel/docker/langs"
	_ "github.com/mutemaniac/steel/docker/langs/golang"
	_ "github.com/mutemaniac/steel/docker/langs/node"
	_ "github.com/mutemaniac/steel/docker/langs/python"
)

func NewLangHelper(runtime string, dir string) (langs.LangHelper, error) {
	if runtime == "" {
		return nil, nil
	}
	return langs.New(runtime, dir)
}
