package cache

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type word struct {
	begin   []byte
	content []byte
	end     []byte
}

func (w *word) reset() {
	w.begin = nil
	w.content = nil
	w.end = nil
}

func (w *word) hasBegin() bool {
	return len(w.begin) == 2 && w.begin[0] == '#' && w.begin[1] == '{'
}

func (w *word) hasEnd() bool {
	return len(w.end) == 1 && w.end[0] == '}'
}

func (w *word) isValid() bool {
	return w.hasBegin() && len(w.content) > 0 && w.hasEnd()
}

func (w *word) add(b byte) {
	if w.hasBegin() && !w.hasEnd() {
		w.content = append(w.content, b)
	}
}

type file struct {
	path    string
	metas   *metas
	content string
}

func newFile(path string) *file {
	return &file{path: path, metas: newMetas()}
}

func (f *file) parseLine(line int, data []byte) error {
	w := &word{}
	count := len(data)
	for col, b := range data {
		switch b {
		case '#':
			if col+1 < count && data[col+1] == '{' {
				w.begin = append(w.begin, b, data[col+1])
			}
		case '}':
			w.end = append(w.end, b)
			if !w.isValid() {
				return fmt.Errorf("end tag format error in file %s(%d,%d)", f.path, line, col+1)
			}
			f.metas.addInvalid(string(w.content), "")
			w.reset()
		case '{':
			if !w.hasBegin() {
				return fmt.Errorf("begin tag format error in file %s(%d,%d)", f.path, line, col+1)
			}
		default:
			w.add(b)
		}
	}

	return nil
}

func (f *file) load() error {
	c, err := os.ReadFile(f.path)
	if err != nil {
		return err
	}

	f.content = string(c)

	fi, err := os.Open(f.path)
	if err != nil {
		return err
	}
	defer fi.Close()

	reader := bufio.NewReader(fi)
	var line = 0
	for {
		buff, err := reader.ReadBytes('\n')
		line++
		if err != nil {
			if err != io.EOF {
				return err
			}

			if err := f.parseLine(line, buff); err != nil {
				return err
			}

			break
		}

		if err := f.parseLine(line, buff); err != nil {
			return err
		}
	}

	return nil
}

func (f *file) flush(key, value string) error {
	if !f.metas.add(key, value) {
		return nil
	}

	return f.parse()
}

func (f *file) parse() error {
	var content = make([]byte, len(f.content))
	copy(content, []byte(f.content))
	tmp := string(content)
	for _, meta := range f.metas.Data {
		if !meta.Valid {
			continue
		}

		tmp = strings.ReplaceAll(tmp, fmt.Sprintf("#{%s}", meta.Key), meta.Value)
	}

	return os.WriteFile(strings.ReplaceAll(f.path, ".schema", ""), []byte(tmp), 0644)
}
