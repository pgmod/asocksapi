package asocksapi

import (
	"encoding/json"
	"fmt"
)

var Countries = []string{
	"SH",
	"RW",
	"SO",
	"YE",
	"IQ",
	"SA",
	"IR",
	"CY",
	"TZ",
	"SY",
	"AM",
	"KE",
	"CD",
	"DJ",
	"UG",
	"CF",
	"SC",
	"JO",
	"LB",
	"KW",
	"OM",
	"QA",
	"BH",
	"AE",
	"IL",
	"TR",
	"ET",
	"ER",
	"EG",
	"SD",
	"GR",
	"BI",
	"EE",
	"LV",
	"AZ",
	"LT",
	"SJ",
	"GE",
	"MD",
	"BY",
	"FI",
	"AX",
	"UA",
	"MK",
	"HU",
	"BG",
	"AL",
	"PL",
	"RO",
	"XK",
	"ZW",
	"ZM",
	"KM",
	"MW",
	"LS",
	"BW",
	"MU",
	"SZ",
	"RE",
	"ZA",
	"YT",
	"MZ",
	"MG",
	"AF",
	"PK",
	"BD",
	"TM",
	"TJ",
	"LK",
	"BT",
	"IN",
	"MV",
	"IO",
	"NP",
	"MM",
	"UZ",
	"KZ",
	"KG",
	"CC",
	"PW",
	"VN",
	"TH",
	"ID",
	"LA",
	"TW",
	"PH",
	"MY",
	"CN",
	"HK",
	"BN",
	"MO",
	"KH",
	"KR",
	"JP",
	"KP",
	"SG",
	"CK",
	"TL",
	"RU",
	"MN",
	"AU",
	"CX",
	"MH",
	"FM",
	"PG",
	"SB",
	"TV",
	"NR",
	"VU",
	"NC",
	"NF",
	"NZ",
	"FJ",
	"LY",
	"CM",
	"SN",
	"CG",
	"PT",
	"LR",
	"CI",
	"GH",
	"GQ",
	"NG",
	"BF",
	"TG",
	"GW",
	"MR",
	"BJ",
	"GA",
	"SL",
	"ST",
	"GI",
	"GM",
	"GN",
	"TD",
	"NE",
	"ML",
	"EH",
	"TN",
	"ES",
	"MA",
	"MT",
	"DZ",
	"FO",
	"DK",
	"IS",
	"GB",
	"CH",
	"SE",
	"NL",
	"AT",
	"BE",
	"DE",
	"LU",
	"IE",
	"MC",
	"FR",
	"AD",
	"LI",
	"JE",
	"IM",
	"GG",
	"SK",
	"CZ",
	"NO",
	"VA",
	"SM",
	"IT",
	"SI",
	"ME",
	"HR",
	"BA",
	"AO",
	"NA",
	"BB",
	"CV",
	"GY",
	"GF",
	"SR",
	"PM",
	"GL",
	"PY",
	"UY",
	"BR",
	"FK",
	"JM",
	"DO",
	"CU",
	"MQ",
	"BS",
	"BM",
	"AI",
	"TT",
	"KN",
	"DM",
	"AG",
	"LC",
	"TC",
	"AW",
	"VG",
	"VC",
	"MS",
	"MF",
	"BL",
	"GP",
	"GD",
	"KY",
	"BZ",
	"SV",
	"GT",
	"HN",
	"NI",
	"CR",
	"VE",
	"EC",
	"CO",
	"PA",
	"HT",
	"AR",
	"CL",
	"BO",
	"PE",
	"MX",
	"PF",
	"PN",
	"KI",
	"TO",
	"WF",
	"WS",
	"MP",
	"GU",
	"PR",
	"VI",
	"UM",
	"AS",
	"CA",
	"US",
	"PS",
	"RS",
	"AQ",
	"SX",
	"CW",
	"BQ",
	"SS",
	"TK",
}

// Создание портов
func (api *Api) CreatePorts(req CreatePortsReq) (*CreateProxyRsp, error) {
	// Данные для POST-запроса
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// Используем метод doPost для создания портов
	m, e := api.doPost("proxy/create-port", jsonData)
	if e != nil {
		return nil, e
	}
	r := &CreateProxyRsp{}
	err = json.Unmarshal(m, &r)
	if err != nil {
		return nil, err
	}

	return r, err
}

// Удаление портов
func (api *Api) DeleteProxy(id int) (bool, error) {
	m, e := api.doDelete("proxy/delete-port", map[string]string{
		"id": fmt.Sprintf("%d", id),
	})
	if e != nil {
		return false, e
	}
	r := DeleteProxyRsp{}

	// fmt.Println(string(m))
	err := json.Unmarshal(m, &r)
	if err != nil {
		return false, err
	}
	return r.Succes, nil
}

// Получение списка стран
func (api *Api) DirCountries() (CountriesRsp, error) {
	b, err := api.doGet("dir/countries")
	if err != nil {
		return CountriesRsp{}, err
	}
	r := CountriesRsp{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return CountriesRsp{}, err
	}
	return r, nil
}
