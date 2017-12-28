package funding

import (
	"fmt"
)

// FundServer servers all requests to the Fund
type FundServer struct {
	commands chan interface{}
	fund     *Fund
}

// NewFundServer creates a new server and imedeately initializes it
func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make creates builtins like channels, maps and slices
		commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) loop() {
	// the built-in "range" clause can interate over channels,
	// amongst other things
	for command := range s.commands {
		// Handle the command

		// command is just an interface{}, but we can check its real type
		switch command.(type) {

		case WithdrawCommand:
			// Add then use a "type assertion" to convert it
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}

// WithdrawCommand a type to pass a command to server to withdraw somw amout
type WithdrawCommand struct {
	Amount int
}

// BalanceCommand  a type of command that contains a channels through wich resposnse is returned from server
type BalanceCommand struct {
	Response chan int
}

// Balance is a member that returns balance of a fund
func (s *FundServer) Balance() int {
	responseChan := make(chan int)
	s.commands <- BalanceCommand{Response: responseChan}
	return <-responseChan
}

// Withdraw excludes given amount from the funds balance
func (s *FundServer) Withdraw(amount int) {
	s.commands <- WithdrawCommand{Amount: amount}
}
