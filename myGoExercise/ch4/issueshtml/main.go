package main

import (
	"log"
	"sort"
	"time"
	"html/template"
	"os"
)


	type rowT  struct {
		Col1   interface{}			// silly struct really - use case would have a well defined dataTypes for struct.  Testing interface{} knowledge
		Col2   interface{}
		Col3   interface{}
	}

        type customSort struct {
		tab  []*rowT				// use pointer. More efficient to copy int64 in swap operation rather than [*data,len,cap] slice metadata.
		less  func(i,j *rowT) bool		//  neater to embed less func in struct so reciever can access it rather than a package func.	
							//   also makes it an anoynmous func (closure) this way so it has state which maybe useful.
	}

	func (tb customSort) Len() int           { return len(tb.tab) }					// methods must be defined at package level
	func (tb customSort) Swap(i,j int )      { tb.tab[i], tb.tab[j] = tb.tab[j], tb.tab[i] }        //  hence its associated type must be defined at package level
	func (tb customSort) Less(i,j int ) bool { return tb.less(tb.tab[i], tb.tab[j]) }
  
	type  Item struct {
		Col1  int
		Col2  string
	}
  
	type  Result struct {
			   Title   string
			   Items   []Item }

func main () {


	loc, _ := time.LoadLocation("Australia/Sydney") 
/*
	var cs []*rowT  = []*rowT{ &rowT{ 10, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },		
	                           &rowT{ 8, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 2, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 6, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 13, "EBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 13, "JBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 13, "CBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 13, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 13, "ABC", time.Date(2018,5,18,14,13,14,29,loc) },
	                           &rowT{ 13, "ABC", time.Date(2018,5,18,8,13,14,29,loc) },
	                           &rowT{ 20, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 17, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 12, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           &rowT{ 4, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
				}

*/
	var cs []*rowT  = []*rowT{ { 10, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },		// Go takes pointer 
	                           { 8, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 2, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 6, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "EBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "JBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "CBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "ABC", time.Date(2018,5,18,14,13,14,29,loc) },
	                           { 13, "ABC", time.Date(2018,5,18,8,13,14,29,loc) },
	                           { 20, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 17, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 12, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 4, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
				}
 	for i,r := range cs { log.Printf("%d   %#v", i, *r) }
	log.Println()

	sort.Sort(customSort{ tab: cs, less: func(i,j *rowT) bool { // use struct literal to initialise customSort arg in sort.Sort
	    var vi,vj interface{}
	    for r := range []int{0,1,2}  {
		switch r {
		case 0 :
			vi=i.Col1
			vj=j.Col1
		case 1:
			vi=i.Col2
			vj=j.Col2
		case 2:
			vi=i.Col3
			vj=j.Col3
		}
		switch v1:=vi.(type) {
		case int:
			v2:=vj.(int)
			if  v1 != v2 { return v1<v2 }
		case float64:
			v2:=vj.(float64)
			if  v1 != v2 { return v1<v2 }
		case string:
			v2:=vj.(string)
			if  v1 != v2 { return v1<v2 }
		case time.Time:
			v2:=vj.(time.Time)
			if  v1 != v2 { return v1.Before(v2) }
		}
	    }
	    return false
	}} )
	var Items []Item
 	for i,r := range cs { log.Printf("%d   %#v", i, *r) }
 	for _,r := range cs { 
		Items  =append(Items, Item{Col1: r.Col1.(int), Col2: r.Col2.(string)} )
	}
        Result_:=Result{Title:"Aazon.com", Items: Items }

 	for i,r := range Items { log.Printf("%d   %#v",i, r )}

	var issueList = template.Must(template.New("issuelist").Parse(`
		<h1>{{.Title}} issues</h1>
		<table>
		<tr style='text-align: left'>
		<th>#</th>
		<th>Col1</th>
		<th>Col2</th>
		</tr>
		{{range .Items}}
		<tr>
		<td>{{.Col1}}</td>
		<td>{{.Col2}}</td>
		</tr>
		{{end}}
		</table>
		`))

	if err := issueList.Execute(os.Stdout, Result_); err != nil {
		log.Fatal(err)
	}
}

//!-
