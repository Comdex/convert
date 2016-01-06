// encode project main.go
package main

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/axgle/mahonia"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	app := cli.NewApp()
	app.Name = "convert"
	app.Usage = "convert is a text file encoding converter  \n   such as you can convert text file's encoding from gbk to utf-8" +
		"\n   support encoding : utf-8, gbk, big5, latin-1, UTF-16, ASCII,, gb18030, SJIS, EUC-JP"
	app.Version = "0.1.0"
	app.Author = "Comdex"
	app.Email = "wcomdex@foxmail.com"

	var srccode string
	var dstcode string

	processfn := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		} else {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			if dstcode == "utf-8" {
				decoder := mahonia.NewDecoder(srccode)
				newstring := decoder.ConvertString(string(bytes))
				b := []byte(newstring)
				ioutil.WriteFile(path, b, 777)
				return nil
			} else {
				decoder := mahonia.NewDecoder(srccode)
				newstring := decoder.ConvertString(string(bytes))
				encoder := mahonia.NewEncoder(dstcode)
				newstring2 := encoder.ConvertString(newstring)
				b := []byte(newstring2)
				ioutil.WriteFile(path, b, 777)
				return nil
			}

		}
	}

	app.Action = func(c *cli.Context) {
		fmt.Println("use -h for help")
	}

	app.Commands = []cli.Command{
		{
			Name:    "convertfile",
			Aliases: []string{"f"},
			Usage:   "use it to convert a file, you can use f -h for help",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "src",
					Value: "",
					Usage: "the file need to be converted",
				},
				cli.StringFlag{
					Name:  "dst",
					Value: "",
					Usage: "the file has converted successful",
				},
				cli.StringFlag{
					Name:  "scode",
					Value: "",
					Usage: "src file encoding",
				},
				cli.StringFlag{
					Name:  "dcode",
					Value: "",
					Usage: "dst file encoding ,the default is utf-8",
				},
			},
			Description: "convert a text file's encoding",
			Action: func(c *cli.Context) {
				src := c.String("src")
				dst := c.String("dst")
				scode := c.String("scode")
				dcode := c.String("dcode")
				
				if src == "" {
					fmt.Println("src can not empty! please use --src")
					return
				}
				if !com.IsFile(src) {
					fmt.Println("src must be a text file!")
					return
				}
				if dst == "" {
					dst = src
				}

				if scode == "" {
					fmt.Println("src encoding can not empty! please use --scode")
					return
				}
				if dcode == "" {
					dcode = "utf-8"
				}
				
				if dcode == "utf-8" {
					bytes, err := ioutil.ReadFile(src)
					if err != nil {
						fmt.Println("read file error")
						return
					}
					fmt.Println("processing.......")
					decoder := mahonia.NewDecoder(scode)
					newstring := decoder.ConvertString(string(bytes))
					b := []byte(newstring)
					ioutil.WriteFile(dst, b, 777)
				} else {
					bytes, err := ioutil.ReadFile(src)
					if err != nil {
						fmt.Println("read file error")
						return
					}
					fmt.Println("processing.......")
					decoder := mahonia.NewDecoder(scode)
					newstring := decoder.ConvertString(string(bytes))
					encoder := mahonia.NewEncoder(dcode)
					newstring2 := encoder.ConvertString(newstring)
					b := []byte(newstring2)
					ioutil.WriteFile(dst, b, 777)
				}
				fmt.Println("process done!")

			},
		},
		{
			Name:    "convertdir",
			Aliases: []string{"d"},
			Usage:   "use it to convert all file in directory, you can use d -h for help",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "src",
					Value: "",
					Usage: "the directory need to be converted",
				},
				cli.StringFlag{
					Name:  "dst",
					Value: "",
					Usage: "the directory has converted successful",
				},
				cli.StringFlag{
					Name:  "scode",
					Value: "",
					Usage: "src directory encoding",
				},
				cli.StringFlag{
					Name:  "dcode",
					Value: "",
					Usage: "dst directory's file encoding ,the default is utf-8",
				},
			},
			Action: func(c *cli.Context) {
				src := c.String("src")
				dst := c.String("dst")
				scode := c.String("scode")
				dcode := c.String("dcode")

				if scode == "" {
					fmt.Println("src encoding can not empty! please use --scode")
					return
				}
				srccode = scode
				if dcode == "" {
					dcode = "utf-8"
				}
				dstcode = dcode
				var err error
				if src == "" || src == "." {
					src, err = os.Getwd()
					if err != nil {
						fmt.Println("can not get work directory")
						return
					}
				}
				if !com.IsDir(src) {
					fmt.Println("src must be a directory!")
					return
				}
				fmt.Println("now start process directory: " + src)
				if dst != "" {

					err := com.CopyDir(src, dst)
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println("processing.......")
					filepath.Walk(dst, processfn)
				} else {
					//if dst is null, process src directory
					fmt.Println("processing.......")
					filepath.Walk(src, processfn)
				}

				fmt.Println("process done!")
			},
		},
	}
	app.Run(os.Args)
}

