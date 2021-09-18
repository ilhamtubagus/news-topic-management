package persistence

// // Basic imports
// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// 	"gorm.io/gorm"
// )

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including a T() method which
// // returns the current testing context
// type ExampleTestSuite struct {
// 	suite.Suite
// 	dbClient *gorm.DB
// }

// // Make sure that VariableThatShouldStartAtFive is set to five
// // before each test
// func (suite *ExampleTestSuite) SetupSuite() {
// 	suite.dbClient
// }

// // All methods that begin with "Test" are run as tests within a
// // suite.
// func (suite *ExampleTestSuite) TestExample() {
// }

// // In order for 'go test' to run this suite, we need to create
// // a normal test function and pass our suite to suite.Run
// func TestExampleTestSuite(t *testing.T) {
// 	suite.Run(t, new(ExampleTestSuite))
// }
