package databases

import (
	sql "database/sql"
	"fmt"
	errors "github.com/crazyfacka/pinky-02/errors"
	models "github.com/crazyfacka/pinky-02/models"
	_ "github.com/mattn/go-sqlite3"
)

type ProductsDB struct {
	file string
}

func NewProductsDB(file string) (sqlite *ProductsDB) {
	sqlite = new(ProductsDB)
	sqlite.file = file
	return sqlite
}

func openFile(sqlite *ProductsDB) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqlite.file)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getProductName(db *sql.DB, productId string) (string, error) {
	sqlStmt := `
		SELECT name FROM product WHERE id = $id
	`
	var name string
	row := db.QueryRow(sqlStmt, productId)
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

func getProductData(db *sql.DB, productId string) (data *models.ProductData, err error) {
	sqlStmt := `
		SELECT description, stock, price FROM product_info pi, product_stock ps WHERE pi.id = $id AND ps.id = $id
	`
	data = models.NewProductData()
	row := db.QueryRow(sqlStmt, productId)
	err = row.Scan(&data.Description, &data.Stock, &data.Price)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (sqlite *ProductsDB) GetProductById(productId string) (product *models.Product, notFound *errors.ProductNotFound) {
	db, err := openFile(sqlite)
	if err != nil {
		return product, errors.NewProductNotFound(err)
	}
	defer db.Close()

	name, err := getProductName(db, productId)
	if err != nil {
		return product, errors.NewProductNotFound(fmt.Errorf("%s not found", productId))
	}

	data, err := getProductData(db, productId)
	if err != nil {
		return product, errors.NewProductNotFound(fmt.Errorf("Could not retrieve data for %s", productId))
	}

	return models.NewProduct(productId, name, data.Description, data.Price, data.Stock), nil
}
