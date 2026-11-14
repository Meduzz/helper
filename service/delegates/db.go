package delegates

import (
	"github.com/Meduzz/helper/service"

	"gorm.io/gorm"
)

type (
	// WithDB defines the interface for services that need a GORM database connection.
	WithDB interface {
		WithDB(db *gorm.DB)
	}

	dbDelegate struct {
		db *gorm.DB
	}
)

var _ service.Delegate = (*dbDelegate)(nil)

// NewDbDelegate creates a new GORM database delegate.
func NewDbDelegate(db *gorm.DB) service.Delegate {
	return &dbDelegate{
		db: db,
	}
}

// Visit visits a service and injects the GORM database if the service implements DbConnection.
func (d *dbDelegate) Visit(svc service.Service) error {
	if svcWithDb, ok := svc.(WithDB); ok {
		svcWithDb.WithDB(d.db)
	}
	return nil
}
