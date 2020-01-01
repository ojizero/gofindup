package gofindup

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

var fs = afero.NewOsFs()

func hasFile(name, dir string, fs afero.Fs) (bool, error) {
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

func findupFrom(name, dir string, fs afero.Fs) (string, error) {
	for {
		found, err := hasFile(name, dir, fs)

		if err != nil {
			return "", err
		}

		if found {
			return filepath.Join(dir, name), nil
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			return "", nil
		}

		dir = parent
	}
}

func findup(name string, fs afero.Fs) (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return findupFrom(name, cwd, fs)
}

// Recursively find a file by walking up parents in the file tree
// starting from a specific directory.
func FindupFrom(name, dir string) (string, error) {
	return findupFrom(name, dir, fs)
}

// Recursively find a file by walking up parents in the file tree
// starting from the current working directory.
func Findup(name string) (string, error) {
	return findup(name, fs)
}
