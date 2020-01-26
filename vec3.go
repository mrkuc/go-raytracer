package main

type Vec3 struct {
	x float32
	y float32
	z float32
}

func (v1 *Vec3) Add(v2 *Vec3) {
	v1.x += v2.x
	v1.y += v2.y
	v1.z += v2.z
}

func (v1 *Vec3) Sub(v2 *Vec3) {
	v1.x -= v2.x
	v1.y -= v2.y
	v1.z -= v2.z
}

func (v1 *Vec3) Mul(v2 *Vec3) {
	v1.x *= v2.x
	v1.y *= v2.y
	v1.z *= v2.z
}

func (v1 *Vec3) Div(v2 *Vec3) {
	v1.x /= v2.x
	v1.y /= v2.y
	v1.z /= v2.z
}
