package debicred

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Constructor struct {
	Id     string
	PubKey string
	SecKey string
}

type Payload struct {
	Code   string  `json:"currency_code"`
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
}

type Response map[string]interface{}

const URL = "https://debicred-wallet-api.herokuapp.com/api/v1"

func Init(id, PubKey, SecKey string) Constructor {
	return Constructor{
		Id:     id,
		PubKey: PubKey,
		SecKey: SecKey,
	}
}

func (customer Constructor) CreateWallet() (Response, error) {

	bearer := fmt.Sprintf("Bearer %s", customer.SecKey)

	// create new request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/wallet", URL), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("customer_id", customer.Id)
	req.Header.Add("public-key", customer.PubKey)

	// send req using http client

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return nil, err
	}

	if response["error"] != false {
		fmt.Println(response)
		return nil, errors.New("An Error Occured")
	}

	return response, nil
}

func (customer Constructor) CreditWallet(payload *Payload, walletId string) (Response, error) {

	postBody, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	respBody := bytes.NewBuffer(postBody)

	url := fmt.Sprintf("%s/wallet/credit/%s", URL, walletId)

	req, err := http.NewRequest("POST", url, respBody)

	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", customer.SecKey)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("customer_id", customer.Id)
	req.Header.Add("public-key", customer.PubKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response["error"] != false {
		return nil, errors.New("An Error Occured")
	}

	return response, nil
}

func (customer Constructor) DebitWallet(payload *Payload, walletId string) (Response, error) {

	postBody, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	respBody := bytes.NewBuffer(postBody)

	url := fmt.Sprintf("%s/wallet/debit/%s", URL, walletId)

	req, err := http.NewRequest("POST", url, respBody)

	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", customer.SecKey)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("customer_id", customer.Id)
	req.Header.Add("public-key", customer.PubKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response["error"] != false {
		return nil, errors.New("An Error Occured")
	}

	return response, nil
}

func (customer Constructor) GetWallet(walletId string) (Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wallet/%s", URL, walletId), nil)
	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", customer.SecKey)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("customer_id", customer.Id)
	req.Header.Add("public-key", customer.PubKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response["error"] != false {
		return nil, errors.New("An Error Occured")
	}

	return response, nil
}

func (customer Constructor) GetCurrencies() (Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/currencies", URL), nil)
	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", customer.SecKey)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("customer_id", customer.Id)
	req.Header.Add("public-key", customer.PubKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response["error"] != false {
		return nil, errors.New("An Error Occured")
	}

	return response, nil
}

func (customer Constructor) GetWalletTransactions(walletId string) (Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wallet/transactions/%s", URL, walletId), nil)
	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", customer.SecKey)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("customer_id", customer.Id)
	req.Header.Add("public-key", customer.PubKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response["error"] != false {
		return nil, errors.New("An Error Occured")
	}

	return response, nil
}
