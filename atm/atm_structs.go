package main

import "fmt"

type ATM struct {
	start bool
	balance int
}
func newTransaction(initBalance int) *ATM {
	sess:=ATM{start:true}
	if initBalance>0{
		sess.balance=initBalance	
	}
	return &sess
}
func (a ATM) CheckBalance() int {
	if a.start{
		return a.balance
	} else {return 0}
}


func (a *ATM) Deposit(amount int) (string,int) {
	if a.start{
		if amount>0{
			a.balance+=amount
			return "Balance:",a.balance
		} else {
			return "Provide a valid number above $0",0}
	} else {return "A new transaction hasn't been initialized",0}	
}

func (a *ATM) Withdraw(amount int) (string,int) {
	if a.start{
		if amount>0{
			if amount<=a.balance{
				a.balance-=amount
				return "Balance:",a.balance
			} else { return "The amount requested exceeds your balance",0}
		} else {return "Provide a valid number above $0",0}
	} else {return "A new transaction hasn't been initialized",0}
	
}
func main(){
	ymf:=newTransaction(1500)
	fmt.Printf("Balance:%d\n",ymf.CheckBalance())
	dep_str,dep_int:=ymf.Deposit(120)
	fmt.Println(dep_str,dep_int)
	with_str,with_int:=ymf.Withdraw(1200)
	fmt.Println(with_str,with_int)
	fmt.Printf("Balance:%d\n",ymf.CheckBalance())
	
}
	