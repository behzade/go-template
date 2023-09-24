package main

import (
	"fmt"
	"os/exec"

	"github.com/behzade/fiber-template/internal/service"
)

func diffDBWithSchema() ([]byte, error) {
	args := getSchemaApplyArgs()
	args = append(args, "--dry-run")

	cmd := exec.Command(
		"atlas",
		args...,
	)

	return cmd.CombinedOutput()
}

func syncDBWithSchema() ([]byte, error) {
	args := getSchemaApplyArgs()
	args = append(args, "--auto-approve")

	cmd := exec.Command(
		"atlas",
		args...,
	)

	return cmd.CombinedOutput()
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
