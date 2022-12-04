package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  packs := readInput("input.txt")

  // Maps each item a-zA-Z to its numerical value
  itemValues := map[byte]int {
    'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,

    'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52,
  }

  fmt.Println("Part 1")
  sum := 0

  // For each rucksack parsed, create a new rucksack struct, find the items they share across compartments, and lookup the value of those shared items
  // Sum the values of the items in each pack, and add it to the total across all packs
  for i := range packs {
    r := newRucksack(packs[i])

    sharedItems := r.SharedItems()
    packSum := 0

    for j := range(sharedItems) {
      packSum += itemValues[sharedItems[j]]
    }
    sum += packSum
  }
  fmt.Println("Sum of shared items across all packs:", sum, "\n")

  fmt.Println("Part 2")
  sum = 0

  // Iterate over three packs at a time
  for i := 0; i < len(packs); i += 3 {
    firstPack := newRucksack(packs[i])
    secondPack := newRucksack(packs[i+1])
    thirdPack := newRucksack(packs[i+2])

    // To save repeat comparisons, remove duplicate item from the slice before iterating over it
    uniqueItems := unique(firstPack.Contents())

    // For each unique item in the slice, see if it's contained in both of the other two packs
    // If it is, look up that item's value and add it to the sum
    for j := range uniqueItems {
      if sliceContains(secondPack.Contents(), uniqueItems[j]) && sliceContains(thirdPack.Contents(), uniqueItems[j]) {
        sum += itemValues[uniqueItems[j]]
        break // Exit from the loop when a match is found since there is only one common item across each set of three packs
      }
    }
  }
  fmt.Println("Sum of shared items across all sets of three packs:", sum)
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

// Helper function which returns whether or not a given byte is contained in the given slice of bytes
func sliceContains(s []byte, b byte) bool {
  for i := range s {
    if s[i] == b {
      return true
    }
  }
  return false
}

// Helper function which takes in a slice of bytes and returns a new slice of bytes which does not contain any duplicate bytes
func unique(s []byte) []byte {
  // Map to store whether a byte has already been found in the slice
  m := make(map[byte]bool)
  var newSlice []byte

  for i := range s {
    b := s[i]

    // If the byte has not already been marked as found in the map, this is the first time seeing that byte in the slice
    // Add the byte to the new slice and mark it as being found in the map
    if m[b] == false {
      newSlice = append(newSlice, b)
      m[b] = true
    }
  }

  return newSlice
}
