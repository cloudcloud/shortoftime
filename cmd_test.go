package main

import (
	"testing"

	"github.com/cloudcloud/shortoftime/config"
	"github.com/namsral/flag"
	"github.com/stretchr/testify/assert"
)

func TestLoadFlags(t *testing.T) {
	assert := assert.New(t)
	args := []string{
		"shortoftime",
		"-debug=false",
		"-config=",
	}
	prefix := "SHORT"
	f := flag.ContinueOnError
	c := &config.Config{}

	conf := loadFlags(args, prefix, f, c)
	assert.NotNil(conf)
}

func TestServe(t *testing.T) {
	//
}
