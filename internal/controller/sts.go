package controller

import (
	"PhotonTrail-backend/internal/global/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetStsToken(c *gin.Context) {
	url := "http://sts:8000/get_post_signature_for_oss_upload"

	stsReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create request failed", err.Error())
		return
	}
	stsReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(stsReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "send request failed", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "read response body failed", err.Error())
		return
	}
	var stsData response.StsData
	if err := json.Unmarshal(body, &stsData); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "json unmarshal failed, body: "+string(body), err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"sts": stsData}, "get sts success")
}
