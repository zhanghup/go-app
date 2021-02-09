package main

import (
	"fmt"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/99designs/gqlgen/plugin/resolvergen"
	"io/ioutil"
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {

	log.SetOutput(ioutil.Discard)

	cfg, err := config.LoadConfigFromDefaultLocations()

	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	options := []api.Option{api.NoPlugins(), api.AddPlugin(resolvergen.New())}
	options = append(options, api.AddPlugin(&modelgen.Plugin{
		MutateHook: func(b *modelgen.ModelBuild) *modelgen.ModelBuild {
			for _, model := range b.Models {
				for _, field := range model.Fields {
					field.Tag += ` xorm:"'` + field.Name + `'"`
				}
			}
			return b
		},
	}))

	err = api.Generate(cfg, options...)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
