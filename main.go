package main

import (
  "encoding/csv"
  "flag"
  "os"

  "github.com/nsf/termbox-go"
)



func main() {
  initializeTerm()
  defer termbox.Close()

  csvFilename := parseArgs()
  go handleKeys()
  go updateScreen(csvFilename)
}


func initializeTerm() {
  if err := termbox.Init(); err != nil {
    panic(err)
  }

  termbox.SetOutputMode(termbox.Output256)
}


func parseArgs() string {
  flag.Parse()

  if len(flag.Args()) != 1 {
    panic("Too many arguments.")
  }

  return flag.Args()[0]
}


func handleKeys() {
  for {
    switch event := termbox.PollEvent(); event.Type {
    case termbox.EventKey:
      switch event.Key {
      case termbox.KeyEsc:
        return
      }
    }
  }
}


func updateScreen(csvFilename string) {
  reader := readCsvFile(csvFilename)
  data := make([][]string, 1)

  for {
    record, err := reader.Read()
    if err.Error() == "EOF" {
      return
    } else if err != nil {
      panic(err)
    }

    data = append(data, record)

    drawGraph(data)
  }
}


func drawGraph(data [][]string) {
  //
  // draw graph of data
  //

  if err := termbox.Clear(' ', termbox.ColorDefault); err != nil {
    panic(err)
  }

  if err := termbox.Flush(); err != nil {
    panic(err)
  }
}


func readCsvFile(csvFilename string) *csv.Reader {
  return csv.NewReader(openFile(csvFilename))
}


func openFile(filename string) *os.File {
  if file, err := os.Open(filename); err != nil {
    panic(err)
  } else {
    return file
  }
}
