# HTTP Server

## Wrapping up

### `http.Handler`

- Implement this interface to create web servers
- Use `http.HandlerFunc` to turn ordinary functions into `http.Handler`s
- Use `httptest.NewRecorder` to pass in as a `ResponseWriter` to let you spy on the responses your handler sends
- Use `http.NewRequest` to construct the requests you expect to come in to your system

### Interfaces, Mocking and DI

- Lets you iteratively build the system up in smaller chunks
- Allows you to develop a handler that needs a storage without needing actual storage
- TDD to drive out the interfaces you need

### Commit sins, then refactor (and then commit to source control)

- You need to treat having failing compilation or failing tests as a red situation that you need to get out of as soon as you can.
- Write just the necessary code to get there. _Then_ refactor and make the code nice.
- By trying to do too many changes whilst the code isn't compiling or the tests are failing puts you at risk of compounding the problems.
- Sticking to this approach forces you to write small tests, which means small changes, which helps keep working on complex systems manageable.
