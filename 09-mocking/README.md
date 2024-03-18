# Mocking

## But isn't mocking evil?

You may have heard mocking is evil. Just like anything in software development it can be used for evil, just like [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself).

People normally get in to a bad state when they don't _listen to their tests_ and are _not respecting the refactoring stage_.

If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should _listen_ to that bad feeling and think about your code. Usually it is a sign of

- The thing you are testing is having to do too many things (because it has too many dependencies to mock)
  - Break the module apart so it does less
- Its dependencies are too fine-grained
  - Think about how you can consolidate some of these dependencies into one meaningful module
- Your test is too concerned with implementation details
  - Favour testing expected behaviour rather than the implementation

Normally a lot of mocking points to _bad abstraction_ in your code.

**What people see here is a weakness in TDD but it is actually a strength**, more often than not poor test code is a result of bad design or put more nicely, well-designed code is easy to test.

### But mocks and tests are still making my life hard!

Ever run into this situation?

- You want to do some refactoring
- To do this you end up changing lots of tests
- You question TDD and make a post on Medium titled "Mocking considered harmful"

This is usually a sign of you testing too much _implementation detail_. Try to make it so your tests are testing _useful behaviour_ unless the implementation is really important to how the system runs.

It is sometimes hard to know _what level_ to test exactly but here are some thought processes and rules I try to follow:

- **The definition of refactoring is that the code changes but the behaviour stays the same**. If you have decided to do some refactoring in theory you should be able to make the commit without any test changes. So when writing a test ask yourself
  - Am I testing the behaviour I want, or the implementation details?
  - If I were to refactor this code, would I have to make lots of changes to the tests?
- Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour. Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.
- I feel like if a test is working with **more than 3 mocks then it is a red flag** - time for a rethink on the design
- Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation. **Be sure you actually care about these details if you're going to spy on them**

#### Can't I just use a mocking framework?

Mocking requires no magic and is relatively simple; using a framework can make mocking seem more complicated than it is. We don't use automocking in this chapter so that we get:

- a better understanding of how to mock
- practice implementing interfaces

In collaborative projects there is value in auto-generating mocks. In a team, a mock generation tool codifies consistency around the test doubles. This will avoid inconsistently written test doubles which can translate to inconsistently written tests.

You should only use a mock generator that generates test doubles against an interface. Any tool that overly dictates how tests are written, or that use lots of 'magic', can get in the sea.

## Wrapping up

### More on TDD approach

- When faced with less trivial examples, break the problem down into "thin vertical slices". Try to get to a point where you have _working software backed by tests_ as soon as you can, to avoid getting in rabbit holes and taking a "big bang" approach.
- Once you have some working software it should be easier to _iterate with small steps_ until you arrive at the software you need.

> "When to use iterative development? You should use iterative development only on projects that you want to succeed."

Martin Fowler.

### Mocking

- **Without mocking important areas of your code will be untested**. In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that _can_ fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.
- Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in **slow feedback loops**.
- By having to spin up a database or a webservice to test something you're likely to have **fragile tests** due to the unreliability of such services.

Once a developer learns about mocking it becomes very easy to over-test every single facet of a system in terms of the _way it works_ rather than _what it does_. Always be mindful about **the value of your tests** and what impact they would have in future refactoring.

In this post about mocking we have only covered **Spies**, which are a kind of mock. Mocks are a type of "test double."

> [Test Double is a generic term for any case where you replace a production object for testing purposes.](https://martinfowler.com/bliki/TestDouble.html)

Under test doubles, there are various types like stubs, spies and indeed mocks! Check out [Martin Fowler's post](https://martinfowler.com/bliki/TestDouble.html) for more detail.
