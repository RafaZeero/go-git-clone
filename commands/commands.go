package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	fmt.Println("You choose init")

	// Check if .zro already exists
	_, err := os.Stat(".zro")
	if !os.IsNotExist(err) {
		log.Fatalf("Already a zro repository")
		return
	}

	// Create .zro folder
	if err := os.Mkdir(".zro", 0755); err != nil {
		log.Fatalf("failed to initialize zro repository")
		return
	}

	// Populate .zro folder with some content
	if err := func() error {
		// Try to create objects folder
		if err := os.Mkdir(filepath.Join(".", ".zro", "objects"), 0755); err != nil {
			return err
		}
		// Try to create index file
		file, err := os.OpenFile(filepath.Join(".", ".zro", "index"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		defer file.Close()

		// Try to create HEAD file
		head, err := os.OpenFile(filepath.Join(".", ".zro", "HEAD"), os.O_CREATE, 0644)
		if err != nil {
			return err
		}

		defer head.Close()

		return nil
	}(); err != nil {
		// If something goes wrong, remove .zro folder
		err2 := os.RemoveAll(".zro")
		if err2 != nil {
			log.Fatalf("failed to initialize zro repository: corrupted .zro folder: %s", err)
			return
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to initialize zro repository: %s", err)
		return
	}
	fmt.Printf("Initialized empty zro repository in %s\n", filepath.Join(pwd, ".zro"))

}
