package account

import "sync"

type Account struct {
	bal    int64
	closed bool
	sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{bal: amount}
}

func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()
	if acc.closed {
		return acc.bal, false
	}

	if acc.bal + amount < 0 {
		return acc.bal, false
	}

	acc.bal += amount
	return acc.bal, true
}

func (acc *Account) Balance() (balance int64, ok bool) {
	acc.RLock()
	defer acc.RUnlock()
	if acc.closed {
		return 0, false
	}
	return acc.bal, true
}

func (acc *Account) Close() (payout int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()
	if acc.closed {
		return 0, false
	}
	acc.closed = true
	return acc.bal, true
}
