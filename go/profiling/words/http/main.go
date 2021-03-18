package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/trace"
	"strings"
	"unicode"

	"net/http"
	_ "net/http/pprof"
)

type bytereader struct {
	buf [1]byte
	r   io.Reader
}

func (b *bytereader) next() (rune, error) {
	_, err := b.r.Read(b.buf[:])
	return rune(b.buf[0]), err
}

func countWords(f *os.File, w io.Writer) {
	br := bytereader{
		r: bufio.NewReader(f),
	}

	words := 0
	inword := false
	for {
		r, err := br.next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "Could not read file %q: %v\n", f.Name(), err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}

	fmt.Fprintf(w, "%q: %d words\n", f.Name(), words)
	return
}

func findTxt(w io.Writer, root, path string, d fs.DirEntry, err error) error {
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
		fmt.Fprintf(w, "could not open file %q: %v\n", path, err)
		return err
	}

	countWords(f, w)

	return nil
}

func startTrace(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Fprintln(w, "Failed to create trace file")
		return
	}

	err = trace.Start(f)
	if err != nil {
		fmt.Fprintln(w, "Failed to start trace:", err)
		return
	}

	fmt.Fprintln(w, "Trace started")
}

func stopTrace(w http.ResponseWriter, r *http.Request) {
	trace.Stop()
	fmt.Fprintln(w, "Trace stopped")
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	fmt.Println("Serving at:", addr)
	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()

	// useless 'words' service
	mux.HandleFunc("/words", func(resp http.ResponseWriter, req *http.Request) {
		root := "../books"
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			return findTxt(resp, root, path, d, err)
		})
		if err != nil {
			fmt.Fprintf(resp, "could not open path %q: %v\n", root, err)
			return
		}
	})

	return serve("127.0.0.1:8082", mux, stop)
}

func serveDebug(stop <-chan struct{}) error {
	mux := http.DefaultServeMux

	// 方法一：
	// http://127.0.0.1:8081/start
	// http://127.0.0.1:8081/stop
	mux.HandleFunc("/start", startTrace)
	mux.HandleFunc("/stop", stopTrace)
	// 方法二：
	// curl -o trace.out http://127.0.0.1:8081/debug/pprof/trace?seconds=5

	return serve("127.0.0.1:8081", mux, stop)
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})

	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error:", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
