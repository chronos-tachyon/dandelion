package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type File struct {
	Ports map[string]*Port `json:"ports"`
}

func (file *File) Reset() {
	for name := range file.Ports {
		delete(file.Ports, name)
	}
}

func (file *File) Merge(other *File) {
	n := len(other.Ports)
	if n > 0 {
		if file.Ports == nil {
			if n < 16 {
				n = 16
			}
			file.Ports = make(map[string]*Port, n)
		}
		for name, port := range other.Ports {
			if port == nil {
				delete(file.Ports, name)
			} else {
				file.Ports[name] = port
			}
		}
	}
}

func (file *File) Load(nodeName string, configHome string, configDirs []string) error {
	fn := func(baseName string) Format {
		var format Format
		rest, ok := format.Identify(baseName)
		if ok && rest == nodeName {
			return format
		}
		return Skip
	}

	var err error
	n := len(configDirs)
	for err == nil && n > 0 {
		n--
		err = file.LoadFiles(configDirs[n], fn)
	}
	if err == nil {
		err = file.LoadFiles(configHome, fn)
	}
	return err
}

func (file *File) LoadFiles(root string, fn func(string) Format) error {
	fn2 := func(fileName string, fileDent fs.DirEntry, err error) error {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		if err != nil {
			return err
		}
		if fileName == root {
			return nil
		}

		fileType := (fileDent.Type() & fs.ModeType)
		if fileType == fs.ModeDir {
			return fs.SkipDir
		}
		if fileType == fs.ModeSymlink {
			fi, err := os.Stat(fileName)
			if errors.Is(err, fs.ErrNotExist) {
				return nil
			}
			if err != nil {
				return fmt.Errorf("failed to stat %q: %w", fileName, err)
			}
			fileType = (fi.Mode() & fs.ModeType)
		}
		if fileType != 0 {
			return nil
		}

		format := fn(filepath.Base(fileName))
		if format == Skip {
			return nil
		}

		log.Debug().Str("path", fileName).Stringer("format", format).Msg("load configuration file")

		raw, err := os.ReadFile(fileName)
		if err != nil {
			return fmt.Errorf("failed to read contents of %q from disk: %w", fileName, err)
		}

		var tmp File
		err = format.Load(&tmp, bytes.NewReader(raw))
		if err != nil {
			return fmt.Errorf("failed to parse contents of %q as %v: %w", fileName, format, err)
		}

		file.Merge(&tmp)
		return nil
	}

	return filepath.WalkDir(root, fn2)
}
