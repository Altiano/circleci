package apiv2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ContextItem struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"createdAt"`
}

type ContextResp struct {
	Items []ContextItem `json:"items" `
}

func (api API) ListContext(nextPageToken string) (ContextResp, error) {
	url := fmt.Sprintf("%s/context?owner-slug=%s&page-token=%s", api.baseURL, api.ownerSlug, nextPageToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ContextResp{}, err
	}

	req.Header.Add("Circle-Token", api.token)

	res, err := api.httpClient.Do(req)
	if err != nil {
		return ContextResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return ContextResp{}, errors.New("not ok: " + string(body))
	}

	result := ContextResp{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return ContextResp{}, err
	}

	return result, nil
}
