package models

import (
	"github.com/volatiletech/null"
	"time"
)

type Session struct {
	SessioId string    `db:"id" json:"sessioId"`
	UserId   string    `db:"uid" json:"userId"`
	Login    time.Time `db:"login_at" json:"loginAt"`
	Logut    time.Time `db:"logout_at" json:"logoutAt"`
}

type User struct {
	ID         string    `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Email      string    `db:"email" json:"email"`
	Password   string    `db:"password" json:"password"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	ArchivedAt null.Time `db:"archived_at" json:"archivedAt"`
}

type Task struct {
	uniqueId      string `db:"uniqueId" json:"uniqueId"`
	TaskID        string `db:"taskID" json:"taskID"`
	TaskName      string `db:"taskName" json:"taskName"`
	Status        bool   `db:"status" json:"status"`
	CompletedUpto string `db:"completed_upTo" json:"completed_upTo"`
}

type Login struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

const ActiveUser string = "active_user"
