package main

import (
	"fmt"

	"github.com/chriscolisse/linalgo/matrix"
)

func main() {
	matA, _ := matrix.NewMatrix([][]float64{
		{1, 0},
		{0, 1},
		{3, 1},
	})

	matA.Transpose_In_Place()

	matA.Update(1, 1, 2.5)

	fmt.Println(matA.Rows())

}
