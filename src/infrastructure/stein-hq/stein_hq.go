package steinhq

import (
	"time"

	stein "github.com/nasrul21/go-stein"
)

func NewSteinHqConnection() *stein.Stein {
	return stein.NewClient(
		"https://api.steinhq.com/v1/storages/62f5ad10bca21f053ea7ff39",
		&stein.Option{
			Timeout: 20 * time.Second,
		},
	)
}
