// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

// This list is duplicated in netlib/lapack/netlib. Keep in sync.
const (
	// Panic strings for bad enumeration values.
	badApplyOrtho      = "lapack: bad ApplyOrtho"
	badBalanceJob      = "lapack: bad BalanceJob"
	badDiag            = "lapack: bad Diag"
	badDirect          = "lapack: bad Direct"
	badEVComp          = "lapack: bad EVComp"
	badEVHowMany       = "lapack: bad EVHowMany"
	badEVJob           = "lapack: bad EVJob"
	badEVSide          = "lapack: bad EVSide"
	badGSVDJob         = "lapack: bad GSVDJob"
	badGenOrtho        = "lapack: bad GenOrtho"
	badLeftEVJob       = "lapack: bad LeftEVJob"
	badMatrixType      = "lapack: bad MatrixType"
	badNorm            = "lapack: bad Norm"
	badPivot           = "lapack: bad Pivot"
	badRightEVJob      = "lapack: bad RightEVJob"
	badSVDJob          = "lapack: bad SVDJob"
	badSchurComp       = "lapack: bad SchurComp"
	badSchurJob        = "lapack: bad SchurJob"
	badSide            = "lapack: bad Side"
	badSort            = "lapack: bad Sort"
	badStoreV          = "lapack: bad StoreV"
	badTrans           = "lapack: bad Trans"
	badUpdateSchurComp = "lapack: bad UpdateSchurComp"
	badUplo            = "lapack: bad Uplo"
	bothSVDOver        = "lapack: both jobU and jobVT are lapack.SVDOverwrite"

	// Panic strings for bad numerical and string values.
	badIfst     = "lapack: ifst out of range"
	badIhi      = "lapack: ihi out of range"
	badIhiz     = "lapack: ihiz out of range"
	badIlo      = "lapack: ilo out of range"
	badIloz     = "lapack: iloz out of range"
	badIlst     = "lapack: ilst out of range"
	badIsave    = "lapack: bad isave value"
	badIspec    = "lapack: bad ispec value"
	badJ1       = "lapack: j1 out of range"
	badJpvt     = "lapack: bad element of jpvt"
	badK1       = "lapack: k1 out of range"
	badK2       = "lapack: k2 out of range"
	badKacc22   = "lapack: invalid value of kacc22"
	badKbot     = "lapack: kbot out of range"
	badKtop     = "lapack: ktop out of range"
	badLWork    = "lapack: insufficient declared workspace length"
	badMm       = "lapack: mm out of range"
	badN1       = "lapack: bad value of n1"
	badN2       = "lapack: bad value of n2"
	badNa       = "lapack: bad value of na"
	badName     = "lapack: bad name"
	badNh       = "lapack: bad value of nh"
	badNw       = "lapack: bad value of nw"
	badPp       = "lapack: bad value of pp"
	badShifts   = "lapack: bad shifts"
	i0LT0       = "lapack: i0 < 0"
	kGTM        = "lapack: k > m"
	kGTN        = "lapack: k > n"
	kLT0        = "lapack: k < 0"
	kLT1        = "lapack: k < 1"
	kdLT0       = "lapack: kd < 0"
	mGTN        = "lapack: m > n"
	mLT0        = "lapack: m < 0"
	mmLT0       = "lapack: mm < 0"
	n0LT0       = "lapack: n0 < 0"
	nGTM        = "lapack: n > m"
	nLT0        = "lapack: n < 0"
	nLT1        = "lapack: n < 1"
	nLTM        = "lapack: n < m"
	nanCFrom    = "lapack: cfrom is NaN"
	nanCTo      = "lapack: cto is NaN"
	nbGTM       = "lapack: nb > m"
	nbGTN       = "lapack: nb > n"
	nbLT0       = "lapack: nb < 0"
	nccLT0      = "lapack: ncc < 0"
	ncvtLT0     = "lapack: ncvt < 0"
	negANorm    = "lapack: anorm < 0"
	negZ        = "lapack: negative z value"
	nhLT0       = "lapack: nh < 0"
	notIsolated = "lapack: block is not isolated"
	nrhsLT0     = "lapack: nrhs < 0"
	nruLT0      = "lapack: nru < 0"
	nshftsLT0   = "lapack: nshfts < 0"
	nshftsOdd   = "lapack: nshfts must be even"
	nvLT0       = "lapack: nv < 0"
	offsetGTM   = "lapack: offset > m"
	offsetLT0   = "lapack: offset < 0"
	pLT0        = "lapack: p < 0"
	recurLT0    = "lapack: recur < 0"
	zeroCFrom   = "lapack: zero cfrom"

	// Panic strings for bad slice lengths.
	badLenAlpha    = "lapack: bad length of alpha"
	badLenBeta     = "lapack: bad length of beta"
	badLenIpiv     = "lapack: bad length of ipiv"
	badLenJpvt     = "lapack: bad length of jpvt"
	badLenK        = "lapack: bad length of k"
	badLenSelected = "lapack: bad length of selected"
	badLenSi       = "lapack: bad length of si"
	badLenSr       = "lapack: bad length of sr"
	badLenTau      = "lapack: bad length of tau"
	badLenWi       = "lapack: bad length of wi"
	badLenWr       = "lapack: bad length of wr"

	// Panic strings for insufficient slice lengths.
	shortA     = "lapack: insufficient length of a"
	shortAB    = "lapack: insufficient length of ab"
	shortAuxv  = "lapack: insufficient length of auxv"
	shortB     = "lapack: insufficient length of b"
	shortC     = "lapack: insufficient length of c"
	shortCNorm = "lapack: insufficient length of cnorm"
	shortD     = "lapack: insufficient length of d"
	shortE     = "lapack: insufficient length of e"
	shortF     = "lapack: insufficient length of f"
	shortH     = "lapack: insufficient length of h"
	shortIWork = "lapack: insufficient length of iwork"
	shortIsgn  = "lapack: insufficient length of isgn"
	shortQ     = "lapack: insufficient length of q"
	shortS     = "lapack: insufficient length of s"
	shortScale = "lapack: insufficient length of scale"
	shortT     = "lapack: insufficient length of t"
	shortTau   = "lapack: insufficient length of tau"
	shortTauP  = "lapack: insufficient length of tauP"
	shortTauQ  = "lapack: insufficient length of tauQ"
	shortU     = "lapack: insufficient length of u"
	shortV     = "lapack: insufficient length of v"
	shortVL    = "lapack: insufficient length of vl"
	shortVR    = "lapack: insufficient length of vr"
	shortVT    = "lapack: insufficient length of vt"
	shortVn1   = "lapack: insufficient length of vn1"
	shortVn2   = "lapack: insufficient length of vn2"
	shortW     = "lapack: insufficient length of w"
	shortWH    = "lapack: insufficient length of wh"
	shortWV    = "lapack: insufficient length of wv"
	shortWi    = "lapack: insufficient length of wi"
	shortWork  = "lapack: insufficient length of work"
	shortWr    = "lapack: insufficient length of wr"
	shortX     = "lapack: insufficient length of x"
	shortY     = "lapack: insufficient length of y"
	shortZ     = "lapack: insufficient length of z"

	// Panic strings for bad leading dimensions of matrices.
	badLdA    = "lapack: bad leading dimension of A"
	badLdB    = "lapack: bad leading dimension of B"
	badLdC    = "lapack: bad leading dimension of C"
	badLdF    = "lapack: bad leading dimension of F"
	badLdH    = "lapack: bad leading dimension of H"
	badLdQ    = "lapack: bad leading dimension of Q"
	badLdT    = "lapack: bad leading dimension of T"
	badLdU    = "lapack: bad leading dimension of U"
	badLdV    = "lapack: bad leading dimension of V"
	badLdVL   = "lapack: bad leading dimension of VL"
	badLdVR   = "lapack: bad leading dimension of VR"
	badLdVT   = "lapack: bad leading dimension of VT"
	badLdW    = "lapack: bad leading dimension of W"
	badLdWH   = "lapack: bad leading dimension of WH"
	badLdWV   = "lapack: bad leading dimension of WV"
	badLdWork = "lapack: bad leading dimension of Work"
	badLdX    = "lapack: bad leading dimension of X"
	badLdY    = "lapack: bad leading dimension of Y"
	badLdZ    = "lapack: bad leading dimension of Z"

	// Panic strings for bad vector increments.
	absIncNotOne = "lapack: increment not one or negative one"
	badIncX      = "lapack: incX <= 0"
	badIncY      = "lapack: incY <= 0"
	zeroIncV     = "lapack: incv == 0"
)
