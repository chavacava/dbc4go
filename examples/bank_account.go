package examples

const (
	maxBalance = 1000
	minBalance = 0
)

type BankAccount struct {
	balance int  // the balance of the account
	closed  bool // is the account closed?
}

// NewBankAccount creates an account with the given initial balance
//@requires initialBalance >= 0 && initialBalance <= maxBalance
//@ensures account.balance == initialBalance
//@ensures !account.closed
func NewBankAccount(initialBalance int) (account BankAccount) {
	// ... implementation ...
}

// Credit the given amount to the account
//@requires !a.closed
//@requires amount > 0 && (a.balance + amount) <= maxBalance
//@ensures a.balance == @old(a.balance) + amount
// Ensure other fields are unchanged:
//@ensures !a.closed
func (a *BankAccount) Credit(amount int) {
	// ... implementation ...
}

// Debit the given amount from the account
//@requires !a.closed
//@requires amount > 0 && (a.balance - amount) >= minBalance
//@ensures a.balance == @old(a.balance) - amount
// Ensure other fields are unchanged:
//@ensures !a.closed
func (a *BankAccount) Debit(amount int) {
	// ... implementation ...
}

// Close the account and returns its payout
//@requires !a.closed
//@ensures payout == @old(a.balance)
//@ensures a.balance == 0
func (a *BankAccount) Close() (payout int) {
	// ... implementation ...
}

// Balance yields the balance of the account
//@ensures balance == a.balance
// Ensure other fields are unchanged:
//@ensures a.balance == @old(a.balance)
//@ensures a.closed == @old(a.closed)
func (a *BankAccount) Balance() (balance int) {
	// ... implementation ...
}
