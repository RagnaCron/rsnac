// Package config
package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    *Config
		wantErr bool
	}{
		{
			name: "valid single argument",
			args: []string{"rename-snake", "/tmp/test"},
			want: &Config{
				Path: "/tmp/test",
			},
			wantErr: false,
		},
		{
			name:    "no arguments",
			args:    []string{"rename-snake"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too many arguments",
			args:    []string{"rename-snake", "/tmp/test", "/tmp/other"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty path argument",
			args: []string{"rename-snake", ""},
			want: &Config{
				Path: "",
			},
			wantErr: false,
		},
	}

	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			got, err := Load()

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}

				if got != nil {
					t.Fatalf("expected nil config on error, got %#v", got)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got == nil {
				t.Fatal("expected config, got nil")
			}

			if got.Path != tt.want.Path {
				t.Fatalf("Path = %q, want %q", got.Path, tt.want.Path)
			}
		})
	}
}
