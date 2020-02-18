package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	DbPruebas struct {
		URLConexion string `yaml:"url_conexion"`
	} `yaml:"db_pruebas"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()

	fmt.Println(c)
	fmt.Println(c.DbPruebas)
	fmt.Println(c.DbPruebas.URLConexion)

}
