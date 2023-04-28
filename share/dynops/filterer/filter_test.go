package filterer_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/realvnc-labs/rport/share/dynops/filterer"
	"github.com/realvnc-labs/rport/share/query"
)

type FiltererTestSuite struct {
	suite.Suite
}

func (suite *FiltererTestSuite) SetupTest() {
	var err error
	suite.NoError(err)
}

func (suite *FiltererTestSuite) TestFilterer_nilOptionsShouldAlwaysBeTrue() {
	filter, err := filterer.CompileFromQueryListOptions[bool](nil)
	suite.NoError(err)
	suite.True(filter.Run(false))
}

func (suite *FiltererTestSuite) TestFilterer_nilFilterShouldBeTrue() {

	options := suite.genOptions(nil)

	filter, err := filterer.CompileFromQueryListOptions[bool](options)
	suite.NoError(err)
	suite.True(filter.Run(false))
}

func (suite *FiltererTestSuite) TestFilterer_emptyFilterShouldBeTrue() {

	options := suite.genOptions([]query.FilterOption{})

	filter, err := filterer.CompileFromQueryListOptions[bool](options)
	suite.NoError(err)
	suite.True(filter.Run(false))
}

type TestStruct struct {
	SomeField string
	SomeInt   int
}

func (suite *FiltererTestSuite) TestFilterer_simpleEquation() {

	options := suite.genOptions([]query.FilterOption{{
		Column:                []string{"SomeField"},
		Operator:              query.FilterOperatorTypeEQ,
		Values:                []string{"some-value"},
		ValuesLogicalOperator: "",
	}})

	filter, err := filterer.CompileFromQueryListOptions[TestStruct](options)
	suite.NoError(err)
	suite.True(filter.Run(TestStruct{SomeField: "some-value"}))
	suite.False(filter.Run(TestStruct{SomeField: "wrong-value"}))
}

func (suite *FiltererTestSuite) TestFilterer_testInt() {

	options := suite.genOptions([]query.FilterOption{{
		Column:                []string{"SomeInt"},
		Operator:              query.FilterOperatorTypeEQ,
		Values:                []string{"5"},
		ValuesLogicalOperator: "",
	}})

	filter, err := filterer.CompileFromQueryListOptions[TestStruct](options)
	suite.NoError(err)
	suite.True(filter.Run(TestStruct{SomeInt: 5}))
	suite.False(filter.Run(TestStruct{SomeInt: 3}))
}

func (suite *FiltererTestSuite) TestFilterer_testGt() {

	options := suite.genOptions([]query.FilterOption{{
		Column:                []string{"SomeInt"},
		Operator:              query.FilterOperatorTypeGT,
		Values:                []string{"15"},
		ValuesLogicalOperator: "",
	}})

	filter, err := filterer.CompileFromQueryListOptions[TestStruct](options)
	suite.NoError(err)
	suite.True(filter.Run(TestStruct{SomeInt: 3}))
	suite.False(filter.Run(TestStruct{SomeInt: 50}))
}

func (suite *FiltererTestSuite) TestFilterer_testMultiple() {

	options := suite.genOptions([]query.FilterOption{
		{
			Column:                []string{"SomeInt"},
			Operator:              query.FilterOperatorTypeEQ,
			Values:                []string{"5"},
			ValuesLogicalOperator: "",
		},
		{
			Column:                []string{"SomeField"},
			Operator:              query.FilterOperatorTypeEQ,
			Values:                []string{"some-value"},
			ValuesLogicalOperator: "",
		},
	})

	filter, err := filterer.CompileFromQueryListOptions[TestStruct](options)
	suite.NoError(err)
	suite.True(filter.Run(TestStruct{SomeInt: 5, SomeField: "some-value"}))
	suite.False(filter.Run(TestStruct{SomeInt: 5, SomeField: "other-value"}))
	suite.False(filter.Run(TestStruct{SomeInt: 3, SomeField: "some-value"}))
}

func (suite *FiltererTestSuite) TestFilterer_errorOnNonExistentField() {

	options := suite.genOptions([]query.FilterOption{
		{
			Column:                []string{"non-existent"},
			Operator:              query.FilterOperatorTypeEQ,
			Values:                []string{"5"},
			ValuesLogicalOperator: "",
		},
	})

	_, err := filterer.CompileFromQueryListOptions[TestStruct](options)
	suite.ErrorContains(err, "field: non-existent does not exist on type: filterer_test.TestStruct on filter: 0")
}

func (suite *FiltererTestSuite) TestFilterer_errorOnNonExistentOperator() {

	options := suite.genOptions([]query.FilterOption{
		{
			Column:                []string{"SomeField"},
			Operator:              "non-existent",
			Values:                []string{"5"},
			ValuesLogicalOperator: "",
		},
	})

	_, err := filterer.CompileFromQueryListOptions[TestStruct](options)
	suite.ErrorContains(err, "invalid string operator: non-existent on filter: 0")
}

func TestFiltererTestSuite(t *testing.T) {
	suite.Run(t, new(FiltererTestSuite))
}

func (suite *FiltererTestSuite) genOptions(filters []query.FilterOption) *query.ListOptions {
	options := &query.ListOptions{
		Sorts:      nil,
		Filters:    filters,
		Fields:     nil,
		Pagination: nil,
	}
	return options
}
