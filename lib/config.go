package gpm

import (
	"path/filepath"
	"runtime"
	"os"
)

type PathS struct {
	Home string
	Root string
	Base string
	Temp string
}

type ConfigS struct {
	Name  string
	Paths PathS
}

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

var Config ConfigS;

func GetConfig() ConfigS {
	if Config.Name != "" {
		return Config
	}
	home := UserHomeDir()
	var name string = "gpm"
	Config.Name = name;
	Config.Paths.Home = home;
	Config.Paths.Root = filepath.Join(Config.Paths.Home, "." + name)
	Config.Paths.Temp = filepath.Join(Config.Paths.Root, "temp")
	Config.Paths.Base = filepath.Join(Config.Paths.Root, name)
	return Config;
}