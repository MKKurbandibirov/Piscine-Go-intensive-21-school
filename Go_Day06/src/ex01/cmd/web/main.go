package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Application struct {
	logger        *log.Logger
	templateCache map[string]*template.Template
}

type Config struct {
	Addr      string
	StaticDir string
	DBHost    string
	DBUser    string
	DBPort    string
	DBName    string
	Pass      string
}

func main() {
	var cfg = new(Config)
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Directory with static files")
	flag.Parse()

	logger := log.New(os.Stderr, "\\u001b[31m[ERROR]\\u001b[0m\\t", log.Ltime|log.Lshortfile)

	// db

	cfg.Addr = "localhost:8888"
	//cfg.DBHost = viper.GetString("db.host")
	//cfg.DBPort = viper.GetString("db.port")
	//cfg.DBUser = viper.GetString("db.user")
	//cfg.DBName = viper.GetString("db.dbname")
	//cfg.Pass = viper.GetString("db.password")

	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		logger.Fatal(err)
	}

	app := &Application{
		logger:        logger,
		templateCache: templateCache,
	}

	server := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: app.logger,
		Handler:  app.routes(cfg),
	}

	err = server.ListenAndServe()
	app.logger.Fatalln(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
