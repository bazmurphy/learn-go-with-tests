package main

// [01]
// a function, CheckWebsites, that checks the status of a list of URLs.
// It returns a map of each URL checked to a boolean value: true for a good response; false for a bad response.
// You also have to pass in a WebsiteChecker which takes a single URL and returns a boolean.
// This is used by the function to check all the websites.

// type WebsiteChecker func(string) bool

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		results[url] = wc(url)
// 	}

// 	return results
// }

// [03]
// The function is in production and being used to check hundreds of websites.
// But has started to get complaints that it's slow, so they've asked you to help speed it up

// [06]
// We can finally talk about concurrency which, for the purposes of the following, means "having more than one thing in progress."
// Instead of waiting for a website to respond before sending a request to the next website,
// we will tell our computer to make the next request while it is waiting.

// Normally in Go when we call a function doSomething() we wait for it to return (even if it has no value to return, we still wait for it to finish).
// We say that this operation is blocking - it makes us wait for it to finish.
// An operation that does not block in Go will run in a separate process called a goroutine.

// Think of a process as reading down the page of Go code from top to bottom,
// going 'inside' each function when it gets called to read what it does.

// When a separate process starts, it's like another reader begins reading inside the function,
// leaving the original reader to carry on going down the page.

// To tell Go to start a new goroutine we turn a function call into a go statement
// by putting the keyword go in front of it: go doSomething().

// Because the only way to start a goroutine is to put go in front of a function call,
// we often use anonymous functions when we want to start a goroutine.

// An anonymous function literal looks just the same as a normal function declaration, but without a name (unsurprisingly).
// You can see one in the body of the for loop.

// Anonymous functions have a number of features which make them useful, two of which we're using above.
// Firstly, they can be executed at the same time that they're declared - this is what the () at the end of the anonymous function is doing.
// (!) Secondly they maintain access to the lexical scope in which they are defined - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

// [07]
// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		// here we use the go keyword
// 		go func() {
// 			results[url] = wc(url)
// 		}()
// 	}

// 	return results
// }

// The body of the anonymous function above is just the same as the loop body was before.
// The only difference is that each iteration of the loop will start a new goroutine, concurrent with the current process (the WebsiteChecker function).
// Each goroutine will add its result to the results map.

// [08]
// We might try and fix this by increasing the time we wait - try it if you like. It won't work.

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		// here we use the go keyword
// 		go func() {
// 			results[url] = wc(url)
// 		}()
// 	}

// 	// we get an empty map (sometimes)
// 	// because none of the go routines have time to return before the results return
// 	time.Sleep(2 * time.Second)

// 	return results
// }

// [09]
// The problem here is that the variable url is reused for each iteration of the for loop - it takes a new value from urls each time.
// But each of our goroutines have a reference to the url variable - they don't have their own independent copy.
// So they're all writing the value that url has at the end of the iteration - the last url.
// Which is why the one result we have is the last url.

// By giving each anonymous function a parameter for the url - u - and then calling the anonymous function with the url as the argument, we make sure that the value of u is fixed as the value of url for the iteration of the loop that we're launching the goroutine in. u is a copy of the value of url, and so can't be changed.

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		go func(u string) {
// 			results[u] = wc(u)
// 		}(url)
// 	}

// 	time.Sleep(2 * time.Second)

// 	return results
// }

// [09]
// when we run the tests now, we sometimes get "fatal error: concurrent map writes"
// Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time.
// Maps in Go don't like it when more than one thing tries to write to them at once, and so fatal error.

// This is a race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over.
// Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the same time.

// Go can help us to spot race conditions with its built in race detector.
// To enable this feature, run the tests with the race flag: go test -race.

// [10]
// We can solve this data race by coordinating our goroutines using channels.
// Channels are a Go data structure that can both receive and send values.
// These operations, along with their details, allow communication between different processes.

// In this case we want to think about the communication between the parent process
// and each of the goroutines that it makes to do the work of running the WebsiteChecker function with the url.

// Alongside the results map we now have a resultChannel, which we make in the same way.
// chan result is the type of the channel - a channel of result.
// The new type, result has been made to associate the return value of the WebsiteChecker with the url being checked - it's a struct of string and bool.
// As we don't need either value to be named, each of them is anonymous within the struct; this can be useful in when it's hard to know what to name a value.

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// Now when we iterate over the urls, instead of writing to the map directly we're sending a result struct for each call to wc to the resultChannel with a send statement.
// This uses the <- operator, taking a channel on the left and a value on the right:

// Send statement
// resultChannel <- result{u, wc(u)}

// The next for loop iterates once for each of the urls.
// Inside we're using a receive expression, which assigns a value received from a channel to a variable.
// This also uses the <- operator, but with the two operands now reversed: the channel is now on the right and the variable that we're assigning to is on the left

// Receive expression
// r := <-resultChannel

// We then use the result received to update the map.

// By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time.
// Although each of the calls of wc, and each send to the result channel, is happening concurrently inside its own process, each of the results is being dealt with one at a time as we take values out of the result channel with the receive expression.

// We have used concurrency for the part of the code that we wanted to make faster, while making sure that the part that cannot happen simultaneously still happens linearly.
// And we have communicated across the multiple processes involved by using channels.
