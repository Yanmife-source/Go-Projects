package main

import "fmt"

var input int
var balance int
var deposit int
var withdrawal int


func main() {
	Loop:
	for {
		
		fmt.Print("1. Check Balance\n2.Deposit Money\n3.Withdraw Money\n4.Transfer funds to another account\n5.Cancel transaction\nSelect an option: ")
		_,err:=fmt.Scan(&input)
		if err!=nil{
			fmt.Println("Kindly enter a number")
			continue
		}

		
		switch input{
			case 1:
				fmt.Printf("Your Account Balance is $%d\n",balance)
			case 2:
				for {
					fmt.Print("Enter amount you would like to deposit: ")
					fmt.Scanln(&deposit)
					if deposit<0 {
						fmt.Println("Kindly enter a valid amount")
						continue
					} else {
						balance+=deposit
						fmt.Println("Thanks for banking with us")
						break 
						
					}
				}
				
			case 3:
				for{
					fmt.Print("Enter amount you would like to withdraw: ")
					fmt.Scanln(&withdrawal)
					
					if withdrawal <0 {
						fmt.Println("Enter a valid amount")
						continue
					} else {
						if withdrawal>balance {
							fmt.Println("Amount requesred exceeds account balance\nTry again but with a smaller amount")
							continue
						} else {
						balance-=withdrawal
						fmt.Print("Thanks for banking with us\n")
						break
						}
					}
				}
			case 4:
				fmt.Println("This feature is coming soon. Anticipate!!")
				break Loop
			case 5:
				break Loop
			default:
				fmt.Println("Kindly enter a value between 1-5")
		}
	}
}
