package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.String()
	d, err := os.Open(s[1:])
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		//		log.Fatal(err)
	}
	defer d.Close()
	entries, err := d.Readdir(-1)
	if err != nil {
		fmt.Fprintln(w, "Error: ", err)
		//		log.Fatal(err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(e.Name(), ".txt") {
			path := filepath.Join(s[1:], e.Name())
			res := Convert(CountFile(path))
			fmt.Fprintln(w, res)
		}

	}
}

func CountFile(path string) (string, int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines, err := CountLines(f)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(path, "\\")
	return str[len(str)-1], lines
}
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}
func Convert(str string, lines int) string {
	d := struct {
		Name  string
		Lines int
	}{Name: str, Lines: lines}
	json, _ := json.Marshal(d)
	return string(json)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
