package model

import (
	"fmt"
	"testing"
)

func TestRegion_Scan(t *testing.T) {
	r := &Region{}
	err := r.Scan([]byte(`{"pos_x": 0, "pos_y": 1, "width": 2, "height": 3}`))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func TestRegion_Value(t *testing.T) {
	r := Region{
		PosX:   1,
		PosY:   2,
		Width:  3,
		Height: 4,
	}
	v, err := r.Value()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v.([]byte)))
}
