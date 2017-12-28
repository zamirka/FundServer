package funding

// FundServer servers all requests to the Fund
type FundServer struct {
	commands chan TransactionCommand
	fund     *Fund
}

// NewFundServer creates a new server and imedeately initializes it
func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make creates builtins like channels, maps and slices
		commands: make(chan TransactionCommand),
		fund:     NewFund(initialBalance),
	}

	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) loop() {
	for transaction := range s.commands {
		//Now we don't need any type-switch mess
		transaction.Transactor(s.fund)
		transaction.Done <- true
	}
}

// Transactor is a callback function which really executes some commands over fund in TransactionCommand
type Transactor func(fund *Fund)

// TransactionCommand is a type of command which can be passed to the server to do some stuff with Fund
type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

// Balance is a member that returns balance of a fund
func (s *FundServer) Balance() int {
	var balance int
	s.Transact(func(f *Fund) {
		balance = f.Balance()
	})
	return balance
}

// Withdraw excludes given amount from the funds balance
func (s *FundServer) Withdraw(amount int) {
	s.Transact(func(f *Fund) {
		f.Withdraw(amount)
	})
}

// Transact is a function that sends a transactionCommand to the server
func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands <- command
	<-command.Done
}
