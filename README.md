# Indie

Opinionated Go boilerplate for the indie hacker or early stage projects.

![indie-logo.png](indie-logo.png)

Why does JS, Rails, Django, and Phoenix have all the fun?

Let's use a Go stack for quickly building ideas.

After you create your project, use or discard what you wish.

This is used as a boilerplate and library.

## Boilerplate

Install the [experimental gonew](https://go.dev/blog/gonew) command.

```sh
go install golang.org/x/tools/cmd/gonew@latest
```

Then in a fresh directory:

```sh
gonew github.com/DavidNix/indie github.com/<YOUR_USER>/<YOUR_PROJECT_NAME>
```

## Library

TODO

## The Stack

- Go (of course)
- [Cobra](https://github.com/spf13/cobra) for cli
- [Echo](https://echo.labstack.com) for web server and router
- [HTMX](https://htmx.org) for dynamic web pages (although I'm considering moving to datastar)
- [Templ](https://github.com/a-h/templ) for HTML templates
- [Overmind](https://github.com/DarthSim/overmind) for local development (a lightweight docker compose that doesn't need docker)
- [Tailwind CSS](https://tailwindcss.com) for the utility-first CSS framework
- [DaisyUI 5](https://daisyui.com) for tailwind components because Daisy doesn't need JS

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

Live reload:

```sh
make watch
```

Caveat: Any ent (data model) changes will require a manual restart.

## Features

### Reasonable Security

The license still stands that this software is provided **as-is with no warranty**.

But I've tried to make reasonable security decisions such as server timeouts, CSRF protection, and secure headers.

## Design Decisions

### Why Echo?

I like Echo's handler signature where you return an error. They also have nice middleware.

I first tried [Fiber](https://github.com/gofiber/fiber) which uses [fasthttp](https://github.com/fasthttp/router) as the
router. Unfortunately, fasthttp has a [nasty race condition](https://twitter.com/davidnix_/status/1720454052973044188)
when using database/sql. Also, Fiber makes you choose between `c.Context()` and `c.UserContext()` which is confusing.

Also, Echo is one of the older Go http frameworks, so has the Lindy Effect.

### No ORM or database library?

BYO database implementation. There's an example of how I like to do migrations with sqlite.

I typically don't like ORMs, but sometimes I use [ent](https://entgo.io) if I need speed of development.
Also, ORMs let you compose dynamic queries a bit easier than compiled solutions like sqlc.

Many prefer [sqlc](https://sqlc.dev). Or, if simple enough, just use `database/sql`.

### No AplineJS?

I found writing vanilla JS suited my needs just fine.
