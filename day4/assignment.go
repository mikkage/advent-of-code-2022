package main

import (
  "strconv"
  "strings"
)

// struct to represent an assignment to a range of sections
// Contains the minimum and maximum section IDs for the range
type assignment struct {
  Min int
  Max int
}

// Creates a new assignmnet struct from a string in the format MIN-MAX
func newAssignment(input string) assignment {
  r := strings.Split(input, "-")
  min, _ := strconv.Atoi(r[0])
  max, _ := strconv.Atoi(r[1])

  return assignment {
    Min: min,
    Max: max,
  }
}

// Returns whether the section ID range in a1 is fully contained in the section ID range of a2
func (a1 *assignment) isFullyContainedBy(a2 assignment) bool {
  return a1.Min >= a2.Min && a1.Max <= a2.Max
}

// Returns whether the section ID range in a1 has any overlap with the section ID range of a2
func (a1 *assignment) hasAnyOverlapWith(a2 assignment) bool {
  // Make a map which maps the section ID to whether it has been covered in the first assignment
  sections := make(map[int]bool)

  // For the range between a1's min and max(inclusive), mark all those IDs as covered
  for i := a1.Min; i <= a1.Max; i++ {
    sections[i] = true
  }

  // For the range between a2's min and max(inclusive), if any of the sections have been covered in the first assignment
  // then we've found overlap and can return true
  for i := a2.Min; i <= a2.Max; i++ {
    // Since this is adding to the value for last assignment, we can save some time here by returning immediately
    // if that value in the map is above one instead of iterating through the map once this loop has completed
    if sections[i] {
      return true
    }
  }

  // If it hasn't returned true by the time the second loop is over, then no values in the map were above one and there is no overlap
  return false
}
