package main

import (
	"context"
	"fmt"

	"github.com/techfort/wyrdtales/models"

	"github.com/labstack/echo"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"github.com/techfort/wyrdtales/api"
)

func main() {
	fmt.Println("Starting wyrdtales....")
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("Error: %+v", r)
			}
			fmt.Println(err.Error())
		}
	}()

	v := viper.New()
	es, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200/"))
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: %v", err.Error()))
		panic(err.Error())
	}
	exists, err := es.IndexExists(models.PostsIndex).Do(ctx)
	if err != nil || !exists {
		createdIndex, err := es.CreateIndex(models.PostsIndex).Do(ctx)
		if !createdIndex.Acknowledged {
			fmt.Println("create the fucking index will ya")
		}
		if err != nil {
			panic(err)
		}
	}
	initEcho(v, es)
}

func initEcho(v *viper.Viper, es *elastic.Client) error {
	e := echo.New()
	e = api.Config(e, v, es)
	return e.Start(":1323")
}
