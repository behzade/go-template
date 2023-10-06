package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const (
	schemaFileLocation  = "/app/schema.sql"
	schemaTemplateFiles = "/app/sql/schema-template/*.tmpl.sql"
)

func RenderSchemaTemplate() error {
	config := GetConfig()

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

type SchemaChanges struct {
	Pending []string `json:"Pending,omitempty"`
	Applied []string `json:"Applied,omitempty"`
}

type ApplySchemaResult struct {
	Changes SchemaChanges `json:"Changes,omitempty"`
}

func (res *ApplySchemaResult) String() string {
	if res.Changes.Applied == nil && res.Changes.Pending == nil {
		return "Up to date"
	}

	var b strings.Builder

	if res.Changes.Pending != nil {
		b.WriteString("Pending Alters:\n")
		for i := range res.Changes.Pending {
			b.WriteString(res.Changes.Pending[i])
			b.WriteString("\n")
		}
	}

	if res.Changes.Applied != nil {
		b.WriteString("Applied Alters:\n")
		for i := range res.Changes.Applied {
			b.WriteString(res.Changes.Applied[i])
			b.WriteString("\n")
		}
	}

	return b.String()
}

func parseApplySchemaResult(data []byte) (*ApplySchemaResult, error) {
	var res ApplySchemaResult
	err := json.Unmarshal(data, &res)
	return &res, err
}

func SyncDBWithSchema(dryRun bool) (*ApplySchemaResult, error) {
	args := getSchemaApplyArgs()

	if dryRun {
		args = append(args, "--dry-run")
	} else {
		args = append(args, "--auto-approve")
	}

	cmd := exec.Command(
		"atlas",
		args...,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Join(errors.New(string(out)), err)
	}
	return parseApplySchemaResult(out)
}

func getSchemaApplyArgs() []string {
	config := GetConfig()
	return []string{
		"schema", "apply",
		"--url", dbURL(config),
		"--schema", config.Database.Name,
		"--dev-url", devDBURL(config),
		"--to", fmt.Sprintf("file://%v", schemaFileLocation),
		"--format", "{{ json . }}",
	}
}

func dbURL(config Config) string {
	return fmt.Sprintf(
		"mysql://%v:%v@%v:%v",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
	)
}

func devDBURL(config Config) string {
	return fmt.Sprintf(
		"mysql://%v:%v@%v_dev:%v",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
	)
}

func SyncEntityWithSchema() ([]byte, error) {
	return exec.Command("sqlc", "generate").CombinedOutput()
}
