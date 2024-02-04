package atapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func GetWithBearer[T any](tokenbearer string, urltarget string) (result T, errormessage string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", urltarget, nil)
	if err != nil {
		errormessage = "http.NewRequest Got error : " + err.Error()
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request : " + err.Error()
		return
	}
	if er := json.Unmarshal(respBody, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}

func GetImageWithBearer(tokenbearer string, urltarget string) (respBody []byte, errormessage string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", urltarget, nil)
	if err != nil {
		errormessage = "http.NewRequest Got error : " + err.Error()
		return
	}
	req.Header.Add("Content-Type", "image/jpeg")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "io.ReadAll(resp.Body) Error occured. Error is :" + err.Error()
		return
	}
	return
}

func GetStructWithBearer[T any](tokenbearer string, structname interface{}, urltarget string) (result T, errormessage string) {
	client := http.Client{}
	v, _ := query.Values(structname)
	req, err := http.NewRequest("GET", urltarget+"?"+v.Encode(), nil)
	if err != nil {
		errormessage = "http.NewRequest Got error : " + err.Error()
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request : " + err.Error()
		return
	}
	if er := json.Unmarshal(respBody, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}

func GetStructWithToken[T any](tokenkey string, tokenvalue string, structname interface{}, urltarget string) (result T, errormessage string) {
	client := http.Client{}
	v, _ := query.Values(structname)
	req, err := http.NewRequest("GET", urltarget+"?"+v.Encode(), nil)
	if err != nil {
		errormessage = "http.NewRequest Got error : " + err.Error()
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(tokenkey, tokenvalue)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request : " + err.Error()
		return
	}
	if er := json.Unmarshal(respBody, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}

func PutStructWithBearer[T any](tokenbearer string, structname interface{}, urltarget string) (result T, errormessage string) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("PUT", urltarget, bytes.NewBuffer(mJson))
	if err != nil {
		errormessage = "http.NewRequest Got error :" + err.Error()
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request : " + err.Error()
		return
	}
	if er := json.Unmarshal(respBody, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}

func PostStructWithBearer[T any](tokenbearer string, structname interface{}, urltarget string) (result T, errormessage string) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("POST", urltarget, bytes.NewBuffer(mJson))
	if err != nil {
		errormessage = "http.NewRequest Got error :" + err.Error()
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request." + err.Error()
		return
	}
	if er := json.Unmarshal(respBody, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}

func PostStructWithToken[T any](tokenkey string, tokenvalue string, structname interface{}, urltarget string) (result T, errormessage string) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("POST", urltarget, bytes.NewBuffer(mJson))
	if err != nil {
		errormessage = "http.NewRequest Got error :" + err.Error()
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(tokenkey, tokenvalue)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request." + err.Error()
		return
	}
	if er := json.Unmarshal(respBody, &result); er != nil {
		errormessage = string(respBody) + "Error Unmarshal from Response : " + er.Error()
	}
	return
}

func PostStruct[T any](structname interface{}, urltarget string) (result T, errormessage string) {
	mJson, _ := json.Marshal(structname)
	resp, err := http.Post(urltarget, "application/json", bytes.NewBuffer(mJson))
	if err != nil {
		errormessage = "Could not make POST request to server : " + err.Error()
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request : " + err.Error()
		return
	}
	if er := json.Unmarshal(body, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}

func Get[T any](urltarget string) (result T, errormessage string) {
	resp, err := http.Get(urltarget)
	if err != nil {
		errormessage = err.Error()
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read data from Response." + err.Error()
		return
	}
	if er := json.Unmarshal(body, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}

func GetStruct[T any](structname interface{}, urltarget string) (result T, errormessage string) {
	v, _ := query.Values(structname)
	resp, err := http.Get(urltarget + "?" + v.Encode())
	if err != nil {
		errormessage = "GetStruct http.get:" + err.Error()
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read data from Response." + err.Error()
		return
	}
	if er := json.Unmarshal(body, &result); er != nil {
		errormessage = "Error Unmarshal from Response." + err.Error()
	}
	return
}
