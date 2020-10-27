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

type Vector []float64

func New(size int) Vector {
	return make(Vector, size)
}

func NewWithValues(values []float64) Vector {
	v := make(Vector, len(values))
	copy(v, values)
	return v
}

// Clone this vector, returning a new Vector.
func (self Vector) Clone() Vector {
	return NewWithValues(self)
}

// Sets the values of this vector.
func (self Vector) Set(values []float64) {
	copy(self, values)
}

// Scale this vector (performs scalar multiplication) by the specified value.
func (self Vector) Scale(value float64) {
	length := len(self)
	for i := 0; i < length; i++ {
		self[i] *= value
	}
}

// Returns the magnitude of this vector.
func (self Vector) Magnitude() float64 {
	result := 0.0
	for _, e := range self {
		result += e * e
	}
	return math.Sqrt(result)
}

// Zeroes this vector
func (self Vector) Zero() {
	for i, _ := range self {
		self[i] = 0.0
	}
}

// Iterates through the elements of this vector and for each element invokes
// the function.
func (self Vector) Do(applyFn func(float64) float64) {
	for i, e := range self {
		self[i] = applyFn(e)
	}
}

// Iterates through the elements of this vector and for each element invokes
// the function with index.
func (self Vector) DoWithIndex(applyFn func(int, float64) float64) {
	for i, e := range self {
		self[i] = applyFn(i, e)
	}
}

// Sums of two vectors, returns the resulting vector.
func Add(v1, v2 Vector) Vector {
	length := min(len(v1), len(v2))
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] + v2[i]
	}
	return result
}

// Difference of two vectors, returns the resulting vector.
func Subtract(v1, v2 Vector) Vector {
	length := min(len(v1), len(v2))
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] - v2[i]
	}
	return result
}

// Dot products of two vectors.
func Dot(v1, v2 Vector) (float64, error) {
	if len(v1) != len(v2) {
		return 0.0, ErrVectorNotSameSize
	}

	length := len(v1)
	result := 0.0
	for i := 0; i < length; i++ {
		result += v1[i] * v2[i]
	}

	return result, nil
}

// Cross products of two vectors.
// Vector dimensionality has to be 3.
func Cross(v1, v2 Vector) (Vector, error) {
	// Early error check to prevent redundant cloning
	if len(v1) != 3 || len(v2) != 3 {
		return nil, ErrVectorInvalidDimension
	}

	result := make(Vector, 3)
	result[0] = v1[1]*v2[2] - v1[2]*v2[1]
	result[1] = v1[2]*v2[0] - v1[0]*v2[2]
	result[2] = v1[0]*v2[1] - v1[1]*v2[0]

	return result, nil
}

func Unit(v Vector) Vector {
	magRec := 1.0 / v.Magnitude()
	unit := v.Clone()
	for i, _ := range unit {
		unit[i] *= magRec
	}
	return unit
}

func Hadamard(v1, v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return nil, ErrVectorInvalidDimension
	}

	length := len(v1)
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] * v2[i]
	}
	return result, nil
}
