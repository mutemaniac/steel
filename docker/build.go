package docker

// Build Build docker image. Paramater may be changed sometimes.
func Build(code string, lang string, image string) error {
	// TODO save file
	fullpath, err := SaveCode(code, "")
	if err != nil {
		return err
	}
	defer DeleteCodeDir(fullpath)
	// TODO dockerfile
	err = GenerateDockerfile(lang, fullpath)
	if err != nil {
		//TODO
		return err
	}
	// TODO docker build
	dockerbuild(image)
	// TODO docker push
	dockerpush(image)

	return nil
}

func dockerbuild(image string) error {
	return nil
}

func dockerpush(image string) error {
	return nil
}
