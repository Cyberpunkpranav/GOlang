package models

import "time"

type User struct {
	Id        int32     `pg:"id,pk"`
	Name      string    `pg:"name"`
	Email     string    `pg:"email,unique"`
	Age       int       `pg:"age"`
	Picture   string    `pg:"picture"`
	Iat       int32     `pg:"iat"`
	Exp       int32     `pg:"exp"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
}
type UserOtherDetails struct {
	Id int32

}
