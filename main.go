package main

import (
  "fmt"
  "strconv"
  "github.com/sleepydog1214/svg-generator/files"
  "github.com/sleepydog1214/svg-generator/svg"
)

func main() {
  // Create the svg file name
  fname := svgfile.CreateName()
  fmt.Println("Creating svg file: " + fname)

  // Open the svg file
  fp := svgfile.Open(fname)

  // Write the svg header
  svg.PrintHeader(fp, fname)

  // Get the number of groups
  groups := svg.RandomInt(3, 6)
  fmt.Println("Creating " + strconv.Itoa(groups) + " groups")

  // Draw the base rectangle
  svg.BaseRect(fp)

  idx := 0
  svg.Rectangle(fp, idx)
  
  // Write the svg footer
  svg.PrintFooter(fp)

  // Close the svg file
  fmt.Println("Closing svg file: " + fname)
  svgfile.Close(fp)
}
