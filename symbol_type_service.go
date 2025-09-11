package binanceth

import (
	"context"
	"encoding/json"
	"net/http"
)

// SymbolTypeInfo represents a symbol and its type
type SymbolTypeInfo struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"` // "GLOBAL" or "SITE", use SymbolTypeGlobal or SymbolTypeSite constants for comparison
}

// SymbolTypeService checks symbol type (GLOBAL or SITE)
type SymbolTypeService struct {
	c *Client
}

// Do send request
func (s *SymbolTypeService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolTypeInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v1/symbolType",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
