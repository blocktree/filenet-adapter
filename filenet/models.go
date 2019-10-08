/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package filenet

import (
	"fmt"

	"github.com/blocktree/openwallet/crypto"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
)

// type Vin struct {
// 	Coinbase string
// 	TxID     string
// 	Vout     uint64
// 	N        uint64
// 	Addr     string
// 	Value    string
// }

// type Vout struct {
// 	N            uint64
// 	Addr         string
// 	Value        string
// 	ScriptPubKey string
// 	Type         string
// }

type Block struct {

	/*
		{
		    "hash":"d72bd1aaa10d47ee68eedcc1e9b99e8d097d93a9a05e0f683456ff1ad72ba846",
		    "confirmations":1,
		    "size":177,
		    "height":2430532,
		    "version":1,
		    "merkleroot":"f41db24c8fbbf3c2540c63ad42bfee33474ac5e4f4b3656c2b5cc4e9c1f0cebd",
		    "txnumber":1,
		    "tx":
		    [
		        "f41db24c8fbbf3c2540c63ad42bfee33474ac5e4f4b3656c2b5cc4e9c1f0cebd"
		    ],
		    "time":1554106110,
		    "nonce":426,
		    "chainwork":"0000000000000000000000000000000000000000000000000000000000251644",
		    "fuel":0,
		    "fuelrate":1,
		    "previousblockhash":"8fc83042ba76bf893d58407704fa73b65dba8484c1b07e282dbf31a89ead507b"
		}
	*/

	Hash                  string
	PrevBlockHash         string
	TransactionMerkleRoot string
	Timestamp             uint64
	Height                uint64
	Transactions          []string
}

type TxTo struct {
	Address string
	Amount uint64
}

type Transaction struct {
	TxID            string
	TimeStamp       uint64
	From            string
	Amount          uint64
	To              []TxTo
	BlockHeight     uint64
	BlockHash       string
}

func NewTransaction(tx string) *Transaction {
	obj := &Transaction{}

	obj.TxID = gjson.Get(tx, "txid").String()
	obj.TimeStamp = gjson.Get(tx, "timestamp").Uint()
	obj.From = gjson.Get(tx, "from").String()
	obj.Amount = gjson.Get(tx, "value").Uint()
	obj.BlockHeight = gjson.Get(tx, "height").Uint()

	if gjson.Get(tx, "tocount").Uint() != 0 {
		for _, to := range gjson.Get(tx, "transferdetails").Array() {
			obj.To = append(obj.To, TxTo{
				Address: to.Get("to").String(),
				Amount:  to.Get("value").Uint(),
			})
		}
	}
	return obj
}

func NewBlock(json *gjson.Result) *Block {
	obj := &Block{}

	obj.Hash = json.Get("hash").String()
	obj.PrevBlockHash = json.Get("prevhash").String()
	obj.TransactionMerkleRoot = json.Get("txroot").String()
	obj.Timestamp = json.Get("timestamp").Uint()
	obj.Height = json.Get("height").Uint()

	if json.Get("transfercount").Uint() != 0 {
		for _, tx := range json.Get("transfers").Array() {
			obj.Transactions = append(obj.Transactions, tx.String())
		}
	}

	return obj
}

//BlockHeader 区块链头
func (b *Block) BlockHeader() *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Hash
	//obj.Confirmations = b.Confirmations
	obj.Merkleroot = b.TransactionMerkleRoot
	obj.Previousblockhash = b.PrevBlockHash
	obj.Height = b.Height
	obj.Time = b.Timestamp
	obj.Symbol = Symbol

	return &obj
}

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	TxID        string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *UnscanRecord {
	obj := UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}
