package config

import (
	"aid.dev/etherms/internal"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
)

type Conf struct {
	Api struct {
		ApiKey  string `yaml:"apiKey"`
		BaseUrl string `yaml:"baseUrl"`
	}
}

func LoadConfig() Conf {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	raw, err := ioutil.ReadFile(path.Join(basePath, "conf.yml"))
	internal.HandleErr(err)
	var conf Conf
	err = yaml.Unmarshal(raw, &conf)
	internal.HandleErr(err)
	return conf
}
