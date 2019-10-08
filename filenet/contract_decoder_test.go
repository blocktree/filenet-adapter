package filenet

import (
	"fmt"
	"testing"

	"github.com/blocktree/openwallet/openwallet"
)

func Test_GetTokenBalanceByAddress(t *testing.T) {

	addr := "Wg6agtC24cva6vviwdszBoUDTwrtgbY8cT"
	contract := openwallet.SmartContract{
		Address:  "3178769-1",
		Decimals: 0,
	}

	balance, err := tw.ContractDecoder.GetTokenBalanceByAddress(contract, addr)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(balance[0])
	}
}
