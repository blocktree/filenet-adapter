package filenet

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

const (
	testNodeAPI = ""
)

func Test_getBlockHeight(t *testing.T) {
	token := ""
	c := NewClient(testNodeAPI, token, true)

	r, err := c.getBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("height:", r)
	}

}

func Test_getBlockByHeight(t *testing.T) {
	token := ""
	c := NewClient(testNodeAPI, token, true)
	r, err := c.getBlockByHeight(383759)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockHash(t *testing.T) {
	token := ""
	c := NewClient(testNodeAPI, token, true)

	height := uint64(9725)

	r, err := c.getBlockHash(height)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getBalance(t *testing.T) {
	token := ""
	c := NewClient(testNodeAPI, token, true)

	address := "3VhHpvwbzG3wG5atoQCGQESrP3XyExTdeVP2HKTE9m3ZhUD9QJqHD94"//"3Psbq3enwAmwXGa2QejWFdd9AwV1rczE6w1GPzs6WTPmJpKbmWghsLB"

	r, err := c.getBalance(address)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getTransaction(t *testing.T) {
	token := ""
	c := NewClient(testNodeAPI, token, true)
	txid := "ad9229bb5846e790f80adb4e525845476d1d23b67ff8db73eb07c9d7b7b9da21" //"9KBoALfTjvZLJ6CAuJCGyzRA1aWduiNFMvbqTchfBVpF"

	r, err := c.getTransaction(txid)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_convert(t *testing.T) {

	amount := uint64(5000000001)

	amountStr := fmt.Sprintf("%d", amount)

	fmt.Println(amountStr)

	d, _ := decimal.NewFromString(amountStr)

	w, _ := decimal.NewFromString("100000000")

	d = d.Div(w)

	fmt.Println(d.String())

	d = d.Mul(w)

	fmt.Println(d.String())

	r, _ := strconv.ParseInt(d.String(), 10, 64)

	fmt.Println(r)

	fmt.Println(time.Now().UnixNano())
}

func Test_getTransactionByAddresses(t *testing.T) {
	addrs := "ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11"

	token := ""
	c := NewClient(testNodeAPI, token, true)
	result, err := c.getMultiAddrTransactions(0, -1, addrs)

	if err != nil {
		t.Error("get transactions failed!")
	} else {
		for _, tx := range result {
			fmt.Println(tx.TxID)
		}
	}
}

func Test_getContractAccountInfo(t *testing.T) {
	regid := "3291379-2" //"1549609-1"
	address := "WPhr838tCoAMu22qvLg7JL6y6c8WESFchQ"

	token := ""
	c := NewClient(testNodeAPI, token, true)

	r, err := c.getContractAccountBalence(regid, address)
	fmt.Println(err)
	fmt.Println(r)
}
