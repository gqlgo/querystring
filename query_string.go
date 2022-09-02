package main

import (
	"bytes"
	"fmt"
	"github.com/mattn/go-zglob"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"github.com/vektah/gqlparser/v2/parser"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := queryStringMain(); err != nil {
		log.Fatalf("error: %+v", err)
	}
}

func queryStringMain() error {
	paths := os.Args[1:]
	if len(paths) == 0 {
		if err := processFile("<standard input>", os.Stdin, os.Stdout); err != nil {
			return fmt.Errorf("failed to process stdin: %w", err)
		}
		return nil
	}

	extractGlobPathMap := make(map[string]struct{}, 0)
	for _, path := range paths {
		extracts, err := zglob.Glob(path)
		if err != nil {
			return fmt.Errorf("failed to glob: %w", err)
		}
		for _, extract := range extracts {
			extractGlobPathMap[extract] = struct{}{}
		}
	}

	var extractGlobPaths []string
	for extractGlobPath := range extractGlobPathMap {
		extractGlobPaths = append(extractGlobPaths, extractGlobPath)
	}
	for _, path := range extractGlobPaths {
		if err := processFile(path, nil, os.Stdout); err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
	}
	return nil
}

func processFile(filename string, in io.Reader, out io.Writer) error {
	if in == nil {
		f, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf(": %w", err)
		}
		defer f.Close()
		in = f
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	queryString, err := queryString(src)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	if len(queryString) > 0 {
		fmt.Fprintf(out, string(queryString))
	}

	return nil
}

func queryString(b []byte) ([]byte, error) {
	stringLiteral, err := stringLiteral(b)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	formatQuery, err := formatQuery(stringLiteral)
	if err != nil {
		return []byte{}, nil

	}

	return formatQuery, nil
}

func stringLiteral(b []byte) ([]byte, error) {
	var queryString []byte
	var start int

	for {
		q, end := singleQueryString(b, start)
		queryString = append(queryString, q...)
		start = start + end + 1

		if start >= len(b) {
			break
		}
	}

	return queryString, nil
}

func singleQueryString(b []byte, start int) ([]byte, int) {
	begin, end := -1, -1
	for i, c := range b[start:] {
		if c == '`' {
			if begin < 0 {
				begin = i
			} else {
				end = i
				break
			}
		}
	}
	if end < 0 {
		return []byte{}, len(b)
	}

	return b[start+begin+1 : start+end], start + end
}

func formatQuery(src []byte) ([]byte, error) {
	source := &ast.Source{Name: "", Input: string(src)}

	query, err := parser.ParseQuery(source)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	var buf bytes.Buffer
	astFormatter := formatter.NewFormatter(&buf)
	astFormatter.FormatQueryDocument(query)
	return buf.Bytes(), nil
}
