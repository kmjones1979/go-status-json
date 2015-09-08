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
	Address string `json:"address"`
	Caches  struct {
		Demo struct {
			Bypass struct {
				Bytes            int `json:"bytes"`
				BytesWritten     int `json:"bytes_written"`
				Responses        int `json:"responses"`
				ResponsesWritten int `json:"responses_written"`
			} `json:"bypass"`
			Cold    bool `json:"cold"`
			Expired struct {
				Bytes            int `json:"bytes"`
				BytesWritten     int `json:"bytes_written"`
				Responses        int `json:"responses"`
				ResponsesWritten int `json:"responses_written"`
			} `json:"expired"`
			Hit struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"hit"`
			MaxSize int `json:"max_size"`
			Miss    struct {
				Bytes            int `json:"bytes"`
				BytesWritten     int `json:"bytes_written"`
				Responses        int `json:"responses"`
				ResponsesWritten int `json:"responses_written"`
			} `json:"miss"`
			Revalidated struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"revalidated"`
			Size  int `json:"size"`
			Stale struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"stale"`
			Updating struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"updating"`
		} `json:"demo"`
		Empty struct {
			Bypass struct {
				Bytes            int `json:"bytes"`
				BytesWritten     int `json:"bytes_written"`
				Responses        int `json:"responses"`
				ResponsesWritten int `json:"responses_written"`
			} `json:"bypass"`
			Cold    bool `json:"cold"`
			Expired struct {
				Bytes            int `json:"bytes"`
				BytesWritten     int `json:"bytes_written"`
				Responses        int `json:"responses"`
				ResponsesWritten int `json:"responses_written"`
			} `json:"expired"`
			Hit struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"hit"`
			MaxSize int `json:"max_size"`
			Miss    struct {
				Bytes            int `json:"bytes"`
				BytesWritten     int `json:"bytes_written"`
				Responses        int `json:"responses"`
				ResponsesWritten int `json:"responses_written"`
			} `json:"miss"`
			Revalidated struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"revalidated"`
			Size  int `json:"size"`
			Stale struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"stale"`
			Updating struct {
				Bytes     int `json:"bytes"`
				Responses int `json:"responses"`
			} `json:"updating"`
		} `json:"empty"`
	} `json:"caches"`
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
		Respawned int64 `json:"respawned"`
	} `json:"processes"`
	Requests struct {
		Current int64 `json:"current"`
		Total   int64 `json:"total"`
	} `json:"requests"`
	ServerZones struct {
		Idle struct {
			Processing int64 `json:"processing"`
			Received   int64 `json:"received"`
			Requests   int64 `json:"requests"`
			Responses  struct {
				OneXx   int64 `json:"1xx"`
				TwoXx   int64 `json:"2xx"`
				ThreeXx int64 `json:"3xx"`
				FourXx  int64 `json:"4xx"`
				FiveXx  int64 `json:"5xx"`
				Total   int64 `json:"total"`
			} `json:"responses"`
			Sent int64 `json:"sent"`
		} `json:"idle"`
		One struct {
			Processing int64 `json:"processing"`
			Received   int64 `json:"received"`
			Requests   int64 `json:"requests"`
			Responses  struct {
				OneXx   int64 `json:"1xx"`
				TwoXx   int64 `json:"2xx"`
				ThreeXx int64 `json:"3xx"`
				FourXx  int64 `json:"4xx"`
				FiveXx  int64 `json:"5xx"`
				Total   int64 `json:"total"`
			} `json:"responses"`
			Sent int64 `json:"sent"`
		} `json:"one"`
		Three struct {
			Processing int64 `json:"processing"`
			Received   int64 `json:"received"`
			Requests   int64 `json:"requests"`
			Responses  struct {
				OneXx   int64 `json:"1xx"`
				TwoXx   int64 `json:"2xx"`
				ThreeXx int64 `json:"3xx"`
				FourXx  int64 `json:"4xx"`
				FiveXx  int64 `json:"5xx"`
				Total   int64 `json:"total"`
			} `json:"responses"`
			Sent int64 `json:"sent"`
		} `json:"three"`
		Two struct {
			Processing int64 `json:"processing"`
			Received   int64 `json:"received"`
			Requests   int64 `json:"requests"`
			Responses  struct {
				OneXx   int64 `json:"1xx"`
				TwoXx   int64 `json:"2xx"`
				ThreeXx int64 `json:"3xx"`
				FourXx  int64 `json:"4xx"`
				FiveXx  int64 `json:"5xx"`
				Total   int64 `json:"total"`
			} `json:"responses"`
			Sent int64 `json:"sent"`
		} `json:"two"`
	} `json:"server_zones"`
	Stream struct {
		ServerZones struct {
			Mysql_frontend struct {
				Connections int `json:"connections"`
				Processing  int `json:"processing"`
				Received    int `json:"received"`
				Sent        int `json:"sent"`
			} `json:"mysql-frontend"`
		} `json:"server_zones"`
		Upstreams struct {
			MysqlBackends []struct {
				Active        int  `json:"active"`
				Backup        bool `json:"backup"`
				ConnectTime   int  `json:"connect_time"`
				Connections   int  `json:"connections"`
				Downstart     int  `json:"downstart"`
				Downtime      int  `json:"downtime"`
				Fails         int  `json:"fails"`
				FirstByteTime int  `json:"first_byte_time"`
				HealthChecks  struct {
					Checks     int  `json:"checks"`
					Fails      int  `json:"fails"`
					LastPassed bool `json:"last_passed"`
					Unhealthy  int  `json:"unhealthy"`
				} `json:"health_checks"`
				ID           int    `json:"id"`
				Received     int    `json:"received"`
				ResponseTime int    `json:"response_time"`
				Selected     int    `json:"selected"`
				Sent         int    `json:"sent"`
				Server       string `json:"server"`
				State        string `json:"state"`
				Unavail      int    `json:"unavail"`
				Weight       int    `json:"weight"`
			} `json:"mysql_backends"`
		} `json:"upstreams"`
	} `json:"stream"`
	//Timestamp int `json:"timestamp"`
	Upstreams struct {
		Appservers []struct {
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
			MaxConns  int `json:"max_conns"`
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
		} `json:"appservers"`
		Demoupstreams []struct {
			Active       int  `json:"active"`
			Backup       bool `json:"backup"`
			Downstart    int  `json:"downstart"`
			Downtime     int  `json:"downtime"`
			Fails        int  `json:"fails"`
			HealthChecks struct {
				Checks     int  `json:"checks"`
				Fails      int  `json:"fails"`
				LastPassed bool `json:"last_passed"`
				Unhealthy  int  `json:"unhealthy"`
			} `json:"health_checks"`
			ID        int `json:"id"`
			MaxConns  int `json:"max_conns"`
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
		} `json:"demoupstreams"`
	} `json:"upstreams"`
	//Version int `json:"version"`
}

func main() {

	for {

		client, err := statsd.NewClient("127.0.0.1:8125", "nginx")
		// handle any errors
		if err != nil {
			log.Fatal(err)
		}

		var status_json string = "http://demo.nginx.com/status"

		// request status json from NGINX Plus
		x, err := http.Get(status_json)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}

		x_dec := json.NewDecoder(x.Body)

		// sleep x seconds
		time.Sleep(time.Millisecond * 1000)

		// re-request json from NGINX Plus
		y, err := http.Get(status_json)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}

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

		// processes
		ngx_proc := (y_data.Processes.Respawned)
		fmt.Println("status.demo.processes.respawned", ngx_proc)
		client.Inc("status.demo.processes.respawned", ngx_proc, 1.0)

		// connections
		ngx_ca := (y_data.Connections.Accepted - x_data.Connections.Accepted)
		fmt.Println("status.demo.connections.accepted", ngx_ca)
		client.Inc("status.demo.connections.accepted", ngx_ca, 1.0)

		ngx_cd := (y_data.Connections.Dropped - x_data.Connections.Dropped)
		fmt.Println("status.demo.connections.dropped", ngx_cd)
		client.Inc("status.demo.connections.dropped", ngx_cd, 1.0)

		ngx_can := (y_data.Connections.Active)
		fmt.Println("status.demo.connections.active", ngx_can)
		client.Inc("status.demo.connections.active", ngx_can, 1.0)

		ngx_cai := (y_data.Connections.Idle)
		fmt.Println("status.demo.connections.idle", ngx_cai)
		client.Inc("status.demo.connections.idle", ngx_cai, 1.0)

		// requests
		ngx_req := (y_data.Requests.Current)
		fmt.Println("status.demo.requests.current", ngx_req)
		client.Inc("status.demo.requests.current", ngx_req, 1.0)

		ngx_reqt := (y_data.Requests.Total - x_data.Requests.Total)
		fmt.Println("status.demo.requests.total", ngx_reqt)
		client.Inc("status.demo.requests.total", ngx_reqt, 1.0)

		// server zones
		ngx_sz_oproc := (y_data.ServerZones.One.Processing)
		fmt.Println("status.demo.serverzones.one.processing", ngx_sz_oproc)
		client.Inc("status.demo.serverzones.one.processing", ngx_sz_oproc, 1.0)

		ngx_sz_orec := (y_data.ServerZones.One.Received - x_data.ServerZones.One.Received)
		fmt.Println("status.demo.serverzones.one.received", ngx_sz_orec)
		client.Inc("status.demo.serverzones.one.received", ngx_sz_orec, 1.0)

		ngx_sz_osent := (y_data.ServerZones.One.Sent - x_data.ServerZones.One.Sent)
		fmt.Println("status.demo.serverzones.one.sent", ngx_sz_osent)
		client.Inc("status.demo.serverzones.one.sent", ngx_sz_osent, 1.0)

		ngx_sz_oreq := (y_data.ServerZones.One.Requests - x_data.ServerZones.One.Requests)
		fmt.Println("status.demo.serverzones.one.requests", ngx_sz_oreq)
		client.Inc("status.demo.serverzones.one.requests", ngx_sz_oreq, 1.0)

		ngx_sz_oresp_t := (y_data.ServerZones.One.Responses.Total - x_data.ServerZones.One.Responses.Total)
		fmt.Println("status.demo.serverzones.one.responses.total", ngx_sz_oresp_t)
		client.Inc("status.demo.serverzones.one.responses.total", ngx_sz_oresp_t, 1.0)

		ngx_sz_oresp_1xx := (y_data.ServerZones.One.Responses.OneXx - x_data.ServerZones.One.Responses.OneXx)
		fmt.Println("status.demo.serverzones.one.responses.1xx", ngx_sz_oresp_1xx)
		client.Inc("status.demo.serverzones.one.responses.1xx", ngx_sz_oresp_1xx, 1.0)

		ngx_sz_oresp_2xx := (y_data.ServerZones.One.Responses.TwoXx - x_data.ServerZones.One.Responses.TwoXx)
		fmt.Println("status.demo.serverzones.one.responses.2xx", ngx_sz_oresp_2xx)
		client.Inc("status.demo.serverzones.one.responses.2xx", ngx_sz_oresp_2xx, 1.0)

		ngx_sz_oresp_3xx := (y_data.ServerZones.One.Responses.ThreeXx - x_data.ServerZones.One.Responses.ThreeXx)
		fmt.Println("status.demo.serverzones.one.responses.3xx", ngx_sz_oresp_3xx)
		client.Inc("status.demo.serverzones.one.responses.3xx", ngx_sz_oresp_3xx, 1.0)

		ngx_sz_oresp_4xx := (y_data.ServerZones.One.Responses.FourXx - x_data.ServerZones.One.Responses.FourXx)
		fmt.Println("status.demo.serverzones.one.responses.4xx", ngx_sz_oresp_4xx)
		client.Inc("status.demo.serverzones.one.responses.4xx", ngx_sz_oresp_4xx, 1.0)

		ngx_sz_oresp_5xx := (y_data.ServerZones.One.Responses.FiveXx - x_data.ServerZones.One.Responses.FiveXx)
		fmt.Println("status.demo.serverzones.one.responses.5xx", ngx_sz_oresp_5xx)
		client.Inc("status.demo.serverzones.one.responses.5xx", ngx_sz_oresp_5xx, 1.0)

		ngx_sz_oresp_err_perc := (ngx_sz_oresp_4xx + ngx_sz_oresp_5xx/ngx_sz_oresp_t)
		fmt.Println("status.demo.serverzones.one.responses.error_percentage", ngx_sz_oresp_err_perc)
		client.Inc("status.demo.serverzones.one.responses.error_percentage", ngx_sz_oresp_err_perc, 1.0)

		// close any open connections before loop restarts
		// close any open connections before loop restarts
		client.Close()
		x.Body.Close()
		y.Body.Close()

	}

}
