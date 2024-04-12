package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

func parseFile(fs fs.FS, filename string) (Post, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer f.Close()
	return newPost(f)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
	}, nil
}
