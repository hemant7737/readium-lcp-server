package index

import (
  "testing"
  _ "github.com/mattn/go-sqlite3"
)

func TestIndexCreation(t *testing.T) {
  idx, err := Open(":memory:")
  if err != nil {
    t.Error("Can't open index")
    t.Error(err)
    t.FailNow()
  }

  p := Package{"test", []byte("1234"), "test.epub"}
  err = idx.Add(p)
  if err != nil {
    t.Error(err)
  }
  _, err = idx.Get("test")
  if err != nil {
    t.Error(err)
  }
}