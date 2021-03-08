package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertError := func(t testing.TB, e error) {
		t.Helper()

		if e == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(20)
		assertBalance(t, got, want)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		wallet.Withdraw(Bitcoin(50))

		got := wallet.Balance()
		want := Bitcoin(50)
		assertBalance(t, got, want)

	})

	t.Run("Overdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(500))
		assertError(t, err)

	})

}
