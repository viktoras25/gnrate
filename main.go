package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

func main() {
	app := cli.App{
		Action: func(c *cli.Context) error {
			subject := "name"
			count := 1
			language := "english"

			args := c.Args()
			argsSlice := args.Slice()

			if !args.Present() {
				return errors.New("at least one argument is required")
			}

			subject, argsSlice = cutLast(argsSlice)

			if len(argsSlice) > 0 {
				argument := ""
				argument, argsSlice = cutLast(argsSlice)

				if isNumeric(argument) {
					count, _ = strconv.Atoi(argument)
				} else {
					language = argument
				}
			}

			if len(argsSlice) > 0 {
				argument := ""
				argument, _ = cutLast(argsSlice)

				if isNumeric(argument) {
					count, _ = strconv.Atoi(argument)
				}
			}

			generate(subject, language, count)

			return nil
		},
	}

	app.Run(os.Args)
}

func cutLast(arguments []string) (string, []string) {
	if len(arguments) == 0 {
		return "", []string{}
	}

	lastElement := arguments[len(arguments)-1]
	remainingSlice := arguments[:len(arguments)-1]

	return lastElement, remainingSlice
}

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)

	return err == nil
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

func generate(subject string, language string, count int) {

	setLocale(language)

	subject = strings.ToLower(subject)

	for i := 0; i < count; i += 1 {
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
