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

func TestSub(t *testing.T) {
	v1 := &Vec3{1, 2, 3}
	v2 := &Vec3{4, 5, 6}
	v1.Sub(v2)
	if !(*v1 == Vec3{-3, -3, -3}) {
		t.Fatal("failed Sub")
	}
}

func TestMul(t *testing.T) {
	v1 := &Vec3{1, 2, 3}
	v2 := &Vec3{4, 5, 6}
	v1.Mul(v2)
	if !(*v1 == Vec3{4, 10, 18}) {
		t.Fatal("failed Mul")
	}
}

func TestDiv(t *testing.T) {
	v1 := &Vec3{1, 2, 3}
	v2 := &Vec3{4, 5, 6}
	v1.Div(v2)
	if !(*v1 == Vec3{0.25, 0.4, 0.5}) {
		t.Fatal("failed Div")
	}
}
