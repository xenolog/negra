package types

import (
	"github.com/juju/juju/network/netplan"
)

// type Netplan netplan.Network
type Netplan struct {
	netplan.Netplan
}
