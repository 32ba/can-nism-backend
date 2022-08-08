package utils

import (
	"context"
	"errors"
	"go-ranking-api/ent"
	"go-ranking-api/ent/token"
	"go-ranking-api/ent/user"
	"go-ranking-api/model"
	"time"
)

func TokenVerifier(typ string, t string) (int, error) {
	switch typ {
	case "access":
		ctx := context.Background()
		id, err := model.DBClient.Token.
			Query().
			Where(
				token.AccessTokenEQ(t),
				token.AccessTokenExpiredAtGTE(time.Now()),
				token.DeletedAtIsNil(),
			).
			OnlyID(ctx)
		return id, err
	case "refresh":
		ctx := context.Background()
		id, err := model.DBClient.Token.
			Query().
			Where(
				token.RefreshTokenEQ(t),
				token.RefreshTokenExpiredAtGTE(time.Now()),
				token.DeletedAtIsNil(),
			).
			OnlyID(ctx)
		return id, err
	default:
		return 0, errors.New("invalid type")
	}
}

func GetUserFromTokenID(id int) (*ent.User, error) {
	ctx := context.Background()
	user, err := model.DBClient.User.
		Query().
		Where(
			user.HasTokenWith(token.IDEQ(id)),
		).
		Only(ctx)
	return user, err
}
