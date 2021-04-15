//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of image-scraper.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package scraper

import (
	"github.com/gocolly/colly"
)

// Data struct to save data
type Data struct {
	Mpn  string `json:"mpn"`
	Link string `json:"link"`
}

// Site ist for scraping the hole site content
func Site(url string) ([]Data, bool) {

	// Define variable to return
	var data []Data

	// To check last site
	exits := false

	// Create colly
	c := colly.NewCollector()

	// Check last site
	c.OnHTML(".cms-listing-col", func(e *colly.HTMLElement) {

		// Check last site
		if e.DOM.Find(".alert-content-container").Children().HasClass("alert-content") {

			// Return error
			exits = true

		}

		// Check placeholder
		if e.DOM.Find(".product-image-placeholder").Children().HasClass("icon-placeholder") {

			// Save mpn to product
			mpn, _ := e.DOM.Find("meta[itemprop=mpn]").Attr("content")

			// Save link to product
			link, _ := e.DOM.Find(".product-image-link").Attr("href")

			// Add to return
			data = append(data, Data{mpn, link})

		}

	})

	// Visit site
	c.Visit(url)

	// Return values
	return data, exits

}
