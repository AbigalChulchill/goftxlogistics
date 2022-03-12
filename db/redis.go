package db

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

var RDB = redis.NewClient(&redis.Options{
	Addr:     "dssq.xyz:6380",
	Password: "coliseum-dealer-entree",
	DB:       0, // use default DB
})

// Test validity of redis connection.
func Test(t *testing.T) {
	ctx := RDB.Context()
	testCommand := RDB.Info(ctx)

	fmt.Printf("t: %v\n", t)
	fmt.Println(testCommand)
}