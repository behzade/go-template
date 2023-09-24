package main

import (
	"os"
	"text/template"

	"github.com/behzade/fiber-template/internal/service"
)

const (
	schemaFileLocation  = "/app/sql/schema.sql"
	schemaTemplateFiles = "/app/sql/schema-template/*.tmpl.sql"
)

func renderTemplate() error {
	config := service.GetConfig()

	t, err := template.ParseGlob(schemaTemplateFiles)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(schemaFileLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(f, "schema.tmpl.sql", map[string]string{
		"schema": config.Database.Name,
	})
	if err != nil {
		return err
	}

	return f.Close()
}
