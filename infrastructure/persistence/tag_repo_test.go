package persistence

// Basic imports
import (
	"database/sql"
	"database/sql/driver"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMockMatcherEqual() (*sql.DB, sqlmock.Sqlmock) {
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
	db, mock := newMockMatcherEqual()
	gdb, err1 := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{}) // open gorm db
	suite.Equal(nil, err1)
	repo := *NewTagRepository(gdb)
	suite.mock = mock
	suite.repo = repo
}

func (suite *TagRepoTestSuite) TestSaveTagSuccess() {
	id := uint64(1)
	tag := entity.Tag{Tag: "investments", CreatedAt: time.Now()}
	const sqlInsert = `INSERT INTO "tags" ("tag","created_at") VALUES ($1,$2) RETURNING "id","created_at","updated_at"`
	suite.mock.ExpectBegin()
	suite.mock.
		ExpectQuery(sqlInsert).WithArgs(tag.Tag, tag.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(1, time.Now(), time.Now()))
	suite.mock.ExpectCommit()

	_, err := suite.repo.SaveTag(&tag)
	suite.Empty(err)
	suite.Equal(id, tag.ID)
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
	const sqlQuery = `SELECT * FROM "tags" WHERE tag = $1 ORDER BY "tags"."id" LIMIT 1`
	const q = "fund"

	suite.mock.ExpectQuery(sqlQuery).
		WithArgs(q).
		WillReturnRows(rows)
	tag, err := suite.repo.GetTag(q)
	//assert tag is not nil
	suite.NotEmpty(tag)
	suite.Empty(err)

	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while getting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}

}

func (suite *TagRepoTestSuite) TestGetTagByIdSuccess() {
	rows := sqlmock.NewRows([]string{"id", "tag", "updated_at", "created_at"}).
		AddRow(1, "fund", nil, time.Now())
	const sqlQuery = `SELECT * FROM "tags" WHERE "tags"."id" = $1 ORDER BY "tags"."id" LIMIT 1`
	const id = 1

	suite.mock.ExpectQuery(sqlQuery).
		WithArgs(id).
		WillReturnRows(rows)
	tag, err := suite.repo.GetTagById(id)
	//assert tag is not nil
	suite.NotEmpty(tag)
	suite.Empty(err)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while getting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}

}

func (suite *TagRepoTestSuite) TestDeleteTagSuccess() {
	// rows := sqlmock.NewRows([]string{"id", "tag", "updated_at", "created_at"}).
	// 	AddRow(1, "fund", nil, time.Now())
	const sqlDelete = `DELETE FROM "tags" WHERE "tags"."id" = $1`
	const id = 1

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(sqlDelete).WithArgs(id).WillReturnResult(driver.RowsAffected(1))
	suite.mock.ExpectCommit()
	err := suite.repo.DeleteTag(id)
	//assert error is empty
	suite.Empty(err)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while getting a row", err)
	}
	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *TagRepoTestSuite) TestGetAllTagSuccess() {
	rows := sqlmock.NewRows([]string{"id", "tag", "updated_at", "created_at"}).
		AddRow(1, "fund", nil, time.Now()).AddRow(1, "investment", nil, time.Now())
	const sqlQuery = `SELECT * FROM "tags"`

	suite.mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
	tag, err := suite.repo.GetAllTag()
	//assert tag is not nil
	suite.Empty(err)
	suite.NotEmpty(tag)
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
