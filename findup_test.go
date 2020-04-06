package gofindup

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

type given struct {
	file string
	base string
}

type expect struct {
	found bool
	path  string
	err   string
}

type finderAssertion struct {
	given
	expect
}

type finderAssertions []finderAssertion

type finderFunc func(string, string, afero.Fs) (bool, error)
type findupFunc func(string, string, afero.Fs) (string, error)

var fakefs = afero.NewMemMapFs()

func fakeReadDir(name string) ([]os.FileInfo, error) {
	return afero.ReadDir(fakefs, name)
}

func init() {
	// Build fake file system

	fakefs.MkdirAll("/test", 0755)

	fakefs.MkdirAll("/test/a", 0755)
	fakefs.MkdirAll("/test/a/ab", 0755)
	afero.WriteFile(fakefs, "/test/a/ab/f.txt", []byte("some mock file"), 0644)
	fakefs.MkdirAll("/test/a/ac", 0755)
	afero.WriteFile(fakefs, "/test/a/ad.txt", []byte("some mock file"), 0644)

	fakefs.MkdirAll("/test/b", 0755)
	fakefs.MkdirAll("/test/b/bc", 0755)
	afero.WriteFile(fakefs, "/test/b/bc/f.txt", []byte("some mock file"), 0644)
	afero.WriteFile(fakefs, "/test/b/bd.txt", []byte("some mock file"), 0644)

	fakefs.MkdirAll("/test/c", 0755)
	fakefs.MkdirAll("/test/c/cd", 0755)
	afero.WriteFile(fakefs, "/test/c/cd/f.txt", []byte("some mock file"), 0644)

	afero.WriteFile(fakefs, "/test/t.txt", []byte("some mock file"), 0644)
	afero.WriteFile(fakefs, "r.txt", []byte("some mock file"), 0644)
}

func TestFindIn(t *testing.T) {
	assertions := finderAssertions{
		{given{"ab", "/test/a"}, expect{found: true, err: ""}},
		{given{"ac", "/test/a"}, expect{found: true, err: ""}},
		{given{"ae", "/test/a"}, expect{found: false, err: ""}},
		{given{"de", "/test/d"}, expect{found: false, err: "open " + filepath.Join("/", "test", "d") + ": file does not exist"}},
	}

	for _, a := range assertions {
		found, err := hasFile(a.given.file, a.given.base, fakeReadDir)

		if a.expect.found {
			assert.True(t, found)
		} else {
			assert.False(t, found)
		}

		if a.expect.err == "" {
			assert.Nil(t, err)
		} else {
			assert.EqualError(t, err, a.expect.err)
		}
	}
}

func TestFindUpFrom(t *testing.T) {
	assertions := finderAssertions{
		{given{"ab", "/test/a"}, expect{path: filepath.Join("/", "test", "a", "ab"), err: ""}},
		{given{"ab", "/test/a/ac"}, expect{path: filepath.Join("/", "test", "a", "ab"), err: ""}},
		{given{"t.txt", "/test/a/ac"}, expect{path: filepath.Join("/", "test", "t.txt"), err: ""}},
		{given{"r.txt", "/test/a/ac"}, expect{path: filepath.Join("/", "r.txt"), err: ""}},
		{given{"r.txt", "/test/d"}, expect{path: "", err: "open " + filepath.Join("/", "test", "d") + ": file does not exist"}},
	}

	for _, a := range assertions {
		found, err := findupFrom(a.given.file, a.given.base, fakeReadDir)

		assert.Equal(t, found, a.expect.path)

		if a.expect.err == "" {
			assert.Nil(t, err)
		} else {
			assert.EqualError(t, err, a.expect.err)
		}
	}
}
