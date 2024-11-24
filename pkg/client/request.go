package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/rs/zerolog/log"
)

func doRequest[T any](c *ApiClient, method, urlPath string, body interface{}) (*ApiResponse[T], error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(urlPath, "?")
	u.Path = path.Join(u.Path, parts[0])
	if len(parts) > 1 {
		u.RawQuery = parts[1]
	}

	url := u.String()
	log.Debug().Str("url", url).Msg("Request")
	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}

	if c.authorizetionHeader != "" {
		req.Header.Set("Authorization", c.authorizetionHeader)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp ApiResponse[T]

	if resp.StatusCode < 400 {
		if err := json.NewDecoder(resp.Body).Decode(&apiResp.Result); err != nil {
			return nil, err
		}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&apiResp.Error); err != nil {
			return nil, err
		}
	}

	// if resp.StatusCode >= 400 || apiResp.Error != nil {
	// 	if apiResp.Error != nil {
	// 		if apiResp.Error.Details != "" {
	// 			return &apiResp, fmt.Errorf("%s: %s", apiResp.Error.Error, apiResp.Error.Details)
	// 		}
	// 		return &apiResp, fmt.Errorf("%s", apiResp.Error.Error)
	// 	}
	// 	return &apiResp, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	// }

	return &apiResp, nil
}
