package leeyaml

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Conf parametros del archivo
type Conf struct {
	DbPruebas struct {
		URLConexion string `yaml:"url_conexion"`
	} `yaml:"db_pruebas"`
}

// GetConf obtiene los parametros del archivo "test.yaml"
func (c *Conf) GetConf() *Conf {

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
