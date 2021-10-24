package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"github.com/fsnotify/fsnotify"
)

type Config struct {
	A        int
	B        string
	filename string
	watcher  *fsnotify.Watcher
}

func NewConfig(filename string) (*Config, error) {
	config := &Config{filename: filename}

	if err := config.loadFile(); err != nil {
		return nil, err
	}
  go config.run()
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

func (c *Config) run() error {
	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()
	if err != nil {
		return err
	}
	if err := watcher.Add(c.filename); err != nil {
		return err
	}
	log.Printf("Watching %s", c.filename)

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				err := c.loadFile()
				if err != nil {
					log.Printf("Error: $s", err)
				} else {
					log.Printf("Refreshed setting from %s", c.filename)
				}
			}
		case err := <-watcher.Errors:
			log.Printf("Error: $s", err)
		}
	}
	return nil
}


func main() {
	config, err := NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", config)
}
