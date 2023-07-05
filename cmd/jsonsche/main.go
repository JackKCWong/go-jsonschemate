package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/invopop/jsonschema"
	"github.com/traefik/yaegi/interp"
	"github.com/twpayne/go-jsonstruct/v2"
	"gopkg.in/yaml.v3"
)

func main() {

	outYaml := flag.Bool("yaml", false, "output schema in yaml, useful for openapi v3")
	flag.Parse()

	var jsonstr []byte
	var err error
	var in io.Reader

	if flag.NArg() == 1 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		in = file
	} else {
		in = os.Stdin
	}

	jsonstr, err = io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	generator := jsonstruct.NewGenerator()
	err = generator.ObserveJSONReader(bytes.NewReader(jsonstr))
	if err != nil {
		panic(err)
	}

	goStruct, err := generator.Generate()
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(goStruct))

	i := interp.New(interp.Options{})

	_, err = i.Eval(string(goStruct))
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v\n", v)

	v, err := i.Eval(`v := T{}`)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v\n", v)

	schema := jsonschema.ReflectFromType(v.Type())

	var out []byte
	if *outYaml {
		out, err = json.MarshalIndent(schema, "", "    ")
		if err != nil {
			panic(err)
		}

		var tmp = make(map[string]interface{})
		err = json.Unmarshal(out, &tmp)
		if err != nil {
			panic(err)
		}

		out, err = yaml.Marshal(tmp)
		if err != nil {
			panic(err)
		}
	} else {
		out, err = json.MarshalIndent(schema, "", "    ")
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("%s\n", string(out))
}
