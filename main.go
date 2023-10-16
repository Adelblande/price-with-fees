package main

// CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))

import (
	"database/sql"
	"fmt"

	"github.com/adelblande/price-with-fees/internal/infra/database"
	"github.com/adelblande/price-with-fees/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	
	if(err != nil) {
		panic(err)
	}

	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput {
		ID: "321",
		Price: 15,
		Tax: 5,
	}

	output, err := uc.Execute(input)

	if(err != nil) {
		panic(err)
	}

	fmt.Println(output)

}