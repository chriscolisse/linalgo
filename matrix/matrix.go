package matrix

import (
	"errors"
	"fmt"
	"math"
)

type Matrix struct {
	rowsize int
	colsize int
	rows    [][]float64
	columns [][]float64
}

// NewMatrix function. Constructor function for building a Matrix struct.
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

// Rows method on a Matrix struct. Returns the matrix's rows (use this to obtain the matrix itself)
func (m *Matrix) Rows() [][]float64 {
	return m.rows
}

// Columns method on a Matrix struct. Returns the matrix's columns
func (m *Matrix) Columns() [][]float64 {
	return m.columns
}

// Rowsize method on a Matrix struct. Returns the matrix's row size
func (m *Matrix) Rowsize() int {
	return m.rowsize
}

// Colsize method on a Matrix struct. Returns the matrix's column size
func (m *Matrix) Colsize() int {
	return m.colsize
}

// Update method on a Matrix struct.
// Provide row index (rowidx) and column index (colidx), as well as the value to insert into that position.
// Performs the update in place.
func (m *Matrix) Update(rowidx, colidx int, newval float64) {
	m.rows[rowidx][colidx] = newval
}

// Dimensions method on a Matrix struct. Returns (rowsize, colsize)
func (m *Matrix) Dimensions() (int, int) {
	return m.rowsize, m.colsize
}

// Is_Square method on a Matrix struct. Returns true if the matrix is square
func (m *Matrix) Is_Square() bool {
	return m.colsize == m.rowsize
}

// Transpose_In_Place method on a Matrix struct. Transposes the matrix in place.
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

// Transpose method on a Matrix struct. Returns the transposed matrix.
func (m *Matrix) Transpose() (Matrix, error) {
	return NewMatrix(m.columns)
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

func (m *Matrix) Rotate(theta float64) (Matrix, error) {
	center_x := float64((m.rowsize - 1) / 2)
	center_y := float64((m.colsize - 1) / 2)
	rotated_mat := make([][]float64, m.rowsize)
	for i := range rotated_mat {
		rotated_mat[i] = make([]float64, m.colsize)
	}
	// fmt.Println(rotated_mat)
	for i, row := range m.rows {
		for j, val := range row {
			var xy []float64 = []float64{float64(i) - center_x, float64(j) - center_y}
			new_x, new_y := rotate_xy(xy, theta)
			new_x, new_y = new_x+center_x, new_y+center_y
			if int(new_x) >= 0 && int(new_x) < m.rowsize && int(new_y) >= 0 && int(new_y) < m.colsize {
				rotated_mat[int(new_x)][int(new_y)] = val
			}
		}
	}
	return NewMatrix(rotated_mat)
}

// func interpolate_matpos() {

// }

func rotate_xy(xy []float64, theta float64) (float64, float64) {
	thetaRad := theta * (math.Pi / 180.0)
	cosTheta := math.Cos(thetaRad)
	sinTheta := math.Sin(thetaRad)
	new_x := float64(xy[0])*cosTheta - float64(xy[1])*sinTheta
	new_y := float64(xy[0])*sinTheta + float64(xy[1])*cosTheta
	return new_x, new_y
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

// simple_matrix_addition adds two matrices together if they are compatible and returns the sum.
func simple_matrix_addition(matA *Matrix, matB *Matrix) (Matrix, error) {
	if rowsA, colsA := matA.Dimensions(); rowsA != matB.rowsize || colsA != matB.colsize {
		return Matrix{}, fmt.Errorf("matrices do not have the same dimensions, cannot perform addition")
	}
	output_matrix := make([][]float64, matA.rowsize)
	for i, row := range matA.rows {
		output_matrix[i] = make([]float64, matA.colsize)
		for j, aval := range row {
			output_matrix[i][j] = aval + matB.rows[i][j]
		}
	}
	return NewMatrix(output_matrix)
}

// matrix_addition takes in n number of matrices and performs the additions if they are compatible. Returns the sum of those additions.
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

// Scalar_Multiply multiplies a matrix by a scalar value. Returns the product matrix.
func Scalar_Multiply(scalar float64, m *Matrix) (Matrix, error) {
	output := make([][]float64, m.rowsize)
	for i, row := range m.rows {
		output[i] = make([]float64, m.colsize)
		for j, val := range row {
			output[i][j] = val * scalar
		}
	}
	return NewMatrix(output)
}

// Performs Gaussian elimination on a matrix
func (m *Matrix) Row_reduce_in_place() (Matrix, error) {
	identity := make([][]float64, m.rowsize)
	return_ident := false
	if m.Is_Square() {
		return_ident = true
		for i := range m.rows {
			identity[i] = make([]float64, m.colsize)
			identity[i][i] = 1
		}
	}
	matrix_data := m.rows
	max_pivot := m.rowsize - 1
	order := func(mdata *[][]float64, level int, max_pivot int) {
		m := *mdata
		if level <= max_pivot {
			if m[level][level] == 0 {
				for i := level + 1; i < len(m); i++ {
					if m[i][level] != 0 {
						m[level], m[i] = m[i], m[level]
						if return_ident {
							identity[level], identity[i] = identity[i], identity[level]
						}
						break
					}
				}
			}
		}
	}

	subtract := func(mdata *[][]float64, level int, max_pivot int) {
		m := *mdata
		if level <= max_pivot {
			divisor := m[level][level]
			if divisor != 0 {
				for i := level + 1; i < len(m); i++ {
					if m[i][level] != 0 {
						factor := m[i][level] / divisor
						for n := range m[i] {
							m[i][n] -= factor * m[level][n]
							if return_ident {
								identity[i][n] -= factor * identity[level][n]
							}
						}
					}
				}
			}
		}

	}
	for i := 0; i < len(matrix_data); i++ {
		order(&matrix_data, i, max_pivot)
		subtract(&matrix_data, i, max_pivot)
	}
	*m, _ = NewMatrix(matrix_data)
	return NewMatrix(identity)
}

func (m *Matrix) Gauss_Jordan_reduction_in_place() (Matrix, error) {
	inverse, _ := m.Row_reduce_in_place()
	inverse_rows := inverse.rows
	is_square := m.Is_Square()
	mrows := m.rows
	rows, cols := len(mrows), len(mrows[0])

	normalize_pivots := func(mdata *[][]float64) {
		m := *mdata
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if m[i][j] != 0 {
					multiplier := 1 / m[i][j]
					for n := range m[i] {
						m[i][n] *= multiplier
						if is_square {
							inverse_rows[i][n] *= multiplier
						}
					}
					break
				}
			}
		}
	}

	backward_elimination := func(mdata *[][]float64) {
		m := *mdata
		for i := len(m) - 1; i >= 0; i-- {
			pivotCol := -1
			for j := 0; j < cols; j++ {
				if m[i][j] != 0 {
					pivotCol = j
					break
				}
			}
			if pivotCol == -1 {
				// No pivot in this row, skip backward elimination
				continue
			}
			for j := 0; j < i; j++ {
				factor := m[j][pivotCol]
				for n := range m[j] {
					m[j][n] -= factor * m[i][n]
					if is_square {
						inverse_rows[j][n] -= factor * inverse_rows[i][n]
					}
				}
			}
		}
	}

	normalize_pivots(&mrows)
	backward_elimination(&mrows)

	*m, _ = NewMatrix(mrows)
	return NewMatrix(inverse_rows)
}
