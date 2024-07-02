# Wrapping up

We've covered a few things from the sync package

- `Mutex` allows us to add locks to our data

- `WaitGroup` is a means of waiting for goroutines to finish jobs

## When to use locks over channels and goroutines?

[We've previously covered goroutines in the first concurrency chapter](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency) which let us write safe concurrent code so why would you use locks? [The go wiki has a page dedicated to this topic; Mutex Or Channel](https://go.dev/wiki/MutexOrChannel)

"A common Go newbie mistake is to over-use channels and goroutines just because it's possible, and/or because it's fun. Don't be afraid to use a sync.Mutex if that fits your problem best. Go is pragmatic in letting you use the tools that solve your problem best and not forcing you into one style of code."

Paraphrasing:

- Use channels when passing ownership of data

- Use mutexes for managing state

## `go vet`

Remember to use go vet in your build scripts as it can alert you to some subtle bugs in your code before they hit your poor users.

## Don't use embedding because it's convenient

Think about the effect embedding has on your public API.

Do you really want to expose these methods and have people coupling their own code to them?

With respect to mutexes, this could be potentially disastrous in very unpredictable and weird ways, imagine some nefarious code unlocking a mutex when it shouldn't be; this would cause some very strange bugs that will be hard to track down.
