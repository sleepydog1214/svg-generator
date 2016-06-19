package check

func Error (err error) () {
  if err != nil {
    panic(err)
  }
}
