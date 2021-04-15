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

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jojojojonas/shopware6-iamge-scraper/scraper"
	"os"
	"time"
)

// Config for starting service
type Config struct {
	Url        string   `json:"url"`
	Categories []string `json:"categories"`
}

func main() {

	// Logging scraping
	fmt.Println("\n\n\n\n--------------------------------------------------------------------\n\n")
	fmt.Println("    ___      ___           ___      ___      ___      ___      ___      ___      ___      ___      ___      ___      ___      ___      ___   \n   /\\  \\    /\\  \\         /\\  \\    /\\  \\    /\\  \\    /\\  \\    /\\__\\    /\\  \\    /\\  \\    /\\__\\    /\\__\\    /\\  \\    /\\  \\    /\\  \\    /\\  \\  \n  _\\:\\  \\  _\\:\\  \\       _\\:\\  \\  /::\\  \\  /::\\  \\  /::\\  \\  /:| _|_  /::\\  \\  /::\\  \\  /:/__/_  /::L_L_  _\\:\\  \\  /::\\  \\  /::\\  \\  /::\\  \\ \n /\\/::\\__\\/\\/::\\__\\     /\\/::\\__\\/:/\\:\\__\\/::\\:\\__\\/::\\:\\__\\/::|/\\__\\/\\:\\:\\__\\/:/\\:\\__\\/::\\/\\__\\/:/L:\\__\\/\\/::\\__\\/::\\:\\__\\/:/\\:\\__\\/::\\:\\__\\\n \\::/\\/__/\\::/\\/__/     \\::/\\/__/\\:\\/:/  /\\:\\:\\/  /\\:\\:\\/  /\\/|::/  /\\:\\:\\/__/\\:\\ \\/__/\\/\\::/  /\\/_/:/  /\\::/\\/__/\\:\\:\\/  /\\:\\/:/  /\\:\\:\\/  /\n  \\/__/    \\/__/         \\:\\__\\   \\::/  /  \\:\\/  /  \\:\\/  /   |:/  /  \\::/  /  \\:\\__\\    /:/  /   /:/  /  \\:\\__\\   \\:\\/  /  \\::/  /  \\:\\/  / \n                          \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/    \\/__/  ")
	fmt.Println("\n\nJ&J Ideenschmiede UG (haftungsbeschränkt)\nFährstraße 31\n21502 Geesthacht\nEmail: info@jj-ideenschmiede.de\nTelefon: + 49 4152 8903730\n\n")
	fmt.Println("--------------------------------------------------------------------\n")

	// Save scraper data
	var data []scraper.Data

	// Open json file
	config, err := os.Open("config.json")
	if err != nil {
		fmt.Println("An error occurred while load config json file: ", err)
	}

	// Close after function ends
	defer config.Close()

	// Decode config
	var decode Config

	err = json.NewDecoder(config).Decode(&decode)
	if err != nil {
		fmt.Println("An error occurred while decoding the json data: ", err)
	}

	// Check each category
	for _, value := range decode.Categories {

		// Logging category
		fmt.Println("Start checking " + value)
		fmt.Println("\n--------------------------------------------------------------------\n")

		// Set page number
		number := 1

		// Check each site
		for {

			// Checking log site
			fmt.Printf("Seite: %d\n", number)

			// Define url
			site := fmt.Sprintf("%s%s/?order=name-desc&p=%d", decode.Url, value, number)

			// Check url
			images, last := scraper.Site(site)

			// Check each broken image
			for _, value := range images {

				// Add broken images
				data = append(data, scraper.Data{value.Mpn, value.Link})

			}

			// Check if this is the lase site
			if last {

				// Stop loop
				break

			} else {

				// Count up
				number++

			}

		}

		// Logging end category
		fmt.Println("--------------------------------------------------------------------\n\n")

	}

	// Check data length
	if len(data) > 0 {

		// Get date
		date := time.Now()

		// Create new file
		file, err := os.Create(fmt.Sprintf("files/images-%s.csv", date.Format("20060102150405")))
		if err != nil {
			fmt.Println("An error occurred while creating an file: ", err)
		}

		// Create writer
		writer := csv.NewWriter(file)

		// Flush writer after function ends
		defer writer.Flush()

		// Write header
		err = writer.Write([]string{"MPN", "URL"})
		if err != nil {
			fmt.Println("An error occurred while creating an header row in file: ", err)
		}

		// Write data
		for _, value := range data {

			// Create new row
			err = writer.Write([]string{value.Mpn, value.Link})
			if err != nil {
				fmt.Println("An error occurred while creating an row in an file: ", err)
			}

		}

		// Logging end category
		fmt.Printf("Scan finished & file was created with filename: images-%s.csv\n", date.Format("20060102150405"))

	} else {

		// Logging end category
		fmt.Println("Scan finished. No missing images were found in this scan.")

	}

	// Print last line for cli
	fmt.Println("\n\n--------------------------------------------------------------------\n\n\n\n")

}
