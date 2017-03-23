package golang

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"path/filepath"

	"github.com/mutemaniac/steel/docker/langs"
)

const tplDockerfile = `FROM {{ .BaseImage }}
WORKDIR /function
ADD . /function/
ENTRYPOINT [{{ .Entrypoint }}]
`
const baseimage = `iron/go`

func init() {
	langs.RegisterLangHelper("golang", new)
}

func new(dir string) (langs.LangHelper, error) {
	return &GoLangHelper{dir: dir}, nil
}

type GoLangHelper struct {
	dir string
}

func (lh *GoLangHelper) Extension() string {
	return ".go"
}
func (lh *GoLangHelper) BaseImage() string {
	return baseimage
}
func (lh *GoLangHelper) Entrypoint() string {
	return "./func"
}

func (lh *GoLangHelper) HasPreBuild() bool {
	return true
}

// PreBuild for Go builds the binary so the final image can be as small as possible
func (lh *GoLangHelper) PreBuild() error {
	pbcmd := fmt.Sprintf("docker run --rm -v %s:/go/src/github.com/x/y -w /go/src/github.com/x/y iron/go:dev go build -o func", lh.dir)
	fmt.Println("Running prebuild command:", pbcmd)
	parts := strings.Fields(pbcmd)
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

func (lh *GoLangHelper) AfterBuild() error {
	return os.Remove(filepath.Join(lh.dir, "func"))
}

func (lh *GoLangHelper) DockerfileTemplate() string {
	return tplDockerfile
}

func (lh *GoLangHelper) SetBaseDir(dir string) {
	lh.dir = dir
}
