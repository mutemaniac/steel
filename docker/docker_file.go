package docker

import (
	"bytes"
	"path/filepath"
	"strings"
	"text/template"

	"os"

	"fmt"

	"github.com/mutemaniac/steel/docker/langs"
)

// GenerateDockerfile Generate the docker file.
// lang -- go python nodejs
// filepath -- code file
func GenerateDockerfile(lang langs.LangHelper, dir string) error {
	dockerfile := filepath.Join(dir, "Dockerfile")
	df, err := os.Create(dockerfile)
	if err != nil {
		return err
	}
	defer df.Close()

	// convert entrypoint string to slice
	epvals := strings.Fields(lang.Entrypoint())
	var buffer bytes.Buffer
	for i, s := range epvals {
		if i > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString("\"")
		buffer.WriteString(s)
		buffer.WriteString("\"")
	}
	fmt.Print(buffer.String())
	t := template.Must(template.New("Dockerfile").Parse(lang.DockerfileTemplate()))
	err = t.Execute(df, struct {
		BaseImage, Entrypoint string
	}{lang.BaseImage(), buffer.String()})

	return err
}
