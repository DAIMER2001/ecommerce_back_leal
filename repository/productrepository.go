package repository

import (
	"context"
	"database/sql"
	"ecommerce/errors"
	"ecommerce/models"
	"fmt"
	"os/exec"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

var (
	psqlCreateProduct = `
		INSERT INTO product(
			id,
			name,
			price,
			accumulation_points
		) values ($1, $2, $3, $4);
	`
	psqlEditProduct = `
		UPDATE product SET name = $1, accumulation_points = $2, price = $3 where id = $4
		
	`
	psqlListProduct = `
	    SELECT * FROM product;
	`
	psqlDeleteProduct = `
		DELETE FROM product WHERE id = $1
	`
)

func (d ProductRepository) CreateProduct(product models.Product) (err error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return
	}
	exec, err := d.db.Exec(
		context.TODO(),
		psqlCreateProduct,
		product.Name,
		product.AccumulationPoints,
		product.Price,
		newUUID,
	)
	if err != nil {
		return errors.ErrSQLSyntax(fmt.Errorf("error in model, %s", err).Error())
	}
	if !exec.Insert() {
		return errors.ErrSQLSyntax(fmt.Errorf("error en process create, %s", product.Name).Error())
	}
	return
}

func (d ProductRepository) UpdateProduct(product models.Product) (err error) {
	exec, err := d.db.Exec(
		context.TODO(),
		psqlEditProduct,
		product.Name,
		product.Price,
		product.AccumulationPoints,
	)
	if err != nil {
		return errors.ErrSQLSyntax(fmt.Errorf("error in model , %s", err).Error())
	}
	if !exec.Insert() {
		return errors.ErrSQLSyntax(fmt.Errorf("error in process update, %s", product.IDProduct).Error())
	}
	return
}

func (d ProductRepository) FindProducts() (products []models.Product, err error) {
	rows, err := d.db.Query(context.Background(), psqlListProduct)
	if err != nil {
		return
	}
	fmt.Println("problemas1")
	for rows.Next() {
		p := models.Product{}
		AccumulationPointsNull := sql.NullInt64{}
		fmt.Println("problemas2")

		err = rows.Scan(
			&p.Name,
			&p.Price,
			&AccumulationPointsNull,
			&p.IDProduct,
			&p.Image,
			&p.Description,
		)
		if err != nil {
			return
		}
		p.AccumulationPoints = AccumulationPointsNull.Int64
		products = append(products, p)
	}
	defer rows.Close()
	return products, nil
}

func (d ProductRepository) DeleteProduct(idProduct string) (err error) {
	_, err = d.db.Exec(context.Background(), psqlDeleteProduct, idProduct)
	if err != nil {
		return err
	}

	return nil
}
