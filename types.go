package asocksapi

type ProxyTypeId int

const (
	RESEDENTIAL ProxyTypeId = 1
	ALL         ProxyTypeId = 2
	MOBILE      ProxyTypeId = 3
	CORPORATE   ProxyTypeId = 4
)

type TypeId int

const (
	KEEP_PROXY        TypeId = 1
	KEEP_CONNECTION   TypeId = 2
	ROTATE_CONNECTION TypeId = 3
)

type ServerPortTypeId int

const (
	DEDICATED ServerPortTypeId = 1
	SHARED    ServerPortTypeId = 0
)
