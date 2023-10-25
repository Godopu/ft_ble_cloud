package config

import (
	"os"

	"github.com/magiconair/properties"
)

const Mode = "debug"

var Params = map[string]interface{}{}

func LoadConfig() {

	p := properties.MustLoadFile("./config.properties", properties.UTF8)
	Params["computing_delay"] = p.GetInt("computing_delay", 200)
}

func CreateInitFile() error {
	f, err := os.Create("./config.properties")
	if err != nil {
		return err
	}
	defer f.Close()

	p := properties.NewProperties()

	for k, v := range Params {
		err = p.SetValue(k, v)
		if err != nil {
			panic(err)
		}
	}

	err = p.SetValue("computing_delay", 200)
	if err != nil {
		panic(err)
	}

	_, err = p.Write(f, properties.UTF8)
	if err != nil {
		panic(err)
	}

	return nil
}
