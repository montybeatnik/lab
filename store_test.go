package lab

import "testing"

func TestCreateTables(t *testing.T) {
	testCases := []struct {
		desc string
		cfg  string
	}{
		{
			desc: "configs",
			cfg:  createConfigsTable,
		},
		{
			desc: "devices",
			cfg:  createDevicesTable,
		},
		{
			desc: "interfaces",
			cfg:  createInterfacesTable,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			store := NewStore()
			if err := store.CreateTable(tC.cfg); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestAddBaseConfigs(t *testing.T) {
	testCases := []struct {
		desc string
		cfg  Config
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestAddDevice(t *testing.T) {
	testCases := []struct {
		desc   string
		device Device
	}{
		{
			desc: "lab-r1",
			device: Device{
				Hostname:    "lab-r1",
				MGMTAddress: "10.0.0.86/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.12.1/24",
						Role:        "core",
						Description: "TO-LAB-R2",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.16.1/24",
						Role:        "core",
						Description: "TO-LAB-R6",
					},
				},
			},
		},
		{
			desc: "lab-r2",
			device: Device{
				Hostname:    "lab-r2",
				MGMTAddress: "10.0.0.23/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.12.2/24",
						Role:        "core",
						Description: "TO-LAB-R1",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.23.2/24",
						Role:        "core",
						Description: "TO-LAB-R3",
					},
				},
			},
		},
		{
			desc: "lab-r3",
			device: Device{
				Hostname:    "lab-r3",
				MGMTAddress: "10.0.0.212/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.34.3/24",
						Role:        "core",
						Description: "TO-LAB-R4",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.23.3/24",
						Role:        "core",
						Description: "TO-LAB-R2",
					},
				},
			},
		},
		{
			desc: "lab-r4",
			device: Device{
				Hostname:    "lab-r4",
				MGMTAddress: "10.0.0.150/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.34.4/24",
						Role:        "core",
						Description: "TO-LAB-R3",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.45.4/24",
						Role:        "core",
						Description: "TO-LAB-R5",
					},
				},
			},
		},
		{
			desc: "lab-r5",
			device: Device{
				Hostname:    "lab-r5",
				MGMTAddress: "10.0.0.87/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.58.5/24",
						Role:        "core",
						Description: "TO-LAB-R8",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.45.5/24",
						Role:        "core",
						Description: "TO-LAB-R4",
					},
				},
			},
		},
		{
			desc: "lab-r6",
			device: Device{
				Hostname:    "lab-r6",
				MGMTAddress: "10.0.0.24/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.67.6/24",
						Role:        "core",
						Description: "TO-LAB-R7",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.16.6/24",
						Role:        "core",
						Description: "TO-LAB-R1",
					},
				},
			},
		},
		{
			desc: "lab-r7",
			device: Device{
				Hostname:    "lab-r7",
				MGMTAddress: "10.0.0.213/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.67.7/24",
						Role:        "core",
						Description: "TO-LAB-R6",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.78.7/24",
						Role:        "core",
						Description: "TO-LAB-R8",
					},
				},
			},
		},
		{
			desc: "lab-r8",
			device: Device{
				Hostname:    "lab-r8",
				MGMTAddress: "10.0.0.149/24",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.58.8/24",
						Role:        "core",
						Description: "TO-LAB-R5",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.78.8/24",
						Role:        "core",
						Description: "TO-LAB-R7",
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			store := NewStore()
			id, err := store.AddDevice(tC.device)
			if err != nil {
				t.Error(err)
			}
			t.Log(id)
		})
	}
}

func TestGetDevice(t *testing.T) {
	store := NewStore()
	devs, err := store.GetDevices()
	if err != nil {
		t.Error(err)
	}
	t.Log(devs)
}
