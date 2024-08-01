package render

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// TemplateRegistry is a custom HTML template renderer for Echo framework
type TemplateRegistry struct {
	templates *template.Template
}

// Render implements the Echo Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var TmplRegistry *TemplateRegistry

type PageData struct {
	Title           string
	Username        string
	PartialTemplate string
	Data            interface{}
}

type TemplateName string

const (
	IndexTemplate          TemplateName = "index"
	SignInTemplate         TemplateName = "sign-in"
	CreateAccountTemplate  TemplateName = "create-account"
	ProfileTemplate        TemplateName = "profile"
	CreateQuestionTemplate TemplateName = "create-question"
	GamePlay               TemplateName = "gameplay"
)

// Initialize templates
func InitTemplates() {
	// Determine the working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	basePath := "ui/templates/"
	partials := []string{
		"base.html",
		"index.html",
		"sign-in-form.html",
		"sign-in-success.html",
	}

	// Create a base layout template
	templates := template.New("base").Funcs(template.FuncMap{})

	// Parse base layout
	templates = template.Must(templates.ParseFiles(dir + "/" + basePath + "base.html"))

	// Parse all partial templates (blocks)
	for _, partial := range partials {
		fmt.Println(partial)
		templates = template.Must(templates.ParseFiles(dir + "/" + basePath + partial))
	}

	// Set up the global template registry
	TmplRegistry = &TemplateRegistry{templates: templates}
}

// Render full or partial templates based on the HX-Request header
func RenderTemplate(c echo.Context, partialTemplate string, data interface{}) error {
	hx := c.Request().Header.Get("Hx-Request") == "true"
	if hx {
		err := c.Render(http.StatusOK, partialTemplate, data)
		if err != nil {
			fmt.Println("Error rendering template from hx-request:", err)
			return err
		}
		return err
	}

	pageData := PageData{
		Title:           "Marren Games - " + partialTemplate,
		PartialTemplate: partialTemplate,
		Data:            data,
	}
	err := c.Render(http.StatusOK, "base", pageData)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
