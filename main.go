package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"github.com/techfort/wyrdtales/api"
)

func main() {
	fmt.Printf("Working out well, %v", "joe")
	e := echo.New()
	v := viper.New()
	es, err := elastic.NewClient()
	if err != nil {
		fmt.Println("Probably ES is not set up yet")
	}
	e = api.Config(e, v, es)
	e.Start(":1323")
}
