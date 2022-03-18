package leetcode

type Bank struct {
	B []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 <= len(b.B) && account2 <= len(b.B) && b.B[account1-1] >= money {
		b.B[account1-1] -= money
		b.B[account2-1] += money
		return true
	}

	return false
}

func (b *Bank) Deposit(account int, money int64) bool {
	if account <= len(b.B) {
		b.B[account-1] += money
		return true
	}

	return false
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if account <= len(b.B) && b.B[account-1] >= money {
		b.B[account-1] -= money
		return true
	}

	return false
}
