package main

import (
	seeder "bwastartup/db/seeders"
	"fmt"
	"os"
)

func main() {
	// Mendapatkan argumen baris perintah
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run seeder.go <function>")
		return
	}
	// Menjalankan fungsi berdasarkan argumen baris perintah
	switch args[1] {
	case "SeedUser":
		// go run db/seeder.go SeedUser
		seeder.SeedUser()
	default:
		fmt.Println("Unknown function")
	}
}
