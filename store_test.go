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

func TestAddDevice(t *testing.T) {
	testCases := []struct {
		desc   string
		device Device
	}{
		{
			desc: "core router",
			device: Device{
				Hostname:    "testlab-r1",
				MGMTAddress: "192.168.1.1/24",
				Interfaces: []Interface{
					{
						IFD:     "ge-0/0/1",
						IFL:     "10",
						VLAN:    "10",
						Address: "10.150.1.0/31",
						Role:    "core",
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
