package utils

import "os"

func SaveFile(bytes []byte, name string, path string) (string, error) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return "", err
	}

	fileName := path + name
	f, err := os.Create(fileName)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func SaveCode(code []byte, username string) (string, error) {
	path := "C:\\project\\go\\My-Exercise\\code\\" + username
	name := "\\main.go"
	return SaveFile(code, name, path)
}
