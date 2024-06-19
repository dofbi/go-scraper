package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Produit struct {
	ID        int    `json:"ID"`
	Nom       string `json:"Nom"`
	Prix      string `json:"Prix"`
	Thumbnail string `json:"Thumbnail"`
}

func main() {
	newProducts := make([]Produit, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("auchan.sn", "www.auchan.sn"),
	)

	collector.OnHTML(".selection-product-featured article", func(h *colly.HTMLElement) {
		produitID, err := strconv.Atoi(h.Attr("data-id-product"))
		if err != nil {
			log.Println("Could not get id")
		}

		produitNom := h.Attr("data-name")

		prixStr := h.ChildText(".product-price-and-shipping .price")

		image := h.ChildAttr("a.thumbnail img", "src")

		Produit := Produit{
			ID: produitID,
			Nom: produitNom,
			Prix:      prixStr,
			Thumbnail: image,
		}

		newProducts = append(newProducts, Produit)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.auchan.sn/")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(newProducts)
}
