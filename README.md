# Learn Go with Tests

My stab at doing the exercises at https://quii.gitbook.io/learn-go-with-tests, along witht my notes.

I am using VS Code, and installed the optional Delve debugger and golangci-lint per the recommendations.

I'm trying to use my own weird version of SemVer here: major revisions correspond to new sections below. Minor revisions indicate a new section was added, and patch versions are for fixes.

## Foundations 
1. [Hello World](hello_world/README.md)
2. [Integers](integers/README.md)
3. [Iteration](iteration/README.md)
4. [Arrays and Slices](arrays-and-slices/README.md)
5. [Structs, Methods, and Interaces](structs-methods-interfaces/README.md)
6. [Pointers and Errors](pointers-and-errors/README.md)
7. [Maps](maps/README.md)
8. [Dependency Injection](dependency-injection/README.md)
9. [Mocking](mocking/README.md)
10. [Concurrency](concurrency/README.md)
11. [Select](select/README.md)
12. [Reflection](reflection/README.md)
13. [Sync](sync/README.md)
14. [Context](context/README.md)
15. [Intro to property based tests](/property-based-tests/README.md)

## Build an application

The assignment:

You have been asked to create a web server where users can track how many games players have won.

- `GET /players/{name}` should return a number indicating the total number of wins
- `POST /players/{name}` should record a win for that name, incrementing for every subsequent POST

We will follow the TDD approach, getting working software as quickly as we can and then making small iterative improvements until we have the solution. By taking this approach we
- Keep the problem space small at any given time
- Don't go down rabbit holes
- If we ever get stuck/lost, doing a revert wouldn't lose loads of work.

### Red, green, refactor
Throughout this book, we have emphasised the TDD process of write a test & watch it fail (red), write the minimal amount of code to make it work (green) and then refactor.
This discipline of writing the minimal amount of code is important in terms of the safety TDD gives you. You should be striving to get out of "red" as soon as you can.

Kent Beck describes it as:
> Make the test work quickly, committing whatever sins necessary in process.

You can commit these sins because you will refactor afterwards backed by the safety of the tests.

### What if you don't do this?
The more changes you make while in red, the more likely you are to add more problems, not covered by tests.
The idea is to be iteratively writing useful code with small steps, driven by tests so that you don't fall into a rabbit hole for hours.

### Chicken and egg

How can we incrementally build this? We can't GET a player without having stored something and it seems hard to know if POST has worked without the GET endpoint already existing.
This is where _mocking_ shines.

GET will need a PlayerStore thing to get scores for a player. This should be an interface so when we test we can create a simple stub to test our code without needing to have implemented any actual storage code.

For POST we can spy on its calls to PlayerStore to make sure it stores players correctly. Our implementation of saving won't be coupled to retrieval.

For having some working software quickly we can make a very simple in-memory implementation and then later we can create an implementation backed by whatever storage mechanism we prefer.

```note
Hey - this is a nifty mix of what might be called "classic" vs "London school" TDD. 
```

1. [Build an HTTP Server](./web-app/http-server.md)