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
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}
	// Используем метод doPost для создания портов
	m, err := api.doPost("proxy/create-port", jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to do Post: %w", err)
	}
	r := &CreateProxyRsp{}
	err = json.Unmarshal(m, &r)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if r.Success {
		return r, nil
	}

	return r, fmt.Errorf("failed to create ports: %s", r.Message)
}

// Удаление портов
func (api *Api) DeleteProxy(id int) (bool, error) {
	m, err := api.doDelete("proxy/delete-port", map[string]string{
		"id": fmt.Sprintf("%d", id),
	})
	if err != nil {
		return false, fmt.Errorf("failed to do delete: %w", err)
	}
	r := DeleteProxyRsp{}

	// fmt.Println(string(m))
	err = json.Unmarshal(m, &r)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	if r.Success {
		return r.Success, nil
	}
	return r.Success, fmt.Errorf("failed to delete proxy: %s", r.Message)

}

// Получение списка стран
func (api *Api) DirCountries() (CountriesRsp, error) {
	b, err := api.doGet("dir/countries")
	if err != nil {
		return CountriesRsp{}, fmt.Errorf("failed to do Get: %w", err)
	}
	r := CountriesRsp{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return CountriesRsp{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return r, nil
}

func (api *Api) ListPorts() (ListPortsRsp, error) {
	b, err := api.doGet("proxy/ports")
	if err != nil {
		return ListPortsRsp{}, err
	}
	r := ListPortsRsp{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return ListPortsRsp{}, err
	}
	return r, nil
}
