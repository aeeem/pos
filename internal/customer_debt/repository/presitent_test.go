package repository_test

import (
	"regexp"
	"testing"

	"pos/internal/customer_debt/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	cleanup := func() {
		db.Close()
	}

	return gdb, mock, cleanup
}

func TestGetDebtDetails(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := repository.NewCustomerDebtPresistentRepository(db)

	expectedID := uint(1)
	rows := sqlmock.NewRows([]string{"id", "customer_id", "trx_id"}).
		AddRow(1, 10, 5)

	// Expect the correct query to be generated
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "customer_debts" WHERE id = $1 ORDER BY "customer_debts"."id" LIMIT 1`)).
		WithArgs(expectedID).
		WillReturnRows(rows)

	_, err := repo.GetDebtDetails(expectedID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
