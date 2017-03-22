package gpm

import (
	"os"
)

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Prepare() {
	config := GetConfig()

	existRoot, existRootErr := pathExist(config.Paths.Root)

	if (existRoot == false || existRootErr != nil) {
		os.Mkdir(config.Paths.Root, 0777)
	}

	existTemp, existTempErr := pathExist(config.Paths.Temp)

	if (existTemp == false || existTempErr != nil) {
		os.Mkdir(config.Paths.Temp, 0777)
	}

	existBase, existBaseErr := pathExist(config.Paths.Base)

	if (existBase == false || existBaseErr != nil) {
		os.Mkdir(config.Paths.Base, 0777)
	}

	existStorage, existStorageErr := pathExist(config.Paths.Storage)

	if (existStorage == false || existStorageErr != nil) {
		os.Mkdir(config.Paths.Storage, 0777)
	}

	existConfig, existConfigErr := pathExist(config.Paths.Config)

	if (existConfig == false || existConfigErr != nil) {
		globalConfig, err := os.Create(config.Paths.Config)
		if err == nil {
			globalConfig.WriteString("{\n}")
		}
	}
}