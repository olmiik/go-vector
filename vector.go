/*
Copyright 2017 Albert Tedja

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vector

import (
	"math"
)

var EPSILON = math.Nextafter(1, 2) - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Vector represents mathematical vector.
type Vector []float64

// New returns vector of specified szie.
func New(size int) Vector {
	return make(Vector, size)
}

// NewWithValues returns vector with specified values.
// The size of new vector is equal to one of array.
func NewWithValues(values []float64) Vector {
	v := make(Vector, len(values))
	copy(v, values)
	return v
}

// Clone this vector, returning a new Vector.
func (v Vector) Clone() Vector {
	return NewWithValues(v)
}

// Set sets the values of this vector.
func (v Vector) Set(values []float64) {
	copy(v, values)
}

// Scale this vector (performs scalar multiplication) by the specified value.
func (v Vector) Scale(value float64) {
	l := len(v)
	for i := 0; i < l; i++ {
		v[i] *= value
	}
}

// Magnitude returns the magnitude of this vector.
func (v Vector) Magnitude() float64 {
	result := 0.0
	for _, e := range v {
		result += e * e
	}
	return math.Sqrt(result)
}

// Zero sets all values to zero.
func (v Vector) Zero() {
	for i := range v {
		v[i] = 0.0
	}
}

// Do iterates over the elements and invokes sepcified function.
func (v Vector) Do(applyFn func(float64) float64) {
	for i, e := range v {
		v[i] = applyFn(e)
	}
}

// DoWithIndex iterates over the elements and invokes sepcified function.
func (v Vector) DoWithIndex(applyFn func(int, float64) float64) {
	for i, e := range v {
		v[i] = applyFn(i, e)
	}
}

// Add adds another vector and returns resutl as new vector.
func (v Vector) Add(other Vector) Vector {
	l := min(len(v), len(other))
	result := make(Vector, l)
	for i := 0; i < l; i++ {
		result[i] = v[i] + other[i]
	}
	return result
}

// Sub substracts another vector and returns result as new vector.
func (v Vector) Sub(other Vector) Vector {
	l := min(len(v), len(other))
	result := make(Vector, l)
	for i := 0; i < l; i++ {
		result[i] = v[i] - other[i]
	}
	return result
}

// Dot computes dot product with another vector.
// Another vector must have the same dimensionality.
func (v Vector) Dot(other Vector) (float64, error) {
	if len(v) != len(other) {
		return 0.0, ErrVectorNotSameSize
	}

	l := len(v)
	result := 0.0
	for i := 0; i < l; i++ {
		result += v[i] * other[i]
	}

	return result, nil
}

// Cross computes cross-product with another vector.
// Vector dimensionality msut be equal to 3
func (v Vector) Cross(other Vector) (Vector, error) {
	// Early error check to prevent redundant cloning
	if len(v) != 3 || len(other) != 3 {
		return nil, ErrVectorInvalidDimension
	}

	result := make(Vector, 3)
	result[0] = v[1]*other[2] - v[2]*other[1]
	result[1] = v[2]*other[0] - v[0]*other[2]
	result[2] = v[0]*other[1] - v[1]*other[0]

	return result, nil
}

// Unit computes unit vector result as new vector.
func Unit(v Vector) Vector {
	magRec := 1.0 / v.Magnitude()
	unit := v.Clone()
	for i := range unit {
		unit[i] *= magRec
	}

	return unit
}

// Hadamard computes Hadamard product with another vector
// and returns result as new vector. Another vector must
// have the same dimensionality.
func (v Vector) Hadamard(other Vector) (Vector, error) {
	if len(v) != len(other) {
		return nil, ErrVectorInvalidDimension
	}

	l := len(v)
	result := make(Vector, l)
	for i := 0; i < l; i++ {
		result[i] = v[i] * other[i]
	}

	return result, nil
}
