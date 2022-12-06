package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  signal := readInput("input.txt")

  // Part 1
  startIndex := 0
  for i := 0; i < len(signal); i++ {
    // Grab the four bytes starting at index i
    signalBlock := []byte(signal[i:i+4])

    // If we remove all duplicate characters from the block and it's still four characters,
    // then all characters are unique and we've found the start of packet marker
    if len(unique(signalBlock)) == 4 {
      // Now that we've found the start of packet marker, we just need to get the index at which
      // the last character was processed before finding it and can drop out of the loop
      startIndex = i + 4
      break
    }
  }
  fmt.Println("Part 1:", startIndex)

  // Part 2
  // Exact same as part one, but looking at blocks of 14 instead of 4
  startIndex = 0
  for i := 0; i < len(signal); i++ {
    signalBlock := []byte(signal[i:i+14])
    if len(unique(signalBlock)) == 14 {
      startIndex = i + 14
      break
    }
  }
  fmt.Println("Part 2:", startIndex)
}

func readInput(filename string) string {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  str := ""

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    str += line
  }

  return str
}

// Helper function which takes in a slice of bytes and returns a slice which only contains one of each character from the input slice
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
