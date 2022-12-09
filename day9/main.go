package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

func main() {
  instructions := readInput("input.txt")

  // Use a silly size for the grid since we don't know how far it'll travel
  const gridSize = 3000
  var grid [gridSize][gridSize]bool

  // Start with the head and tail in the middle of the grid
  startLocation := gridSize / 2
  headX := startLocation
  headY := startLocation
  tailX := startLocation
  tailY := startLocation
  // The starting location counts as somnewhere that the tail has visited
  grid[startLocation][startLocation] = true

  for i := range instructions {
    // Get the direction the head moves and the magnitude
    strSplit := strings.Split(instructions[i], " ")
    direction := strSplit[0]
    distance, _ := strconv.Atoi(strSplit[1])

    // Defaults to 0 so we only have to set one of these to move the correct direction
    var xDirection int
    var yDirection int

    // Depending on the direction, figure out how the head will move
    switch direction {
    case "U":
      yDirection = 1
    case "D":
      yDirection = -1
    case "R":
      xDirection = 1
    case "L":
      xDirection = -1
    }

    // Move the distance, one unit at a time
    for i := 0; i < distance; i++ {
      // Store the location of the head before moving
      prevHeadX := headX
      prevHeadY := headY

      // Move the head
      headX += xDirection
      headY += yDirection

      // If the difference in the X or Y coordinates of the head and tail are greater than one, the tail
      // has to catch up. To catch up, the tail moves to the previous location of the head
      if math.Abs(float64(headX - tailX)) > 1 || math.Abs(float64(headY - tailY)) > 1 {
        tailX = prevHeadX
        tailY = prevHeadY
      }

      // Mark the current position of the tail, whether it moved or not this iteration as visited
      grid[tailY][tailX] = true
    }
  }

  totalVisited := 0
  for i := 0; i < gridSize; i++ {
    for j := 0; j < gridSize; j++ {
      if grid[i][j] {
        totalVisited += 1
      }
    }
  }

  fmt.Println("Part 1:", totalVisited)
}


func readInput(filename string) []string {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var instructions []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    instructions = append(instructions, line)
  }

  return instructions
}
