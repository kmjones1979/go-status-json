package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	//"sync"
	"time"
)

type StatusJSON struct {
	//Address     string   `json:"address"`
	Caches      struct{} `json:"caches"`
	Connections struct {
		Accepted int `json:"accepted"`
		Active   int `json:"active"`
		Dropped  int `json:"dropped"`
		Idle     int `json:"idle"`
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

	var status_json string = "http://demo.nginx.com/status"

	// load status json
	x, err := http.Get(status_json)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer x.Body.Close()

	x_dec := json.NewDecoder(x.Body)

	//	for {
	//		var x_data StatusJSON
	//		if err := x_dec.Decode(&x_data); err == io.EOF {
	//			break
	//		} else if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		//fmt.Println(x_data.Connections.Accepted)
	//		x_ca := x_data.Connections.Accepted
	//
	//	}

	// sleep 1 second
	time.Sleep(time.Millisecond * 5000)

	// load status json into second variable
	y, err := http.Get(status_json)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer y.Body.Close()

	y_dec := json.NewDecoder(y.Body)

	//	for {
	//		var y_data StatusJSON
	//		if err := y_dec.Decode(&y_data); err == io.EOF {
	//			break
	//		} else if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		//fmt.Println(y_data.Connections.Accepted)
	//		y_ca := y_data.Connections.Accepted
	//
	//	}

	for {

		var x_data StatusJSON
		if err := x_dec.Decode(&x_data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(x_data.Connections.Accepted)
		x_ca := x_data.Connections.Accepted

		var y_data StatusJSON
		if err := y_dec.Decode(&y_data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(y_data.Connections.Accepted)
		y_ca := y_data.Connections.Accepted

		z := (y_ca - x_ca)
		fmt.Println(z)

	}
}
