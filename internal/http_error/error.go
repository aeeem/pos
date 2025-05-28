package http_error

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type HttpErrorCode struct {
	HTTPErrorCode int
	Message       string
}

func CheckError(err error) (httpErr HttpErrorCode) {
	log.Info(err)
	log.Info()
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			httpErr.HTTPErrorCode = 400
			httpErr.Message = fmt.Sprintf("Item with name %s already exists", pgErr.ConstraintName)
			return
		}
	}
	return
}
