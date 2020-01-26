package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	v1 := &Vec3{1, 2, 3}
	v2 := &Vec3{4, 5, 6}
	v1.Add(v2)
	if !(*v1 == Vec3{5, 7, 9}) {
		t.Fatal("failed Add")
	}
}
