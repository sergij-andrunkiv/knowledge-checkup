package view

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template = nil

const TEMPLATE_PATH = "templates/"

// Функція парсингу шаблонів
func GetTpl() *template.Template {
	if tpl == nil {
		tpl = template.Must(ParseTemplates(TEMPLATE_PATH))
	}

	return tpl
}

// Кастомна функція рекурсивного масового парсингу шаблонів
func ParseTemplates(path string) (*template.Template, error) {
	templ := template.New("")
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ, nil
}

// Відобразити сторінку помилки
func ErrorPage(wr io.Writer, message string) {
	GetTpl().ExecuteTemplate(wr, "error_redirect_page.html", map[string]string{"ErrorMessage": message})
}
