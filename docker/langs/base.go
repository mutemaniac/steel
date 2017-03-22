package langs

import "fmt"

// GetLangHelper returns a LangHelper for the passed in language
func GetLangHelper(lang string) (LangHelper, error) {
	switch lang {
	case "go":
		return &GoLangHelper{}, nil
	case "node":
		return &NodeLangHelper{}, nil
	case "python":
		return &PythonHelper{}, nil
	}
	return nil, fmt.Errorf("No language helper found for %v", lang)
}

type LangHelper interface {
	Entrypoint() string
	HasPreBuild() bool
	PreBuild() error
	AfterBuild() error
	Extension() string
}

func RegisterLangHelper(name string, external string, new func()) {

}

func New(name string) (LangHelper, error) {
	return nil, nil
}

func GetExternal(name string) (string, error) {
	return "", nil
}
