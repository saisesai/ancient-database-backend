package model

import (
	"fmt"
	"testing"
)

func TestNewArtwork(t *testing.T) {
	c, err := NewArtwork()
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}

func TestFindArtworks(t *testing.T) {
	cs, err := FindArtworks("ti_ming = ?", "")
	if err != nil {
		panic(err)
	}
	for i, v := range cs {
		fmt.Println(i, v)
	}
}

func TestArtwork_Save(t *testing.T) {
	cs, err := FindArtworks(1)
	if err != nil {
		panic(err)
	}
	cs[0].Cover = "https://114514.com/homo.jpg"
	cs[0].Pages = append(cs[0].Pages, 1, 2, 3)
	err = cs[0].Save()
}

func TestArtwork_Delete(t *testing.T) {
	a, err := NewArtwork()
	if err != nil {
		panic(err)
	}
	p1, err := NewPage(a.ID)
	if err != nil {
		panic(err)
	}
	p2, err := NewPage(a.ID)
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p1.ID, "c11")
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p1.ID, "c12")
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p2.ID, "c21")
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p2.ID, "c22")
	if err != nil {
		panic(err)
	}
	err = a.Delete()
	if err != nil {
		panic(err)
	}
}

func TestArtwork_Delete2(t *testing.T) {
	a, err := NewArtwork()
	if err != nil {
		panic(err)
	}
	p1, err := NewPage(a.ID)
	if err != nil {
		panic(err)
	}
	p2, err := NewPage(a.ID)
	if err != nil {
		panic(err)
	}
	c11, err := NewChar(p1.ID, "c11")
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p1.ID, "c12")
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p2.ID, "c21")
	if err != nil {
		panic(err)
	}
	_, err = NewChar(p2.ID, "c22")
	if err != nil {
		panic(err)
	}
	err = p2.Delete()
	if err != nil {
		panic(err)
	}
	err = c11.Delete()
	if err != nil {
		panic(err)
	}
}
