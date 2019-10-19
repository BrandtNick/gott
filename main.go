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
	s, isStr := data.(string)
	if isStr {
		return base64.StdEncoding.EncodeToString([]byte(s))

	}

	b, isByte := data.([]byte)
	if isByte {
		return base64.StdEncoding.EncodeToString([]byte(b))
	}

	return "Oops, translation failed"
}

func toHex(data interface{}) string {
	s, isStr := data.(string)

	bs, err := base64.StdEncoding.DecodeString(s)
	if err == nil {
		return hex.EncodeToString(bs)
	}

	if isStr {
		return hex.EncodeToString([]byte(s))
	}

	return "Oops, translation failed"
}

func main() {
	app := cli.NewApp()
	app.Name = "Go Type Translator (GOTT)"
	app.Usage = "Translate, decode, encode. \n \n Supports: \n - Base64 \n - Hex \n - Text \n"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Nick Brandt",
			Email: "niiick@live.se",
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command does not exist: %q\n", command)
	}
	app.Commands = []cli.Command{
		{
			Name:      "translate",
			Aliases:   []string{"t"},
			Usage:     "Translate",
			ArgsUsage: "[argument]",
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
					Usage: "Copy result to clipboard on execution",
				},
			},
			Action: func(c *cli.Context) error {
				var input string
				var data interface{}
				var res string

				if c.NArg() > 0 {
					input = c.Args().First()
					fmt.Println("Input: ", input)
				} else {
					fmt.Println("No arguments provided")
					return nil
				}

				// Text to base64
				if c.String("from") == "text" && c.String("to") == "base64" {
					data = input
					res = toBase64(data)
				}
				// Hex to base64
				if c.String("from") == "hex" && c.String("to") == "base64" {
					data, err := hex.DecodeString(input)
					if err != nil {
						fmt.Println(err)
					}
					res = toBase64(data)
				}

				// Text to hex
				if c.String("from") == "text" && c.String("to") == "hex" {
					data = input
					res = toHex(data)
				}
				// Base64 to hex
				if c.String("from") == "base64" && c.String("to") == "hex" {
					data = input
					res = toHex(data)
				}

				// TODO: Copies the result to clipboard
				if c.Bool("copy") && res != "" {
					fmt.Println("Copied: ", res)
				}

				fmt.Println("Result: ", res)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
