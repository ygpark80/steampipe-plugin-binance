package main

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/mitchellh/go-homedir"
)

type Config struct {
	Connection ConnectionConfig `hcl:"connection,block"`
}

type ConnectionConfig struct {
	Name      string `hcl:"name,label"`
	Plugin    string `hcl:"plugin"`
	ApiKey    string `hcl:"api_key"`
	ApiSecret string `hcl:"api_secret"`
}

func GetConfig() ConnectionConfig {
	homedir, _ := homedir.Dir()
	filename := filepath.Join(homedir, ".steampipe/config/binance.spc")
	src, _ := os.ReadFile(filename)
	file, _ := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})

	var config Config
	gohcl.DecodeBody(file.Body, nil, &config)

	return config.Connection
}
