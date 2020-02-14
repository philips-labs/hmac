package alerts

import (
	"encoding/json"
	"time"

	"github.com/cloudfoundry-community/gautocloud"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/postgresql"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"
)

// PGPayload is an entry in the alerts table
type PGPayload struct {
	ID        int64          `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	Payload   types.JSONText `db:"payload"`
}

// PGStorer is an PostgreSQL implementation of the Storer interface
type PGStorer struct {
	svc *dbtype.PostgresqlDB
	db  *sqlx.DB
}

// Store stores a payload in the psql DB
func (p *PGStorer) Store(payload Payload) error {
	db := p.db

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	tx := db.MustBegin()

	tx.MustExec("INSERT INTO alerts (payload) VALUES ($1)", payloadJSON)
	tx.Commit()

	return nil
}

// NewPGStorer retruns a psql storer
func NewPGStorer() (*PGStorer, error) {
	storer := &PGStorer{}

	err := gautocloud.Inject(&storer.svc)

	if err != nil {
		return nil, err
	}

	storer.db = sqlx.NewDb(storer.svc.DB, "postgres")
	err = storer.db.Ping()
	if err != nil {
		return nil, err
	}

	return storer, nil

}