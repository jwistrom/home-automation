package client

import (
	"encoding/json"
	"github.com/jwistrom/home-automation/tv-background/types"
	"io"
	"log"
	"net/http"
	"strconv"
)

type TvBackgroundLightClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewTvBackgroundLightClient(ip string, port int) *TvBackgroundLightClient {
	return &TvBackgroundLightClient{
		baseUrl:    "http://" + ip + ":" + strconv.Itoa(port) + "/",
		httpClient: &http.Client{},
	}
}

func (client *TvBackgroundLightClient) GetStatus() (types.Status, error) {
	log.Println("Sending GET request to " + client.baseUrl)
	response, err := http.Get(client.baseUrl)
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return types.Status{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return types.Status{}, err
	}

	var result types.Status
	err = json.Unmarshal(body, &result)
	if err != nil {
		return types.Status{}, err
	}

	return result, nil

}

func (client *TvBackgroundLightClient) SwitchState() (types.SwitchState, error) {
	response, err := client.doPut(client.baseUrl + "state")
	if err != nil {
		return false, err
	}

	responseBody, err := strconv.ParseBool(response)
	if err != nil {
		return false, err
	}

	return types.SwitchState(responseBody), nil

}

func (client *TvBackgroundLightClient) SetMode(mode types.TvBackgroundLightMode) error {
	_, err := client.doPut(client.baseUrl + "mode?mode=" + strconv.Itoa(int(mode)))
	if err != nil {
		return err
	}

	return nil
}

func (client *TvBackgroundLightClient) SetSpeed(speed int) error {
	_, err := client.doPut(client.baseUrl + "speed?speed=" + strconv.Itoa(speed))
	if err != nil {
		return err
	}

	return nil
}

func (client *TvBackgroundLightClient) doPut(path string) (string, error) {
	request, err := http.NewRequest(http.MethodPut, path, nil)
	if err != nil {
		return "", err
	}

	log.Println("Sending PUT request to " + request.URL.String())
	response, err := client.httpClient.Do(request)

	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return "", err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}
