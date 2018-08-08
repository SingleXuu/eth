package account

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func NewAcct(pass string) {
	cli, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		//log.Fatal("connect to geth error:",err)
		fmt.Println("connect to geth error :", err)
	}

	defer cli.Close()
	var account string
	err = cli.Call(&account, "personal_newAccount:", pass)
	if nil != err {
		fmt.Println("call personal_newAccount error:", err)
	}
	fmt.Println("account=", account)
}
