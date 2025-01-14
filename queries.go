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
	device_id,
) VALUES (
	?, ?, ?, ?, ?
);	
	`

	getConfigsQuery = `-- get many 
SELECT
	id,
	created_at,
	desc,
	path,
	type
FROM configs;
	`

	createDevicesTable = ` -- create devices
CREATE TABLE IF NOT EXISTS devices (	
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATETIME,
	hostname TEXT,
	mgmt_address TEXT,
	UNIQUE(hostname, mgmt_address)
	-- can add this back if necessary
	-- interfaces BLOB
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
	mgmt_address
) VALUES (
	?, ?, ?
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

	getDeviceQuery = ` --get one
SELECT
	d.id,
	d.hostname,
	d.mgmt_address,
	i.id,
	i.ifd,
	i.ifl,
	i.vlan,
	i.address,
	i.role,
	i.desc
FROM devices d
JOIN interfaces i 
ON d.id = i.device_id;`
)
