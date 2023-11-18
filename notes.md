# TODO

* Find type of data when parsed so we can remove a few any type.
* Rerun barbershop when source dir content is modified (see fsnotify
  pkg).
* Add error message when JSON file is invalid (empty file)

# Notes

## `fmt.Print` vs `log.Fatal`

Print + Exit is used instead of Fatal only when users
of the command might see the error. Fatal is used for
errors that should be caught during testing. See:
https://google.github.io/styleguide/go/decisions.html#dont-panic

## Initialization statement

`if` initialization statement is used to handle errors only when a
function returns an error and nothing else, or when the variables are
local to the conditional.

## Functions vs methods

> Here is a rule of thumb that may guide you in deciding to use a
> method or a function. Methods for what they do, functions for what
> they return.

See : https://dave.cheney.net/practical-go/presentations/gophercon-israel.html#_function_vs_methods

## Error message

Avoid words like error, failed, cannot, won't, etc. - it is clear to
the reader of the log message that if it occurred, something did not
happen.
