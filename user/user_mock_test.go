package user

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

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
func TestGetUserByID(t *testing.T) {

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "tc1-normal",
			args: args{5},
			want: &User{
				ID:    5,
				Email: "ravif@tokopedia.com",
				Phone: "0000",
				Name:  "rafi ahmad",
			},
			wantErr: false,
		},
		{
			name:    "tc2-err",
			args:    args{5},
			want:    nil,
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"})

			mk := mock.ExpectQuery("SELECT id, name, email, phone FROM training_user").
				WithArgs(tt.args.id)

			if tt.wantErr {
				mk.WillReturnError(errors.New("err db"))
			} else {
				rows.AddRow(tt.want.ID, tt.want.Name, tt.want.Email, tt.want.Phone)
				mk.WillReturnRows(rows)
			}

			got, err := GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
