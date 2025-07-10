package base

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	Port int
	Dir  string
	Help bool
)

func ParseFlags() error {
	flag.StringVar(&Dir, "dir", "./data", "the directory of static file to host")
	flag.BoolVar(&Help, "help", false, "show help")
	flag.IntVar(&Port, "port", 8080, "port to serve on")

	flag.Parse()

	if Help {
		PrintHelp()
		os.Exit(0)
	}

	log.Printf("Serving files from directory: %s", Dir)
	log.Printf("Listening on port: %d", Port)

	if Port < 0 || Port > 65535 {
		log.Fatal("Invalid port number")
	}

	if Dir == "" {
		log.Fatal("Error: directory is required")
	}

	if strings.Contains(Dir, "home/") || strings.Contains(Dir, "arch/") || strings.Contains(Dir, "go.mod") || strings.Contains(Dir, "main.go") || strings.Contains(Dir, "README.md") || strings.Contains(Dir, "base/") || strings.Contains(Dir, "handlers/") || strings.Contains(Dir, "helpers/") || strings.Contains(Dir, "launch/") {
		log.Fatalf("directory is not allowed. Please provide a valid directory.")
	}

	return nil
}

func PrintHelp() {
	fmt.Println(`Simple Storage Service.

**Usage:**
    triple-s [-port <N>] [-dir <S>]  
    triple-s --help

**Options:**
- --help     Show this screen.
- --port N   Port number
- --dir S    Path to the directory`)
}
