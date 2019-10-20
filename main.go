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

type translator struct {
	dataType string
	data     interface{}
}

func (t translator) textToBase64() string {
	// string type assertion
	s, isString := t.data.(string)
	if isString {
		return base64.StdEncoding.EncodeToString([]byte(s))
	}
	return ""
}

func (t translator) hexToBase64() string {
	s := t.data.(string)
	data, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (t translator) textToHex() string {
	s, isString := t.data.(string)

	bs, err := base64.StdEncoding.DecodeString(s)
	if err == nil {
		return hex.EncodeToString(bs)
	}

	if isString {
		return hex.EncodeToString([]byte(s))
	}
	return ""
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
				var T translator
				var res string

				if c.NArg() > 0 {
					T.data = c.Args().First()
					fmt.Println("Input: ", T.data)
				} else {
					fmt.Println("No arguments provided")
					return nil
				}

				// Text to base64
				if c.String("from") == "text" && c.String("to") == "base64" {
					res = T.textToBase64()
				}

				// Hex to base64
				if c.String("from") == "hex" && c.String("to") == "base64" {
					res = T.hexToBase64()
				}

				// Text to hex
				if c.String("from") == "text" && c.String("to") == "hex" {
					res = T.textToHex()
				}

				// Base64 to hex
				if c.String("from") == "base64" && c.String("to") == "hex" {
					res = T.textToHex()
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
