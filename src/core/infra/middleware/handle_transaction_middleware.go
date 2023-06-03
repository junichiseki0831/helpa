package middleware

import (
	"helpa/src/support/smperr"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

var (
	transactionRequests = []string{
		"POST",
		"PUT",
		"DELETE",
	}
)

func isTransactionMethod(method string) bool {

	for _, tMethod := range transactionRequests {
		if strings.HasPrefix(method, tMethod) {
			return true
		}
	}
	return false
}

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:secrets@tcp(db:3306)/helpa")
	if err != nil {
		return nil, smperr.Internalf("Failed to initialize database: %+v", err)
	}
	return db, nil
}

func DBInterceptor(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("start interceptor")
		method := c.Request.Method
		var tx *sqlx.Tx
		if isTransactionMethod(method) {
			log.Println("start transaction")
			transaction, err := db.Beginx()
			if err != nil {
				c.Error(err)
				return
			}
			tx = transaction

			c.Set("tx", transaction)
		} else {
			c.Set("db", db)
		}

		c.Next()

		if len(c.Errors) > 0 {
			if tx != nil {
				tx.Rollback()
			}
			log.Println("rollback")
			return
		}
		if tx != nil {
			if err := tx.Commit(); err != nil {
				c.Error(smperr.Internal("fail to handle transaction commit"))
				return
			}
			log.Println("committed")
		}
	}
}
