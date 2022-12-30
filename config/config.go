package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type File struct {
	path       string
	userConfig any
}

func New(fp string, c any) (*File, error) {
	ap, err := filepath.Abs(fp)
	if err != nil {
		return nil, err
	}
	f := &File{
		path:       ap,
		userConfig: c,
	}
	return f, f.loadData()
}

func (f *File) loadData() error {
	cf, err := os.OpenFile(f.path,os.O_RDONLY,0644)
	if err != nil {
		return err
	}
	defer cf.Close()
	jd := json.NewDecoder(cf)
	return jd.Decode(f.userConfig)

}
