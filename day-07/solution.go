/*
--- Day 7: Recursive Circus ---

Wandering further through the circuits of the computer, you come upon a tower of programs that have gotten themselves into a bit of trouble. A recursive algorithm has gotten out of hand, and now they're balanced precariously in a large tower.

One program at the bottom supports the entire tower. It's holding a large disc, and on the disc are balanced several more sub-towers. At the bottom of these sub-towers, standing on the bottom disc, are other programs, each holding their own disc, and so on. At the very tops of these sub-sub-sub-...-towers, many programs stand simply keeping the disc below them balanced but with no disc of their own.

You offer to help, but first you need to understand the structure of these towers. You ask each program to yell out their name, their weight, and (if they're holding a disc) the names of the programs immediately above them balancing on that disc. You write this information down (your puzzle input). Unfortunately, in their panic, they don't do this in an orderly fashion; by the time you're done, you're not sure which program gave which information.

For example, if your list is the following:

pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
...then you would be able to recreate the structure of the towers that looks like this:

                gyxo
              /
         ugml - ebii
       /      \
      |         jptl
      |
      |         pbga
     /        /
tknk --- padx - havc
     \        \
      |         qoyq
      |
      |         ktlj
       \      /
         fwft - cntj
              \
                xhth
In this example, tknk is at the bottom of the tower (the bottom program), and is holding up ugml, padx, and fwft. Those programs are, in turn, holding up other programs; in this example, none of those programs are holding up any other programs, and are all the tops of their own towers. (The actual tower balancing in front of you is much larger.)

Before you're ready to help them, you need to make sure your information is correct. What is the name of the bottom program?

--- Part Two ---

The programs explain the situation: they can't get down. Rather, they could get down, if they weren't expending all of their energy trying to keep the tower balanced. Apparently, one program has the wrong weight, and until it's fixed, they're stuck here.

For any program holding a disc, each program standing on that disc forms a sub-tower. Each of those sub-towers are supposed to be the same weight, or the disc itself isn't balanced. The weight of a tower is the sum of the weights of the programs in that tower.

In the example above, this means that for ugml's disc to be balanced, gyxo, ebii, and jptl must all have the same weight, and they do: 61.

However, for tknk to be balanced, each of the programs standing on its disc and all programs above it must each match. This means that the following sums must all be the same:

ugml + (gyxo + ebii + jptl) = 68 + (61 + 61 + 61) = 251
padx + (pbga + havc + qoyq) = 45 + (66 + 66 + 66) = 243
fwft + (ktlj + cntj + xhth) = 72 + (57 + 57 + 57) = 243
As you can see, tknk's disc is unbalanced: ugml's stack is heavier than the other two. Even though the nodes above ugml are balanced, ugml itself is too heavy: it needs to be 8 units lighter for its stack to weigh 243 and keep the towers balanced. If this change were made, its weight would be 60.

Given that exactly one program is the wrong weight, what would its weight need to be to balance the entire tower?
*/


package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil { panic(e); }
}

type TreeNode struct {
    Name string
    Weight int
    Children []*TreeNode
    ChildrenWeight int
    TowerWeight int
}

func buildTree(node *TreeNode, nodesMap map[string]*TreeNode) {
    if len(node.Children) == 0 { return; }

    for i, _ := range node.Children {
        node.Children[i] = nodesMap[node.Children[i].Name]
        buildTree(node.Children[i], nodesMap)
        node.ChildrenWeight += node.Children[i].Weight + node.Children[i].ChildrenWeight
        node.TowerWeight = node.Weight + node.ChildrenWeight
    }
}

func findExpectedWeight(node *TreeNode, eWeight int) (int, bool) {
    if len(node.Children) == 0 {return 0, false; }

    nWeight := 0
    badNodeId := -1
    weightMap := make(map[int][]int)
    for i, child := range node.Children {
        weightMap[child.TowerWeight] = append(weightMap[child.TowerWeight], i)
    }
    for k, v := range weightMap {
        if len(v) == 1 {
            badNodeId = v[0]
        } else {
            nWeight = k
        }
    }
    if badNodeId == -1 { return eWeight - node.ChildrenWeight, true; }
    return findExpectedWeight(node.Children[badNodeId], nWeight)
}

func main () {
    file, err := os.Open("input.txt")
    check(err)
    defer file.Close()

    var root *TreeNode
    childList := make(map[string]struct{})
    nodesMap := make(map[string]*TreeNode)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " -> ")
        s := strings.Split(line[0], " ")
        nodeName := s[0]
        nodeWeight, _ := strconv.Atoi(strings.Trim(s[1], "()"))
        nodesMap[nodeName] = &TreeNode{Name: nodeName, Weight: nodeWeight}

        if len(line) < 2 { continue; }
        for _, child := range strings.Split(line[1], ", ") {
            childList[child] = struct{}{}
            nodesMap[nodeName].Children = append(nodesMap[nodeName].Children, &TreeNode{Name: child})
        }
    }

    for name, _ := range nodesMap {
        if _, ok := childList[name]; !ok { root = nodesMap[name]; }
    }

    buildTree(root, nodesMap)
    weight, _ := findExpectedWeight(root, 0)

    fmt.Printf("[Part 1] Bottom program:  %s\n", root.Name)
    fmt.Printf("[Part 2] Expected weight: %d\n", weight)
}
