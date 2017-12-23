package funding

// Fund stores some amount inside and allows operations on it
type Fund struct {
	balance int
}

// NewFund returns a pointer to a new Fund instance
func NewFund(initialBalance int) *Fund {
	return &Fund{
		balance: initialBalance,
	}
}
