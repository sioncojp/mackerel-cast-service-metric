package mcsm

import (
	"bytes"
	"fmt"
	"net/http"
)

// Data ...This type is for POST data.
type Data struct {
	Name  string
	Value float64
	Time  int64
}

// postURLString ...Output mackerel url with servicename.
func postURLString(sn string) string {
	return fmt.Sprintf("https://mackerel.io/api/v0/services/%s/tsdb", sn)
}

// jsonString ...Output json string to use POST.
func (d *Data) jsonString() string {
	return fmt.Sprintf(`[{"name": "%s", "value": %f, "time": %d}]`, d.Name, d.Value, d.Time)
}

// HTTPPost ...Post mackerel API.
func (config *Config) HTTPPost(d Data, url string) error {
	jsonStr := []byte(d.jsonString())

	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonStr),
	)

	req.Header.Set("X-Api-Key", config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("http request: %d", resp.StatusCode)
}
