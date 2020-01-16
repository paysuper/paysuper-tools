package string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
)

type ToolsTestSuite struct {
	suite.Suite
	log *zap.Logger
}

func Test_Tools(t *testing.T) {
	suite.Run(t, new(ToolsTestSuite))
}

func (suite *ToolsTestSuite) SetupTest() {
	var err error

	suite.log, err = zap.NewProduction()
	assert.NoError(suite.T(), err)
}

func (suite *ToolsTestSuite) Test_MaskBankCardNumber() {
	assert.EqualValues(suite.T(), "1234567890", MaskBankCardNumber("1234567890"))
	assert.EqualValues(suite.T(), "123456******1234", MaskBankCardNumber("1234560000001234"))
	assert.EqualValues(suite.T(), "123456*1234", MaskBankCardNumber("12345601234"))
}
