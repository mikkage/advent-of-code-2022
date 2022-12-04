package main

import (
  "strings"
)

type rucksack struct {
  FirstCompartment string
  SecondCompartment string
}

// Create a new rucksack struct from a string
// The first half of characters is placed in the first compartment and the second half in the second compartment
func newRucksack(input string) rucksack {
  dividingIndex := len(input) / 2

  return rucksack {
    FirstCompartment: input[0:dividingIndex],
    SecondCompartment: input[dividingIndex:],
  }
}

// Returns a list of bytes of the items shared between both compartments in a rucksack
func(r *rucksack) SharedItems() []byte {
  var sharedItems []byte

  // For each item in the first compartment, check if it exists in the second
  // If it is in both, do a check to make sure that item isn't already in the list of shared items, then add it if not
  for i := range r.FirstCompartment {
    if strings.Contains(r.SecondCompartment, string(r.FirstCompartment[i])) {
      if !sliceContains(sharedItems, r.FirstCompartment[i]) {
        sharedItems = append(sharedItems, r.FirstCompartment[i])
      }
    }
  }

  return sharedItems
}

// Concatenates the contents of the first and second compartment and returns it as a list of bytes
func (r *rucksack) Contents() []byte {
  return []byte(r.FirstCompartment + r.SecondCompartment)
}
