package coingecko

import "encoding/json"

type SearchAPI struct {
	client *Client
}

func (s *SearchAPI) Search(query string) (map[string]interface{}, error) {
	params := map[string]string{
		"query": query,
	}

	data, err := s.client.get("/search", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SearchAPI) Trending() (map[string]interface{}, error) {
	data, err := s.client.get("/search/trending", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
