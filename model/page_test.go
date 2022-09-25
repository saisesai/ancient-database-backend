package model

import (
	"fmt"
	"testing"
)

func TestNewPage(t *testing.T) {
	p, err := NewPage(123)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

func TestFindPages(t *testing.T) {
	_, err := FindPages("id == 0")
	if err.Error() != "page not found" {
		panic("failed to make page not found error!")
	}
	ps, err := FindPages("id != 0")
	if err != nil {
		panic(err)
	}
	for i, p := range ps {
		fmt.Println(i, p)
	}
}

func TestPage_Save(t *testing.T) {
	ps, err := FindPages("id == 1")
	if err != nil {
		panic(err)
	}
	ps[0].Chars = append(ps[0].Chars, 233, 123, 43242)
	err = ps[0].Save()
	if err != nil {
		panic(err)
	}
}

func TestPage_Delete(t *testing.T) {
	p, err := NewPage(0)
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p.ID, "1")
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p.ID, "2")
	if err != nil {
		panic(err)
	}
	err = p.Reload()
	if err != nil {
		panic(err)
	}
	err = p.Delete()
	if err != nil {
		panic(err)
	}
}
