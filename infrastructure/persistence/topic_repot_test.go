package persistence

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TopicRepoTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo TopicRepoImpl
}

// before each test
func (suite *TopicRepoTestSuite) SetupTest() {
	db, mock := newMockMatcherEqual()
	gdb, err1 := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{}) // open gorm db
	suite.Equal(nil, err1)
	repo := *NewTopicRepository(gdb)
	suite.mock = mock
	suite.repo = repo
}

func (suite *TopicRepoTestSuite) TestSaveTopicSuccess() {
	id := uint64(1)
	topic := entity.Topic{Topic: "investments", CreatedAt: time.Now()}
	const sqlInsert = `INSERT INTO "topics" ("topic","created_at") VALUES ($1,$2) RETURNING "created_at","id"`
	suite.mock.ExpectBegin()
	suite.mock.
		ExpectQuery(sqlInsert).WithArgs(topic.Topic, topic.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(time.Now(), id))
	suite.mock.ExpectCommit()

	_, err := suite.repo.SaveTopic(&topic)
	suite.Empty(err)
	suite.Equal(id, topic.ID)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while inserting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}
func (suite *TopicRepoTestSuite) TestGetTopicSuccess() {
	rows := sqlmock.NewRows([]string{"id", "tag", "updated_at", "created_at"}).
		AddRow(1, "investments", nil, time.Now())
	const sqlQuery = `SELECT * FROM "topics" WHERE topic = $1 ORDER BY "topics"."id" LIMIT 1`
	const q = "investments"

	suite.mock.ExpectQuery(sqlQuery).
		WithArgs(q).
		WillReturnRows(rows)
	topic, err := suite.repo.GetTopic(q)
	//assert tag is not nil
	suite.NotEmpty(topic)
	suite.Empty(err)

	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while getting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}

}

func (suite *TopicRepoTestSuite) TestGetTopicByIdSuccess() {
	rows := sqlmock.NewRows([]string{"id", "tag", "updated_at", "created_at"}).
		AddRow(1, "fund", nil, time.Now())
	const sqlQuery = `SELECT * FROM "topics" WHERE "topics"."id" = $1 ORDER BY "topics"."id" LIMIT 1`
	const id = 1

	suite.mock.ExpectQuery(sqlQuery).
		WithArgs(id).
		WillReturnRows(rows)
	topic, err := suite.repo.GetTopicById(id)
	//assert tag is not nil
	suite.NotEmpty(topic)
	suite.Empty(err)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while getting a row", err)
	}

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *TopicRepoTestSuite) TestDeleteTopicSuccess() {
	const sqlDelete = `DELETE FROM "topics" WHERE "topics"."id" = $1`
	const id = 1

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(sqlDelete).WithArgs(id).WillReturnResult(driver.RowsAffected(1))
	suite.mock.ExpectCommit()
	err := suite.repo.DeleteTopic(id)
	//assert error is empty
	suite.Empty(err)
	if err != nil {
		suite.T().Errorf("error '%v' was not expected, while getting a row", err)
	}
	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}
func (suite *TopicRepoTestSuite) TestGetAllTopicSuccess() {
	rows := sqlmock.NewRows([]string{"id", "topic", "updated_at", "created_at"}).
		AddRow(1, "fund", nil, time.Now()).AddRow(1, "investment fund", nil, time.Now())
	const sqlQuery = `SELECT * FROM "topics"`

	suite.mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
	tag, err := suite.repo.GetAllTopic()
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
func TestTopicRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TopicRepoTestSuite))
}
