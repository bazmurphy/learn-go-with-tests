# Concurrency

## Wrapping up

This exercise has been a little lighter on the TDD than usual. In a way we've
been taking part in one long refactoring of the `CheckWebsites` function; the
inputs and outputs never changed, it just got faster. But the tests we had in
place, as well as the benchmark we wrote, allowed us to refactor `CheckWebsites`
in a way that maintained confidence that the software was still working, while
demonstrating that it had actually become faster.

In making it faster we learned about

- _goroutines_, the basic unit of concurrency in Go, which let us manage more
  than one website check request.
- _anonymous functions_, which we used to start each of the concurrent processes
  that check websites.
- _channels_, to help organize and control the communication between the
  different processes, allowing us to avoid a _race condition_ bug.
- _the race detector_ which helped us debug problems with concurrent code

### Make it fast

One formulation of an agile way of building software, often misattributed to Kent
Beck, is:

> [Make it work, make it right, make it fast][wrf]

Where 'work' is making the tests pass, 'right' is refactoring the code, and
'fast' is optimizing the code to make it, for example, run quickly. We can only
'make it fast' once we've made it work and made it right. We were lucky that the
code we were given was already demonstrated to be working, and didn't need to be
refactored. We should never try to 'make it fast' before the other two steps
have been performed because

> [Premature optimization is the root of all evil][popt]
> -- Donald Knuth

[DI]: dependency-injection.md
[wrf]: http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast
[godoc_race_detector]: https://blog.golang.org/race-detector
[popt]: http://wiki.c2.com/?PrematureOptimization
