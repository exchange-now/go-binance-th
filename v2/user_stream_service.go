package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// StartUserStreamService create listen key for user stream service
type StartUserStreamService struct {
	c *Client
}

// Do send request
// https://www.binance.th/api-docs/en/?go#create-a-listenkey-user_stream
func (s *StartUserStreamService) Do(ctx context.Context, opts ...RequestOption) (siteListenKey, globalListenKey string, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v1/listenKey",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", "", err
	}

	//j, err := newJSON(data)
	//if err != nil {
	//	return "", "", err
	//}
	//listenKey = j.Get("listenKey").MustString()
	//return listenKey, nil
	type ResponseStruct []struct {
		ListenKey string `json:"listenKey"`
		Type      string `json:"type"`
	}
	var resp ResponseStruct
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return "", "", err
	}
	var siteKey, globalKey string
	for _, item := range resp {
		if item.Type == "SITE" {
			siteKey = item.ListenKey
		} else if item.Type == "GLOBAL" {
			globalKey = item.ListenKey
		} else {
			return "", "", fmt.Errorf("unknown listen key type: %s", item.Type)
		}
	}
	return siteKey, globalKey, nil
}

// KeepaliveUserStreamService update listen key
type KeepaliveUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveUserStreamService) ListenKey(listenKey string) *KeepaliveUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseUserStreamService delete listen key
type CloseUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseUserStreamService) ListenKey(listenKey string) *CloseUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}
