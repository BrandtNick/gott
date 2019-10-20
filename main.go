/*
Go Type Translator

Translate, decode or encode any of these types:
- Base64
- Hex
- Text
- Binary
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
	input  interface{}
	output string
}

func (t translator) textToBase64() string {
	s := t.input.(string)
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (t translator) hexToBase64() string {
	s := t.input.(string)
	data, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (t translator) textToHex() string {
	s := t.input.(string)
	return hex.EncodeToString([]byte(s))
}

func (t translator) base64ToHex() string {
	s := t.input.(string)

	ds, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(ds)
}

func (t translator) textToBinary() string {
	var b string
	s := t.input.(string)
	for i := 0; i < len(s); i++ {
		b += fmt.Sprintf("%08b", byte(s[i]))
	}
	return b
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

				if c.NArg() > 0 {
					T.input = c.Args().First()
					fmt.Println("Input: ", T.input)
				} else {
					fmt.Println("No arguments provided")
					return nil
				}

				// Text to base64
				if c.String("from") == "text" && c.String("to") == "base64" {
					T.output = T.textToBase64()
				}

				// Hex to base64
				if c.String("from") == "hex" && c.String("to") == "base64" {
					T.output = T.hexToBase64()
				}

				// Text to hex
				if c.String("from") == "text" && c.String("to") == "hex" {
					T.output = T.textToHex()
				}

				// Base64 to hex
				if c.String("from") == "base64" && c.String("to") == "hex" {
					T.output = T.base64ToHex()
				}

				// Text to binary
				if c.String("from") == "text" && c.String("to") == "binary" {
					T.output = T.textToBinary()
				}

				// TODO: Copies the result to clipboard
				if c.Bool("copy") && T.output != "" {
					fmt.Println("Copied: ", T.output)
				}

				fmt.Println("Output: ", T.output)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
