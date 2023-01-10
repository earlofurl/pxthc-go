package pxthc

import (
	"context"
	"github.com/earlofurl/pxthc/sqlc"
)

type UomService interface {
	FindAllUoms(ctx context.Context) ([]*sqlc.Uom, error)
	FindUomByID(ctx context.Context, id int64) (*sqlc.Uom, error)
	FindUomByName(ctx context.Context, name string) (*sqlc.Uom, error)
}
