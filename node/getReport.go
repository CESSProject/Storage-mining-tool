/*
   Copyright 2022 CESS (Cumulus Encrypted Storage System) authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (n *Node) GetReport(c *gin.Context) {
	var (
		err error
	)
	val, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	var report Report

	str := strings.ReplaceAll(string(val), "\\", "")
	str = strings.TrimPrefix(str, "\"")
	str = strings.TrimSuffix(str, "\"")
	strs := strings.Split(str, "|")
	if len(strs) == 4 {
		report.Cert = strs[2]
		report.Ias_sig = strs[1]
		report.Quote = strs[0]
		report.Quote_sig = strs[3]
	}
	Ch_Report <- report
	c.JSON(http.StatusOK, nil)
	return
}

func GetReportReq(localIp string, sgxPort int) error {
	callbackurl := fmt.Sprintf("http://%v:%d/report", localIp, sgxPort)
	param := map[string]string{
		"callback_url": callbackurl,
	}
	data, err := json.Marshal(param)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:80/get_report", bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	cli := http.Client{
		Transport: globalTransport,
	}

	_, err = cli.Do(req)
	if err != nil {
		return err
	}

	return nil
}
