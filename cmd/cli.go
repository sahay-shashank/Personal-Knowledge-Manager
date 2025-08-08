package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sahay-shashank/Personal-Knowledge-Manager/internal/config"
	note "github.com/sahay-shashank/Personal-Knowledge-Manager/internal/note"
)

func NewCli() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()
	if len(flag.Args()) < 2 {
		showHelp()
		os.Exit(1)
	}
	cfg, usedPath, err := config.ParseConfig("pkm", configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration from file \"%s\": %s", usedPath, err)
	}
	cmd := flag.Args()[0]
	args := flag.Args()[1:]
	switch cmd {
	case "new":
		note.HandleNew(cfg, args)
	}
}

func showHelp() {
	fmt.Println(`pkm: A Personal Knowledge Manager based on Zettelkasten

Usage:
  pkm <command> [arguments] [flags]

Available Command:
  new			Create a new note
  edit			Edit an existing note
  delete		Delete a note
  show			Display a note's content
  list			List all the notes
  search		A full text search across all the notes
  tags			List all the tags in use
  links			Show all the backward and forward links
  sync			Rebuild search index
  config		Get or set the current configuration
  help			Show help information

Global Flags:
  --config		Path to a configuration file

Use "pkm help <command> for detailed information regarding a specific command.

	`)
}

/*
Yet to implement:
  edit			Edit an existing note
  delete		Delete a note
  show			Display a note's content
  list			List all the notes
  search		A full text search across all the notes
  tags			List all the tags in use
  links			Show all the backward and forward links
  sync			Rebuild search index
  config		Get or set the current configuration
  help			Show help information
*/
