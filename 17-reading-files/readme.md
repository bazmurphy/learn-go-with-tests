# Reading Files

## Wrapping up

`fs.FS`, and the other changes in Go 1.16 give us some elegant ways of reading data from file systems and testing them simply.

If you wish to try out the code "for real":

- Create a `cmd` folder within the project, add a `main.go` file
- Add the following code

```go
import (
	blogposts "github.com/quii/fstest-spike"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
```

- Add some markdown files into a `posts` folder and run the program!

Notice the symmetry between the production code

```go
posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
```

And the tests

```go
posts, err := blogposts.NewPostsFromFS(fs)
```

This is when consumer-driven, top-down TDD _feels correct_.

A user of our package can look at our tests and quickly get up to speed with what it's supposed to do and how to use it. As maintainers, we can be _confident our tests are useful because they're from a consumer's point of view_. We're not testing implementation details or other incidental details, so we can be reasonably confident that our tests will help us, rather than hinder us when refactoring.

By relying on good software engineering practices like [**dependency injection**](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection) our code is simple to test and re-use.

When you're creating packages, even if they're only internal to your project, prefer a top-down consumer driven approach. This will stop you over-imagining designs and making abstractions you may not even need and will help ensure the tests you write are useful.

The iterative approach kept every step small, and the continuous feedback helped us uncover unclear requirements possibly sooner than with other, more ad-hoc approaches.

### Writing?

It's important to note that these new features only have operations for _reading_ files. If your work needs to do writing, you'll need to look elsewhere. Remember to keep thinking about what the standard library offers currently, if you're writing data you should probably look into leveraging existing interfaces such as `io.Writer` to keep your code loosely-coupled and re-usable.

### Further reading

- This was a light intro to `io/fs`. [Ben Congdon has done an excellent write-up](https://benjamincongdon.me/blog/2021/01/21/A-Tour-of-Go-116s-iofs-package/) which was a lot of help for writing this chapter.
- [Discussion on the file system interfaces](https://github.com/golang/go/issues/41190)
