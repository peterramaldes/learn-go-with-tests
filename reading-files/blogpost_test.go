package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/peterramaldes/learn-go-with-tests/reading-files"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello-world.md": {Data: []byte("Title: Post 1")},
			"hello-foo.md":   {Data: []byte("Title: Post 2")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("expected %d posts, got %d posts", len(fs), len(posts))
		}

	})

	t.Run("err test", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Error("expect an error")
		}
	})

	t.Run("should set the title", func(t *testing.T) {
		fs := fstest.MapFS{
			"post1.md": {Data: []byte("Title: Post 1")},
			"post2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, got, want)
	})

	t.Run("parsing the body", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1`
			secondBody = `Title: Post 2
Description: Description 2`
		)

		fs := fstest.MapFS{
			"post1.md": {Data: []byte(firstBody)},
			"post2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewPostsFromFS(fs)

		got := posts[0]
		want := blogposts.Post{Title: "Post 1", Description: "Description 1"}
		assertPost(t, got, want)

	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
