/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package configs

import "time"

const (
	TokenTCESS = 1000000000000000000
	// the time to wait for the event, in seconds
	TimeToWaitEvent = time.Duration(time.Second * 30)
	// Default config file
	DefaultConfigFile = "conf.yaml"
	//
	DefaultWorkspace = "/"
	//
	DefaultServicePort = 4001
	//
	DefaultRpcAddr = "wss://testnet-rpc.cess.cloud/ws/"

	//
	DefaultBootNodeAddr = "_dnsaddr.boot-miner-testnet.cess.cloud"
	//
	DefaultDeossAddr = "https://deoss-pub-gateway.cess.cloud/"

	//
	MinTagFileSize = 600000

	//
	FileMode = 0755
)

const (
	Err_tee_Busy         = "is being fully calculated"
	Err_ctx_exceeded     = "context deadline exceeded"
	Err_file_not_fount   = "no such file"
	Err_miner_not_exists = "the miner not exists"
)

const (
	DevNet  = "devnet"
	TestNet = "testnet"
	MainNet = "mainnet"
)

const (
	Unregistered = iota
	UnregisteredPoisKey
	Registered
)
