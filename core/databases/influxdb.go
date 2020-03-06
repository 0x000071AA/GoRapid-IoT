package databases

import (
	"fmt"
	"net/url"
	"time"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

func InfluxDbClient(host string, port int) client.Client {
	var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond

	addr, errURL := url.Parse(fmt.Sprintf("http://%s:%s", host, port))

	if errURL != nil {
		fmt.Println("Error parsing url: ", errURL.Error())
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr.String(),
		Username: "",
		Password: "",
	})

	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}

	_, _, errPing := c.Ping(waitFiveHundredMillisections)

	if errPing != nil {
		fmt.Println("Could not ping InfluxDB Client: ", errPing.Error())
	}

	return c
}

func Query(c client.Client, queryString string) ([]client.Result, error, error) {

	q := client.NewQuery(queryString, "mydb", "")
	response, err := c.Query(q)
	if err != nil && response.Error() != nil {
		return nil, err, response.Error()
	}
	return response.Results, nil, nil
}

func Write(c client.Client) {

}

func CloseConnection(c client.Client) {
	c.Close()
}
