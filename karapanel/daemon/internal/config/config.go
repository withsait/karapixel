package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Auth     AuthConfig     `yaml:"auth"`
	Servers  []MCServer     `yaml:"servers"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type AuthConfig struct {
	Secret   string   `yaml:"secret"`
	Users    []User   `yaml:"users"`
}

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"` // bcrypt hash
	Role     string `yaml:"role"`     // admin, viewer
}

type MCServer struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`        // velocity, paper, purpur
	ServiceName string `yaml:"serviceName"` // systemd service name
	WorkDir     string `yaml:"workDir"`     // /opt/karapixel/servers/velocity
	JarFile     string `yaml:"jarFile"`     // velocity.jar
	RconPort    int    `yaml:"rconPort"`
	RconPass    string `yaml:"rconPass"`
	MaxRAM      string `yaml:"maxRam"`      // 6G
	MinRAM      string `yaml:"minRam"`      // 2G
}

var cfg *Config

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg = &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func Get() *Config {
	return cfg
}

func GetServersConfigPath() string {
	exe, _ := os.Executable()
	return filepath.Join(filepath.Dir(exe), "configs", "servers.yml")
}
