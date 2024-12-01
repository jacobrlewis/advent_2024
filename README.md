# Advent of Code 2024

[Problems here](https://adventofcode.com/2024)

My solutions to AoC 2024.

Just for fun!

Solutions can be tested by downloading your own input and example files and saving into each day's directory. The creators ask that inputs are not uploaded publicly. Save your problem input as `input.txt`, and an example file as `example.txt`.

# Running the code

This project is set up to receive command line flags for each problem.
* `-d` day
* `-p` part (1 or 2)
* `--example` to use the `exmaple.txt` instead of `input.txt`

Examples:

```sh
# day 1 part 1, example file
go run main.go -d=1 -p=1 --example

# day 3, part 2, real input file
go run main.go -d=3 -p=2
```