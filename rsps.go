package asocksapi

type CountriesRsp struct {
	Success   bool `json:"success"`
	Countries []struct {
		ID   int    `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"countries"`
}

type CreateProxyRsp struct {
	Succes bool        `json:"success"`
	Data   []UserProxy `json:"data"`
}

type DeleteProxyRsp struct {
	Succes bool `json:"success"`
}

type UserProxy struct {
	Id          int         `json:"id"`
	AuthTypeId  int         `json:"auth_type_id"`
	TypeId      TypeId      `json:"type_id"`
	CountryCode string      `json:"country_code"`
	StateId     int         `json:"state_id"`
	CityId      int         `json:"city_id"`
	Asn         int         `json:"asn"`
	Password    string      `json:"password"`
	ProxyTypeId ProxyTypeId `json:"proxy_type_id"`
	Name        string      `json:"name"`
	Port        uint16      `json:"port"`
	Login       string      `json:"login"`
	Server      string      `json:"server"`
	Country     string      `json:"country"`
	State       string      `json:"state"`
	City        string      `json:"city"`
	// Status           string      `json:"status"`
	Extip            string `json:"extip"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	LastConnectedAgo int    `json:"last_connected_ago"`
}

type ListPortsRsp struct {
	Success bool `json:"success"`
	Message struct {
		CountProxies int `json:"countProxies"`
		Pagination   struct {
			Page       int `json:"page"`
			PageCount  int `json:"pageCount"`
			PageSize   int `json:"pageSize"`
			TotalCount int `json:"totalCount"`
		} `json:"pagination"`
		Proxies []UserProxy `json:"proxies"`
	} `json:"message"`
}
