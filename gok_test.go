package gok

import (
  "testing"
)

func TestLex(t *testing.T) {

  var blocks chan Block = make(chan Block)

  go Lex("# xyz\n## abc def\n  \t\n## ghi\r### jkl\n## mno\r\npqr", blocks)

  block := <- blocks
  if block.Type != HEADER_ONE || block.Value != "xyz" {
    t.Errorf("FAILURE %d", block.Type)
  }

  block = <- blocks
  if block.Type != HEADER_TWO || block.Value != "abc def" {
    t.Errorf("FAILURE %d", block.Type)
  }

  block = <- blocks
  if block.Type != BLANK || block.Value != "  \t" {
    t.Errorf("FAILURE %d", block.Type)
  }

  block = <- blocks
  if block.Type != HEADER_TWO || block.Value != "ghi" {
    t.Errorf("FAILURE %d", block.Type)
  }

  block = <- blocks
  if block.Type != HEADER_THREE || block.Value != "jkl" {
    t.Errorf("FAILURE %d", block.Type)
  }

  block = <- blocks
  if block.Type != HEADER_TWO || block.Value != "mno" {
    t.Errorf("FAILURE %d", block.Type)
  }

  block = <- blocks
  if block.Type != PARAGRAPH || block.Value != "pqr" {
    t.Errorf("FAILURE %d", block.Type)
  }
}
