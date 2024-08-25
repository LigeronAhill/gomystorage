package gomystorage

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/LigeronAhill/gomystorage/models"
)

func (c ApiClient) SearchEntity(entity models.Entity, search string) ([]byte, error) {
	uri := c.baseUrl.JoinPath(entity.Endpoint())
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Add("search", search)
	req.Header.Add(http.CanonicalHeaderKey("Authorization"), c.token)
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
