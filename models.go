package lab

import "time"

type Config struct {
	ID        int
	CreatedAt time.Time
	Desc      string
	Path      string
	Type      string
}

type Interface struct {
	ID int
	// IFD physical interface device. This is fairly specific to Juniper
	// devices, but I like it.
	// Example: "ge-0/0/0"
	IFD string
	// IFL is a logical interface. This is usually 0 when no VLAN is
	// in use, and if there is a VLAN, it normally matches the VLAN value.
	// This, of course, is a guideline and not always the case.
	// Examples: [0, 20]
	IFL string
	// VLAN is a 802.1q 12 bit value that represents a broadcast domain.
	VLAN string
	// Address this should be in a family, but I'm cutting corners here.
	// This will come back to byte (see what I did there) me.
	Address string
	// Role determines the functionality of the interface.
	// Examples: ["core", "access"]
	Role        string
	Description string
}

type Device struct {
	ID          int
	Hostname    string
	MGMTAddress string
	Interfaces  []Interface
}
