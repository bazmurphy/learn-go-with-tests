package blogposts_test

import (
	"blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

// Write the test first

// We should keep scope as small and useful as possible.
// If we prove that we can read all the files in a directory, that will be a good start.
// This will give us confidence in the software we're writing.
// We can check that the count of []Post returned is the same as the number of files in our fake file system.

// func TestNewBlogPosts(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte("hi")},
// 		"hello-world2.md": {Data: []byte("hola")},
// 	}

// 	posts := blogposts.NewPostsFromFS(fs)

// 	if len(posts) != len(fs) {
// 		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 	}
// }

// Notice that the package of our test is blogposts_test.
// Remember, when TDD is practiced well we take a consumer-driven approach: we don't want to test internal details because consumers don't care about them. By appending _test to our intended package name, we only access exported members from our package - just like a real user of our package.

// We've imported testing/fstest which gives us access to the fstest.MapFS type.
// Our fake file system will pass fstest.MapFS to our package.
// "A MapFS is a simple in-memory file system for use in tests, represented as a map from path names (arguments to Open) to information about the files or directories they represent."

// This feels simpler than maintaining a folder of test files, and it will execute quicker.

// Finally, we codified the usage of our API from a consumer's point of view, then checked if it creates the correct number of posts.

// ----------

// Let's change our API (via our tests first, naturally) so that it can return an error.

// func TestNewBlogPosts(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte("hi")},
// 		"hello-world2.md": {Data: []byte("hola")},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(posts) != len(fs) {
// 		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 	}
// }

// ----------

// We'll start with the first line in the proposed blog post schema, the title field.
// We need to change the contents of the test files so they match what was specified,
// and then we can make an assertion that it is parsed correctly.

// func TestNewBlogPosts(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte("Title: Post 1")},
// 		"hello-world2.md": {Data: []byte("Title: Post 2")},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(posts) != len(fs) {
// 		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 	}

// 	got := posts[0]
// 	want := blogposts.Post{Title: "Post 1"}

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %+v, want %+v", got, want)
// 	}
// }

// ----------

// We should take care of our tests too.
// We're going to be making assertions on Posts a lot, so we should write some code to help with that.

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// func TestNewBlogPosts(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte("Title: Post 1")},
// 		"hello-world2.md": {Data: []byte("Title: Post 2")},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(posts) != len(fs) {
// 		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 	}

// 	got := posts[0]
// 	want := blogposts.Post{Title: "Post 1"}

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %+v, want %+v", got, want)
// 	}

// 	assertPost(t, posts[0], blogposts.Post{
// 		Title:       "Post 1",
// 		Description: "Description 1",
// 	})
// }

// func TestNewBlogPosts(t *testing.T) {
// 	const (
// 		firstBody = `Title: Post 1
// Description: Description 1`
// 		secondBody = `Title: Post 2
// Description: Description 2`
// 	)

// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte(firstBody)},
// 		"hello-world2.md": {Data: []byte(secondBody)},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assertPost(t, posts[0], blogposts.Post{
// 		Title:       "Post 1",
// 		Description: "Description 1",
// 	})
// }

// For brevity, I will not go through the TDD steps, but here's the test with tags added.

// func TestNewBlogPosts(t *testing.T) {
// 	const (
// 		firstBody = `Title: Post 1
// Description: Description 1
// Tags: tdd, go`
// 		secondBody = `Title: Post 2
// Description: Description 2
// Tags: rust, borrow-checker`
// 	)

// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte(firstBody)},
// 		"hello-world2.md": {Data: []byte(secondBody)},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assertPost(t, posts[0], blogposts.Post{
// 		Title:       "Post 1",
// 		Description: "Description 1",
// 		Tags:        []string{"tdd", "go"},
// 	})
// }

// Change the test data to have the separator, and a body with a few newlines to check we grab all the content.

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}
