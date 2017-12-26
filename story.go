package cyoa

import (
	"encoding/json"
	"io"
)

// JSONStory takes reader and decodes json into Story.
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story is a map of chapter names to Chapters.
type Story map[string]Chapter

// Chapter struct has Title, Paragraphs slice, and Options slice.
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option Struct has Text and Chapter.
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
