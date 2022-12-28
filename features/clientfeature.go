package features

import (
	"ecommerce/models"
)

type ClientFeatures struct {
	client models.ClientRepository
}

func NewClientFeatures(repository models.ClientRepository) ClientFeatures {
	return ClientFeatures{client: repository}
}

func (d ClientFeatures) CreateClient(c models.ClientCreate) (err error) {
	encText, err := Encrypt(c.Password, MySecret)
	c.Password = encText
	return d.client.CreateClient(c)
}

func (d ClientFeatures) UpdatePointsClient(accumulation_points int, idClient string) (err error) {
	return d.client.UpdateAccumulationPointsClient(accumulation_points, idClient)
}

func (d ClientFeatures) DeleteClient(idClient string) (err error) {
	return d.client.DeleteClient(idClient)
}

func (d ClientFeatures) AuthClient(name, password string) (m models.Client, err error) {
	encText, err := Encrypt(password, MySecret)
	if err != nil {
		return
	}
	return d.client.AuthClient(name, encText)
}

func (d ClientFeatures) FindByIdClient(idClient string) (m models.Client, err error) {
	return d.client.FindByIdClient(idClient)
}

func (d ClientFeatures) FindByNameClient(name string) (m models.Client, err error) {
	return d.client.FindByNameClient(name)
}

func (d ClientFeatures) FindAllClient() (m []models.ClientCreate, err error) {
	return d.client.FindAllClient()
}
