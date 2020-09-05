package main

import (
	"fmt"
	"log"
	"time"

	"github.com/genjidb/genji/database"

	"github.com/genjidb/genji/document"

	"github.com/genjidb/genji"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	fmt.Println("こんにちわーるど")

	db, err := genji.Open("genji.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// create table
	{
		if err := db.Exec("create table orders"); err != nil {
			if err == database.ErrTableAlreadyExists {
				log.Println(err) // テーブルがあればそれはそれでいい
			} else {
				log.Fatalln(err)
			}
		}
	}

	// insert
	{
		if err := db.Exec("insert into orders values ?", &Order{
			ID:                 "unique order id",
			Code:               "security's order code",
			Symbol:             Symbol{Code: "1320", Exchange: ExchangeT},
			ExecutionType:      ExecutionTypeMarket,
			Quantity:           3,
			ContractedQuantity: 2,
			Status:             StatusPart,
			Contracts: []Contract{
				{
					PositionCode:   "security's position code 1",
					Price:          22500,
					Quantity:       1,
					ContractedTime: time.Date(2020, 9, 7, 9, 0, 1, 0, time.Local),
				},
				{
					PositionCode:   "security's position code 2",
					Price:          22510,
					Quantity:       1,
					ContractedTime: time.Date(2020, 9, 7, 9, 0, 3, 0, time.Local),
				},
			},
		}); err != nil {
			log.Fatalln(err)
		}
	}

	// select
	{
		res, err := db.Query("select * from orders")
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Close() // これ忘れると終了時にデッドロックが起こってパニくる

		err = res.Iterate(func(d document.Document) error {
			var o Order
			if err := document.StructScan(d, &o); err != nil {
				return err
			}
			log.Printf("%+v\n", o)
			return nil
		})
	}
}

type Symbol struct {
	Code     string
	Exchange Exchange
}

type Exchange string

const (
	ExchangeUnspecified Exchange = ""
	ExchangeT           Exchange = "T"
	ExchangeM           Exchange = "M"
	ExchangeF           Exchange = "F"
	ExchangeS           Exchange = "S"
)

type OrderID string

type Order struct {
	ID                 OrderID
	Code               string
	Symbol             Symbol
	ExecutionType      ExecutionType
	Quantity           float64
	ContractedQuantity float64
	Status             Status
	Contracts          []Contract
}

type ExecutionType string

const (
	ExecutionTypeUnspecified ExecutionType = ""
	ExecutionTypeMarket      ExecutionType = "market"
	ExecutionTypeLimit       ExecutionType = "limit"
)

type Status string

const (
	StatusUnspecified Status = ""
	StatusNew         Status = "new"
	StatusInOrder     Status = "in_order"
	StatusPart        Status = "part"
	StatusDone        Status = "done"
	StatusInCancel    Status = "in_cancel"
	StatusCanceled    Status = "canceled"
)

type Contract struct {
	PositionCode   string
	Price          float64
	Quantity       float64
	ContractedTime time.Time
}
