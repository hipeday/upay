package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Database *Database `yaml:"database"`
	Server   *Server   `yaml:"server"`
	Logger   *Logging  `yaml:"logger"`
}

type Database struct {
	MySQL *MySQL `yaml:"mysql"`
}

type MySQL struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int16  `yaml:"port"`
}

type Server struct {
	IP   string `yaml:"ip"`
	Port int16  `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Logging struct {
	// Encoding can be one "json" or "console". Defaults to "console"
	Encoding string `yaml:"encoding"`

	// Level configures the log level
	Level string `yaml:"level"`

	// Colors configures if color output should be enabled
	Colors *bool `yaml:"colors"`

	// time format
	TimeFormat string `yaml:"time_format"`
}

var (
	cfg  Config
	once sync.Once
)

func GetCfg() Config {
	once.Do(func() {
		loadConfig()
	})
	return cfg
}

func loadConfig() {
	var (
		env        = flag.String("e", "", "Set the environment")
		configFile = "conf/config%v.yaml"
	)

	flag.Parse()

	if env != nil && (*env) != "" {
		configFile = fmt.Sprintf(configFile, "-"+*env)
	}
	contains := strings.Contains(configFile, "%v")
	if contains {
		configFile = fmt.Sprintf(configFile, "")
	}
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Failed to close config file: %v", err)
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
}
