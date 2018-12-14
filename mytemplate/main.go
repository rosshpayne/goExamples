package main
   
import (
	"log"
	"os"
	"strings"
	"text/template"
	"fmt"
)

func main() {
	const (
//
//		this uses block to define a template and execute it - anything between {{block}} and the last {{end}} will be a template called "list".  Block execute automatically
//
//		master_  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{if 2}}{{println "madeit"}}{{end}}{{if not .}}{{println "-  **EMPTY"}}{{else}}{{"ABC"}}{{end}}{{end}}{{end}}`
//
//		this uses define rather than block - anything between {{define}} and the last {{end}} will be a template called "list" which we can execute 
//                using {{template "list" .}}
//
		master_  = `Names:{{define "list"}}{{"\n"}}{{range .}}{{if 2}}{{println "madeit"}}{{end}}{{if not .}}{{println "-  **EMPTY"}}{{else}}{{"ABC\n"}}{{end}}{{end}}{{end}}{{template "list" .}}`
//
//		this deines list but doesn't execute it - no {{template "list" .}}
//
//		master_  = `Names:{{define "list"}}{{"\n"}}{{range .}}{{if 2}}{{println "madeit"}}{{end}}{{if not .}}{{println "-  **EMPTY"}}{{else}}{{"ABC"}}{{end}}{{end}}{{end}}`
//
		//overlay = `{{define "list"}} {{join . "!"}}{{end}} `
		overlay = `{{define "list2"}}{{"Hello ."}}{{end}}` //{{template "list" .}}{{template "list2"}}`      // must specify template to execute the named template
		//overlay = `{{define "list"}} {{join . "!"}}{{end}}`// {{template "list" .}}`
		//overlay = `{{define "list"}}{{range .}}{{if  .}}{{.}}{{end}}{{end}}{{end}}`// {{template "list" .}}`
//		overlay = `{{define "list"}}{{range .}}{{.}}{{end}}{{end}}`// {{template "list" .}}`
	)
	var (
		//funcs     = template.FuncMap{"join": strings.Join}                          // static function FuncMap.:q!
		funcs     = template.FuncMap{"join": strings.Join}                          // static function FuncMap.:q!
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "", "Star-Lord"}    // go creates an array of strings and guardians is a slice on that array
		positions = []int{1,2,3,0,5,6,7,8}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master_)     // create tpl from string "master" NB not identifier master
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)	// overlay templated added to existing template
	if err != nil {
		log.Fatal(err)
	}
        fmt.Println(masterTmpl.DefinedTemplates())

	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
        fmt.Println(overlayTmpl.DefinedTemplates())
//
	/*if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, positions); err != nil {
		log.Fatal(err)
	}*/
	if err := overlayTmpl.ExecuteTemplate(os.Stdout, "list2",  positions); err != nil {
		log.Fatal(err)
	}
}
