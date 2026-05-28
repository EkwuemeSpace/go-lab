package main

import (
	"testing"
)

func TestParseBannerFile(t *testing.T) {
	data, err := readBannerFile("banners/standard.txt")
	if err != nil {
		t.Skip("banner file not found — skipping test")
	}

	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{name: "valid file", content: data, wantErr: false},
		{name: "empty file", content: "", wantErr: true},
		{name: "corrupt file", content: "bad content", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charMap, err := parseBannerFile(tt.content)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.wantErr {
				if len(charMap) != 95 {
					t.Errorf("expected 95 characters, got %d", len(charMap))
				}
				rows, ok := charMap['H']
				if !ok {
					t.Errorf("expected charMap to contain 'H'")
				}
				if len(rows) != 8 {
					t.Errorf("expected 8 rows for 'H', got %d", len(rows))
				}
			}
		})
	}
}
