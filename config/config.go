package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

type Conf struct {
	API struct {
		APIKey  string `yaml:"apiKey"`
		BaseURL string `yaml:"baseURL"`
	} `yaml:"API"`
	DB struct {
		Host            string `yaml:"host"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		ConnMaxLifetime int    `yaml:"connMaxLifetime"`
	} `yaml:"DB"`
}

func Load() Conf {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	raw, err := ioutil.ReadFile(path.Join(basePath, "conf.yml"))
	HandleErr(err)
	var conf Conf
	err = yaml.Unmarshal(raw, &conf)
	HandleErr(err)
	return conf
}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
