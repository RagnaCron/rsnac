// Package normalize
package normalize

import "testing"

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "simple filename",
			in:   "Test Hello.txt",
			want: "test_hello.txt",
		},
		{
			name: "uppercase extension",
			in:   "MyFile.TXT",
			want: "myfile.txt",
		},
		{
			name: "multiple dots",
			in:   "archive.tar.gz",
			want: "archive_tar.gz",
		},
		{
			name: "no extension",
			in:   "My File",
			want: "my_file",
		},
		{
			name: "hidden file",
			in:   ".gitignore",
			want: ".gitignore",
		},
		{
			name: "hidden file with extension",
			in:   ".My Config.json",
			want: "my_config.json",
		},
		{
			name: "mixed separators",
			in:   "My---File___Name.txt",
			want: "my_file_name.txt",
		},
		{
			name: "leading separators",
			in:   "---My File.txt",
			want: "my_file.txt",
		},
		{
			name: "trailing separators",
			in:   "My File---.txt",
			want: "my_file.txt",
		},
		{
			name: "numbers preserved",
			in:   "File 123 Name.txt",
			want: "file_123_name.txt",
		},
		{
			name: "only separators",
			in:   "---___.txt",
			want: ".txt",
		},
		{
			name: "empty string",
			in:   "",
			want: "",
		},
		{
			name: "extension only",
			in:   ".txt",
			want: ".txt",
		},
		{
			name: "multiple spaces",
			in:   "My   File   Name.txt",
			want: "my_file_name.txt",
		},
		{
			name: "tabs and newlines",
			in:   "My\tFile\nName.txt",
			want: "my_file_name.txt",
		},
		{
			name: "non ascii chars",
			in:   "Héllo 世界.txt",
			want: "h_llo.txt",
		},
		{
			name: "already normalized",
			in:   "already_normalized.txt",
			want: "already_normalized.txt",
		},
		{
			name: "trailing dot",
			in:   "filename.",
			want: "filename.",
		},
		{
			name: "double extension uppercase",
			in:   "Backup.TAR.GZ",
			want: "backup_tar.gz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSnakeCase(tt.in)

			if got != tt.want {
				t.Fatalf("ToSnakeCase(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "simple words",
			in:   "hello world",
			want: "hello_world",
		},
		{
			name: "multiple spaces",
			in:   "hello   world",
			want: "hello_world",
		},
		{
			name: "dashes",
			in:   "hello-world",
			want: "hello_world",
		},
		{
			name: "mixed separators",
			in:   "hello---___***world",
			want: "hello_world",
		},
		{
			name: "leading separators",
			in:   "---hello",
			want: "hello",
		},
		{
			name: "trailing separators",
			in:   "hello---",
			want: "hello",
		},
		{
			name: "leading and trailing separators",
			in:   "---hello---",
			want: "hello",
		},
		{
			name: "numbers preserved",
			in:   "hello123world",
			want: "hello123world",
		},
		{
			name: "numbers separated",
			in:   "hello 123 world",
			want: "hello_123_world",
		},
		{
			name: "only separators",
			in:   "---___***",
			want: "",
		},
		{
			name: "empty string",
			in:   "",
			want: "",
		},
		{
			name: "single underscore collapse",
			in:   "hello_____world",
			want: "hello_world",
		},
		{
			name: "already normalized",
			in:   "hello_world",
			want: "hello_world",
		},
		{
			name: "uppercase ascii ignored",
			in:   "HELLO WORLD",
			want: "",
		},
		{
			name: "mixed case",
			in:   "hello WORLD",
			want: "hello",
		},
		{
			name: "non ascii chars",
			in:   "héllo世界",
			want: "h_llo",
		},
		{
			name: "dots",
			in:   "hello.world.test",
			want: "hello_world_test",
		},
		{
			name: "tabs and newlines",
			in:   "hello\tworld\nagain",
			want: "hello_world_again",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(tt.in)

			if got != tt.want {
				t.Fatalf("normalize(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
