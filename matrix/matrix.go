package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	rowsize int
	colsize int
	rows    [][]float64
	columns [][]float64
}

func NewMatrix(data [][]float64) (Matrix, error) {
	var rowsize int = len(data)
	var colsize int = len(data[0])
	// _rows := [][]float64{}
	_rows := make([][]float64, 0, rowsize)
	_columns := make([][]float64, colsize)

	for _, row := range data {
		if len(row) != colsize {
			return Matrix{}, errors.New("column lengths do not match, invalid matrix")
		}
		_rows = append(_rows, row)

		for i, value := range row {
			_columns[i] = append(_columns[i], value)

		}
	}

	return Matrix{rowsize: rowsize, colsize: colsize, rows: _rows, columns: _columns}, nil
}

func (m *Matrix) Rows() [][]float64 {
	return m.rows
}
func (m *Matrix) Columns() [][]float64 {
	return m.columns
}
func (m *Matrix) Rowsize() int {
	return m.rowsize
}
func (m *Matrix) Colsize() int {
	return m.colsize
}

func (m *Matrix) Update(rowidx, colidx int, newval float64) {
	m.rows[rowidx][colidx] = newval
}

func (m *Matrix) Dimensions() (int, int) {
	return m.rowsize, m.colsize
}

func (m *Matrix) Is_Square() bool {
	return m.colsize == m.rowsize
}

func (m *Matrix) Transpose_In_Place() {
	if m.rowsize == m.colsize {
		for i := 0; i < m.rowsize; i++ {
			for j := i + 1; j < m.colsize; j++ {
				m.rows[i][j], m.rows[j][i] = m.rows[j][i], m.rows[i][j]
			}
		}
	} else {

		t := make([][]float64, m.colsize)
		for i := range t {
			t[i] = make([]float64, m.rowsize)
		}

		for i, row := range m.rows {
			for j, val := range row {
				t[j][i] = val
			}
		}
		m.rows = t
		m.rowsize, m.colsize = m.colsize, m.rowsize
	}

	m.columns = make([][]float64, m.colsize)
	for _, row := range m.rows {
		for i, value := range row {
			m.columns[i] = append(m.columns[i], value)

		}
	}

}

func (m *Matrix) Transpose() (Matrix, error) {

	t := make([][]float64, m.colsize)
	for i := 0; i < m.colsize; i++ {
		t[i] = make([]float64, m.rowsize)
	}

	for i, row := range m.rows {
		for j, val := range row {
			t[j][i] = val
		}
	}
	return NewMatrix(t)
}

// get_diagonal returns the diagonal of a square matrix.
// Returns the main diagonal if the main_diagonal parameter is set to true, and the secondary diagonal if set to false. Defaults to true.
func (m *Matrix) Get_Diagonal(main_diagonal ...bool) ([]float64, error) {
	is_main := true
	if len(main_diagonal) > 0 {
		is_main = main_diagonal[0]
	}

	if !m.Is_Square() {
		return nil, fmt.Errorf("Matrix is not square, cannot get diagonal")
	}
	output := make([]float64, m.rowsize)
	for row := range m.rows {
		if is_main {
			output[row] = m.rows[row][row]
		} else {
			output[row] = m.rows[row][m.colsize-1-row]
		}
	}
	return output, nil
}

// matrix_multiply takes in n number of matrices and performs the multiplication if they are compatible. Returns the product of those multiplications.
func Matrix_Multiply(matrices ...Matrix) (Matrix, error) {
	if len(matrices) == 0 {
		return Matrix{}, fmt.Errorf("no matrices to multiply")
	}
	if len(matrices) == 1 {
		return matrices[0], nil
	}
	m, err := simple_matrix_multiply(&matrices[0], &matrices[1])
	if err != nil {
		return Matrix{}, err
	}
	if len(matrices[2:]) > 0 {
		matrices = append([]Matrix{m}, matrices[2:]...)
		return Matrix_Multiply(matrices...)
	}

	return m, nil
}

// simple_matrix_multiple multiplies two matrices together if they are compatible and returns the product.
func simple_matrix_multiply(matA *Matrix, matB *Matrix) (Matrix, error) {
	if matA.colsize != matB.rowsize {
		return Matrix{}, errors.New("cannot multiply matrices, incompatible row and columns sizes")
	}
	product_matrix := make([][]float64, matA.rowsize)
	for i, row := range matA.rows {
		product_matrix[i] = make([]float64, matB.colsize)
		for j := range matB.columns {
			for n, r := range row {
				product_matrix[i][j] += r * matB.columns[j][n]
			}
		}
	}

	return NewMatrix(product_matrix)
}

func simple_matrix_addition(matA *Matrix, matB *Matrix) (Matrix, error) {
	rowsA, colsA := matA.Dimensions()
	rowsB, colsB := matB.Dimensions()
	if rowsA != rowsB || colsA != colsB {
		return Matrix{}, fmt.Errorf("matrices do not have the same dimensions, cannot perform addition")
	}
	output_matrix := make([][]float64, rowsA)
	for i, row := range matA.rows {
		output_matrix[i] = make([]float64, colsA)
		for j, aval := range row {
			output_matrix[i][j] = aval + matB.rows[i][j]
		}
	}
	return NewMatrix(output_matrix)
}

func Matrix_Addition(matrices ...Matrix) (Matrix, error) {
	if len(matrices) == 0 {
		return Matrix{}, fmt.Errorf("no matrices input as arguments to the function. Cannot perform addition")
	}
	if len(matrices) == 1 {
		return matrices[0], nil
	}
	m, err := simple_matrix_addition(&matrices[0], &matrices[1])
	if err != nil {
		return Matrix{}, err
	}
	if len(matrices[2:]) > 1 {
		return Matrix_Addition(append([]Matrix{m}, matrices[2:]...)...)
	}
	return m, nil
}
