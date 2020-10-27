# Vector

[![Build Status](https://travis-ci.org/atedja/go-vector.svg?branch=master)](https://travis-ci.org/atedja/go-vector)

Simple N-dimensional Vector math library. Uses standard `[]float64` type.

### Quick Example

```go
package main

import (
  "fmt"
  "github.com/atedja/go-vector"
)

func main() {
  v1 := vector.New(4)
  v2 := vector.NewWithValues([]float64{0.0, 1.0, 2.0, 3.0})

  result := v1.Add(v2)
}
```

### Basic Usage

#### Creating Vectors

    var v vector.Vector
    v = vector.New(3)
    v = vector.NewWithValues([]float64{0.0, 1.0, 2.0})

#### Scale Vectors

    v.Scale(2.0)

#### Dot and Cross Products

    v1 := vector.NewWithValues([]float64{0.0, 1.0, 2.0})
    v2 := vector.NewWithValues([]float64{2.0, -1.0, 4.0})
    cross, _ := v1.Cross(v2)
    dot, _ := v1.Dot(v2)

#### And more!

`Add`, `Subtract`, `Scale`, `Resize`, `Zero`, `Magnitude`, `Unit`, `Hadamard`
