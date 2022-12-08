package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  grid := readInput("input.txt")

  // Part 1
  fmt.Println(countVisibleNodes(grid))

  // Part 2
  fmt.Println(findHighestScenicScore(grid))
}


func readInput(filename string) [][]int {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var grid [][]int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
      // Each line is a row of numbers
      var row []int
      chars := strings.Split(line, "")
      for j := range chars {
        val, _ := strconv.Atoi(string(chars[j]))
        row = append(row, val)
      }

      // Add each row to form the full grid
      grid = append(grid, row)
  }

  return grid
}

// Part 1
// Counts the total number of visible trees from the the outer edges of the grid
func countVisibleNodes(grid [][]int) int {
  gridSize := len(grid[0])
  visibleCount := 0

  for y := range grid {
    for x := range grid[y] {
      // Quick dropout to catch all of the nodes at the edge of the grid
      if (x == 0 || y == 0 || x == gridSize-1 || y == gridSize-1) {
        visibleCount += 1
      } else {
        // Create a list of all nodes above the current one
        var nodesAbove []int
        for i := y - 1; i >= 0; i-- {
          nodesAbove = append(nodesAbove, grid[i][x])
        }
        // If all nodes above are lower than the current, then it can be seen from the edge
        // and we can stop checking any other directions
        if allElementsLowerThan(grid[y][x], nodesAbove) {
          visibleCount += 1
          continue
        }

        // Do the same as above, but for nodes below the current
        var nodesBelow []int
        for i := y + 1; i < gridSize; i++ {
          nodesBelow = append(nodesBelow, grid[i][x])
        }
        if allElementsLowerThan(grid[y][x], nodesBelow) {
          visibleCount += 1
          continue
        }

        // Do the same as above, but for nodes to the left of the current one
        var nodesLeft []int
        for i := x - 1; i >= 0; i-- {
          nodesLeft = append(nodesLeft, grid[y][i])
        }
        if allElementsLowerThan(grid[y][x], nodesLeft) {
          visibleCount += 1
          continue
        }

        // Do the same as above, but for nodes to the right of the current one
        var nodesRight []int
        for i := x + 1; i < gridSize; i++ {
          nodesRight = append(nodesRight, grid[y][i])
        }
        if allElementsLowerThan(grid[y][x], nodesRight) {
          visibleCount += 1
          continue
        }
      }
    }
  }

  return visibleCount
}

// Part 2
// Goes through each node in the grid and finds the highest scenic score
func findHighestScenicScore(grid [][]int) int {
  gridSize := len(grid[0])
  highestScore := 0

  for y := range grid {
    for x := range grid[y] {
      // Nodes on the edge will always have at least one view distance of 0 so they can be ignored
      if (x == 0 || y == 0 || x == gridSize-1 || y == gridSize-1) {
        continue
      } else {
        // Get a list of all nodes above the current one
        var nodesAbove []int
        for i := y - 1; i >= 0; i-- {
          nodesAbove = append(nodesAbove, grid[i][x])
        }

        // Get a list of all nodes below the current one
        var nodesBelow []int
        for i := y + 1; i < gridSize; i++ {
          nodesBelow = append(nodesBelow, grid[i][x])
        }

        // Get a list of all nodes to the left of the current one
        var nodesLeft []int
        for i := x - 1; i >= 0; i-- {
          nodesLeft = append(nodesLeft, grid[y][i])
        }

        // Get a list of all nodes to the left of the current one
        var nodesRight []int
        for i := x + 1; i < gridSize; i++ {
          nodesRight = append(nodesRight, grid[y][i])
        }

        // Calculate the viewing distance in each direction from the current node
        upViewingDistance := viewingDistance(grid[y][x], nodesAbove)
        downViewingDistance := viewingDistance(grid[y][x], nodesBelow)
        leftViewingDistance := viewingDistance(grid[y][x], nodesLeft)
        rightViewingDistance := viewingDistance(grid[y][x], nodesRight)

        // The score is calculated by multiplying the view distance in all directions
        score := upViewingDistance * downViewingDistance * leftViewingDistance * rightViewingDistance

        // Update the highest score if the current one is higher
        if score > highestScore {
          highestScore = score
        }
      }
    }
  }

  return highestScore
}

// Returns whether all of the elements in array s are lower than n
func allElementsLowerThan(n int, s []int) bool {
  for i := range s {
    if s[i] >= n {
      return false
    }
  }

  return true
}

// Gets the viewing distance given the height at the current node, and a list of other heights
func viewingDistance(startingHeight int, otherHeights []int) int {
  distance := 0

  for i := range otherHeights {
    // Increment the total viewable distance by one no matter what since even if the height is greater
    // than the current, it would still count as viewable
    distance += 1
    if startingHeight <= otherHeights[i] {
      // Once the height is equal or greater than the starting height, we can't see anything passed it
      // so we've reach the longest viewable distance
      return distance
    }
  }

  // If we exit the loop without returning, then all heights were lower and the entire line of sight is viewable
  return distance
}
