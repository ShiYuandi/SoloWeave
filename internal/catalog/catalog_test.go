package catalog

import "testing"

func TestPresetIsDraftAndCustomizable(t *testing.T) {
	all, err := List()
	if err != nil || len(all) < 3 {
		t.Fatalf("catalog: %v %v", all, err)
	}
	cfg, err := Preset("api-go")
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Project.Status != "draft" || cfg.Backend == nil || cfg.Backend.Language != "go" {
		t.Fatalf("preset: %#v", cfg)
	}
	if _, err := Preset("unknown"); err == nil {
		t.Fatal("unknown preset accepted")
	}
}
