# Golang Filtering Subtests Example

Golang's `-run` option for the `test` subcommand also allows filtering subtests.
Here's a quick example:

```bash
# Run all tests and subtests
$ go test -v
=== RUN   TestFib
=== RUN   TestFib/Fibonacci_valid_input
=== RUN   TestFib/Fibonacci_invalid_input
--- PASS: TestFib (0.00s)
    --- PASS: TestFib/Fibonacci_valid_input (0.00s)
    --- PASS: TestFib/Fibonacci_invalid_input (0.00s)
=== RUN   TestGreet
=== RUN   TestGreet/Hello_World
    main_test.go:28: wrong output `Hello, World!`
--- FAIL: TestGreet (0.00s)
    --- FAIL: TestGreet/Hello_World (0.00s)
FAIL
exit status 1
FAIL    github.com/wperron/go-subtests-example  0.001s

# Filter top-level test with the pattern "Fib"
$ go test -v -run Fib
=== RUN   TestFib
=== RUN   TestFib/Fibonacci_valid_input
=== RUN   TestFib/Fibonacci_invalid_input
--- PASS: TestFib (0.00s)
    --- PASS: TestFib/Fibonacci_valid_input (0.00s)
    --- PASS: TestFib/Fibonacci_invalid_input (0.00s)
PASS
ok      github.com/wperron/go-subtests-example  0.001s

# Filter subtests with pattern "_valid" in top-level test with pattern "Fib"
$ go test -v -run Fib/_valid
=== RUN   TestFib
=== RUN   TestFib/Fibonacci_valid_input
--- PASS: TestFib (0.00s)
    --- PASS: TestFib/Fibonacci_valid_input (0.00s)
PASS
ok      github.com/wperron/go-subtests-example  0.001s
```