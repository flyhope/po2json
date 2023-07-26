package main

import (
	"encoding/json"
	"flag"
	"github.com/leonelquinteros/gotext"
	"os"
)

var (
	paramInput     = flag.String("input", "", "/path/to/go/pkg")
	paramOutput    = flag.String("output", "", "/path/to/go/pkg")
	paramAddExport = flag.Bool("add_export", false, "if true, auto add export default for js")
)

func main() {
	flag.Parse()

	po := gotext.NewPo()
	po.ParseFile(*paramInput)
	translations := po.GetDomain().GetTranslations()
	result := make(map[string]string, len(translations))
	for k, v := range translations {
		value := v.Get()
		if k == "" || value == "" || value == k {
			continue
		}
		result[k] = v.Get()
	}
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fo, errFo := os.Create(*paramOutput)
	if errFo != nil {
		panic(errFo)
	}
	if *paramAddExport {
		if _, errWrite := fo.WriteString("export default "); errWrite != nil {
			panic(errWrite)
		}

	}

	_, errWrite := fo.Write(output)
	if errWrite != nil {
		panic(errWrite)
	}
}
