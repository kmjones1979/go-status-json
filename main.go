package main

import (
	"encoding/json"
	"fmt"
	"github.com/cactus/go-statsd-client/statsd"
	"io"
	"log"
	"net/http"
	"time"
)

type StatusJSON struct {
	//Address     string   `json:"address"`
	Caches      struct{} `json:"caches"`
	Connections struct {
		Accepted int64 `json:"accepted"`
		Active   int64 `json:"active"`
		Dropped  int64 `json:"dropped"`
		Idle     int64 `json:"idle"`
	} `json:"connections"`
	//Generation    int    `json:"generation"`
	//LoadTimestamp int    `json:"load_timestamp"`
	//NginxVersion  string `json:"nginx_version"`
	Processes struct {
		Respawned int `json:"respawned"`
	} `json:"processes"`
	Requests struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"requests"`
	ServerZones struct{} `json:"server_zones"`
	Timestamp   int      `json:"timestamp"`
	Upstreams   struct {
		Backend []struct {
			Active       int  `json:"active"`
			Backup       bool `json:"backup"`
			Downstart    int  `json:"downstart"`
			Downtime     int  `json:"downtime"`
			Fails        int  `json:"fails"`
			HealthChecks struct {
				Checks    int `json:"checks"`
				Fails     int `json:"fails"`
				Unhealthy int `json:"unhealthy"`
			} `json:"health_checks"`
			ID        int `json:"id"`
			Received  int `json:"received"`
			Requests  int `json:"requests"`
			Responses struct {
				OneXx   int `json:"1xx"`
				TwoXx   int `json:"2xx"`
				ThreeXx int `json:"3xx"`
				FourXx  int `json:"4xx"`
				FiveXx  int `json:"5xx"`
				Total   int `json:"total"`
			} `json:"responses"`
			Selected int    `json:"selected"`
			Sent     int    `json:"sent"`
			Server   string `json:"server"`
			State    string `json:"state"`
			Unavail  int    `json:"unavail"`
			Weight   int    `json:"weight"`
		} `json:"backend"`
	} `json:"upstreams"`
	//Version int `json:"version"`
}

func main() {

	for {
		var status_json string = "http://demo.nginx.com/status"

		// first create a client
		client, err := statsd.NewClient("127.0.0.1:8125", "nginx")
		// handle any errors
		if err != nil {
			log.Fatal(err)
		}
		// make sure to clean up
		defer client.Close()

		// request status json from NGINX Plus
		x, err := http.Get(status_json)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		defer x.Body.Close()

		x_dec := json.NewDecoder(x.Body)

		// sleep x seconds
		time.Sleep(time.Millisecond * 5000)

		// re-request json from NGINX Plus
		y, err := http.Get(status_json)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		defer y.Body.Close()

		y_dec := json.NewDecoder(y.Body)

		// loop through both to get diff

		var x_data StatusJSON
		if err := x_dec.Decode(&x_data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		var y_data StatusJSON
		if err := y_dec.Decode(&y_data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		ngx_ca := (y_data.Connections.Accepted - x_data.Connections.Accepted)
		fmt.Println("status.demo.connections.accepted", ngx_ca)
		client.Inc("status.demo.connections.accepted", ngx_ca, 5.0)

		ngx_cd := (y_data.Connections.Dropped - x_data.Connections.Dropped)
		fmt.Println("status.demo.connections.dropped", ngx_cd)
		client.Inc("status.demo.connections.dropped", ngx_cd, 5.0)

		ngx_can := (y_data.Connections.Active)
		fmt.Println("status.demo.connections.active", ngx_can)
		client.Inc("status.demo.connections.active", ngx_can, 5.0)

		ngx_cai := (y_data.Connections.Idle)
		fmt.Println("status.demo.connections.idle", ngx_cai)
		client.Inc("status.demo.connections.idle", ngx_cai, 5.0)
	}
}
