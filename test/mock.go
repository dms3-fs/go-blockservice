package bstest

import (
	. "github.com/dms3-fs/go-blockservice"

	bitswap "github.com/dms3-fs/go-bitswap"
	tn "github.com/dms3-fs/go-bitswap/testnet"
	delay "github.com/dms3-fs/go-fs-delay"
	mockrouting "github.com/dms3-fs/go-fs-routing/mock"
)

// Mocks returns |n| connected mock Blockservices
func Mocks(n int) []BlockService {
	net := tn.VirtualNetwork(mockrouting.NewServer(), delay.Fixed(0))
	sg := bitswap.NewTestSessionGenerator(net)

	instances := sg.Instances(n)

	var servs []BlockService
	for _, i := range instances {
		servs = append(servs, New(i.Blockstore(), i.Exchange))
	}
	return servs
}
