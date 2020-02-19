package main

import (
	"fmt"

	"github.com/OrlandoHdz/yamltest/leeyaml"
)

func main() {
	var c leeyaml.Conf
	c.GetConf()

	fmt.Println(c)
	fmt.Println(c.DbPruebas)
	fmt.Println(c.DbPruebas.URLConexion)

}
