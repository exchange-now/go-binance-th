package binanceth

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type symbolTypeServiceTestSuite struct {
	baseTestSuite
}

func TestSymbolTypeService(t *testing.T) {
	suite.Run(t, new(symbolTypeServiceTestSuite))
}

func (s *symbolTypeServiceTestSuite) TestSymbolType() {
	data := []byte(`{"type":"GLOBAL"}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSymbolTypeService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("GLOBAL", res.Type)
}
