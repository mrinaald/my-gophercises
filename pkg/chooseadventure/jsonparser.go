package chooseadventure

import (
	"encoding/json"
	"io"
	"os"
)

type Story map[string]Chapter

type Choice struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Choice `json:"options"`
}

func ParseJson(r io.Reader) (Story, error) {
	jsonDecoder := json.NewDecoder(r)
	var story Story
	if err := jsonDecoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

func ParseJsonFromFile(filename string) (Story, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return ParseJson(jsonFile)
}
