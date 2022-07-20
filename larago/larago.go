package larago

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type Larago struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   config
	Routes   *chi.Mux
}

type config struct {
	port     string
	renderer string
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

	// create loggers and application env
	l.InfoLog, l.ErrorLog = l.startLoggers()
	l.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	l.Version = version
	l.RootPath = rootPath
	l.Routes = l.routes().(*chi.Mux)

	// framework config
	l.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
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

// ListenAndServe starts the web server
func (l Larago) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     l.ErrorLog,
		Handler:      l.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	l.InfoLog.Printf("Server started at: %s:%s", os.Getenv("SERVER_NAME"), os.Getenv("PORT"))
	err := srv.ListenAndServe()
	l.ErrorLog.Fatal(err)
}

func (l *Larago) checkDotEnv(path string) error {
	err := l.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return nil
	}

	return nil
}

func (l *Larago) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
