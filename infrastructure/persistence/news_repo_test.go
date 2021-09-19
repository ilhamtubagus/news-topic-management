package persistence

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

type NewsRepoTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo NewsRepoImpl
}

func newMockRegex() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// before each test
func (suite *NewsRepoTestSuite) SetupTest() {
	db, mock := newMockRegex()
	gdb, err1 := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{}) // open gorm db
	suite.Equal(nil, err1)
	repo := *NewNewsRepository(gdb)
	suite.mock = mock
	suite.repo = repo
}

func (suite *NewsRepoTestSuite) TestSaveNewsSuccess() {
	id := uint64(1)
	const sql = `INSERT INTO "news" (.+) RETURNING`
	news := entity.News{Title: "Safe Investments", Author: "fian", Status: "published", Content: "Safe investments content", TopicID: 1}
	// const sqlInsert = `INSERT INTO "news" ("title","author","status","content","topic_id") VALUES ($1,$2,$3,$4,$5) RETURNING "updated_at","published_at","created_at","deleted_at","id"`
	suite.mock.ExpectBegin()
	suite.mock.
		ExpectQuery(sql).WithArgs(news.Title, news.Author, news.Status, news.Content, news.TopicID).
		WillReturnRows(sqlmock.NewRows([]string{"deleted_at", "published_at", "created_at", "updated_at", "id"}).AddRow(nil, time.Now(), time.Now(), time.Now(), id))
	suite.mock.ExpectCommit()

	_, err := suite.repo.SaveNews(&news)
	suite.Empty(err)
	suite.Equal(id, news.ID)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while inserting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestNewsRepoTestSuite(t *testing.T) {
	suite.Run(t, new(NewsRepoTestSuite))
}
