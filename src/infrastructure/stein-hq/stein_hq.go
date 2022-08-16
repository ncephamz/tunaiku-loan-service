package steinhq

import (
	"time"

	stein "github.com/nasrul21/go-stein"
)

func NewSteinHqConnection() *stein.Stein {
	return stein.NewClient(
		"https://api.steinhq.com/v1/storages/62f5ad10bca21f053ea7ff39",
		&stein.Option{
			Username: "ncephamz@gmail.com",
			Password: "26feb1994",
			Timeout:  20 * time.Second,
		},
	)
}
