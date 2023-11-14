# Indie

Opinionated Go boilerplate for the indie hacker or early stage projects.

![indie-logo.png](indie-logo.png)

Why does JS, Rails, Django, and Phoenix have all the fun?

Let's use a Go stack for quickly building ideas.

**Important**: This is not a framework. It's boilerplate. A template.

After you create your project, use or discard what you wish.

## The Stack

- Go (duh)
- [Cobra](https://github.com/spf13/cobra) for cli
- [Echo](https://echo.labstack.com) for web server and router
- [HTMX](https://htmx.org) for dynamic web pages
- [Templ](https://github.com/a-h/templ) for HTML templates
- [Ent](https://entgo.io) for database/ORM
- [Air](https://github.com/cosmtrek/air) for live reload
- [Testify](https://github.com/stretchr/testify) for test matchers

Fat free! No npm, npx, yarn, pnpm, webpack, and whatever else the Front End World conjures up.

## Use as Project Template

Install the [experimental gonew](https://go.dev/blog/gonew) command.

```sh
go install golang.org/x/tools/cmd/gonew@latest
```

Then in a fresh directory:

```sh
gonew github.com/DavidNix/indie github.com/<YOUR_USER>/<YOUR_PROJECT_NAME>
```

## Local development

All funneled through `make`.

To see what you can do:

```sh
make
```

Then (assumes you have homebrew installed):

```sh
make setup
```

Generate code:
```sh
make gen
```

Run the server:

```sh
make run
```

Live reload:

```sh
make watch
```
Caveat: Any ent (data model) changes will require a manual restart.

# Features

## Development Speed

Using ent allows automatic migrations. At scale, this is bad. But for iterating quickly, it's great.

Ent lets us use an in-memory sqlite database for unit tests. This is a huge win for speed.

## Reasonable Security

The license still stands that this software is provided **as-is with no warranty**.

But I've tried to make reasonable security decisions such as server timeouts, CSRF protection, and secure headers.

# Design Decisions

## Why Echo?

I first tried [Fiber](https://github.com/gofiber/fiber) which uses [fasthttp](https://github.com/fasthttp/router) as the
router. Unfortunately, fasthttp has a [nasty race condition](https://twitter.com/davidnix_/status/1720454052973044188)
when using database/sql. Also, Fiber makes you choose between `c.Context()` and `c.UserContext()` which is confusing.

Also, Echo is one of the older Go http frameworks, so hopefully has the Lindy Effect.

## Wait, an ORM?!

Those who know me will be shocked I'm using an ORM. (I typically despise them.)

But hear me out. In this context (getting a project off the ground at light speed), it's a good fit:

- Validation out of the box.
- Unit tests with in-memory sqlite.
- Automatic migrations.

If your project grows and becomes more complex, you should move off Ent (the ORM).