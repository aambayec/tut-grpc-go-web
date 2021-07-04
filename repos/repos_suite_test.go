package repos_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aambayec/tut-grpc-go-web/repos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	err error
	db *xorm.Engine
	dbSql *sql.DB
	mock sqlmock.Sqlmock
	gr repos.GlobalRepository

	truncateUsers = func() {
		// delete this line if you won't use mock db
		mock.ExpectQuery("TRUNCATE users").
			WillReturnRows(sqlmock.NewRows([]string{}))

		_, err = db.Query("TRUNCATE users")
		Ω(err).To(BeNil())
	}

	clearDatabase = func () {
		if db == nil {
			Fail("unable to run test because database is missing")
		}

		truncateUsers()
	}
)

var _ = BeforeSuite(func () {
	// pass connection string here if you won't use mockdb, eg. root:Ulyanin123@tcp(localhost:3306)/tut_grpc
	db, err = xorm.NewEngine("mysql", "")
	Ω(err).To(BeNil())

	// delete these 3 lines if you won't use mock db
	dbSql, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	Ω(err).To(BeNil())
	db.DB().DB = dbSql

	gr = repos.GlobalRepo(db)
})

var _ = AfterSuite(func () {
	err = mock.ExpectationsWereMet()
	Ω(err).To(BeNil())
})

func TestRepos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repos Suite")
}
