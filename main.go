package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// hash()
	var cmd string

	if len(os.Args) == 1 {
		cmd = "help"
	} else {
		cmd = strings.ToLower(os.Args[1])
	}

	switch cmd {
	case "init":
		fmt.Println("You choose init")
	case "add":
		fmt.Println("You choose add")

		if len(os.Args) < 3 {
			fmt.Println("Select file")
			return
		}
		h := sha256.New()

		f, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		bs := h.Sum(nil)

		fmt.Println("file:", os.Args[2])
		fmt.Printf("%x\n", bs)
	case "commit":
		fmt.Println("You choose commit")
	case "status":
		fmt.Println("You choose status")
	case "help":
		fallthrough
	default:
		instructions()
	}

}

func instructions() {
	fmt.Print("You can choose between these options\n\n\n")

	fmt.Println("init    --  Start your repo")
	fmt.Println("add     --  Add a file to your repo")
	fmt.Println("commit  --  Save your changes")
	fmt.Println("status  --  Current state of your files")
	fmt.Print("\n\n")
	fmt.Println("-----------------")
}

func hash() {
	const (
		input1 = "The tunneling gopher digs downwards, "
		input2 = "unaware of what he will find."
	)

	h := sha256.New()

	h.Write([]byte(input1))

	bs := h.Sum(nil)

	fmt.Println("string:", input1)
	fmt.Printf("%x\n", bs)
}
