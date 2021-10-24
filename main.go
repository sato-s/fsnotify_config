package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	A        int
	B        string
	filename string
}

func NewConfig(filename string) (*Config, error) {
	config := &Config{filename: filename}

	if err := config.loadFile(); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) loadFile() error {
	data, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), &c)
	return err
}

func main() {
	config, err := NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", config)
}
