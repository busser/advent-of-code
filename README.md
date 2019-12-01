# Solutions to the Advent of Code

## Testing solutions

```bash
# Run all tests.
go test ./...

# Test a single day.
go test ./2018/day01

# See test details.
go test ./2018/day01 -v

# Run a specific subtest.
go test ./2018/day01 -run=TestDay/PartOne
go test ./2018/day01 -run=TestDay/PartTwo

# Run a speed benchmark.
go test ./2018/day01 -bench=.

# Run a memory benchmark.
go test ./2018/day01 -bench=. -benchmem
```

## Generating code

Much of the code written is the same for each day: functions `PartOne` and
`PartTwo` solve the two parts of the day's puzzle, and both are tested and
benchmarked. I also often have some functions that are shared by `PartOne` and
`PartTwo`, like a `readInput` function that provides the puzzle input in a
structured form. Beyond the code, every day's solution has its own directory,
and the puzzle input is stored inside `testdata/input.txt`.

All of these files can be generated to some extent. Parts of the `PartOne`,
`PartTwo`, and `readInput` functions can be pre-written and the puzzle's input
can be downloaded.

In order to generate the boilerplate code and download the puzzle's input, use
the following command:

```bash
# This will create files for December 1st 2019.
go run gen/main.go --day 1 --year 2019 --cookie <your-cookie>
```

You should replace `<your-cookie>` with the actual value of your Advent of Code
session cookie. If you inspect the cookies created by the website, you should
find one called `session` if you are signed in. Use its value in the command.

For more information on what can be done regarding code generation, run:

```bash
go run gen/main.go --help
```
