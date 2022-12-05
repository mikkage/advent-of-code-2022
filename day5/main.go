package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
  "strings"
)

func main() {
  stacks, instructions := readInput("input.txt")

  // Part 1
  for i := range instructions {
    instruction := instructions[i]
    // Instruction order is in the same order as in the input:
    // Number of crates to move, where it's moved from, then where it's moved to
    amount := instruction[0]
    sourceIndex := instruction[1]
    destinationIndex := instruction[2]

    // 'amount' times, do...
    for j := 0; j < amount; j++ {
      // Get the index of the last element in the list, and then store the value at that index
      topElementIndex := len(stacks[sourceIndex]) - 1
      topElement := stacks[sourceIndex][topElementIndex]

      // 'Pop' the element from the list by setting the slice's range from 0 to the index of the last element(not inclusive of last element)
      stacks[sourceIndex] = stacks[sourceIndex][:topElementIndex]

      // Push the element to the stack it's being moved to
      stacks[destinationIndex] += string(topElement)
    }
  }

  // Build the result
  result := ""

  // Need to iterate from 1 to 9 to get the crates in the correct order
  for i := 1; i <= 9; i++ {
    // Append the top item in the stack to the result
    result += string(stacks[i][len(stacks[i]) - 1])
  }

  fmt.Println("Part 1")
  fmt.Println("Result:", result, "\n")

  fmt.Println("Part 2")
  stacks, instructions = readInput("input.txt")
  result = ""

  for i := range instructions {
    // Instruction order is in the same order as in the input:
    instruction := instructions[i]
    amount := instruction[0]
    sourceIndex := instruction[1]
    destinationIndex := instruction[2]

    // To move multiple crates at once, we need to get the elements between the top of the stack and
    // the amount of boxes we want to move before that
    stackHeight := len(stacks[sourceIndex])
    moveElementIndex := stackHeight - amount

    // Get the elements that will be moved by taking the slice from the lowest crate being moved to the top crate
    topElements := stacks[sourceIndex][moveElementIndex:stackHeight]

    // Remove the elements from the stack that we just copied
    stacks[sourceIndex] = stacks[sourceIndex][:stackHeight - amount]

    // Push the elements to the stack they're being moved to
    stacks[destinationIndex] += topElements
  }

  for i := 1; i <= 9; i++ {
    // Append the top item in the stack to the result
    result += string(stacks[i][len(stacks[i]) - 1])
  }

  fmt.Println("Result:", result, "\n")
}

func readInput(filename string) (map[int]string, [][]int) {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  // Temporary storage for parsing the crates. Since they need to be stored in the opposite order that they're
  // given in the file, they will be read into here and processed once all have been read
  var crates []string

  // Instructions are represented by a list of integers and will be stored in a list of instructions
  var instructions [][]int

  instructionRegex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    // Separate the different types of input per line
    // If the line has a '[', it's a line that describes a crate
    // If the line has 'move' in it, it describes an instruction
    if strings.Contains(line, "[") {
      // Just add the line to a temporary list of crates, processing has to finish when all are done
      crates = append(crates, line)
    }
    if strings.Contains(line, "move") {
      // Using the regex above, extract the three numbers from the the line of input
      match := instructionRegex.FindAllStringSubmatch(line, -1)

      // Convert the strings into ints
      moveAmount, _ := strconv.Atoi(match[0][1])
      source, _ := strconv.Atoi(match[0][2])
      destination, _ := strconv.Atoi(match[0][3])

      // Create an 'instruction' using the three numbers by making a list containing the three numbers
      instruction := []int { moveAmount, source, destination }

      // Add the created instruction to the list of all instructions
      instructions = append(instructions, instruction)
    }
  }

  // Since the crates are given from top to bottom, but we need to add the bottom ones to the stacks first,
  // we need to reverse the order of the lines of the input to add each crate to its stack in the right order
  var revCrates []string
  for i := len(crates) - 1; i >= 0; i-- {
    revCrates = append(revCrates, crates[i])
  }

  stacks := make(map[int]string)

  for i := range revCrates {
    // Iterate over four characters of the string at a time
    for j := 0; j < len(revCrates[i]); j += 4 {
      // If the current four characters are all whitespace, there is no crate in this position and we can skip it
      if revCrates[i][j:j+4] == "    " {
        continue
      }
      // Otherwise, these four characters describe a crate and we should add it
      // Since we didn't use the numbers below each stack from the input, we need to figure out the stack it's in
      // using the inner loop index. Since each iteration of the loop increments the index by 4, we can simply add four
      // to the index, then divide it by four to get the number of the stack it's in.
      // For example:
      //  First iteration: (0 + 4) / 4 = 1
      //  Second iteration: (4 + 4) / 4 = 2
      stackIndex := (j + 4) / 4

      // We can finally push the crate to the top of the stack it belongs to
      stacks[stackIndex] += string(revCrates[i][j+1])
    }
  }

  // Finally ready to return all of the parsed data
  // stacks contains a map of ints to a list of strings. The key is the number of the stack, and the value is the stack itself
  // instructions contains a list of instructions, which are a list of ints
  return stacks, instructions
}
