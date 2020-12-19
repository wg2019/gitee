package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	err := Load()
	if err != nil {
		log.Fatalf("Load config fail, err: %+v", err)
	}
	app.HideHelp = true
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var app = &cli.App{
	Name:        AppName,
	Usage:       "码云代码库工具",
	Description: "啊吧啊吧啊吧",
	Commands: []*cli.Command{
		{
			Name:    DeleteCommand,
			Aliases: []string{"d"},
			Usage:   "删除代码库",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    ConfigFlag,
					Aliases: []string{"c"},
					Usage:   "配置文件，格式: \".ini\"",
					Value:   "info.conf",
				},
			},
			Action: Delete,
		},
	},
}
