package docker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mutemaniac/steel/docker/langs"
)

const sourcefileName = "func"

//SaveCode return file fullpath && error
func SaveCode(dir string, lang langs.LangHelper, source string) (string, error) {
	fmt.Println("Enter SaveCode.")
	if source == "" {
		return "", errors.New("source is null")
	}
	sourcefile := filepath.Join(dir, sourcefileName+lang.Extension())
	sf, err := os.Create(sourcefile)
	if err != nil {
		return "", err
	}
	defer sf.Close()
	sf.WriteString(source)
	return sourcefile, nil
}
