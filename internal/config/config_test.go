// Package config
package config

import (
	"flag"
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
			name: "valid positional argument",
			args: []string{"rename-snake", "/tmp/test"},
			want: &Config{
				Path:   "/tmp/test",
				DryRun: false,
			},
			wantErr: false,
		},
		{
			name: "valid dry run flag",
			args: []string{"rename-snake", "-d", "/tmp/test"},
			want: &Config{
				Path:   "/tmp/test",
				DryRun: true,
			},
			wantErr: false,
		},
		{
			name:    "no positional argument",
			args:    []string{"rename-snake"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "dry run without positional argument",
			args:    []string{"rename-snake", "-d"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too many positional arguments",
			args:    []string{"rename-snake", "/tmp/test", "/tmp/other"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "dry run with too many positional arguments",
			args:    []string{"rename-snake", "-d", "/tmp/test", "/tmp/other"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty positional argument",
			args: []string{"rename-snake", ""},
			want: &Config{
				Path:   "",
				DryRun: false,
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
			flag.CommandLine = flag.NewFlagSet(tt.args[0], flag.ContinueOnError)

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

			if got.DryRun != tt.want.DryRun {
				t.Fatalf("DryRun = %v, want %v", got.DryRun, tt.want.DryRun)
			}
		})
	}
}
