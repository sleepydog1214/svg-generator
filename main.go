package main

import (
  "fmt"
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

  // Write the svg footer
  svg.PrintFooter(fp)
  
  // Close the svg file
  fmt.Println("Closing svg file: " + fname)
  svgfile.Close(fp)
}
