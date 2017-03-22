package langs

import "github.com/mutemaniac/steel/docker/langs"

func init() {
	langs.RegisterLangHelper("node", new)
}

func new() (langs.LangHelper, error) {
	return &NodeLangHelper{}, nil
}

type NodeLangHelper struct {
}

func (lh *NodeLangHelper) Entrypoint() string {
	return "node func.js"
}
func (lh *NodeLangHelper) Extension() string {
	return ".js"
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
