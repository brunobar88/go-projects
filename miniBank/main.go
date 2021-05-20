package main

import (
	"bank/account"
	"fmt"
)

func PayTicket(account CheckAccount, tiketValue float64) {
	account.Withdraw(tiketValue)
}

type CheckAccount interface {
	Withdraw(value float64) string
}

func main() {
	accountTeste := account.SavingsAccount{}
	accountTeste2 := account.CurrentAcount{}

	accountTeste.Deposit(200)
	PayTicket(&accountTeste, 56)

	accountTeste2.Deposit(510)
	PayTicket(&accountTeste2, 100)

	fmt.Println(accountTeste.GetBalance())
	fmt.Println(accountTeste2.GetBalance())
}
