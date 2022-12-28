package models

import "github.com/google/uuid"

type Client struct {
	IDClient           uuid.UUID `json:"id,omitempty"`
	Name               string    `json:"name"`
	Password           string    `json:"password,omitempty"`
	AccumulationPoints int64     `json:"accumulation_points"`
	Role               string    `json:"role,omitempty"`
}

type ClientCreate struct {
	Name               string `json:"name"`
	Password           string `json:"password,omitempty"`
	AccumulationPoints int64  `json:"accumulation_points"`
	Role               string `json:"role,omitempty"`
}

type ClientAuth struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}

type ClientRepository interface {
	AuthClient(name, password string) (m Client, err error)
	CreateClient(client ClientCreate) (err error)
	UpdateAccumulationPointsClient(accumulation_points int, idClient string) (err error)
	DeleteClient(clientId string) (err error)
	FindByIdClient(clientId string) (m Client, err error)
	FindByNameClient(name string) (m Client, err error)
	FindAllClient() (m []ClientCreate, err error)
}
