package criptos

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Cripto struct {
	Name   string
	Price  string
	Ticker string
}

func (cr Cripto) GetCriptoList() []string {
	c := colly.NewCollector(
		colly.AllowedDomains("coinmarketcap.com", "www.coinmarketcap.com"),
	)

	criptoList := []string{}

	c.OnHTML("p.coin-item-name", func(e *colly.HTMLElement) {
		criptoList = append(criptoList, e.Text)
	})

	err := c.Visit("https://coinmarketcap.com/pt-br/")

	if err != nil {
		panic(err)
	}

	return criptoList
}

func (cr Cripto) GetCriptoInfo(name string) Cripto {
	cr.Name = name

	fmt.Println("Buscando informações sobre", name)

	c := colly.NewCollector(
		colly.AllowedDomains("coinmarketcap.com", "www.coinmarketcap.com"),
	)

	c.OnHTML(`[data-test="text-cdp-price-display"]`, func(e *colly.HTMLElement) {
		cr.Price = e.Text
	})

	err := c.Visit("https://coinmarketcap.com/pt-br/currencies/" + name + "/")

	if err != nil {
		panic(err)
	}

	return cr
}
