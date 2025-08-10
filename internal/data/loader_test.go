package data_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/xaaha/address-api/internal/data"
)

func writeJSONFile(t *testing.T, dir, filename string, v any) string {
	t.Helper()
	path := filepath.Join(dir, filename)
	dataBytes, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal test data: %v", err)
	}
	if err := os.WriteFile(path, dataBytes, 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	return path
}

func TestReadJSON(t *testing.T) {
	tests := []struct {
		name     string
		setupDir func(t *testing.T) string
		want     []data.Address
		wantErr  bool
	}{
		{
			name: "single valid JSON file",
			setupDir: func(t *testing.T) string {
				dir := t.TempDir()
				addresses := []data.Address{
					{
						ID:          1,
						Name:        "Glaner Hotel Café",
						Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
						Phone:       "+376879444",
						CountryCode: "AD",
						Country:     "Andorra",
					},
				}
				writeJSONFile(t, dir, "test.json", addresses)
				return dir
			},
			want: []data.Address{
				{
					ID:          1,
					Name:        "Glaner Hotel Café",
					Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
					Phone:       "+376879444",
					CountryCode: "AD",
					Country:     "Andorra",
				},
			},
			wantErr: false,
		},
		{
			name: "multiple valid JSON files",
			setupDir: func(t *testing.T) string {
				dir := t.TempDir()
				writeJSONFile(t, dir, "a.json", []data.Address{
					{
						ID:          1,
						Name:        "Glaner Hotel Café",
						Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
						Phone:       "+376879444",
						CountryCode: "AD",
						Country:     "Andorra",
					},
				})
				writeJSONFile(t, dir, "b.json", []data.Address{
					{
						ID:          2,
						Name:        "Hotel Magic",
						Address:     "Av. Doctor Mitjavila, 3, AD500 Andorra la Vella, Andorra",
						Phone:       "+376876900",
						CountryCode: "AD",
						Country:     "Andorra",
					},
				})
				return dir
			},
			want: []data.Address{
				{
					ID:          1,
					Name:        "Glaner Hotel Café",
					Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
					Phone:       "+376879444",
					CountryCode: "AD",
					Country:     "Andorra",
				},
				{
					ID:          2,
					Name:        "Hotel Magic",
					Address:     "Av. Doctor Mitjavila, 3, AD500 Andorra la Vella, Andorra",
					Phone:       "+376876900",
					CountryCode: "AD",
					Country:     "Andorra",
				},
			},
			wantErr: false,
		},
		{
			name: "non-JSON files are ignored",
			setupDir: func(t *testing.T) string {
				dir := t.TempDir()
				writeJSONFile(t, dir, "valid.json", []data.Address{
					{
						ID:          1,
						Name:        "Glaner Hotel Café",
						Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
						Phone:       "+376879444",
						CountryCode: "AD",
						Country:     "Andorra",
					},
				})
				os.WriteFile(filepath.Join(dir, "note.txt"), []byte("not json"), 0644)
				return dir
			},
			want: []data.Address{
				{
					ID:          1,
					Name:        "Glaner Hotel Café",
					Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
					Phone:       "+376879444",
					CountryCode: "AD",
					Country:     "Andorra",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid JSON format in file",
			setupDir: func(t *testing.T) string {
				dir := t.TempDir()
				os.WriteFile(filepath.Join(dir, "bad.json"), []byte("{not valid json]"), 0644)
				return dir
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "directory does not exist",
			setupDir: func(_ *testing.T) string {
				return filepath.Join(os.TempDir(), "nonexistent-dir-12345")
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setupDir(t)
			got, err := data.ReadJSON(dir)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Fatalf("ReadJSON() length = %d, want %d", len(got), len(tt.want))
				}
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("ReadJSON()[%d] = %+v, want %+v", i, got[i], tt.want[i])
					}
				}
			}
		})
	}
}
