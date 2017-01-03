package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	var c Conf
	c = &Config{}

	assert.Equal(c.GetListener(), "")
	assert.Equal(c.GetRootDir(), "")
	assert.Equal(c.IsDebug(), false)
}

func TestConfigFile(t *testing.T) {
	assert := assert.New(t)

	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}

	var c Conf
	c = &Config{ConfigFile: cwd + "/config_test.json"}

	cn, err := c.LoadFromFile()
	assert.Nil(err)

	assert.Equal(true, cn.IsDebug())
	assert.Equal("madness", cn.(*Config).Database.User)
}
