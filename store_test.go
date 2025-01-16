package lab

import (
	"testing"
	"time"
)

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
			desc: "lab-r1",
			device: Device{
				Hostname:    "lab-r1",
				MGMTAddress: "10.0.0.86",
				Loopback:    "10.1.0.1/32",
				ISO:         "49.0001.0100.0001.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.12.1/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R2",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.16.1/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R6",
					},
				},
			},
		},
		{
			desc: "lab-r2",
			device: Device{
				Hostname:    "lab-r2",
				MGMTAddress: "10.0.0.23",
				Loopback:    "10.1.0.2/32",
				ISO:         "49.0001.0100.0002.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.12.2/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R1",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.23.2/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R3",
					},
				},
			},
		},
		{
			desc: "lab-r3",
			device: Device{
				Hostname:    "lab-r3",
				MGMTAddress: "10.0.0.212",
				Loopback:    "10.1.0.3/32",
				ISO:         "49.0001.0100.0003.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.34.3/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R4",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.23.3/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R2",
					},
				},
			},
		},
		{
			desc: "lab-r4",
			device: Device{
				Hostname:    "lab-r4",
				MGMTAddress: "10.0.0.150",
				Loopback:    "10.1.0.4/32",
				ISO:         "49.0001.0100.0004.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.34.4/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R3",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.45.4/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R5",
					},
				},
			},
		},
		{
			desc: "lab-r5",
			device: Device{
				Hostname:    "lab-r5",
				MGMTAddress: "10.0.0.87",
				Loopback:    "10.1.0.5/32",
				ISO:         "49.0001.0100.0005.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.58.5/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R8",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.45.5/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R4",
					},
				},
			},
		},
		{
			desc: "lab-r6",
			device: Device{
				Hostname:    "lab-r6",
				MGMTAddress: "10.0.0.24",
				Loopback:    "10.1.0.6/32",
				ISO:         "49.0001.0100.0006.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.67.6/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R7",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.16.6/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R1",
					},
				},
			},
		},
		{
			desc: "lab-r7",
			device: Device{
				Hostname:    "lab-r7",
				MGMTAddress: "10.0.0.213",
				Loopback:    "10.1.0.7/32",
				ISO:         "49.0001.0100.0007.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.67.7/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R6",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.78.7/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R8",
					},
				},
			},
		},
		{
			desc: "lab-r8",
			device: Device{
				Hostname:    "lab-r8",
				MGMTAddress: "10.0.0.149",
				Loopback:    "10.1.0.8/32",
				ISO:         "49.0001.0100.0008.00",
				Interfaces: []Interface{
					{
						IFD:         "ge-0/0/0",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.58.8/24",
						Role:        "infrastructure",
						Description: "TO-LAB-R5",
					},
					{
						IFD:         "ge-0/0/1",
						IFL:         "0",
						VLAN:        "0",
						Address:     "172.16.78.8/24",
						Role:        "infrastructure",
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

func TestGetDevices(t *testing.T) {
	store := NewStore()
	devs, err := store.GetDevices()
	if err != nil {
		t.Error(err)
	}
	t.Log(devs)
}

func TestAddBaseConfigs(t *testing.T) {
	testCases := []struct {
		desc string
		cfg  Config
	}{
		{
			desc: "lab-r1",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r1-base.conf",
				Device:    Device{ID: 1},
			},
		},
		{
			desc: "lab-r2",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r2-base.conf",
				Device:    Device{ID: 2},
			},
		},
		{
			desc: "lab-r3",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r3-base.conf",
				Device:    Device{ID: 3},
			},
		},
		{
			desc: "lab-r4",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r4-base.conf",
				Device:    Device{ID: 4},
			},
		},
		{
			desc: "lab-r5",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r5-base.conf",
				Device:    Device{ID: 5},
			},
		},
		{
			desc: "lab-r6",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r6-base.conf",
				Device:    Device{ID: 6},
			},
		},
		{
			desc: "lab-r7",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r7-base.conf",
				Device:    Device{ID: 7},
			},
		},
		{
			desc: "lab-r8",
			cfg: Config{
				CreatedAt: time.Now(),
				Type:      BaseCfg,
				Path:      "./docs/configs/lab-r8-base.conf",
				Device:    Device{ID: 8},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			store := NewStore()
			_, err := store.AddConfig(tC.cfg)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetDeviceInterfaces(t *testing.T) {
	store := NewStore()
	dev, err := store.GetDevicesInterfaces(Device{ID: 1})
	if err != nil {
		t.Error(err)
	}
	if len(dev.Interfaces) != 2 {
		t.Errorf("got %v; expected %v", len(dev.Interfaces), 2)
	}
	t.Logf("%+v\n", dev)
}
