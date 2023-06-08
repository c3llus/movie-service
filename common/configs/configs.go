package configs

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var ConfigValues Config // Singleton

// GetConfig gets the application config.
func GetConfig() Config {
	// TODO: add config values on GSM
	err := readConfig(&ConfigValues, "./files/configs/{MTENV}-config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return ConfigValues
}

// readConfig reads the configuration from the given paths.
func readConfig(dest interface{}, paths ...string) error {
	for _, path := range paths {

		path = replacePathByEnv(path)

		// check if this path is exist
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		// load config
		ext := filepath.Ext(path)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		switch {
		case ext == ".yaml" || ext == ".yml":
			return yaml.Unmarshal(content, dest)
		}
	}
	return errors.New("no config file found")
}

func replacePathByEnv(path string) string {
	env := os.Getenv(`MTENV`)
	if env == "" {
		env = "dev"
	}

	return strings.Replace(path, "{MTENV}", env, -1)
}
