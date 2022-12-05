package usecase

import (
    "cartola-backend/internal/domain/entity"
    "cartola-backend/internal/domain/repository"
    "cartola-backend/pkg/uow"
    "context"
)

type AddPlayerInput struct {
    ID string
    Name string
    InitialPrice float64
}

type AddPlayerUseCase struct {
    Uow uow.UowInterface
}

func (a *AddPlayerUseCase) Execute(ctx context.Context, input AddPlayerInput) error{
    playerRepository := a.getPlayerRepository(ctx)
    player := entity.NewPlayer(input.ID, input.Name, input.InitialPrice)
    err := playerRepository.Create(ctx, player)
    if err != nil {
        return err
    }
    return a.Uow.CommitOrRollback()
}

func (a *AddPlayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface{
    playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
    if err != nil {
        panic(err)
    }
    return playerRepository.(repository.PlayerRepositoryInterface)
}