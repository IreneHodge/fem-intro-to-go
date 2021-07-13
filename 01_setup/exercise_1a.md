# Exercise 1A: Find Stuff

## Goals

- Familiarize yourself with the golang.org website and resources

## Setup

- Visit `golang.org`

## Directions

Answer the following questions

1. Read about `for loops` in the _Effective Go_ document

- What kind of loop doesn’t exist in Go?

There is no do-while loop in Go. There are three forms, only one of which has semicolons.

// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

2. Read about the `fmt` _package_

- What does `fmt.Println()` return?
  integer of bytes written and any write error encountered

Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.

Package fmt implements formatted Input/Output with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

3. Find a _blog post_ about the recent release of Go 1.13

- What are some of the new features?
  Checks on database
  cool stuff around erro
