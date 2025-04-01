package go_routine_banking_example

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

type BankAccount struct {
	balance    int
	customerId int
	lock       sync.Mutex
}

func (a *BankAccount) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *BankAccount) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("not enough money in account")
		return
	}
	a.balance = a.balance - v
	fmt.Printf("%d withdrawn. Remaining Balance %d\n", v, a.balance)
}

func GoRoutineBankingExample() {

	var acct BankAccount
	acct.balance = 100
	pl("Balance:", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(12)
	}
	time.Sleep(2 * time.Second)

}
