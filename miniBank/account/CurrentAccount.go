package account

import "personal-projects/miniBank/client"

type CurrentAcount struct {
	AccountOwner                client.AccountOwner
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (account *CurrentAcount) Withdraw(withdrawValue float64) string {
	if account.balance > withdrawValue && withdrawValue > 0 {
		account.balance -= withdrawValue
		return "Withdraw success"
	}

	return "insufficient balance"
}

func (account *CurrentAcount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		account.balance += depositValue
		return "deposit success!", account.balance
	}

	return "failed deposit!", account.balance
}

func (currentAcount *CurrentAcount) Transfer(transferValue float64, targetAccount *CurrentAcount) bool {
	if currentAcount.balance > transferValue && transferValue > 0 {
		currentAcount.Withdraw(transferValue)
		targetAccount.Deposit(transferValue)
		return true
	}
	return false
}

func (account *CurrentAcount) GetBalance() float64 {
	return account.balance
}
