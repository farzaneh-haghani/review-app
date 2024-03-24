package repository

import (
	"context"

	"github.com/farzaneh-haghani/review-app/internal/entity"
)

type CommonBehaviorRepository[T entity.DBModel] interface{
	ByID(ctx context.Context, id uint)(T,error)
	ByField(ctx context.Context, field string,id uint)(T,error)
	Save(ctx context.Context, model T)error
}

type UserRepository interface{
	CommonBehaviorRepository(entity.User)
}