/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package node

import (
	"sync"

	"github.com/CESSProject/cess-bucket/configs"
	"github.com/CESSProject/cess-bucket/pkg/cache"
	"github.com/CESSProject/cess-bucket/pkg/confile"
	"github.com/CESSProject/cess-bucket/pkg/logger"
	"github.com/CESSProject/p2p-go/core"
	"github.com/CESSProject/sdk-go/core/sdk"
)

type Bucket interface {
	Run()
}

type Node struct {
	confile.Confile
	logger.Logger
	cache.Cache
	sdk.SDK
	core.P2P
	Lock  *sync.RWMutex
	Peers map[string]string
}

// New is used to build a node instance
func New() *Node {
	return &Node{
		Lock:  new(sync.RWMutex),
		Peers: make(map[string]string, 10),
	}
}

func (n *Node) Run() {
	go n.TaskMgt()
	configs.Ok("Start successfully")
	select {}
}

func (n *Node) PutPeer(peerid, addr string) {
	n.Lock.Lock()
	n.Peers[peerid] = addr
	n.Lock.Unlock()
}

func (n *Node) Has(peerid string) bool {
	n.Lock.RLock()
	_, ok := n.Peers[peerid]
	n.Lock.RUnlock()
	return ok
}

func (n *Node) GetPeerAddr(peerid string) string {
	n.Lock.RLock()
	defer n.Lock.RUnlock()
	addr, ok := n.Peers[peerid]
	if ok {
		return addr
	}
	return ""
}
