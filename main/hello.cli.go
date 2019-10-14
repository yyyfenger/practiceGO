package main
import (
	"fmt"
	"os"
	"gopkg.in/urfave/cli.v1"
)
func main(){
	app:=cli.NewApp()
	app.Name="hello_cli"
	app.Usage="Print hello world"
	app.Flags=[]cli.Flag{
		cli.StringFlag{
			Name:"name,n",
			Value:"world",
			Usage:"who to say hello to.",
		},
	}
	app.Action=func(c *cli.Context) error{
		name:=c.GlobalString("name")
		fmt.Printf("hello %s!\n",name)
		return nil
	}
	app.Run(os.Args)
}