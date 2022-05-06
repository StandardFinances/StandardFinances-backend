package db

import (
	"context"
	"github.com/google/uuid"
	"main/ent"
	"main/ent/user"
	"time"
)

func GetUserByPhone(instance *ent.Client, phone int64) *ent.User {
	getuserresult, err := instance.User.Query().Where(user.Phone(phone)).Only(context.Background())

	if err != nil {
		getuserresult, _ = instance.User.Create().SetPhone(phone).SetUID(uuid.New().String()).SetRegtime(time.Now()).Save(context.Background())
	}
	return getuserresult
}
