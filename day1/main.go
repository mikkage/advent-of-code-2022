package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  sums := parseInventories(file)

  fmt.Println("Part 1")
  fmt.Println("The largest sum is:", max(sums), "\n")

  fmt.Println("Part 2")
  fmt.Println("The three largest inventories are:", largestElements(sums, 3))
  fmt.Println("The sum of the three largest inventories is:", sum(largestElements(sums, 3)))
}

func parseInventories(file *os.File) []int {
  var sums []int
  var currentSum int

  // Scan each line of the file
  // If the line is empty, we've hit the divider between two inventories, and should add the current sum to the list of inventory sums, and reset the current sum to 0
  // If the line isn't empty, convert the value into an integer and add it to the current sum
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      sums = append(sums, currentSum)
      currentSum = 0
    } else {
      result, _ := strconv.Atoi(line)
      currentSum = currentSum + result
    }
  }

  // Since the sum is only added when a blank line is found, the last sum won't get added in the loop above, add it here
  sums = append(sums, currentSum)

  return sums
}

// Returns the N largest values from a list of integers
func largestElements(elements []int, amount int) []int {
  sort.Ints(elements)
  return elements[len(elements)-amount:]
}

func sum(elements []int) int {
  sum := 0

  for i := range elements {
    sum += elements[i]
  }

  return sum
}

// Returns the largest value from a list of integers
func max(elements []int) int {
  var max = elements[0]

  for i := range elements {
    if elements[i] > max {
      max = elements[i]
    }
  }

  return max
}
