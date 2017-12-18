/*
--- Day 4: High-Entropy Passphrases ---

A new system policy has been put in place that requires all accounts to use a passphrase instead of simply a password. A passphrase consists of a series of words (lowercase letters) separated by spaces.

To ensure security, a valid passphrase must contain no duplicate words.

For example:

aa bb cc dd ee is valid.
aa bb cc dd aa is not valid - the word aa appears more than once.
aa bb cc dd aaa is valid - aa and aaa count as different words.
The system's full passphrase list is available as your puzzle input. How many passphrases are valid?

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

For added security, yet another system policy has been put in place. Now, a valid passphrase must contain no two words that are anagrams of each other - that is, a passphrase is invalid if any word's letters can be rearranged to form any other word in the passphrase.

For example:

abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
Under this new system policy, how many passphrases are valid?
*/


package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main () {
    file, err := os.Open("input.txt")
    check(err)
    defer file.Close()

    var answer1, answer2 int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        map_puzzle1 := make(map[string] int)
        map_puzzle2 := make(map[string] int)
        answer1 += 1
        answer2 += 1

        line := scanner.Text()
        for _, word := range strings.Split(line, " ") {
            map_puzzle1[word] += 1
            letters := strings.Split(word, "")
            sort.Strings(letters)
            map_puzzle2[strings.Join(letters, "")] += 1
        }

        for _, v := range map_puzzle1 {
            if v > 1 {
                answer1 -= 1
                break
            }
        }

        for _, v := range map_puzzle2 {
            if v > 1 {
                answer2 -= 1
                break
            }
        }
    }

    fmt.Printf("[Part 1] Valid: %d\n", answer1)
    fmt.Printf("[Part 2] Valid: %d\n", answer2)
}
