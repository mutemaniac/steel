package node

import "github.com/mutemaniac/steel/docker/langs"

const tplDockerfile = `FROM {{ .BaseImage }}
WORKDIR /function
ADD . /function/
ENTRYPOINT [{{ .Entrypoint }}]
`

func init() {
	langs.RegisterLangHelper("node", new)
}

func new(dir string) (langs.LangHelper, error) {
	return &NodeLangHelper{dir: dir}, nil
}

type NodeLangHelper struct {
	dir string
}

func (lh *NodeLangHelper) Entrypoint() string {
	return "node func.js"
}
func (lh *NodeLangHelper) Extension() string {
	return ".js"
}
func (lh *NodeLangHelper) BaseImage() string {
	return "iron/node"
}

func (lh *NodeLangHelper) HasPreBuild() bool {
	return false
}

// PreBuild for Go builds the binary so the final image can be as small as possible
func (lh *NodeLangHelper) PreBuild() error {
	return nil
}

func (lh *NodeLangHelper) AfterBuild() error {
	return nil
}

func (lh *NodeLangHelper) DockerfileTemplate() string {
	return tplDockerfile
}
func (lh *NodeLangHelper) SetBaseDir(dir string) {
	lh.dir = dir
}
