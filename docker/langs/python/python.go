package python

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"path/filepath"

	"github.com/mutemaniac/steel/common"
	"github.com/mutemaniac/steel/docker/langs"
)

const tplDockerfile = `FROM {{ .BaseImage }}
WORKDIR /function
ADD . /function/
ENTRYPOINT [{{ .Entrypoint }}]
`

func init() {
	langs.RegisterLangHelper("python", new)
}

func new(dir string) (langs.LangHelper, error) {
	return &PythonHelper{dir: dir}, nil
}

type PythonHelper struct {
	dir string
}

func (lh *PythonHelper) Entrypoint() string {
	return "python2 func.py"
}
func (lh *PythonHelper) Extension() string {
	return ".py"
}
func (lh *PythonHelper) BaseImage() string {
	return "iron/python:2"
}
func (lh *PythonHelper) HasPreBuild() bool {
	return common.Exists(filepath.Join(lh.dir, "requirements.txt"))
}

// PreBuild for Go builds the binary so the final image can be as small as possible
func (lh *PythonHelper) PreBuild() error {
	// wd, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }

	pbcmd := fmt.Sprintf("docker run --rm -v %s:/worker -w /worker iron/python:2-dev pip install -t packages -r requirements.txt", lh.dir)
	fmt.Println("Running prebuild command:", pbcmd)
	parts := strings.Fields(pbcmd)
	fmt.Println("parts: %v", parts)
	head := parts[0]
	parts = parts[1:len(parts)]
	cmd := exec.Command(head, parts...)
	cmd.Dir = lh.dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running docker build: %v", err)
	}
	return nil
}

func (lh *PythonHelper) AfterBuild() error {
	return nil
}

func (lh *PythonHelper) DockerfileTemplate() string {
	return tplDockerfile
}
func (lh *PythonHelper) SetBaseDir(dir string) {
	lh.dir = dir
}
