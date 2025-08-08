package note

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/sahay-shashank/Personal-Knowledge-Manager/internal/config"
	"github.com/sahay-shashank/Personal-Knowledge-Manager/internal/utility"
)

func HandleNew(cfg *config.Config, args []string) {
	title := args[0]
	id := uuid.New().String()
	filename := fmt.Sprintf("%s_%s.md", title, id)
	now := time.Now().UTC()
	note := Note{
		ID:       id,
		Title:    title,
		Created:  now,
		Modified: now,
		Tags:     []string{},
	}
	content := buildMetadata(note)
	os.MkdirAll(cfg.StorageLocation, 0755)
	filePath := filepath.Join(cfg.StorageLocation, filename)
	os.WriteFile(filePath, []byte(content), 0644)

	err := utility.OpenEditor(cfg.Editor, filePath)
	if err != nil {
		log.Fatalf("Process ended with error: %s", err)
	}

	// to update the modified field
	updatedNote, _ := loadNoteFromFile(filePath)
	saveNoteToFile(filePath, updatedNote)

	fmt.Printf("Note Created: %s\n", filename)
}
