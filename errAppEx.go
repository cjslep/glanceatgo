// Computes the eigenvalue factorization of a Hermitian matrix.
func EigHerm(a Const) (*Mat, []float64, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, nil, err
	}
	if err := errNonHerm(a); err != nil {
		return nil, nil, err
	}
	return eigHerm(cloneMat(a), DefaultTri)
}