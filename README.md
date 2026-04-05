# Timer in Golang

## Goal

This project is a command-line timer written in Golang.  
The program starts a countdown from a number of seconds given as an argument, then displays the remaining time every second until it reaches zero.

## What I used

- `fmt` for displaying the timer
- `os` for reading command-line arguments and exiting on error
- `strconv` for converting the argument to an integer
- `time` for handling the countdown
- loops and integers to make the timer work

## You need Golang

If you do not have Golang installed, run:

```bash
sudo apt update
sudo apt install golang-go
```

## How to use it

Run the program with one argument:

```bash
go run . <seconds>
```

- `<seconds>`: number of seconds for the timer, must be a strictly positive integer

## Example

```bash
go run . 5
```

You should see something like:

```text
Timer: 5s
Timer: 4s
Timer: 3s
Timer: 2s
Timer: 1s
Timer: 0s
```

## Error handling

If you provide an invalid argument, the program prints an error message and exits.

Examples of invalid input:

```bash
go run .
go run . abc
go run . 0
go run . -3
```

## Notes

- The timer decreases by one second on each loop.
- The program expects exactly one argument.
- The value must be a positive non-zero integer.

---

Enjoy using the timer!
