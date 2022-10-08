package conf

import (
	"errors"
	"io/ioutil"

	"bitbucket.org/btcid/startrack/internal/config"
	"gopkg.in/yaml.v2"
)

func LoadConfig(path string, config *config.Config) error {

	if path == "" {
		return errors.New("Path Not Set")
	}

	b, err := ioutil.ReadFile(path)
	parserYaml(b, config)
	return err
}

func parserYaml(b []byte, config *config.Config) {
	_ = yaml.Unmarshal(b, &config)
}
