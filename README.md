# Indie
An opinionated Go stack for the indie hacker or early stage projects.

Why does JS have all the fun?

Let's use a Go stack for quickly building ideas.

Why not Rails, Django, or Phoenix?

Sure, if you enjoy the hell that is dynamically typed languages.

## The Stack
- Go (duh)
- [Viper](https://github.com/spf13/viper) and [Cobra](https://github.com/spf13/cobra) for configuration and CLI
- [Fiber](https://gofiber.io) for web server
- [HTMX](https://htmx.org) for dynamic web pages
- [Ent](https://entgo.io) for ORM

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
```
make setup
```

## OMG an ORM?!?!

Those who know me will be shocked I'm using an ORM. (I typically despise them.)

But hear me out. In this context (getting a project off the ground at light speed), it's a good fit:
- Validation out of the box.
- Unit tests with in-memory sqlite.
- Automatic migrations.

If your project grows and becomes more complex, you should move off Ent (the ORM).