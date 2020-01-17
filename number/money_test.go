package number

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

func (suite *ToolsTestSuite) Test_ToPrecise() {
	assert.EqualValues(suite.T(), 10.123454, ToPrecise(10.123454321))
	assert.EqualValues(suite.T(), 10.123456, ToPrecise(10.123455555))
	assert.EqualValues(suite.T(), 10.123457, ToPrecise(10.123456789))
}

func (suite *ToolsTestSuite) Test_GetPercentPartFromAmount() {
	assert.EqualValues(suite.T(), 20, GetPercentPartFromAmount(120, 0.2))
	assert.EqualValues(suite.T(), 3, GetPercentPartFromAmount(33, 0.1))
}

func (suite *ToolsTestSuite) Test_FormatAmount() {
	assert.EqualValues(suite.T(), 10.12, FormatAmount(10.12))
	assert.EqualValues(suite.T(), 10.12, FormatAmount(10.123))
	assert.EqualValues(suite.T(), 10.12, FormatAmount(10.1234))
	assert.EqualValues(suite.T(), 10.12, FormatAmount(10.12345))
}

func (suite *ToolsTestSuite) Test_ToFixed() {
	assert.EqualValues(suite.T(), 0, ToFixed(0, 0))
	assert.EqualValues(suite.T(), 1, ToFixed(1, 0))
	assert.EqualValues(suite.T(), -1, ToFixed(-1, 0))

	assert.EqualValues(suite.T(), 19, ToFixed(19, 0))
	assert.EqualValues(suite.T(), 19, ToFixed(19.1, 0))
	assert.EqualValues(suite.T(), 19, ToFixed(19.19, 0))
	assert.EqualValues(suite.T(), 19, ToFixed(19.199, 0))
	assert.EqualValues(suite.T(), 19, ToFixed(19.1999, 0))

	assert.EqualValues(suite.T(), 19, ToFixed(19, 1))
	assert.EqualValues(suite.T(), 19.1, ToFixed(19.1, 1))
	assert.EqualValues(suite.T(), 19.1, ToFixed(19.19, 1))
	assert.EqualValues(suite.T(), 19.1, ToFixed(19.199, 1))
	assert.EqualValues(suite.T(), 19.1, ToFixed(19.1999, 1))

	assert.EqualValues(suite.T(), 19, ToFixed(19, 2))
	assert.EqualValues(suite.T(), 19.1, ToFixed(19.1, 2))
	assert.EqualValues(suite.T(), 19.19, ToFixed(19.19, 2))
	assert.EqualValues(suite.T(), 19.19, ToFixed(19.199, 2))
	assert.EqualValues(suite.T(), 19.19, ToFixed(19.1999, 2))
}
