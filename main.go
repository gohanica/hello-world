package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

//main関数
func main() {

//　github連携
//
// 俺も追加したよー
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("html/*.html")),
	}
	e.Renderer = t
	//e.POST("/hello", Hello)

	e.GET("/hello:id", Hello)
	e.Logger.Fatal(e.Start(":8080"))

}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)

}
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
  return c.String(http.StatusOK, id)

func Hello(c echo.Context) error {
	
	return c.Render(http.StatusOK, "hello", "Worldaaaaaaaaa"+id)
}
