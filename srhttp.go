package srhttp

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Client :
type Client struct {
	URL string
	//Method  string
	Headers map[string]string

	Payload io.Reader
}

// Get :
func (c *Client) Get() ([]byte, error) {
	body, err := doRequest("GET", c.URL, c.Headers)
	if err != nil {
		logrus.Fatal("Error: ", err)
		return nil, err
	}
	return body, nil
}

// Post :
func (c *Client) Post() ([]byte, error) {
	body, err := doRequest("POST", c.URL, c.Headers)
	if err != nil {
		logrus.Fatal("Error: ", err)
		return nil, err
	}
	return body, nil
}

func doRequest(method, url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		logrus.Fatal("Error: ", err)
		return nil, err
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)
	if err != nil {
		logrus.Fatal("Error: ", err)
		return nil, err
	}
	defer res.Body.Close()

	// if res.StatusCode != 200 {
	// 	logrus.Println("Error: did not get right response(200).")
	// 	return nil, nil
	// }

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Fatal("Error: ", err)
		return nil, err
	}
	return body, nil
}
