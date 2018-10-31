package tiled

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	m, err := Parse("tilemaptest.json")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(m) // use -v flag for go test
}
