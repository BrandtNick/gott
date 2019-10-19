/*
Go Type Translator

Translate, decode or encode any of these types:
- Base64
- Hex
- Text
- Bytes
- Decimal
*/
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func toBase64(data interface{}) string {
	s, isString := data.(string)
	if isString {
		return base64.StdEncoding.EncodeToString([]byte(s))

	}

	b, isByte := data.([]byte)
	if isByte {
		return base64.StdEncoding.EncodeToString([]byte(b))
	}

	return "Oops, translation failed"
}

func main() {
	app := cli.NewApp()
	app.Name = "Go Type Translator (GOTT)"
	app.Usage = "Translate, decode, encode. \n \n Supports: \n - Base64 \n - Hex \n - Text \n"

	app.Commands = []cli.Command{
		{
			Name:    "translate",
			Aliases: []string{"t"},
			Usage:   "Translate",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "from, f",
					Usage: "From type",
				},
				cli.StringFlag{
					Name:  "to, t",
					Usage: "Target type",
				},
				cli.BoolFlag{
					Name:  "copy, c",
					Usage: "Copy result on exec",
				},
			},
			Action: func(c *cli.Context) error {
				var input string
				var data interface{}
				var result string

				if c.NArg() > 0 {
					input = c.Args().First()
				} else {
					fmt.Println("No arguments provided")
					return nil
				}

				// Text to base64
				if c.String("from") == "text" && c.String("to") == "base64" {
					fmt.Println("Type:", data)
					data = input
					result = toBase64(data)
					fmt.Println("Result: ", result)
				}

				// Hex tp base64
				if c.String("from") == "hex" && c.String("to") == "base64" {
					fmt.Println("Type:", input)
					data, err := hex.DecodeString(input)
					if err != nil {
						fmt.Println(err)
					}

					result = toBase64(data)
					fmt.Println("Result: ", result)
				}

				// Copies the result to clipboard
				if c.Bool("copy") && result != "" {
					fmt.Println("Copied: ", result)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
