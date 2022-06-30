package wallet

import (
	"Vico1993/Wallet/service"
	"Vico1993/Wallet/util"
	"log"
)

var (
	PURCHASE = "purchase"
	EXCHANGE = "exchange"
	EARN = "earn"
)

type Operation struct {
	Date 			string 		`json:"date"`
	Quantity 		float64 	`json:"quantity"`
	Unit 			string 		`json:"unit"`
	UnitPrice		float64 	`json:"unitPrice"`
	From 			string  	`json:"from"`
	FromQuantity	float64		`json:"fromQuantity"`
	Price 			float64		`json:"price"`
	Fiat 			string		`json:"fiat"`
	OType			string		`json:"otype"`
	Tag 			[]string	`json:"tag"`
}

func NewOperation(
	date string,
	quantity float64,
	unit string,
	unitPrice float64,
	from string,
	fromQuantity float64,
	price float64,
	fiat string,
	oType string,
	tag ...string,
) Operation {
	if unitPrice == 0 {
		unitPrice = 1 * price / quantity
	}

	return Operation{
		Date: date,
		Quantity: quantity,
		Unit: unit,
		UnitPrice: unitPrice,
		From: from,
		FromQuantity: fromQuantity,
		Price: price,
		Fiat: fiat,
		OType: oType,
		Tag: tag,
	}
}

func (o Operation) WithProfit() []string {
	currentPrice, err := service.GetAssetPrice(o.Unit)
	if err != nil {
		log.Printf("Error procession Unit: %s - %s", o.Unit, err.Error())
		currentPrice = 0
	}

	return []string{
		o.Unit,
		util.FormatFloat(o.Quantity),
		util.FormatFloat(o.UnitPrice),
		util.FormatFloat(o.Price),
		util.FormatFloat(o.GetCurrentPrice(currentPrice)),
		util.FormatFloat(o.GetProfit(o.GetCurrentPrice(currentPrice))) + "%",
	}
}

func (o Operation) GetCurrentPrice(currentPrice float64) float64 {
	return currentPrice * o.Quantity
}

func (o Operation) getCurrentUnitPrice() float64 {
	value, err := service.GetAssetPrice(o.Unit)
	if err != nil {
		log.Printf("Impossible to retrieve operation current value: %s - %s", o.Unit, err.Error())
		value = 0
	}

	return value
}

func (o Operation) GetProfit(latestPrice float64) float64 {
	return calculProfit(o.Price, latestPrice)
}