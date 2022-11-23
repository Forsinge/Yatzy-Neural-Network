package matrix

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	data       []float64
	rows       int
	cols       int
	indexer    func(m *Matrix, row, col int) int
	transposed bool
}

func New(rows, cols, len int) *Matrix {
	if len != rows*cols {
		panic("NEW: mismatching matrix dimensions")
	}

	return &Matrix{
		data:       make([]float64, len),
		rows:       rows,
		cols:       cols,
		indexer:    index,
		transposed: false,
	}
}

func NewFrom(rows, cols int, data []float64) *Matrix {
	if len(data) != rows*cols {
		panic("NEW: mismatching matrix dimensions")
	}

	return &Matrix{
		data:       data,
		rows:       rows,
		cols:       cols,
		indexer:    index,
		transposed: false,
	}
}

func NewRand(rows, cols int) *Matrix {
	data := make([]float64, rows*cols)
	for i := range data {
		data[i] = rand.Float64() - 0.5
	}

	return &Matrix{
		data:       data,
		rows:       rows,
		cols:       cols,
		indexer:    index,
		transposed: false,
	}
}

func index(m *Matrix, row, col int) int {
	return row*m.cols + col
}

func index_t(m *Matrix, row, col int) int {
	return col*m.rows + row
}

func (m *Matrix) Get(row, col int) float64 {
	return m.data[m.indexer(m, row, col)]
}

func (m *Matrix) Set(row, col int, val float64) {
	m.data[m.indexer(m, row, col)] = val
}

func (m *Matrix) Dims() (int, int) {
	return m.rows, m.cols
}

func (m *Matrix) MaxDim() int {
	if m.rows > m.cols {
		return m.rows
	} else {
		return m.cols
	}
}

func (m *Matrix) Print() {
	for row := 0; row < m.rows; row += 1 {
		for col := 0; col < m.cols; col += 1 {
			fmt.Print(m.Get(row, col), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *Matrix) Data() []float64 {
	return m.data
}

func Copy(m *Matrix) *Matrix {
	data := make([]float64, len(m.data))
	copy(data, m.data)
	return &Matrix{
		data:       data,
		rows:       m.rows,
		cols:       m.cols,
		indexer:    m.indexer,
		transposed: m.transposed,
	}
}

func (m *Matrix) Transpose() {
	m.transposed = !m.transposed
	rows := m.rows
	m.rows = m.cols
	m.cols = rows

	if m.transposed {
		m.indexer = index_t
	} else {
		m.indexer = index
	}
}

// Performs computations in-place
func (m *Matrix) Apply(fn func(val float64) float64) {
	for i := range m.data {
		m.data[i] = fn(m.data[i])
	}
}

func (m *Matrix) Sub(other *Matrix) {
	if m.rows != other.rows || m.cols != other.cols {
		panic("SUB: mismatching matrix dimensions")
	}
	for i := range m.data {
		m.data[i] = m.data[i] - other.data[i]
	}
}

// Performs computations on a new matrix and returns it
func Apply(m *Matrix, fn func(val float64) float64) *Matrix {
	result := Copy(m)
	for i := range result.data {
		result.data[i] = fn(result.data[i])
	}
	return result
}

func Mul(m *Matrix, n *Matrix) *Matrix {
	if m.cols != n.rows {
		panic("MUL: mismatching matrix dimensions")
	}

	rows := m.rows
	cols := n.cols
	result := Matrix{
		data:       make([]float64, rows*cols),
		rows:       rows,
		cols:       cols,
		indexer:    index,
		transposed: false,
	}

	for row := 0; row < rows; row += 1 {
		for col := 0; col < cols; col += 1 {
			dot := 0.0
			for i := 0; i < m.cols; i += 1 {
				dot += m.Get(row, i) * n.Get(i, col)
			}
			result.Set(row, col, dot)
		}
	}

	return &result
}

func ElemOp(m *Matrix, n *Matrix, fn func(a float64, b float64) float64) *Matrix {
	if m.rows != n.rows || m.cols != n.cols {
		panic("MULELEM: mismatching matrix dimensions")
	}

	rows := m.rows
	cols := m.cols

	result := Matrix{
		data:       make([]float64, rows*cols),
		rows:       rows,
		cols:       cols,
		indexer:    index,
		transposed: false,
	}

	for row := 0; row < rows; row += 1 {
		for col := 0; col < cols; col += 1 {
			val := fn(m.Get(row, col), n.Get(row, col))
			result.Set(row, col, val)
		}
	}

	return &result
}

func ElemMul(m *Matrix, n *Matrix) *Matrix {
	fn := func(a, b float64) float64 { return a * b }
	return ElemOp(m, n, fn)
}

func ElemSub(m *Matrix, n *Matrix) *Matrix {
	fn := func(a, b float64) float64 { return a - b }
	return ElemOp(m, n, fn)
}

func ElemAdd(m *Matrix, n *Matrix) *Matrix {
	fn := func(a, b float64) float64 { return a + b }
	return ElemOp(m, n, fn)
}

func (m *Matrix) Sum() float64 {
	sum := 0.0
	for i := range m.data {
		sum += m.data[i]
	}
	return sum
}
