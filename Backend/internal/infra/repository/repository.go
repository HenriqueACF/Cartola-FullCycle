package repository

import (
    "cartola-backend/internal/infra/db"
    "database/sql"
)

type Repository struct {
    dbConn *sql.DB
    *db.Queries
}

