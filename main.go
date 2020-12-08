package main

// Create wallet

import (
	"fmt"

	debicred "github.com/debicred/debicred-go-sdk/sdk"
)

func main() {
	id := "ba6145c57ddb42af90b7ceecf033948b"
	pub := "6c08bbcbd6d8a42720698bb91de305e7d179d31ac4bc3b4b0c5ed71132d640d2"
	sec := "9771a2b1a4dcf649a372fcd355691d6262600b6e98d2ec2cb48d24da9b9cf42e1ffc87cf8f9671bedd9ea325153b9aea23c77a9327237d5faa3cc1c6"
	wallet := "d71be42763f04f7a8f3c04cd9c6b6cf6"

	customer := debicred.Init(id, pub, sec)

	// payload := &debicred.Payload{
	// 	Code:   "USD",
	// 	Symbol: "$",
	// 	Amount: 100.4,
	// }

	res, err := customer.GetWalletTransactions(wallet)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
