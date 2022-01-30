package azkaban

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Server string
}

type StatusResp struct {
	Version           string
	Pid               string
	InstallationPath  string
	UsedMemory        int64
	Xmx               int64
	IsDatabaseUp      bool
	ExecutorStatusMap map[string]struct {
		Id       int32
		Host     string
		Port     int16
		IsActive bool
	}
}

func (c *Client) Status() (*StatusResp, error) {

	/*
		raw := `{
			"version" : "3.91.0-134-g68e7c718",
			"pid" : "19233",
			"installationPath" : "/data/azkaban/azkaban-web-server/lib/azkaban-web-server-3.91.0-134-g68e7c718.jar",
			"usedMemory" : 260057872,
			"xmx" : 3817865216,
			"isDatabaseUp" : true,
			"executorStatusMap" : {
			  "34" : {
				"id" : 34,
				"host" : "header-01",
				"port" : 12321,
				"isActive" : true
			  },
			  "35" : {
				"id" : 35,
				"host" : "header-02",
				"port" : 12321,
				"isActive" : true
			  }
			}
		  }`
	*/
	// todo 超时
	resp, err := http.Get(c.Server + "/status")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	statusResp := &StatusResp{}
	err = json.Unmarshal(respBody, statusResp)
	if err != nil {
		return nil, err
	}
	return statusResp, nil
}
