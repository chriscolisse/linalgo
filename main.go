package main

// import (
// 	// "encoding/csv"
// 	"fmt"
// 	// "os"
// 	// "strconv"
// 	// "strings"

// 	// "reflect"
// 	// "time"

// 	"github.com/chriscolisse/linalgo/matrix"
// )

// func main() {
// 	data := [][]float64{
// 		{1, 2, 3},
// 		{2, 3, 4},
// 		{3, 6, 5},
// 	}
// 	matA, _ := matrix.NewMatrix(data)
// 	aTrans, _ := matA.Transpose()
// 	rotA, _ := matA.Rotate(10)
// 	fmt.Println(matA.Columns())
// 	fmt.Println(aTrans.Transpose())
// 	fmt.Println(rotA.Rows())
// 	// if len(os.Args) < 2 {
// 	// 	fmt.Println("Expected 'add' or 'mul' subcommands")
// 	// 	//os.Exit(1)
// 	// }
// 	// start := time.Now()
// 	// switch os.Args[1] {
// 	// case "add":
// 	// 	addCommand(os.Args[2:])
// 	// case "mul":
// 	// 	mulCommand(os.Args[2:])
// 	// case "transpose":
// 	// 	transposeCommand(os.Args[2:])
// 	// default:
// 	// 	fmt.Println("Expected 'add' or 'mul' subcommands")
// 	// 	os.Exit(1)
// 	// }
// 	// duration := time.Since(start)
// 	// fmt.Printf("function executed in %v", duration)
// }

// // func addCommand(args []string) {
// // 	if len(args) != 2 {
// // 		fmt.Printf("Usage: add <matrix1> <matrix2>", "len args recieved: %v", len(args))
// // 		os.Exit(1)
// // 	}

// // 	matrix1, _ := parseMatrix(args[0])
// // 	matrix2, _ := parseMatrix(args[1])

// // 	_, err := matrix.Matrix_Addition(matrix1, matrix2)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		os.Exit(1)
// // 	}

// // 	// fmt.Println("Result of addition:")
// // 	// fmt.Println(result.Rows())
// // }

// // func mulCommand(args []string) {
// // 	if len(args) != 2 {
// // 		fmt.Println("Usage: mul <matrix1> <matrix2>")
// // 		os.Exit(1)
// // 	}

// // 	matrix1, _ := parseMatrix(args[0])
// // 	matrix2, _ := parseMatrix(args[1])
// // 	// fmt.Println(matrix1.Rowsize(), matrix1.Rows(), matrix2.Rowsize(), matrix2.Rows())

// // 	result, err := matrix.Matrix_Multiply(matrix1, matrix2)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		os.Exit(1)
// // 	}

// // 	fmt.Println("Result of multiplication:")
// // 	fmt.Println(result.Rows(), result.Columns())
// // }

// // func transposeCommand(args []string) {
// // 	mat, err := parseMatrix(args[0])
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		os.Exit(1)
// // 	}
// // 	res, _ := mat.Transpose()
// // 	fmt.Println("Result of transpose:")
// // 	fmt.Println(mat.Columns())
// // 	fmt.Println("transpose function result: ", res.Rows(), reflect.DeepEqual(mat.Columns(), res.Columns()))

// // }

// // func parseMatrix(input string) (matrix.Matrix, error) {
// // 	if strings.HasSuffix(input, ".csv") {
// // 		return readMatrixFromCSV(input)
// // 	} else {
// // 		return readMatrixFromString(input)
// // 	}
// // }

// // func readMatrixFromCSV(filename string) (matrix.Matrix, error) {
// // 	file, err := os.Open(filename)
// // 	if err != nil {
// // 		fmt.Println("Error opening file:", err)
// // 		os.Exit(1)
// // 	}
// // 	defer file.Close()

// // 	reader := csv.NewReader(file)
// // 	records, err := reader.ReadAll()
// // 	if err != nil {
// // 		fmt.Println("Error reading CSV:", err)
// // 		os.Exit(1)
// // 	}

// // 	data := make([][]float64, len(records))
// // 	for i, row := range records {
// // 		data[i] = make([]float64, len(row))
// // 		for j, value := range row {
// // 			data[i][j], err = strconv.ParseFloat(value, 64)
// // 			if err != nil {
// // 				fmt.Println("Error parsing value:", err)
// // 				os.Exit(1)
// // 			}
// // 		}
// // 	}

// // 	return matrix.NewMatrix(data)
// // }

// // func readMatrixFromString(s string) (matrix.Matrix, error) {
// // 	rows := strings.Split(s, ";")
// // 	data := make([][]float64, len(rows))

// // 	for i, row := range rows {
// // 		cols := strings.Split(row, ",")
// // 		data[i] = make([]float64, len(cols))
// // 		for j, col := range cols {
// // 			val, err := strconv.ParseFloat(col, 64)
// // 			if err != nil {
// // 				return matrix.Matrix{}, err
// // 			}
// // 			data[i][j] = val
// // 		}
// // 	}

// // 	return matrix.NewMatrix(data)
// // }

import (
	"fmt"

	"github.com/chriscolisse/linalgo/matrix"
)

func main() {
	m, _ := matrix.NewMatrix([][]float64{
		{0, 1, 2, 3, 10},
		{0, 0, 0, 4, 5},
		{1, 0, 3, 4, 6},
		{0, 2, 0, 0, 7},
		{3, 1, 4, 0, 2},
	})
	n := m
	m_inv, _ := m.Gauss_Jordan_reduction_in_place()
	// m.Row_reduce_in_place()
	fmt.Println(m.Rows())
	fmt.Println(m_inv.Rows())

	validate, _ := matrix.Matrix_Multiply(m_inv, n)

	fmt.Println("\n\n\n\n\n", validate.Rows())
}
