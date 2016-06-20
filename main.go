package main

import (
  "os"
  "fmt"
  "strconv"
  "github.com/sleepydog1214/svg-generator/files"
  "github.com/sleepydog1214/svg-generator/svg"
  "github.com/sleepydog1214/svg-generator/defs"
)

var numRect, numEllipse, numPolygon, numPolyline int
var idx = 0

// groupVals generate number of shapes in a groupVals
func groupVals() () {
  numRect = svg.RandomInt(defs.ShapeParams["ShapeLB"], defs.ShapeParams["ShapeUB"])
  numEllipse = svg.RandomInt(defs.ShapeParams["ShapeLB"], defs.ShapeParams["ShapeUB"])
  numPolygon = svg.RandomInt(defs.ShapeParams["ShapeLB"], defs.ShapeParams["ShapeUB"])
  numPolyline = svg.RandomInt(defs.ShapeParams["ShapeLB"], defs.ShapeParams["ShapeUB"])
}

// drawRectangles draw random number of rectangles
func drawRectangles(fp *os.File) () {
  for i:= 0; i < numRect; i++ {
    svg.Rectangle(fp, idx)
    idx++
  }
}

// drawEllipses draw random number of ellipses
func drawEllipses(fp *os.File) () {
  for i:= 0; i < numEllipse; i++ {
    svg.Ellipse(fp, idx)
    idx++
  }
}

// drawPolygons draw random number of polygons
func drawPolygons(fp *os.File) () {
  for i:= 0; i < numPolygon; i++ {
    svg.Polygon(fp, idx)
    idx++
  }
}

// drawPolylines draw random number of polylines
func drawPolylines(fp *os.File) () {
  for i:= 0; i < numPolyline; i++ {
    svg.Polyline(fp, idx)
    idx++
  }
}

// main Draw a random svg file
func main() {
  // Create the svg file name
  fname := svgfile.CreateName()
  fmt.Println("Creating svg file: " + fname)

  // Open the svg file
  fp := svgfile.Open(fname)

  // Write the svg header
  svg.PrintHeader(fp, fname)

  // Get the number of groups
  groups := svg.RandomInt(defs.ShapeParams["GroupLB"], defs.ShapeParams["GroupUB"])

  // Draw the base rectangle
  svg.BaseRect(fp)
  fmt.Println("Creating base rectangle")

  fmt.Println("Creating " + strconv.Itoa(groups) + " groups")
  for i:= 0; i < groups; i++ {
    switch i {
    case 0:
      groupVals()

      svg.StartGroup(fp, idx, "1")
      idx++

      drawRectangles(fp)
      drawEllipses(fp)
      drawPolygons(fp)
      drawPolylines(fp)

      svg.EndGroup(fp)
    case 1:
      groupVals()

      svg.StartGroup(fp, idx, "0.85")
      idx++

      drawEllipses(fp)
      drawRectangles(fp)
      drawPolylines(fp)
      drawPolygons(fp)

      svg.EndGroup(fp)
    case 2:
      groupVals()

      svg.StartGroup(fp, idx, "0.80")
      idx++

      drawPolylines(fp)
      drawEllipses(fp)
      drawRectangles(fp)
      drawPolygons(fp)

      svg.EndGroup(fp)
    case 3:
      groupVals()

      svg.StartGroup(fp, idx, "0.75")
      idx++

      drawPolygons(fp)
      drawRectangles(fp)
      drawEllipses(fp)
      drawPolylines(fp)

      svg.EndGroup(fp)
    case 4:
      groupVals()

      svg.StartGroup(fp, idx, "0.70")
      idx++

      drawPolylines(fp)
      drawEllipses(fp)
      drawPolygons(fp)
      drawRectangles(fp)

      svg.EndGroup(fp)
    }
  }

  // Write the svg footer
  svg.PrintFooter(fp)

  // Close the svg file
  fmt.Println("Closing svg file: " + fname)
  svgfile.Close(fp)
}
