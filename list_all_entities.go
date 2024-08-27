package gomystorage

import (
	"gomystorage/models"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// ListAllEntities retrieves an entity from the API.
//
// entity - the entity to be retrieved, which must implement the models.Entity interface.
// Returns a byte slice containing the response body and an error if the request fails.
func (c ApiClient) ListAllEntities(entity models.Entity, limit, offset int) ([]byte, error) {
	uri := c.baseUrl.JoinPath(entity.Endpoint())
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}
	if limit < 1 {
		limit = 1
	}
	if limit > 1000 {
		limit = 1000
	}
	if offset < 0 {
		offset = 0
	}
	q := url.Values{}
	q.Add("limit", strconv.Itoa(limit))
	q.Add("offset", strconv.Itoa(offset))
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
