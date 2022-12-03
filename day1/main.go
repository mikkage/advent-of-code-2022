package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  var sums []int
  var currentSum int

  file, err := os.Open("input.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

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

  // When finished with the above loop, the sums of each inventory will be in a list and we just have to get the largest value from it
  fmt.Println("The largest sum is:", max(sums))
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
