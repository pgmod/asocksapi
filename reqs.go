package asocksapi

type CreatePortsReq struct {
	CountryCode      string           `json:"country_code"`
	City             *string          `json:"city"`
	State            *string          `json:"state"`
	Asn              *int             `json:"asn"`
	TypeId           TypeId           `json:"type_id"`
	ProxyTypeId      ProxyTypeId      `json:"proxy_type_id"`
	Name             string           `json:"name"`
	ServerPortTypeId ServerPortTypeId `json:"server_port_type_id"`
	Count            string           `json:"count"`
}