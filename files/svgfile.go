package svgfile

import (
  "time"
  "os"
  "strings"
  "github.com/sleepydog1214/svg-generator/check"
)

// CreateName create the svg generated file name
func CreateName () (fname string) {
  temp := time.Now()
  ts := temp.Format(time.Stamp)
  ts = strings.Replace(ts, " ", "", -1)
  ts = strings.Replace(ts, ":", "", -1)
  fname = "tsart-generated-" + ts + ".svg"
  return fname
}

// Open opens the svg file for writing
func Open (fname string) (fp *os.File) {
  fp, err := os.Create(fname)
  check.Error(err)
  return
}

// Close closes the svg file
func Close (fp *os.File) () {
  err := fp.Close()
  check.Error(err)
}
