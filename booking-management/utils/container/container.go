package container

import (
	"database/sql"
)

type Containers struct {
	Adapters     Adapters
	Repositories Repositories
}

type Adapters struct {
	Db *sql.DB
}

type Repositories struct {
}
