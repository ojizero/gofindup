package gofindup

import (
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
	err   string
}

type assertion struct {
	given  given
	expect expect
}

type assertions []assertion

type finderFunc func(string, string, afero.Fs) (bool, error)

var fakefs = afero.NewMemMapFs()

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

func testAssertions(t *testing.T, as assertions, fn finderFunc) {
	for _, a := range as {
		found, err := fn(a.given.file, a.given.base, fakefs)

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

func TestFindIn(t *testing.T) {
	assertions := assertions{
		{given{"ab", "/test/a"}, expect{true, ""}},
		{given{"ac", "/test/a"}, expect{true, ""}},
		{given{"ae", "/test/a"}, expect{false, ""}},
		{given{"de", "/test/d"}, expect{false, "open " + filepath.Join("/", "test", "d") + ": file does not exist"}},
	}

	testAssertions(t, assertions, findIn)
}

func TestFindUpFrom(t *testing.T) {
	assertions := assertions{
		{given{"ab", "/test/a"}, expect{true, ""}},
		{given{"ab", "/test/a/ac"}, expect{true, ""}},
		{given{"t.txt", "/test/a/ac"}, expect{true, ""}},
		{given{"r.txt", "/test/a/ac"}, expect{true, ""}},
		{given{"r.txt", "/test/d"}, expect{false, "open " + filepath.Join("/", "test", "d") + ": file does not exist"}},
	}

	testAssertions(t, assertions, findupFrom)
}
