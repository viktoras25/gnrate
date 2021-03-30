package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

func main() {
	app := cli.App{
		Name:        "gnrate",
		Description: "Genrate fake data",
		Authors: []*cli.Author{
			{
				Name: "Viktoras Bezaras",
			},
		},
		Usage:     "Generates fake data",
		UsageText: "gnrate [count] [language] subject",
		Action: func(c *cli.Context) error {
			configuration, err := parse(c.Args().Slice())
			if err != nil {
				fmt.Println(err.Error())
			}

			generate(configuration)

			return nil
		},
	}

	app.Run(os.Args)
}

func setLocale(language string) {
	language = strings.ToLower(language)

	languagesMap := map[string]map[string]interface{}{
		"en":       locales.En,
		"english":  locales.En,
		"en_us":    locales.En_US,
		"british":  locales.En_GB,
		"de":       locales.De,
		"german":   locales.De,
		"austrian": locales.De_AT,
		"ru":       locales.Ru,
		"russian":  locales.Ru,
		"es":       locales.Es,
		"spanish":  locales.Es,
		"fr":       locales.Fr,
		"french":   locales.Fr,
		"it":       locales.It,
		"italian":  locales.It,
		"jp":       locales.Ja,
		"japanese": locales.Ja,
		"ko":       locales.Ko,
		"korean":   locales.Ko,
	}

	if value, ok := languagesMap[language]; ok {
		faker.Locale = value
	}
}

func generate(c *Configuration) {

	setLocale(c.language)

	subject := strings.ToLower(c.subject)

	for i := 0; i < c.count; i += 1 {
		switch subject {
		case "name", "names":
			fmt.Println(faker.Name())
		case "firstname", "firstnames":
			fmt.Println(faker.Name().FirstName())
		case "address", "addresses":
			fmt.Println(faker.Address())
		case "email", "emails":
			fmt.Println(faker.Internet().FreeEmail())
		case "url", "urls":
			fmt.Println(faker.Internet().Url())
		case "phone", "phones":
			fmt.Println(faker.PhoneNumber())
		case "company", "companies":
			fmt.Println(faker.Company().Name() + " " + faker.Company().Suffix())
		case "sentence", "sentences", "lorem":
			fmt.Println(faker.Lorem().Paragraph(1))
		}
	}
}
