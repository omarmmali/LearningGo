package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(10)

		got := wallet.Balance()
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := 20
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(100)
		
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}