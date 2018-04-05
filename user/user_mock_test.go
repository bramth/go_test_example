package user

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var mock sqlmock.Sqlmock

func InitMock() {
	var err error
	var dbx *sql.DB
	dbx, mock, err = sqlmock.New()
	if err != nil {
		panic(err.Error())
	}

	db = sqlx.NewDb(dbx, "postgres")

}

func TestMain(m *testing.M) {
	InitMock()

	os.Exit(m.Run())
}

func TestGetUserMock(t *testing.T) {
	tc := []struct {
		user User
		err  error
	}{
		{
			user: User{
				ID:    76,
				Name:  "asdf",
				Email: "a@gmail.com",
				Phone: "56789",
			},
			err: nil,
		},
		{
			user: User{},
			err:  fmt.Errorf("err db"),
		},
	}

	for _, tt := range tc {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"})
		rows.AddRow(tt.user.ID, tt.user.Name, tt.user.Email, tt.user.Phone)

		mock.ExpectQuery("SELECT id, name, email, phone FROM training_user").WithArgs(tt.user.ID).WillReturnRows(rows).WillReturnError(tt.err)

		user, err := GetUserByID(tt.user.ID)
		if err != tt.err {
			t.Error(err.Error())
		}

		tempt := tt.user
		if !assert.Equal(t, *user, tempt) {
			t.Error("not equal")
		}
	}

	// t.Log(user)
}
