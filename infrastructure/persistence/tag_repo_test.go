package persistence

// Basic imports
import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TagRepoTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo TagRepoImpl
}

// before each test
func (suite *TagRepoTestSuite) SetupTest() {
	db, mock := NewMock()
	gdb, err1 := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{}) // open gorm db
	suite.Equal(nil, err1)
	repo := *NewTagRepository(gdb)
	suite.mock = mock
	suite.repo = repo
}

func (suite *TagRepoTestSuite) TestSaveTagSuccess() {

	tag := entity.Tag{Tag: "investments", CreatedAt: time.Now()}
	const sqlInsert = `INSERT INTO "tags" ("tag","created_at") VALUES ($1,$2) RETURNING "id","created_at","updated_at"`
	suite.mock.ExpectBegin()
	suite.mock.
		ExpectQuery(sqlInsert).WithArgs(tag.Tag, tag.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(1, time.Now(), time.Now()))
	suite.mock.ExpectCommit()

	_, err := suite.repo.SaveTag(&tag)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while inserting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *TagRepoTestSuite) TestGetTagSuccess() {
	rows := sqlmock.NewRows([]string{"id", "tag", "updated_at", "created_at"}).
		AddRow(1, "fund", nil, time.Now())
	const sqlSearch = `SELECT * FROM "tags" WHERE tag = $1 ORDER BY "tags"."id" LIMIT 1`
	const q = "fund"

	suite.mock.ExpectQuery(sqlSearch).
		WithArgs(q).
		WillReturnRows(rows)
	tag, err := suite.repo.GetTag(q)
	//assert tag is not nil
	suite.NotEqual(nil, tag)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while getting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestTagRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TagRepoTestSuite))
}
