package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  // Maps the inputs expected from the file to what move they represent
  // X, Y, and Z mappings are only valid for part 1
  inputs := map[string]string {
    "A": "rock",
    "X": "rock",
    "B": "paper",
    "Y": "paper",
    "C": "scissors",
    "Z": "scissors",
  }

  // Maps each move and game outcome to how many points each is worth
  scores := map[string]int {
    "rock": 1,
    "paper": 2,
    "scissors": 3,

    "loss": 0,
    "win": 6,
    "tie": 3,
  }

  totalScore := 0
  strategyGuide := readInput("input.txt")

  // Part 1
  // For each game in the strategy guide, get the moves the characters map to, play the game with those moves, then add the point values for the your move and the game outcome to the total
  for i := range strategyGuide {
    // When each line of the strategy guide is split on a space, the first element will be the move your opponent will make and the second will be the move you will make
    moves := strings.Split(strategyGuide[i], " ")
    opponentMove := inputs[moves[0]]
    yourMove := inputs[moves[1]]

    rpsResult := playRps(yourMove, opponentMove)
    totalScore += scores[rpsResult]
    totalScore += scores[yourMove]
  }

  fmt.Println("Part 1")
  fmt.Println("Total score across all RPS games:", totalScore, "\n")

  // Part 2
  fmt.Println("Part 2")
  totalScore = 0

  // Modify the inputs map to reflect the change in what X, Y, and Z mean for part 2
  // These now correspond to whether you lose, tie, or win rather than rock, paper, or scissors
  inputs["X"] = "lose"
  inputs["Y"] = "tie"
  inputs["Z"] = "win"

  // For each game in the strategy guide:
  // Get the move the first character maps to and the expected outcome of the game
  // Using those, get the move you should make to result in the expected outcome
  // Add the point values for the your move and the game outcome to the total
  for i := range strategyGuide {
    // When each line of the strategy guide is split on a space, the first element will be the move your opponent will make and the second will be the required outcome of the game
    moves := strings.Split(strategyGuide[i], " ")
    opponentMove := inputs[moves[0]]
    outcome := inputs[moves[1]]

    yourMove := playRpsReverse(opponentMove, outcome)
    totalScore += scores[outcome]
    totalScore += scores[yourMove]
  }

  fmt.Println("Total score across all RPS games:", totalScore, "\n")
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

// Takes in "rock", "paper", or "scissors" as the player's and opponent's selection, and returns "win", "lose", or "tie" from the player's perspective
func playRps(playerChoice, opponentChoice string) string {
  if playerChoice == opponentChoice {
    return "tie"
  }
  if playerChoice == "rock" && opponentChoice == "scissors" {
    return "win"
  }
  if playerChoice == "paper" && opponentChoice == "rock" {
    return "win"
  }
  if playerChoice == "scissors" && opponentChoice == "paper" {
    return "win"
  }
  return "lose"
}

// Plays RPS, but kind of in reverse. Instead of taking in two RPS inputs and returning the result, take in the opponent's move
// and the expected outcome to return the move your should make to get that outcome
func playRpsReverse(opponentChoice, outcome string) string {
  // If the game has to tie, simply return the opponent's choice
  if outcome == "tie" {
    return opponentChoice
  }

  var yourChoice string

  // If the expected outcome is to win, use a mapping of the opponent's choice to the move that would beat it to get the correct choice
  if outcome == "win" {
    opponentLosingMatchups := map[string]string {
      "rock": "paper",
      "paper": "scissors",
      "scissors": "rock",
    }
    yourChoice = opponentLosingMatchups[opponentChoice]
  }

  // If the expected outcome is to lose, use a mapping of the opponent's choice to the move that would lose to it to ge the correct choice
  if outcome == "lose" {
    opponentWinningMatchups := map[string]string {
      "rock": "scissors",
      "paper": "rock",
      "scissors": "paper",
    }
    yourChoice = opponentWinningMatchups[opponentChoice]
  }

  return yourChoice
}
