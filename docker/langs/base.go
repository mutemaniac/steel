package langs

import "errors"

// GetLangHelper returns a LangHelper for the passed in language
// func GetLangHelper(lang string) (LangHelper, error) {
// 	switch lang {
// 	case "go":
// 		return &GoLangHelper{}, nil
// 	case "node":
// 		return &NodeLangHelper{}, nil
// 	case "python":
// 		return &PythonHelper{}, nil
// 	}
// 	return nil, fmt.Errorf("No language helper found for %v", lang)
// }

type LangHelper interface {
	Entrypoint() string
	HasPreBuild() bool
	PreBuild() error
	AfterBuild() error
	Extension() string
	DockerfileTemplate() string
	BaseImage() string
	SetBaseDir(dir string)
}
type newLangHelperFunc func(dir string) (LangHelper, error)

var langHelpers = map[string](newLangHelperFunc){}

func RegisterLangHelper(name string, new newLangHelperFunc) {
	langHelpers[name] = new
}

func New(runtime string, dir string) (LangHelper, error) {
	if runtime == "" {
		return nil, nil
	}
	f, ok := langHelpers[runtime]
	if !ok {
		return nil, errors.New("The language " + runtime + " not support")
	}
	return f(dir)
}
