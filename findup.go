package gofindup

import (
	"path/filepath"

	"github.com/spf13/afero"
)

var fs = afero.NewOsFs()

func findIn(name, dir string, fs afero.Fs) (bool, error) {
	files, err := afero.ReadDir(fs, dir)

	if err != nil {
		return false, err
	}

	for _, f := range files {
		if name == f.Name() {
			return true, nil
		}
	}

	return false, nil
}

func findupFrom(name, dir string, fs afero.Fs) (bool, error) {
	for {
		found, err := findIn(name, dir, fs)

		if err != nil {
			return false, err
		}

		if found {
			return true, nil
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			return false, nil
		}

		dir = parent
	}
}

func findup(name string, fs afero.Fs) (bool, error) {
	return findupFrom(name, ".", fs)
}

// Recursively find a file by walking up parents in the file tree
// starting from a specific directory.
func FindupFrom(name, dir string) (bool, error) {
	return findupFrom(name, dir, fs)
}

// Recursively find a file by walking up parents in the file tree
// starting from the current working directory.
func Findup(name string) (bool, error) {
	return findup(name, fs)
}
