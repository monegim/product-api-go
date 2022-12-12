package data

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/monegim/product-api-go/data/model"
)

type Connection interface {
	IsConnected() (bool, error)
	GetCoffees(*int) (model.Coffees, error)
}

type PostgresSQL struct {
	db *sqlx.DB
}

// New creates a new connection to the database
func New(connection string) (Connection, error) {
	db, err := sqlx.Connect("postgres", connection)
	if err != nil {
		return nil, err
	}
	return &PostgresSQL{db}, nil
}

// IsConnected checks the connection to the database and returns an error if not connected
func (c *PostgresSQL) IsConnected() (bool, error) {
	err := c.db.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetCoffees returns all coffees from the database
func (c *PostgresSQL) GetCoffees(coffeeid *int) (model.Coffees, error) {
	cos := model.Coffees{}

	if coffeeid != nil {
		err := c.db.Select(&cos, "SELECT * FROM coffees WHERE id = $1", &coffeeid)
		if err != nil {
			return nil, err
		}
	} else {
		err := c.db.Select(&cos, "SELECT * FROM coffees", &coffeeid)
		if err != nil {
			return nil, err
		}
	}
	// fetch the ingredients for each coffee
	for n, cof := range cos {
		i := []model.CoffeeIngredient{}
		err := c.db.Select(&i, "SELECT ingredient_id FROM coffee_ingredients WHERE coffee_id=$1 AND quantity > 0", cof.ID)
		if err != nil {
			return nil, err
		}
		cos[n].Ingredients = i
	}
	return cos, nil
}
