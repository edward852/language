package main

import (
	"bufio"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/pkg/profile"
)

type bytereader struct {
	buf [1]byte
	r   io.Reader
}

func (b *bytereader) next() (rune, error) {
	_, err := b.r.Read(b.buf[:])
	return rune(b.buf[0]), err
}

func countWords(f *os.File) {
	br := bytereader{
		r: bufio.NewReader(f),
		//r: f,
	}

	words := 0
	inword := false
	for {
		r, err := br.next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", f.Name(), err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}

	log.Printf("%q: %d words\n", f.Name(), words)
	return
}

func findTxt(root, path string, d fs.DirEntry, err error) error {
	if path == root || err != nil {
		return err
	}

	if d.IsDir() {
		return fs.SkipDir
	}

	if !strings.HasSuffix(d.Name(), ".txt") {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		log.Printf("could not open file %q: %v", path, err)
		return err
	}

	countWords(f)

	return nil
}

func main() {
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	err := filepath.WalkDir(os.Args[1], func(path string, d fs.DirEntry, err error) error {
		return findTxt(os.Args[1], path, d, err)
	})
	if err != nil {
		log.Fatalf("could not open dir %q: %v", os.Args[1], err)
	}
}
