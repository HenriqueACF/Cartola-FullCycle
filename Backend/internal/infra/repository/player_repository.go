package repository

import (
    "cartola-backend/internal/domain/entity"
    "cartola-backend/internal/infra/db"
    "context"
    "database/sql"
)

type PlayerRepository struct {
    Repository
}

func NewPlayerRepository(dbConn *sql.DB) *PlayerRepository{
    return &PlayerRepository{
        Repository: Repository{
            dbConn: dbConn,
            Queries: db.New(dbConn),
        },
    }
}

func (r *PlayerRepository) Create(ctx context.Context, player *entity.Player) error {
    err := r.Queries.CreatePlayer(ctx, db.CreatePlayerParams{
        ID: player.ID,
        Name: player.Name,
        Price: player.Price,
    })
    return err
}

func (r *PlayerRepository) FindByID(ctx context.Context, id string) (*entity.Player, error){
    player, err := r.Queries.FindPlayerById(ctx, id)
    if err != nil {
        return nil, err
    }
    
    return &entity.Player{
        ID: player.ID,
        Name: player.Name,
        Price: player.Price
    }, nil
}