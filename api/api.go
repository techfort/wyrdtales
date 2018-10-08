package api

import (
	"net/http"

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
	return e
}
