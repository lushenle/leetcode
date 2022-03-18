package leetcode

import (
	"reflect"
	"testing"
)

func TestBank_Deposit(t *testing.T) {
	type fields struct {
		B []int64
	}
	type args struct {
		account int
		money   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "testDeposit01",
			fields: fields{B: []int64{10, 20, 25}},
			args:   args{account: 1, money: 30},
			want:   true,
		},
		{
			name:   "testDeposit01",
			fields: fields{B: []int64{10, 20, 25}},
			args:   args{account: 4, money: 30},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bank{
				B: tt.fields.B,
			}
			if got := b.Deposit(tt.args.account, tt.args.money); got != tt.want {
				t.Errorf("Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Transfer(t *testing.T) {
	type fields struct {
		B []int64
	}
	type args struct {
		account1 int
		account2 int
		money    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "testTransfer01",
			fields: fields{B: []int64{10, 20, 30, 40, 50}},
			args:   args{account1: 4, account2: 1, money: 25},
			want:   true,
		},
		{
			name:   "testTransfer02",
			fields: fields{B: []int64{10, 20, 30, 40, 50}},
			args:   args{account1: 4, account2: 1, money: 52},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bank{
				B: tt.fields.B,
			}
			if got := b.Transfer(tt.args.account1, tt.args.account2, tt.args.money); got != tt.want {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Withdraw(t *testing.T) {
	type fields struct {
		B []int64
	}
	type args struct {
		account int
		money   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "testWithdraw01",
			fields: fields{B: []int64{10, 20, 30, 40, 50}},
			args:   args{account: 4, money: 25},
			want:   true,
		},
		{
			name:   "testWithdraw02",
			fields: fields{B: []int64{10, 20, 30, 40, 50}},
			args:   args{account: 1, money: 25},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bank{
				B: tt.fields.B,
			}
			if got := b.Withdraw(tt.args.account, tt.args.money); got != tt.want {
				t.Errorf("Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	type args struct {
		balance []int64
	}
	tests := []struct {
		name string
		args args
		want Bank
	}{
		{
			name: "testConstructor",
			args: args{balance: []int64{10, 20, 30}},
			want: Bank{B: []int64{10, 20, 30}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.balance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
