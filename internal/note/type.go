package note

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Note struct {
	ID       string    `yaml:"id"`
	Title    string    `yaml:"title"`
	Created  time.Time `yaml:"created"`
	Modified time.Time `yaml:"modified"`
	Tags     []string  `yaml:"tags"`
	Content  string    `yaml:"-"`
}

func buildMetadata(note Note) string {
	meta, err := yaml.Marshal(note)
	if err != nil {
		return fmt.Sprintf("---\nid: %s\ntitle: %q\n---\n\n", note.ID, note.Title)
	}
	return fmt.Sprintf("---\n%s---\n\n", string(meta))
}

func loadNoteFromFile(path string) (*Note, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	str := string(data)

	// Ensure it starts with front-matter delimiter
	if !strings.HasPrefix(str, "---\n") {
		return nil, errors.New("missing front-matter")
	}
	parts := strings.SplitN(str, "---\n", 3)
	if len(parts) < 3 {
		return nil, errors.New("invalid front-matter format")
	}

	yamlPart := parts[1]
	contentPart := parts[2]

	contentPart = strings.TrimPrefix(contentPart, "\n")

	var note Note
	if err := yaml.Unmarshal([]byte(yamlPart), &note); err != nil {
		return nil, err
	}
	note.Content = contentPart
	log.Print(note)
	return &note, nil
}

func saveNoteToFile(path string, note *Note) error {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	// Update modified timestamp
	note.Modified = time.Now().UTC()

	// Marshal metadata
	metaBytes, err := yaml.Marshal(note)
	if err != nil {
		return err
	}
	full := strings.Builder{}
	full.WriteString("---\n")
	full.Write(metaBytes)
	full.WriteString("---\n\n")
	full.WriteString(note.Content)

	return os.WriteFile(path, []byte(full.String()), 0644)
}
