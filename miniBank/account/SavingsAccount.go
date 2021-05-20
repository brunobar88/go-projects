package account

import "personal-projects/miniBank/client"

type SavingsAccount struct {
	Titular                                client.AccountOwner
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (account *SavingsAccount) Withdraw(withdrawValue float64) string {
	if account.balance > withdrawValue && withdrawValue > 0 {
		account.balance -= withdrawValue
		return "Withdraw success"
	}

	return "insufficient balance"
}

func (account *SavingsAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		account.balance += depositValue
		return "deposit success!", account.balance
	}

	return "failed deposit!", account.balance
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.balance
}
