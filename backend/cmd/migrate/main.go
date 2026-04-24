package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	migrationPath := filepath.Join("internal", "db", "migrations", "0001_init.sql")
	if _, err := os.Stat(migrationPath); err != nil {
		panic(err)
	}
	fmt.Println("migration ready:", migrationPath)
}
