package examples

const (
	maxBalance = 1000
	minBalance = 0
)

// BankAccount represents a bank account
// @invariant [balance of open accounts is always at least the minBalance] /
// !BankAccount.closed ==> BankAccount.balance >= minBalance
// @invariant [balance of open accounts is always at most the maxBalance] /
// !BankAccount.closed ==> BankAccount.balance <= maxBalance
// @invariant [balance of closed accounts is 0] BankAccount.closed ==> BankAccount.balance == 0
type BankAccount struct {
	balance int  // the balance of the account
	closed  bool // is the account closed?
}

// NewBankAccount creates an account with the given initial balance
// @requires initialBalance >= minBalance && initialBalance <= maxBalance
// @ensures account.balance == initialBalance
// @ensures !account.closed
func NewBankAccount(initialBalance int) (account BankAccount) {
	// ... implementation ...
}

// Credit the given amount to the account
// @requires [can not credit a closed account] !a.closed
// @requires amount > 0 && (a.balance + amount) <= maxBalance
// @let initialBalance := a.balance
// @ensures a.balance == initialBalance + amount
// Ensure other fields are unchanged:
// @unmodified a.closed
func (a *BankAccount) Credit(amount int) {
	// ... implementation ...
}

// Debit the given amount from the account
// @requires !a.closed
// @requires amount > 0 && (a.balance - amount) >= minBalance
// @let initialBalance := a.balance
// @ensures a.balance == initialBalance - amount
// Ensure other fields are unchanged:
// @unmodified a.closed
func (a *BankAccount) Debit(amount int) {
	// ... implementation ...
}

// Close the account and returns its payout
// @requires !a.closed
// @ensures payout == @old{a.balance}
// @ensures a.closed
// @ensures a.balance == 0
func (a *BankAccount) Close() (payout int) {
	// ... implementation ...
}

// Balance yields the balance of the account
// @ensures balance == a.balance
// Ensure fields are unchanged:
// @unmodified a.balance
// @unmodified a.closed
func (a *BankAccount) Balance() (balance int) {
	// ... implementation ...
}
