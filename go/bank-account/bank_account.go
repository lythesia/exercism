package account

import "sync"

type Account struct {
	balance int64
	avail   bool
	lck     *sync.Mutex
}

func Open(init int64) *Account {
	if init < 0 {
		return nil
	}
	return &Account{init, true, &sync.Mutex{}}
}

func (a *Account) Close() (int64, bool) {
	a.lck.Lock()
	defer a.lck.Unlock()
	if !a.avail {
		return 0, false
	}
	ans := a.balance
	a.balance = 0
	a.avail = false
	return ans, true
}

func (a *Account) Balance() (int64, bool) {
	a.lck.Lock()
	defer a.lck.Unlock()
	if !a.avail {
		return 0, false
	}
	return a.balance, a.avail
}

func (a *Account) Deposit(amt int64) (int64, bool) {
	a.lck.Lock()
	defer a.lck.Unlock()
	if !a.avail || -amt > a.balance {
		return 0, false
	}
	a.balance += amt
	return a.balance, true
}
