# Go-Projects

A collection of small projects built while learning Go from scratch — no prior tutorials, learning by building and debugging as I go (coming from a Python background).

## Purpose

This repo is my Go learning log. Each folder is a self-contained project that pushed me to learn a new piece of the language — control flow, structs, error handling, and eventually concurrency and interfaces as I get there.

## Projects

### 1. Banking System CLI Simulator (atm.go)
A command-line banking simulator supporting balance checks, deposits, withdrawals, and (soon) transfers between accounts.

**Concepts covered:**
- Variables and package-level state (`var` vs `:=`)
- Control flow: `for`, `switch`, labeled `break`/`continue`
- User input handling (`fmt.Scan`, `fmt.Scanln`) and basic input validation
- Error checking with Go's `error` return pattern

**Run it:**
```bash
go run Go-projects/atm.go
```

---

*More projects added as I go — each one below will follow the same format: what it does, what I learned building it, and how to run it.*

## Why Go

Coming from Python, picking Go up by building small, complete projects rather than following tutorials — this repo tracks that progress and doubles as a reference for concepts I've had to work through myself.