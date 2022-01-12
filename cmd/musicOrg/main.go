package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	var app = cli.NewApp()
	defineAppConfig(app)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func defineAppConfig(app *cli.App) {

	author := &cli.Author{
		Name: "Zapi",
	}

	app.Name = "Music Renamer"
	app.Description = "This app rename music from author - song to song"
	app.Authors = []*cli.Author{author}
	app.Version = "1.0.0"
	app.Usage = "Rename music naming"
	app.Action = Rename
	// app.Commands = []*cli.Command{
	// 	{
	// 		Name:    "path",
	// 		Aliases: []string{"p"},
	// 		Usage:   "Path to songs",
	// 	},
	// }
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Aliases:  []string{"p"},
			Value:    ".",
			Usage:    "Path to songs",
			Required: true,
		},
	}
}

//Rename rename from author - song to song
func Rename(c *cli.Context) error {

	dir := c.String("path") + "/"

	files, err := ioutil.ReadDir(c.String("path"))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		//Only for .mp3 files
		if !f.IsDir() &&
			strings.HasSuffix(f.Name(), ".mp3") {

			arr := strings.Split(f.Name(), " - ")
			fmt.Printf("Lenght: %d, Name: %s\n", len(arr), f.Name())
			//whose format is author - song
			if len(arr) == 2 {
				err := os.Rename(dir+f.Name(), dir+strings.TrimSpace(arr[1]))
				if err != nil {
					return err
				}
				fmt.Printf("Original: %s renamed to %s", f.Name(), strings.TrimSpace(arr[1]))
				fmt.Println()
			}
			fmt.Printf("[WARN] Song: %s Not renamed", f.Name())
		}
	}
	return nil
}
