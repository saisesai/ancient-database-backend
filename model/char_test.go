package model

import (
	"fmt"
	"testing"
)

func TestNewChar(t *testing.T) {
	c, err := NewChar(0, "a")
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}

func TestFindChar(t *testing.T) {
	_, err := FindChars("id = 0")
	if err.Error() != "char not found" {
		panic("failed to make char not found error!")
	}
	chars, err := FindChars("id != ?", 0)
	if err != nil {
		panic(err)
	}
	for i, c := range chars {
		fmt.Println(i, c)
	}
}

func TestChar_Save(t *testing.T) {
	c, err := FindChars(1)
	if err != nil {
		panic(err)
	}
	c[0].Region.PosX = 0
	c[0].Region.PosY = 1
	c[0].Region.Width = 2
	c[0].Region.Height = 3
	err = c[0].Save()
	if err != nil {
		panic(err)
	}
}

func TestChar_Delete(t *testing.T) {
	c, err := NewChar(0, "1")
	if err != nil {
		panic(err)
	}
	err = c.Delete()
	if err != nil {
		panic(err)
	}
}
