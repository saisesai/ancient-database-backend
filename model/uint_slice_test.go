package model

import (
	"fmt"
	"testing"
)

func TestUintSlice_Scan(t *testing.T) {
	us := &UintSlice{}
	err := us.Scan([]byte("[0,1,2]"))
	if err != nil {
		panic(err)
	}
	fmt.Println(us)
}

func TestUintSlice_Value(t *testing.T) {
	us := UintSlice{1, 2, 3}
	v, err := us.Value()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v.([]byte)))
}

func TestUintSlice_FindAndDelete(t *testing.T) {
	us := &UintSlice{1, 2, 2, 2, 3}
	us.FindAndDelete(2)
	fmt.Println(us)
}

func TestUintSlice_ToIntSlice(t *testing.T) {
	us := &UintSlice{1, 2, 3}
	fmt.Println(us.ToIntSlice())
}
