package configs

var conf = &Configs{
	Logger:      &logger{},
	Server:      &server{},
	Application: &application{},
}

var confineMap = map[string]interface{}{
	"conf/application.json": conf.Application,
	"conf/logger.json":      conf.Logger,
	"conf/server.json":      conf.Server,
}

// Configs : Holds all the Application Configs.
type Configs struct {
	Logger      *logger
	Server      *server
	Application *application
}

type application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type logger struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
