/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package tutorial_simple_client_test

import (
	// go std libs
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	// 3rd-party libs
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	// CESS libs
	cess "github.com/CESSProject/cess-go-sdk"
	cessConfig "github.com/CESSProject/cess-go-sdk/config"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	p2pgo "github.com/CESSProject/p2p-go"
)

const DEFAULT_WAIT_TIME = time.Second * 15
const P2P_PORT = 4001

func TestMain(m *testing.M) {
	fmt.Println("Get into TestMain")

	godotenv.Load("../.env.testnet")
	os.Exit(m.Run())
}

func TestSimpleClient(t *testing.T) {
	_, err := cess.New(
		cessConfig.CharacterName_Client,
		cess.ConnectRpcAddrs(strings.Split(os.Getenv("RPC_ADDRS"), " ")),
		cess.Mnemonic(os.Getenv("MY_MNEMONIC")),
		cess.TransactionTimeout(time.Duration(DEFAULT_WAIT_TIME)),
	)

	assert.NoError(t, err)
}

func TestDeOSS(t *testing.T) {
	conn, err := cess.New(
		cessConfig.CharacterName_Deoss,
		cess.ConnectRpcAddrs(strings.Split(os.Getenv("RPC_ADDRS"), " ")),
		cess.Mnemonic(os.Getenv("MY_MNEMONIC")),
		cess.TransactionTimeout(time.Duration(DEFAULT_WAIT_TIME)),
	)
	assert.NoError(t, err)

	bootnodes := make([]string, 0)

	for _, node := range strings.Split(os.Getenv("BOOTSTRAP_NODES"), " ") {
		addrs, err := utils.ParseMultiaddrs(node)
		if err != nil {
			continue
		}
		bootnodes = append(bootnodes, addrs...)
	}

	p2p, err := p2pgo.New(
		context.Background(),
		p2pgo.ListenPort(P2P_PORT),
		p2pgo.Workspace("./"),
		p2pgo.BootPeers(bootnodes),
	)

	assert.NoError(t, err)

	fmt.Printf("p2p:\n%+v\n\n", p2p)
	fmt.Printf("conn:\n%+v\n\n", conn)

	txHash, _, err := conn.Register(conn.GetRoleName(), p2p.GetPeerPublickey(), "", 0)

	assert.NoError(t, err)

	fmt.Printf("txHash:\n%+v\n\n", txHash)
}
