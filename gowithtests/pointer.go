package main

import "fmt"
import "errors"

// Bitcoin type is an alias to integer type
type Bitcoin int

// Wallet defines a crypto wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit puts $$ on the wallet
func (wallet *Wallet) Deposit(amount Bitcoin) error {
	//fmt.Printf("address of balance in Deposit is %v \n", &wallet.balance)
	wallet.balance += amount
	return nil
}

// Withdraw draws $$ from the wallet
func (wallet *Wallet) Withdraw(amount Bitcoin) error {

	if wallet.balance < amount {
		return errors.New("insufficient balance")
	}

	if wallet.balance <= 0 {
		return errors.New("insufficient balance")
	}

	wallet.balance -= amount

	return nil

}

// Balance shows the remaining $$ on the wallet
func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// func main() {
// 	w := Wallet{Bitcoin(10)}
// 	w.Deposit(Bitcoin(10))
// 	w.Withdraw(Bitcoin(1000))
// 	fmt.Println(w.balance)
// }
