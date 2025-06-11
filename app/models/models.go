package models

import "time"

type User struct {
	Id        int32     `pg:"id,pk"`
	Name      string    `pg:"name"`
	Email     string    `pg:"email,unique"`
	Age       int       `pg:"age"`
	Picture   string    `pg:"picture"`
	IAT       int32     `pg:"iat"`
	EXP       int32     `pg:"exp"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
}
type UserDetails struct {
	Id int32
}
type Places struct {
	Id int16 `pg:"id,pk"`
}
