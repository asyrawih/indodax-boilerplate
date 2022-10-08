package conf

import (
	"testing"

	"bitbucket.org/btcid/startrack/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	var config config.Config
	t.Run("Should Correct Loadconfig", func(t *testing.T) {
		path := "../../etc/startrack.yaml"

		err := LoadConfig(path, &config)
		assert.NoError(t, err)
	})

	t.Run("Invalid Path", func(t *testing.T) {
		path := ""
		err := LoadConfig(path, &config)
		assert.Error(t, err)
	})

}
