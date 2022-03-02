package account

import "sync"

// Define the Account type here.

type Account struct {
	sync.Mutex
	closed  bool
	balance int64
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{closed: false, balance: amount}
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.closed {
		return balance, ok
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed || a.balance+amount < 0 {
		return balance, ok
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return payout, ok
	}

	payout = a.balance
	a.balance -= payout
	a.closed = true
	return payout, true
}
