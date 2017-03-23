package gpm

import (
	"os"
	"path/filepath"
	"runtime"
)

type PathS struct {
	Home    string
	Root    string
	Base    string
	Temp    string
	Storage string
	Config  string
}

type ConfigS struct {
	Name  string
	Paths PathS
	init  bool
}

var Config ConfigS

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func GetConfig() ConfigS {
	if Config.init == true {
		return Config
	}
	home := UserHomeDir()

	var name string = "gpm.go"
	if os.Getenv("GO_ENV") == "DEVELOPMENT" {
		Config.Paths.Home = filepath.Join("./", ".home")
	} else {
		Config.Paths.Home = home
	}
	Config.Name = name
	Config.Paths.Root = filepath.Join(Config.Paths.Home, "."+Config.Name)
	Config.Paths.Temp = filepath.Join(Config.Paths.Root, "temp")
	Config.Paths.Storage = filepath.Join(Config.Paths.Root, "storage")
	Config.Paths.Base = filepath.Join(Config.Paths.Home, name)
	Config.Paths.Config = filepath.Join(Config.Paths.Root, name+".conf.json")
	Config.init = true
	return Config
}
