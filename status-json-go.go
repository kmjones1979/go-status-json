package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type StatusJson struct {
	Version       int    `json:"version"`
	NginxVersion  string `json:"nginx_version"`
	Address       string `json:"address"`
	Generation    int    `json:"generation"`
	LoadTimestamp int64  `json:"load_timestamp"`
	Timestamp     int64  `json:"timestamp"`
	Processes     struct {
		Respawned int `json:"respawned"`
	} `json:"processes"`
	Connections struct {
		Accepted int `json:"accepted"`
		Dropped  int `json:"dropped"`
		Active   int `json:"active"`
		Idle     int `json:"idle"`
	} `json:"connections"`
	Requests struct {
		Total   int `json:"total"`
		Current int `json:"current"`
	} `json:"requests"`
	ServerZones struct {
		StatusDashboard struct {
			Processing int `json:"processing"`
			Requests   int `json:"requests"`
			Responses  struct {
				OneXx   int `json:"1xx"`
				TwoXx   int `json:"2xx"`
				ThreeXx int `json:"3xx"`
				FourXx  int `json:"4xx"`
				FiveXx  int `json:"5xx"`
				Total   int `json:"total"`
			} `json:"responses"`
			Received int `json:"received"`
			Sent     int `json:"sent"`
		} `json:"status_dashboard"`
		NginxPlusExternal struct {
			Processing int `json:"processing"`
			Requests   int `json:"requests"`
			Responses  struct {
				OneXx   int `json:"1xx"`
				TwoXx   int `json:"2xx"`
				ThreeXx int `json:"3xx"`
				FourXx  int `json:"4xx"`
				FiveXx  int `json:"5xx"`
				Total   int `json:"total"`
			} `json:"responses"`
			Received int `json:"received"`
			Sent     int `json:"sent"`
		} `json:"nginx-plus_external"`
	} `json:"server_zones"`
	Upstreams struct {
		NginxPlus80 []struct {
			ID           int    `json:"id"`
			Server       string `json:"server"`
			Backup       bool   `json:"backup"`
			Weight       int    `json:"weight"`
			State        string `json:"state"`
			Active       int    `json:"active"`
			Requests     int    `json:"requests"`
			HeaderTime   int    `json:"header_time"`
			ResponseTime int    `json:"response_time"`
			Responses    struct {
				OneXx   int `json:"1xx"`
				TwoXx   int `json:"2xx"`
				ThreeXx int `json:"3xx"`
				FourXx  int `json:"4xx"`
				FiveXx  int `json:"5xx"`
				Total   int `json:"total"`
			} `json:"responses"`
			Sent         int `json:"sent"`
			Received     int `json:"received"`
			Fails        int `json:"fails"`
			Unavail      int `json:"unavail"`
			HealthChecks struct {
				Checks     int  `json:"checks"`
				Fails      int  `json:"fails"`
				Unhealthy  int  `json:"unhealthy"`
				LastPassed bool `json:"last_passed"`
			} `json:"health_checks"`
			Downtime  int   `json:"downtime"`
			Downstart int   `json:"downstart"`
			Selected  int64 `json:"selected"`
		} `json:"nginx-plus_80"`
	} `json:"upstreams"`
	Caches struct {
		Mycache struct {
			Size    int   `json:"size"`
			MaxSize int64 `json:"max_size"`
			Cold    bool  `json:"cold"`
			Hit     struct {
				Responses int `json:"responses"`
				Bytes     int `json:"bytes"`
			} `json:"hit"`
			Stale struct {
				Responses int `json:"responses"`
				Bytes     int `json:"bytes"`
			} `json:"stale"`
			Updating struct {
				Responses int `json:"responses"`
				Bytes     int `json:"bytes"`
			} `json:"updating"`
			Revalidated struct {
				Responses int `json:"responses"`
				Bytes     int `json:"bytes"`
			} `json:"revalidated"`
			Miss struct {
				Responses        int `json:"responses"`
				Bytes            int `json:"bytes"`
				ResponsesWritten int `json:"responses_written"`
				BytesWritten     int `json:"bytes_written"`
			} `json:"miss"`
			Expired struct {
				Responses        int `json:"responses"`
				Bytes            int `json:"bytes"`
				ResponsesWritten int `json:"responses_written"`
				BytesWritten     int `json:"bytes_written"`
			} `json:"expired"`
			Bypass struct {
				Responses        int `json:"responses"`
				Bytes            int `json:"bytes"`
				ResponsesWritten int `json:"responses_written"`
				BytesWritten     int `json:"bytes_written"`
			} `json:"bypass"`
		} `json:"mycache"`
	} `json:"caches"`
	Stream struct {
		ServerZones struct {
			Rabbitmq struct {
				Processing  int `json:"processing"`
				Connections int `json:"connections"`
				Received    int `json:"received"`
				Sent        int `json:"sent"`
			} `json:"rabbitmq"`
		} `json:"server_zones"`
		Upstreams struct {
			Rabbitmq7000 []struct {
				ID           int    `json:"id"`
				Server       string `json:"server"`
				Backup       bool   `json:"backup"`
				Weight       int    `json:"weight"`
				State        string `json:"state"`
				Active       int    `json:"active"`
				Connections  int    `json:"connections"`
				Sent         int    `json:"sent"`
				Received     int    `json:"received"`
				Fails        int    `json:"fails"`
				Unavail      int    `json:"unavail"`
				HealthChecks struct {
					Checks    int `json:"checks"`
					Fails     int `json:"fails"`
					Unhealthy int `json:"unhealthy"`
				} `json:"health_checks"`
				Downtime  int `json:"downtime"`
				Downstart int `json:"downstart"`
				Selected  int `json:"selected"`
			} `json:"rabbitmq_7000"`
		} `json:"upstreams"`
	} `json:"stream"`
}

func average(xs []float64) float64 {
	panic("Not Implemented")
}

func main() {
	var s_json string = "http://45.55.15.196:9000/status"
	response, err := http.Get(s_json)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	wg := new(sync.WaitGroup)

	for {
		var m StatusJson
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(m.ServerZones.StatusDashboard.Requests)
		var sz_sd_reqs int = m.ServerZones.StatusDashboard.Requests
		xs := float64(x)
		fmt.Println(xs)
	}
	wg.Wait()
}
