package svg

import (
  "os"
  "github.com/sleepydog1214/svg-generator/check"
)

func PrintHeader (fp *os.File, fname string) () {
  header := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
             <svg>`

  _, err := fp.WriteString(header)
  check.Error(err)
}

func PrintFooter (fp *os.File) () {
  footer := `<\svg>`

  _, err := fp.WriteString(footer)
  check.Error(err)
}
