package catalog

import (
	"encoding/json"
	"fmt"

	"github.com/ShiYuandi/SoloWeave/internal/bundle"
	"github.com/ShiYuandi/SoloWeave/internal/config"
)

type Item struct {
	ID          string        `json:"id"`
	Description string        `json:"description"`
	Config      config.Config `json:"config"`
}

func List() ([]Item, error) {
	data, err := bundle.Files.ReadFile("assets/catalogs/stacks.json")
	if err != nil {
		return nil, err
	}
	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	for _, item := range items {
		if err := config.Validate(item.Config); err != nil {
			return nil, fmt.Errorf("catalog %s: %w", item.ID, err)
		}
	}
	return items, nil
}

func Preset(id string) (config.Config, error) {
	items, err := List()
	if err != nil {
		return config.Config{}, err
	}
	for _, item := range items {
		if item.ID == id {
			return item.Config, nil
		}
	}
	return config.Config{}, fmt.Errorf("unknown preset: %s", id)
}
