package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRenderString(t *testing.T) {
	// setup — load charMap once for all tests
	data, err := readBannerFile("banners/standard.txt")
	if err != nil {
		t.Skip("banner file not found — skipping test")
	}
	charMap, err := parseBannerFile(data)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "single char", input: "H", wantErr: false},
		{name: "multiple chars", input: "Hi", wantErr: false},
		{name: "invalid char", input: "Hello🌍", wantErr: true},
		{name: "newline in input", input: `Hello\nWorld`, wantErr: false},
		{name: "empty segment", input: `\n`, wantErr: false},
		{name: "space character", input: " ", wantErr: false},
		{name: "numbers", input: "123", wantErr: false},
		{name: "special chars", input: "!@#", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := renderString(&buf, charMap, tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.wantErr && buf.Len() == 0 {
				t.Errorf("expected output, got empty buffer")
			}
			if tt.name == "single char" {
				lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
				if len(lines) != 8 {
					t.Errorf("expected 8 lines for single char, got %d", len(lines))
				}
			}
		})
	}
}
