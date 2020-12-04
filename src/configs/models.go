package configs

var conf = &Configs{
	Application: &application{},
	Logger:      &logger{},
	Path:        &path{},
	Server:      &server{},
}

var confineMap = map[string]interface{}{
	"conf/application.json": conf.Application,
	"conf/logger.json":      conf.Logger,
	"conf/path.json":        conf.Path,
	"conf/server.json":      conf.Server,
}

// Configs : Holds all the Application Configs.
type Configs struct {
	Application *application
	Logger      *logger
	Path        *path
	Server      *server
}

type application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type logger struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type path struct {
	Cover string `json:"cover"`
}

type server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
