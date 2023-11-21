package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/helper"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient = &http.Client{}

func LoginRest(ctx *gin.Context) {

	address, port := Find("USER-MANAGEMENT")
	url := fmt.Sprintf("http://%s:%d/login/rest", address, port)
	req, err := http.NewRequest("POST", url, ctx.Request.Body)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	ctx.JSON(http.StatusCreated, req)

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()
	resp.Cookies()
	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	for _, cookie := range resp.Cookies() {
		ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	}
	//ctx.JSON(http.StatusOK, string(body))
	ctx.Data(http.StatusOK, "application/json", body)

}

func Logout(ctx *gin.Context) {

	address, port := Find("USER-MANAGEMENT")
	url := fmt.Sprintf("http://%s:%d/user/logout", address, port)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	for _, cookie := range ctx.Request.Cookies() {
		req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")

	ctx.JSON(http.StatusCreated, req)

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()
	resp.Cookies()
	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	for _, cookie := range resp.Cookies() {
		ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	}
	//ctx.JSON(http.StatusOK, string(body))
	ctx.Data(http.StatusOK, "application/json", body)

}

func GetUserInfo(ctx *gin.Context) {

	address, port := Find("USER-MANAGEMENT")
	url := fmt.Sprintf("http://%s:%d/facade/users/get-user-information", address, port)

	var accessToken string
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name == "access-token" {
			accessToken = cookie.Value
		}
		//ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	}
	userName, clientId := helper.Extract(accessToken)
	// Set headers
	reqBody := fmt.Sprintf("{\"username\":\"%s\",\"clientId\":\"%s\"}", userName, clientId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))

	for _, cookie := range ctx.Request.Cookies() {
		req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}

	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()
	resp.Cookies()
	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	for _, cookie := range resp.Cookies() {
		ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	}
	//ctx.JSON(http.StatusOK, string(body))
	ctx.Data(http.StatusOK, "application/json", body)

}

func CheckAccess(ctx *gin.Context) bool {

	address, port := Find("USER-MANAGEMENT")
	url := fmt.Sprintf("http://%s:%d/check-access", address, port)

	var checkAccessReq models.CheckAccessRequest
	checkAccessReq.Url = ctx.Request.URL.Path

	reqBody := fmt.Sprintf("{\"url\":\"%s\"}", checkAccessReq.Url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))

	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	for _, cookie := range ctx.Request.Cookies() {
		req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
		//ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	var checkAccess models.CheckAccess
	getJson(resp, &checkAccess)

	if checkAccess.ResponseCode == 0 {
		return true
	}

	return false
}
func getJson(res *http.Response, target interface{}) error {

	if err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(&target)
}
