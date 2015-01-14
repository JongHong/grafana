package models

import (
	"time"
)

const (
	ROLE_READ_WRITE RoleType = "ReadWrite"
	ROLE_READ                = "Read"
)

type RoleType string

type Collaborator struct {
	Id             int64
	AccountId      int64    `xorm:"not null unique(uix_account_id_for_account_id)"`
	Role           RoleType `xorm:"not null"`
	CollaboratorId int64    `xorm:"not null unique(uix_account_id_for_account_id)"`

	Created time.Time
	Updated time.Time
}

type RemoveCollaboratorCommand struct {
	CollaboratorId int64
	AccountId      int64
}

type AddCollaboratorCommand struct {
	Email          string   `json:"email" binding:"required"`
	AccountId      int64    `json:"-"`
	CollaboratorId int64    `json:"-"`
	Role           RoleType `json:"-"`
}

func NewCollaborator(accountId int64, collaboratorId int64, role RoleType) *Collaborator {
	return &Collaborator{
		AccountId:      accountId,
		CollaboratorId: collaboratorId,
		Role:           role,
		Created:        time.Now(),
		Updated:        time.Now(),
	}
}
