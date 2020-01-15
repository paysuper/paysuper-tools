package tools

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

func (suite *ToolsTestSuite) Test_ToFixed() {

	assert.EqualValues(suite.T(), ToFixed(0, 0), 0)
	assert.EqualValues(suite.T(), ToFixed(1, 0), 1)
	assert.EqualValues(suite.T(), ToFixed(-1, 0), -1)

	assert.EqualValues(suite.T(), ToFixed(19, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.1, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.19, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.199, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.1999, 0), 19)

	assert.EqualValues(suite.T(), ToFixed(19, 1), 19)
	assert.EqualValues(suite.T(), ToFixed(19.1, 1), 19.1)
	assert.EqualValues(suite.T(), ToFixed(19.19, 1), 19.1)
	assert.EqualValues(suite.T(), ToFixed(19.199, 1), 19.1)
	assert.EqualValues(suite.T(), ToFixed(19.1999, 1), 19.1)

	assert.EqualValues(suite.T(), ToFixed(19, 2), 19)
	assert.EqualValues(suite.T(), ToFixed(19.1, 2), 19.1)
	assert.EqualValues(suite.T(), ToFixed(19.19, 2), 19.19)
	assert.EqualValues(suite.T(), ToFixed(19.199, 2), 19.19)
	assert.EqualValues(suite.T(), ToFixed(19.1999, 2), 19.19)

	assert.EqualValues(suite.T(), ToFixed(19, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.5, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.59, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.599, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.5999, 0), 19)

	assert.EqualValues(suite.T(), ToFixed(19, 1), 19)
	assert.EqualValues(suite.T(), ToFixed(19.5, 1), 19.5)
	assert.EqualValues(suite.T(), ToFixed(19.59, 1), 19.5)
	assert.EqualValues(suite.T(), ToFixed(19.599, 1), 19.5)
	assert.EqualValues(suite.T(), ToFixed(19.5999, 1), 19.5)

	assert.EqualValues(suite.T(), ToFixed(19, 2), 19)
	assert.EqualValues(suite.T(), ToFixed(19.5, 2), 19.5)
	assert.EqualValues(suite.T(), ToFixed(19.59, 2), 19.59)
	assert.EqualValues(suite.T(), ToFixed(19.599, 2), 19.59)
	assert.EqualValues(suite.T(), ToFixed(19.5999, 2), 19.59)

	assert.EqualValues(suite.T(), ToFixed(19, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.9, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.99, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.999, 0), 19)
	assert.EqualValues(suite.T(), ToFixed(19.9999, 0), 19)

	assert.EqualValues(suite.T(), ToFixed(19, 1), 19)
	assert.EqualValues(suite.T(), ToFixed(19.9, 1), 19.9)
	assert.EqualValues(suite.T(), ToFixed(19.99, 1), 19.9)
	assert.EqualValues(suite.T(), ToFixed(19.999, 1), 19.9)
	assert.EqualValues(suite.T(), ToFixed(19.9999, 1), 19.9)

	assert.EqualValues(suite.T(), ToFixed(19, 2), 19)
	assert.EqualValues(suite.T(), ToFixed(19.9, 2), 19.9)
	assert.EqualValues(suite.T(), ToFixed(19.99, 2), 19.99)
	assert.EqualValues(suite.T(), ToFixed(19.999, 2), 19.99)
	assert.EqualValues(suite.T(), ToFixed(19.9999, 2), 19.99)
}
