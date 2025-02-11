package lab

var (
	createConfigsTable = ` --create configs
CREATE TABLE IF NOT EXISTS configs (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATETIME,
	cfg_path TEXT,
	type TEXT,
	desc TEXT,
	device_id INTEGER,
	FOREIGN KEY (device_id) REFERENCES devices(id)
);`

	insertConfigQuery = ` -- add one
INSERT INTO configs (
	created_at,
	cfg_path,
	type,
	desc,
	device_id
) VALUES (
	?, ?, ?, ?, ?
);	
	`

	getConfigsQuery = `-- get many 
SELECT
	c.id,
	c.created_at,
	c.desc,
	c.cfg_path,
	c.type,
	d.hostname,
	d.mgmt_address
FROM configs c
JOIN devices d
ON c.device_id = d.id;
	`

	createDevicesTable = ` -- create devices
CREATE TABLE IF NOT EXISTS devices (	
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATETIME,
	hostname TEXT,
	mgmt_address TEXT,
	loopback TEXT,
	iso TEXT,
	UNIQUE(hostname, mgmt_address)
);`

	createInterfacesTable = ` -- create interfaces
CREATE TABLE IF NOT EXISTS interfaces (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATETIME,
	device_id INTEGER,
	ifd TEXT,
	ifl TEXT,
	vlan TEXT,
	address TEXT,
	role TEXT,
	desc TEXT,
	UNIQUE(ifd, ifl, device_id)
	FOREIGN KEY (device_id) REFERENCES devices(id)
);`

	insertDevice = ` -- add one
INSERT INTO devices (
	created_at,
	hostname,
	mgmt_address,
	loopback,
	iso
) VALUES (
	?, ?, ?, ?, ?
);`
	insertInterface = ` -- add one
INSERT INTO interfaces (
	created_at,
	device_id,
	ifd,
	ifl,
	vlan,
	address,
	role,
	desc
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?
);`

	getDevicesQuery = ` --get one
SELECT
	id,
	hostname,
	mgmt_address,
	loopback,
	iso
FROM devices`

	getDeviceByHostnameQuery = ` --get one
SELECT
	id,
	hostname,
	mgmt_address,
	loopback,
	iso
FROM devices
WHERE hostname = ?`

	getDevicesInterfacesByDeviceIDQuery = ` --get one
SELECT
	i.id,
	i.ifd,
	i.ifl,
	i.vlan,
	i.address,
	i.role,
	i.desc
FROM interfaces i
JOIN devices d 
ON d.id = i.device_id
WHERE d.id = ?
ORDER BY i.ifd;`
)
