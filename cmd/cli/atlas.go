package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/behzade/fiber-template/internal/service"
)

type Changes struct {
	Pending []string `json:"Pending,omitempty"`
	Applied []string `json:"Applied,omitempty"`
}

type ApplySchemaResult struct {
	Changes Changes `json:"Changes,omitempty"`
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

func diffDBWithSchema() (*ApplySchemaResult, error) {
	args := getSchemaApplyArgs()
	args = append(args, "--dry-run")

	cmd := exec.Command(
		"atlas",
		args...,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return parseApplySchemaResult(out)
}

func syncDBWithSchema() (*ApplySchemaResult, error) {
	args := getSchemaApplyArgs()
	args = append(args, "--auto-approve")

	cmd := exec.Command(
		"atlas",
		args...,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return parseApplySchemaResult(out)
}

func getSchemaApplyArgs() []string {
	config := service.GetConfig()
	return []string{
		"schema", "apply",
		"--url", dbURL(config),
		"--schema", config.Database.Name,
		"--dev-url", devDBURL(config),
		"--to", fmt.Sprintf("file://%v", schemaFileLocation),
		"--format", "{{ json . }}",
	}
}

func dbURL(config service.Config) string {
	return fmt.Sprintf(
		"mysql://%v:%v@%v:%v",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
	)
}

func devDBURL(config service.Config) string {
	return fmt.Sprintf(
		"mysql://%v:%v@%v_dev:%v",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
	)
}
