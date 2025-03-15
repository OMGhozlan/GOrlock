package pointerr

import (
	"fmt"
	"reflect"
	"testing"
)

func assertValueMessage(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestWallet(t *testing.T) {
	checkWalletBalance := func(t testing.TB, wallet *Wallet, want any) {
		t.Helper()
		got := (*wallet).Balance()
		// fmt.Printf("Current balance is %d \n", (*wallet).Balance())
		// fmt.Printf("Address of balance in test is %v \n", (*wallet).balance)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	checkWalletBalanceString := func(t testing.TB, wallet *Wallet, want any) {
		t.Helper()
		got := (*wallet).Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	checkForError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	wallet := Wallet{}
	testSuite := []struct {
		name   string
		wallet *Wallet
		want   Bitcoin
	}{
		{name: "Deposting 10 BTC, wanting 10 in Balance", wallet: &wallet, want: Bitcoin(10)},
		{name: "Deposting another 10 BTC, wanting 20 in Balance", wallet: &wallet, want: Bitcoin(20)},
		{name: "Withdrawing 10 BTC, wanting 10 in Balance", wallet: &wallet, want: Bitcoin(10)},
		{name: "Withdrawing insufficent funds, wanting error", wallet: &wallet, want: Bitcoin(10)},
	}
	for testNum, test := range testSuite {
		t.Run(test.name, func(t *testing.T) {
			switch testNum {
			case 0:
				test.wallet.Deposit(Bitcoin(10))
				checkWalletBalance(t, test.wallet, test.want)
			case 1:
				test.wallet.Deposit(Bitcoin(10))
				checkWalletBalanceString(t, test.wallet, test.want)
			case 2:
				err := test.wallet.Withdraw(Bitcoin(10))
				checkWalletBalanceString(t, test.wallet, test.want)
				checkForError(t, err, ErrInsufficientFunds)
			case 3:
				err := test.wallet.Withdraw(Bitcoin(500))
				checkForError(t, err, ErrInsufficientFunds)
				checkWalletBalanceString(t, test.wallet, test.want)
			default:
			}
			defer fmt.Println(test.name, "done")
		})
	}
	// wallet.Deposit(Bitcoin(10))
	// checkWalletBalance(t, wallet, Bitcoin(10))
	// wallet.Deposit(Bitcoin(10))
	// checkWalletBalanceString(t, wallet, Bitcoin(20))
	// wallet.Withdraw(Bitcoin(10))
	// checkWalletBalanceString(t, wallet, Bitcoin(10))
}
