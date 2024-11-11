package asocksapi

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
)

const (
	baseApiURL = "https://api.asocks.com/v2/"
)

type Api struct {
	ApiKey string
	Client *http.Client
}

// Конструктор для создания нового API-клиента
func NewApi(apiKey string) *Api {

	return &Api{
		ApiKey: apiKey,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (api *Api) doDelete(cmd string, params map[string]string) ([]byte, error) {

	// Создаем параметры запроса
	queryParams := url.Values{}
	queryParams.Add("apiKey", api.ApiKey)
	for k, v := range params {
		queryParams.Add(k, v)
	}

	// Формируем URL
	finalURL := baseApiURL + cmd + "?" + queryParams.Encode()

	req, _ := http.NewRequest("DELETE", finalURL, nil)
	// Выполняем GET-запрос
	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return body, err
}

// Метод для выполнения POST-запросов
func (api *Api) doPost(cmd string, body []byte) ([]byte, error) {

	// Создаем параметры запроса
	queryParams := url.Values{}
	queryParams.Add("apiKey", api.ApiKey)

	// Выполняем POST-запрос
	resp, err := api.Client.Post(baseApiURL+cmd+"?"+queryParams.Encode(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	return body, err
}

func (api *Api) doGet(cmd string) ([]byte, error) {

	// Создаем параметры запроса
	queryParams := url.Values{}
	queryParams.Add("apiKey", api.ApiKey)

	// Выполняем POST-запрос
	resp, err := api.Client.Get(baseApiURL + cmd + "?" + queryParams.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return body, err
}
