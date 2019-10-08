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

var (
	tw *WalletManager
)

func init() {
	tw = NewWalletManager()
	tw.Config.RpcUser = "fn"
	tw.Config.RpcPassword = "fn_wallet_2019"
	token := BasicAuth("fn", "fn_wallet_2019")
	tw.Client = NewClient("http://47.91.224.127:20026", token, true)
}
