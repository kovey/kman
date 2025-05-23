package cache

import (
	"fmt"
	"os"
	"strings"
)

type files struct {
	files []*file
	dir   string
}

func newFiles(path string) *files {
	return &files{dir: path}
}

func (f *files) parse() error {
	for _, fi := range f.files {
		if err := fi.parse(); err != nil {
			return err
		}
	}

	return nil
}

func (f *files) load() error {
	stat, err := os.Stat(f.dir)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		if !strings.Contains(f.dir, ".schema.") {
			return fmt.Errorf("file name format error")
		}

		fi := newFile(f.dir)
		if err := fi.load(); err != nil {
			return err
		}

		f.files = append(f.files, fi)
		return nil
	}

	fis, err := os.ReadDir(f.dir)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if !strings.Contains(fi.Name(), ".schema.") {
			continue
		}

		ff := newFile(fmt.Sprintf("%s/%s", f.dir, fi.Name()))
		if err := ff.load(); err != nil {
			return err
		}

		f.files = append(f.files, ff)
	}

	return nil
}

func (f *files) add(key, value string) {
	for _, fi := range f.files {
		if !fi.metas.has(key) {
			continue
		}

		fi.metas.add(key, value)
	}
}

func (f *files) flush(key, value string) error {
	for _, fi := range f.files {
		if !fi.metas.has(key) {
			continue
		}

		if err := fi.flush(key, value); err != nil {
			return err
		}
	}

	return nil
}
