package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  assignments := readInput("input.txt")

  fullyOverlappingAssignments := 0
  anyOverlappingAssignments := 0

  for i := range assignments {
    assignmentPair := strings.Split(assignments[i], ",")
    firstAssignment := newAssignment(assignmentPair[0])
    secondAssignment := newAssignment(assignmentPair[1])

    // Part 1
    if firstAssignment.isFullyContainedBy(secondAssignment) || secondAssignment.isFullyContainedBy(firstAssignment) {
      fullyOverlappingAssignments += 1
    }

    // Part 2
    if firstAssignment.hasAnyOverlapWith(secondAssignment) {
      anyOverlappingAssignments += 1
    }
  }

  fmt.Println("Total number of fully overlapping assigment pairs:", fullyOverlappingAssignments)
  fmt.Println("Total number of at least partially overlapping assignment pairs:", anyOverlappingAssignments)
}

func readInput(filename string) []string {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var output []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    output = append(output, line)
  }

  return output
}
