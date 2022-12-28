package repository

import (
	"context"
	"database/sql"
	"ecommerce/errors"
	"ecommerce/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ClientRepository struct {
	db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{db: db}
}

var (
	psqlCreateClient = `
		INSERT INTO client(
			id,
			name,
			password,
			accumulation_points,
			role
		) values ($1, $2, $3, $4, $5);
	`

	psqlFindByIdClient = `
	    SELECT name, accumulation_points FROM client WHERE id = $1;
	`

	psqlFindAllClients = `
	    SELECT name, accumulation_points, role FROM client;
	`

	psqlDeleteClient = `
		DELETE FROM client WHERE id = $1;
	`

	psqlEditAccumulationPointsClient = `
		UPDATE client SET accumulation_points = accumulation_points + $1 where id = $2;
	`

	psqlAuthClient = `SELECT * FROM client c WHERE c.name = $1;`
)

func (d ClientRepository) CreateClient(client models.ClientCreate) (err error) {
	tx, err := d.db.BeginTx(context.TODO(), pgx.TxOptions{})

	if err != nil {
		return errors.ErrSQLSyntax(fmt.Sprintf("%s", err))
	}
	if err != nil {
		return
	}
	exec, err := tx.Exec(
		context.TODO(),
		psqlCreateClient,
		uuid.New(),
		client.Name,
		client.Password,
		0,
		client.Role,
	)
	if err != nil {
		return
	}
	if exec.RowsAffected() != 1 {
		return errors.ErrSQLSyntax(fmt.Errorf("psql: expected 1 row affected, got %d", exec.RowsAffected()).Error())
	}
	err = tx.Commit(context.TODO())
	return
}

func (d ClientRepository) UpdateAccumulationPointsClient(accumulation_points int, idClient string) (err error) {
	tx, err := d.db.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return errors.ErrSQLSyntax(fmt.Sprintf("%s", err))
	}
	exec, err := tx.Exec(
		context.TODO(),
		psqlEditAccumulationPointsClient,
		accumulation_points,
		idClient,
	)
	if err != nil {
		return err
	}
	if exec.RowsAffected() != 1 {
		return errors.ErrSQLSyntax(fmt.Errorf("psql: expected 1 row affected, got %d", exec.RowsAffected()).Error())
	}
	err = tx.Commit(context.TODO())
	return
}

func (d ClientRepository) DeleteClient(idClient string) (err error) {
	_, err = d.db.Exec(context.Background(), psqlDeleteClient, idClient)
	if err != nil {
		return err
	}

	return nil
}

func (d ClientRepository) FindByIdClient(idClient string) (m models.Client, err error) {
	rows, err := d.db.Query(context.Background(), psqlFindByIdClient, idClient)
	if err != nil {
		return m, errors.ErrSQLSyntax(fmt.Sprintf("%s", err))
	}
	for rows.Next() {
		AccumulationPointsNull := sql.NullInt64{}
		err = rows.Scan(
			&m.Name,
			&AccumulationPointsNull,
		)
		if err != nil {
			return m, errors.ErrSQLSyntax(fmt.Sprintf("%s", err))
		}
		m.AccumulationPoints = AccumulationPointsNull.Int64

		if err != nil {
			return m, errors.ErrSQLSyntax(fmt.Sprintf("%s", err))
		}
	}
	defer rows.Close()
	return m, nil
}

func (d ClientRepository) FindAllClient() (clients []models.ClientCreate, err error) {
	rows, err := d.db.Query(context.Background(), psqlFindAllClients)
	if err != nil {
		return clients, errors.ErrSQLSyntax(fmt.Sprintf("%s", err))
	}
	for rows.Next() {
		m := models.ClientCreate{}
		AccumulationPointsNull := sql.NullInt64{}
		err = rows.Scan(
			&m.Name,
			&AccumulationPointsNull,
			&m.Role,
		)
		if err != nil {
			return
		}
		m.AccumulationPoints = AccumulationPointsNull.Int64
		clients = append(clients, m)
		if err != nil {
			return clients, errors.ErrSQLSyntax(fmt.Sprintf("%s", err))
		}
	}
	fmt.Println("problemas3")
	defer rows.Close()
	return
}

func (d ClientRepository) AuthClient(name, password string) (m models.Client, err error) {

	row := d.db.QueryRow(context.Background(), psqlAuthClient, name)
	if err != nil {
		return
	}
	AccumulationPointsNull := sql.NullInt64{}
	err = row.Scan(
		&m.IDClient,
		&m.Name,
		&m.Password,
		&AccumulationPointsNull,
		&m.Role,
	)
	m.AccumulationPoints = AccumulationPointsNull.Int64
	if err != nil {
		return
	}
	if m.Password != password {
		return m, errors.ErrSQLSyntax(fmt.Sprintf("usuario no encontrado %s", err))
	} else {
		m.Password = ""
	}

	return m, nil
}
