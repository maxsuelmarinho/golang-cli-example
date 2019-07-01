package main

import(
	"os"
	"os/exec"
	"log"
	"net"
	"fmt"
	"bytes"
	"github.com/urfave/cli"
	"runtime"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Author = "Maxsuel"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"
	app.Version = "0.0.1"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name: "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name: "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				
				fmt.Println(cname)				
				return nil
			},
		},
		{
			Name: "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}			
				return nil
			},
		},
		{
			Name: "echo",
			Aliases: []string{"e"},
			Usage: "Echos a message",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "message",
					Value: "default message",
				},
			},
			Action: func(c *cli.Context) error {
				cmd := exec.Command("echo", c.String("message"))				

				cmdOutput := &bytes.Buffer{}
				cmd.Stdout = cmdOutput
				err := cmd.Run()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}

				fmt.Print(string(cmdOutput.Bytes()))
				return nil
			},
		},
		{
			Name: "ls",
			Usage: "Execute ls command",
			Flags: []cli.Flag{
			},
			Action: func(c *cli.Context) error {

				if runtime.GOOS == "windows" {
					fmt.Println("Can't execute this command on a windows machine")
					return nil
				}
				out, err := exec.Command("ls", "-ltr").Output()

				if err != nil {
					fmt.Printf("%s", err)
				}

				fmt.Println(string(out[:]))

				return nil
			},
		},
	}

	
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}