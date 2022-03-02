package account

import "sync"

// Define the Account type here.

type Account struct {
	mu      sync.Mutex
	open    bool
	balance int64
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{open: true, balance: amount}
}

func (a *Account) Balance() (int64, bool) {
	if !a.open {
		return 0, false
	}

	if a.balance >= 0 {
		return a.balance, true
	} else {
		return 0, false
	}
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open || a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}

	payout := a.balance
	a.balance -= payout
	a.open = false
	return payout, true
}
