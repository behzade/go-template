package main

import "os/exec"

func syncEntityWithSchema() ([]byte, error) {
	return exec.Command("sqlc", "generate").CombinedOutput()
}

