package main

import (
	"fmt"
	"github.com/kwanCCC/tiny/args"
)

type Config struct {
	APPName string `default:"app name"`
	PortNum uint   `default:"8080"`

	//DB struct {
	//	Name     string
	//	User     string `default:"root"`
	//	Password string `required:"true" env:"DBPassword"`
	//	Port     uint   `default:"3306"`
	//}

	//Contacts []struct {
	//	Name  string
	//	Email string `required:"true"`
	//}
}

func main() {
	argParser := args.NewParser()
	argParser.Version = `0.0.1`
	argParser.Helptext = `--config=/a/b/c`
	argParser.NewString("config c")
	argParser.Parse()
	fmt.Println(argParser)
	configFile := argParser.GetString(`c`)
	fmt.Println(configFile)
	//config := Config{}
	//configor.Load(&config, )
}
