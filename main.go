package main

import (
  "gopkg.in/yaml.v2"
	"io/ioutil"
  "fmt"
)

type Config struct {
  A int
  B string
}

func NewConfig(filename string) (*Config) {
	data, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }

  c := Config{}
	err = yaml.Unmarshal([]byte(data), &c)
  if err != nil {
    panic(err)
  }
  return &c
}

func main(){
  config := NewConfig("config.yaml")
  fmt.Printf("%+v", config)

}

