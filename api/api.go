package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/techfort/wyrdtales/repository"
	"github.com/techfort/wyrdtales/usecases"

	"github.com/olivere/elastic"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

// Context is the API context
type Context struct {
	echo.Context
	Viper   *viper.Viper
	Elastic *elastic.Client
}

// GetRoutes returns all the mapped routes
func GetRoutes() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		"test": func(c echo.Context) error {
			return c.JSONBlob(http.StatusOK, []byte(`{"message":"ok"}`))
		},
		"posts/story/:id": func(c echo.Context) error {
			id := c.Param("id")
			cc := c.(*Context)
			repo := repository.NewRepoFactory(cc.Elastic)
			postrepo := repo.GetPostRepository()
			pu := usecases.NewPostUseCase(postrepo)
			post, err := pu.GetPost(id)
			if err != nil {
				return cc.JSONBlob(http.StatusInternalServerError, []byte(fmt.Sprintf("%v", err.Error())))
			}
			jp, err := json.Marshal(post)
			if err != nil {
				return cc.JSONBlob(http.StatusInternalServerError, []byte("error marshaling"))
			}
			return cc.JSONBlob(http.StatusOK, jp)
		},
		"posts/testsave": func(c echo.Context) error {
			cc := c.(*Context)
			repo := repository.NewRepoFactory(cc.Elastic)
			postrepo := repo.GetPostRepository()
			pu := usecases.NewPostUseCase(postrepo)
			post := `{ "authorid": "joe", "title": "some post again", "body": "some body", "category": "story", "status": "draft"}`
			ir, err := pu.SavePost(post)
			fmt.Println(fmt.Sprintf("IR: %+v", ir))
			return err
		},
	}
}

// PostRoutes returns POST routes
func PostRoutes() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		"posts/story": func(c echo.Context) error {
			cc := c.(*Context)
			repo := repository.NewRepoFactory(cc.Elastic)
			postrepo := repo.GetPostRepository()
			pu := usecases.NewPostUseCase(postrepo)
			req := cc.Request()
			bytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				return cc.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
			}
			ir, err := pu.SavePost(string(bytes))
			if err != nil {
				return cc.JSONBlob(http.StatusInternalServerError, []byte(err.Error()))
			}
			return cc.JSONBlob(http.StatusOK, []byte(fmt.Sprintf(`{ "message": "post saved with id %v"`, ir.Id)))
		},
	}
}

// Config returns an echo instance
func Config(e *echo.Echo, v *viper.Viper, es *elastic.Client) *echo.Echo {
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c, v, es}
			return h(cc)
		}
	})
	for route, handler := range GetRoutes() {
		e.GET(route, handler)
	}
	for route, handler := range PostRoutes() {
		e.POST(route, handler)
	}
	return e
}
