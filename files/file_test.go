package files

import (
	"github.com/spf13/afero"
	"os"
	"testing"
	"time"
)

func TestFileInfo_GetSize(t *testing.T) {
	type fields struct {
		Listing   *Listing
		Fs        afero.Fs
		Path      string
		Name      string
		Size      int64
		Extension string
		ModTime   time.Time
		Mode      os.FileMode
		IsDir     bool
		IsSymlink bool
		Type      string
		Subtitles []string
		Content   string
		Checksums map[string]string
		Token     string
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "empty directory",
			fields: fields{
				IsDir: true,
			},
			want: 0,
		},
		{
			name: "normal directory",
			fields: fields{
				IsDir: true,
				Listing: &Listing{
					Items: []*FileInfo{
						&FileInfo{
							IsDir: false,
							Size: 32,
						},
					},
				},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FileInfo{
				Listing:   tt.fields.Listing,
				Fs:        tt.fields.Fs,
				Path:      tt.fields.Path,
				Name:      tt.fields.Name,
				Size:      tt.fields.Size,
				Extension: tt.fields.Extension,
				ModTime:   tt.fields.ModTime,
				Mode:      tt.fields.Mode,
				IsDir:     tt.fields.IsDir,
				IsSymlink: tt.fields.IsSymlink,
				Type:      tt.fields.Type,
				Subtitles: tt.fields.Subtitles,
				Content:   tt.fields.Content,
				Checksums: tt.fields.Checksums,
				Token:     tt.fields.Token,
			}
			if got := s.GetSize(tt.args.size); got != tt.want {
				t.Errorf("DirSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
