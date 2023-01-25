# Advent of Code 2022

Solution for Advent of Code 2022 written in Go

## Build the solution

    $ go build

Only tested with Go-lang version 1.18 and later, may work with earlier version.

## Build program

    $ go build

## Solve a single day

    $ ./aoc2022 DAY

## Solve multiple days sequentially

    $ ./aoc2022 DAYx DAYy DAYz...

solving all days can be done like this on *NIX platforms

    $ ./aoc2022 $(seq 1 25)

or using a pre-built image

    $ ./aoc2022 DAY

## Run tests

### All

    $ go test

### Specific day

    $ go test -run Day7

## Run benchmarks

### All

    $ go test -bench .

### Specific day

    $ go test -bench Day11
