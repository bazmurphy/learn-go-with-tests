# Context

## Wrapping up

### What we've covered

- How to test a HTTP handler that has had the request cancelled by the client.
- How to use context to manage cancellation.
- How to write a function that accepts `context` and uses it to cancel itself by using goroutines, `select` and channels.
- Follow Google's guidelines as to how to manage cancellation by propagating request scoped context through your call-stack.
- How to roll your own spy for `http.ResponseWriter` if you need it.

### What about `context.Value` ?

[Michal Štrba](https://faiface.github.io/post/context-should-go-away-go2/) and I have a similar opinion.

> If you use ctx.Value in my (non-existent) company, you’re fired

Some engineers have advocated passing values through `context` as it _feels convenient_.

Convenience is often the cause of bad code.

The problem with `context.Values` is that it's just an untyped map so you have no type-safety and you have to handle it not actually containing your value. You have to create a coupling of map keys from one module to another and if someone changes something things start breaking.

In short, **if a function needs some values, put them as typed parameters rather than trying to fetch them from** `**context.Value**`. This makes it statically checked and documented for everyone to see.

#### But...

On other hand, it can be helpful to include information that is orthogonal to a request in a context, such as a trace id. Potentially this information would not be needed by every function in your call-stack and would make your functional signatures very messy.

[Jack Lindamood says **Context.Value should inform, not control**](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)

> The content of context.Value is for maintainers not users. It should never be required input for documented or expected results.

### Additional material

- I really enjoyed reading [Context should go away for Go 2 by Michal Štrba](https://faiface.github.io/post/context-should-go-away-go2/). His argument is that having to pass `context` everywhere is a smell, that it's pointing to a deficiency in the language in respect to cancellation. He says it would better if this was somehow solved at the language level, rather than at a library level. Until that happens, you will need `context` if you want to manage long running processes.
