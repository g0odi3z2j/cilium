// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/internal/asm/f32"
	"gonum.org/v1/gonum/internal/asm/f64"
)

// TODO(Kunde21):  Merge these methods back into level2double/level2single when Sgemv assembly kernels are merged into f32.

// Dgemv computes
//  y = alpha * A * x + beta * y    if tA = blas.NoTrans
//  y = alpha * A^T * x + beta * y  if tA = blas.Trans or blas.ConjTrans
// where A is an m×n dense matrix, x and y are vectors, and alpha and beta are scalars.
func (Implementation) Dgemv(tA blas.Transpose, m, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	if tA != blas.NoTrans && tA != blas.Trans && tA != blas.ConjTrans {
		panic(badTranspose)
	}
	if m < 0 {
		panic(mLT0)
	}
	if n < 0 {
		panic(nLT0)
	}
	if lda < max(1, n) {
		panic(badLdA)
	}
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	// Set up indexes
	lenX := m
	lenY := n
	if tA == blas.NoTrans {
		lenX = n
		lenY = m
	}

	// Quick return if possible
	if m == 0 || n == 0 {
		return
	}

	if (incX > 0 && (lenX-1)*incX >= len(x)) || (incX < 0 && (1-lenX)*incX >= len(x)) {
		panic(shortX)
	}
	if (incY > 0 && (lenY-1)*incY >= len(y)) || (incY < 0 && (1-lenY)*incY >= len(y)) {
		panic(shortY)
	}
	if len(a) < lda*(m-1)+n {
		panic(shortA)
	}

	// Quick return if possible
	if alpha == 0 && beta == 1 {
		return
	}

	if alpha == 0 {
		// First form y = beta * y
		if incY > 0 {
			Implementation{}.Dscal(lenY, beta, y, incY)
		} else {
			Implementation{}.Dscal(lenY, beta, y, -incY)
		}
		return
	}

	// Form y = alpha * A * x + y
	if tA == blas.NoTrans {
		f64.GemvN(uintptr(m), uintptr(n), alpha, a, uintptr(lda), x, uintptr(incX), beta, y, uintptr(incY))
		return
	}
	// Cases where a is transposed.
	f64.GemvT(uintptr(m), uintptr(n), alpha, a, uintptr(lda), x, uintptr(incX), beta, y, uintptr(incY))
}

// Sgemv computes
//  y = alpha * A * x + beta * y    if tA = blas.NoTrans
//  y = alpha * A^T * x + beta * y  if tA = blas.Trans or blas.ConjTrans
// where A is an m×n dense matrix, x and y are vectors, and alpha and beta are scalars.
//
// Float32 implementations are autogenerated and not directly tested.
func (Implementation) Sgemv(tA blas.Transpose, m, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	if tA != blas.NoTrans && tA != blas.Trans && tA != blas.ConjTrans {
		panic(badTranspose)
	}
	if m < 0 {
		panic(mLT0)
	}
	if n < 0 {
		panic(nLT0)
	}
	if lda < max(1, n) {
		panic(badLdA)
	}
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}

	// Quick return if possible.
	if m == 0 || n == 0 {
		return
	}

	// Set up indexes
	lenX := m
	lenY := n
	if tA == blas.NoTrans {
		lenX = n
		lenY = m
	}
	if (incX > 0 && (lenX-1)*incX >= len(x)) || (incX < 0 && (1-lenX)*incX >= len(x)) {
		panic(shortX)
	}
	if (incY > 0 && (lenY-1)*incY >= len(y)) || (incY < 0 && (1-lenY)*incY >= len(y)) {
		panic(shortY)
	}
	if len(a) < lda*(m-1)+n {
		panic(shortA)
	}

	// Quick return if possible.
	if alpha == 0 && beta == 1 {
		return
	}

	// First form y = beta * y
	if incY > 0 {
		Implementation{}.Sscal(lenY, beta, y, incY)
	} else {
		Implementation{}.Sscal(lenY, beta, y, -incY)
	}

	if alpha == 0 {
		return
	}

	var kx, ky int
	if incX < 0 {
		kx = -(lenX - 1) * incX
	}
	if incY < 0 {
		ky = -(lenY - 1) * incY
	}

	// Form y = alpha * A * x + y
	if tA == blas.NoTrans {
		if incX == 1 && incY == 1 {
			for i := 0; i < m; i++ {
				y[i] += alpha * f32.DotUnitary(a[lda*i:lda*i+n], x[:n])
			}
			return
		}
		iy := ky
		for i := 0; i < m; i++ {
			y[iy] += alpha * f32.DotInc(x, a[lda*i:lda*i+n], uintptr(n), uintptr(incX), 1, uintptr(kx), 0)
			iy += incY
		}
		return
	}
	// Cases where a is transposed.
	if incX == 1 && incY == 1 {
		for i := 0; i < m; i++ {
			tmp := alpha * x[i]
			if tmp != 0 {
				f32.AxpyUnitaryTo(y, tmp, a[lda*i:lda*i+n], y[:n])
			}
		}
		return
	}
	ix := kx
	for i := 0; i < m; i++ {
		tmp := alpha * x[ix]
		if tmp != 0 {
			f32.AxpyInc(tmp, a[lda*i:lda*i+n], y, uintptr(n), 1, uintptr(incY), 0, uintptr(ky))
		}
		ix += incX
	}
}
