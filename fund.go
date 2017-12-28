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

// Balance is a member that renurns balance of a fund
func (f *Fund) Balance() int {
	return f.balance
}

// Withdraw excludes given amount from the funds balance
func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
