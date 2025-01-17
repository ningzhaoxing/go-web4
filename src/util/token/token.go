package token

import (
	"database/sql"
	"net/http"
	db2 "userManageSystem-blog/src/dao/db/userDb"
	user2 "userManageSystem-blog/src/model/params/sql"
	"userManageSystem-blog/src/model/user"
	errors2 "userManageSystem-blog/src/pkg/errors"
)

type Token struct {
	r *http.Request
}

func NewToken(r *http.Request) *Token {
	return &Token{r: r}
}

// GetUser 获取user
func (t *Token) GetUser(db *sql.DB) (*user.User, error) {
	claim := t.r.Context().Value("claim")
	if claim == nil {
		return nil, errors2.ErrNotFoundUserInToken
	}
	u := claim.(*CustomClaims).User

	us, err := db2.QueryUserByEmail(user2.UserSqlParam{
		Db:   db,
		User: u,
	})
	if err != nil {
		return nil, err
	}
	return us, nil
}
