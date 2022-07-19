package larago

import (
	"fmt"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Larago struct {
	AppName string
	Debug   bool
	Version string
}

func (l *Larago) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers", "migrations", "views", "models", "public", "temp", "middleware", "logs",
		},
	}

	err := l.Init(pathConfig)

	if err != nil {
		return err
	}

	err = l.checkDotEnv(rootPath)

	if err != nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath + "/.env")

	if err != nil {
		return err
	}

	return nil
}

func (l *Larago) Init(p initPaths) error {
	root := p.rootPath

	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := l.CreateDirIfNotExists(root + "/" + path)

		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Larago) checkDotEnv(path string) error {
	err := l.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return nil
	}

	return nil
}
