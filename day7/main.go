package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

func main() {
  outputLines := readInput("input.txt")

  // Use a stack to keep track of the level of nested directories we're in.
  // Change to a nested directory by pushing, cgo up a level by popping
  var path []string

  // Map a directory path to the total size of its files and files nested in directories
  // The path stack is joined with a "/" to create the key. For example, if the stack is ["abc", "123", "foo"],
  // the key "abc/123/foo" will contain the total size of that directory
  var dirs = make(map[string]int)

  for i := range outputLines {
    line := outputLines[i]
    // The line is a command(ls, cd) if it starts with $, but we only have to care about cd
    if strings.HasPrefix(line, "$") {
      if strings.Contains(line, "cd") {
        // When it's a cd command, we just need to get where it's changing to
        cdCmdSplit := strings.Split(line, " ")
        cdDirectory := cdCmdSplit[2]

        // cd / does nothing as it's always the first command and never used after
        switch cdDirectory {
        case "/":
        case "..":
          // Going up a directory just means we pop from the path stack
          path = path[:len(path) - 1]
        default:
          // Otherwise we changing to another directory in the current directory, so we push the dir
          // onto the path stack
          path = append(path, cdDirectory)
        }
      }
    } else {
      // If the command doesn't start with $, it's output from an ls command
      // We can ignore any lines that have 'dir' as it gets handled by keeping dirs in the path stack
      if !strings.HasPrefix(line, "dir") {
        // This has to be a file, so get the size
        fileOutput := strings.Split(line, " ")
        fileSize, _ := strconv.Atoi(fileOutput[0])

        // For each level in the path stack, add the file size to it
        // For example, if the path stack is ["a", "b", "c"], it will add the file size to
        // /a, /a/b, and /a/b/c
        for i := range path {
          p := strings.Join(path[:i], "/")
          dirs[p] += fileSize
        }
      }
    }
  }

  // Once the above loop is finished, each key will be a path of a directory mapped to its size
  // and we can just sum up the ones under the required size
  sum := 0
  for _, v := range dirs {
    if v < 100000 {
      sum += v
    }
  }

  // Part 1
  fmt.Println(sum)

  // Part 2
  totalSize := 70000000
  requiredSpace := 30000000
  free := totalSize - dirs[""]

  // To get the minimumm amount of space we need to free up, need to subract the current free
  // space from the space required for the update
  requiredSpace = requiredSpace - free

  // Go through and get the size of each directory. If it's gte the amount of space we need
  // to free up, then add it to a list of all potential dirs
  var deletionCandidates []int
  for _, v := range dirs {
    if v >= requiredSpace {
      deletionCandidates = append(deletionCandidates, v)
    }
  }

  // Sort the list of all directory sizes in ascending order and the smallest directory we can
  // delete is first in the list after
  sort.Ints(deletionCandidates)
  fmt.Println(deletionCandidates[0])
}

func readInput(filename string) []string {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var lines []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    lines = append(lines, line)
  }

  return lines
}
