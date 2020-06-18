package main

import (
  "fmt"
  "regexp"
  "strings"
)

func main() {
  Parse("# xyz\n## abc def\n  \t\n## ghi\r### jkl\n## mno\r\npqr")
}

type BlockType int

const (
  BLANK BlockType = iota
  HEADER_ONE
  HEADER_TWO
  HEADER_THREE
  HEADER_FOUR
  HEADER_FIVE
  HEADER_SIX
  PARAGRAPH
  EOF
)

type Block struct {
  Type BlockType
  Value string
}

func Parse(source string) {
  blocks := make(chan Block)
  go Lex(source, blocks)

  out := ""
  for block := <- blocks; block.Type != EOF; block = <- blocks {
    switch block.Type {
    case HEADER_ONE:
      out = strings.Join([]string{out, "<h1>", block.Value, "</h1>"}, "")
    case HEADER_TWO:
      out = strings.Join([]string{out, "<h2>", block.Value, "</h2>"}, "")
    case HEADER_THREE:
      out = strings.Join([]string{out, "<h3>", block.Value, "</h3>"}, "")
    case HEADER_FOUR:
      out = strings.Join([]string{out, "<h4>", block.Value, "</h4>"}, "")
    case HEADER_FIVE:
      out = strings.Join([]string{out, "<h5>", block.Value, "</h5>"}, "")
    case HEADER_SIX:
      out = strings.Join([]string{out, "<h6>", block.Value, "</h6>"}, "")
    case PARAGRAPH:
      out = strings.Join([]string{out, "<p>", block.Value, "</p>"}, "")
    default:
      out = strings.Join([]string{out, block.Value}, "")
    }
    out = strings.Join([]string{out, "\n"}, "")
  }
  fmt.Println(out)
}

var (
  endOfLine = regexp.MustCompile("\r\n|\r|\n")
  blank = regexp.MustCompile("^( |\t)*$")
)

func split(s string) []string {
  return endOfLine.Split(s, -1) 
}

func Lex(src string, ch chan Block) {
  for _, ln := range split(src) {
    ch <- lex(ln)
  }
  ch <- Block{ Type: EOF, Value: "" }
}

func lex(line string) Block {
  switch {
  case blank.MatchString(line):
    return Block{ Type: BLANK, Value: line }
  case line == "#" || len(line) > 1 && line[0:2] == "# ":
    return Block{ Type: HEADER_ONE, Value: strings.TrimPrefix(line, "# ") }
  case line == "##" || len(line) > 2 && line[0:3] == "## ":
    return Block{ Type: HEADER_TWO, Value: strings.TrimPrefix(line, "## ") }
  case line == "###" || len(line) > 3 && line[0:4] == "### ":
    return Block{ Type: HEADER_THREE, Value: strings.TrimPrefix(line, "### ") }
  case line == "####" || len(line) > 4 && line[0:5] == "#### ":
    return Block{ Type: HEADER_FOUR, Value: strings.TrimPrefix(line, "#### ") }
  case line == "#####" || len(line) > 5 && line[0:6] == "##### ":
    return Block{ Type: HEADER_FIVE, Value: strings.TrimPrefix(line, "##### ") }
  case line == "######" || len(line) > 6 && line[0:7] == "###### ":
    return Block{ Type: HEADER_SIX, Value: strings.TrimPrefix(line, "###### ") }
  default:
    return Block{ Type: PARAGRAPH, Value: line }
  }
}
