/*
--- Day 11: Hex Ed ---

Crossing the bridge, you've barely reached the other side of the stream when a program comes up to you, clearly in distress. "It's my child process," she says, "he's gotten lost in an infinite grid!"

Fortunately for her, you have plenty of experience with infinite grids.

Unfortunately for you, it's a hex grid.

The hexagons ("hexes") in this grid are aligned such that adjacent hexes can be found to the north, northeast, southeast, south, southwest, and northwest:

  \ n  /
nw +--+ ne
  /    \
-+      +-
  \    /
sw +--+ se
  / s  \
You have the path the child process took. Starting where he started, you need to determine the fewest number of steps required to reach him. (A "step" means to move from the hex you are in to any adjacent hex.)

For example:

ne,ne,ne is 3 steps away.
ne,ne,sw,sw is 0 steps away (back where you started).
ne,ne,s,s is 2 steps away (se,se).
se,sw,se,sw,sw is 3 steps away (s,s,sw).

--- Part Two ---

How many steps away is the furthest he ever got from his starting position?

*/


package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil { panic(e); }
}

func abs(i int) int {
    if i < 0 { i *= -1; }
    return i
}

func dist(x, y int) int {
    if abs(x) > abs(y) {
        return abs(x)
    } else {
        return (abs(y) - abs(x))/2 + abs(x)
    }
}

func main () {
    file, err := os.Open("input.txt")
    check(err)
    defer file.Close()

    var x, y int
    var answer1, answer2 int


    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        for _, direction := range strings.Split(scanner.Text(), ",") {
            switch direction {
            case "n":
                y += 2
            case "ne":
                x += 1
                y += 1
            case "se":
                x += 1
                y -= 1
            case "s":
                y -= 2
            case "sw":
                x -= 1
                y -= 1
            case "nw":
                x -= 1
                y += 1
            }
            answer1 = dist(x, y)
            if answer2 < answer1 { answer2 = answer1 }
        }
    }

    fmt.Printf("[Part 1] Steps: %d\n", answer1)
    fmt.Printf("[Part 2] Steps: %d\n", answer2)
}
