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

### 2. ATM Transaction Simulator
A command-line ATM simulator supporting balance checks, deposits, and withdrawals on a single account session.

**Concepts covered:**
- Struct methods with value receivers (`CheckBalance`) vs pointer receivers (`Deposit`, `Withdraw`)
- Constructor-style functions returning a pointer (`newTransaction`)
- Multiple return values (returning a status message alongside the balance)
- Variable scope and why declaring a variable inside an `if` block hides it from the rest of the function
- Organizing a multi-file package (`atm.go` + `atm_struct.go`) inside its own subfolder

**Run it:**
```bash
go run ./atm
```

*More projects added as I go — each one below will follow the same format: what it does, what I learned building it, and how to run it.*

## Why Go

Coming from Python, picking Go up by building small, complete projects rather than following tutorials — this repo tracks that progress and doubles as a reference for concepts I've had to work through myself.