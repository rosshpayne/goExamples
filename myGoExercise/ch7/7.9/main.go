package main

/* presume this is used by a single user.  Otherwise we would have to maintain state between individual clients.

*/
import (
	"net/http"
	"log"
	"time"
	"html/template"
	"strconv"
	"context"
	"sort"
	"sync"
	)

	const ( maxSortCols_ int = 2 )

	type  Sess struct {

	sync.Mutex					// Mutex protects these 3 members only
	sorting   bool
	colIndex  int
	colSort   [3]int

	sortCh	  chan struct{}
	cancelC	  context.CancelFunc
	}

	//var state_ map[cookie]Sess


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
                Items   []Item 
	}


	var loc, _ = time.LoadLocation("Australia/Sydney") 

	var cs []*rowT  = []*rowT{ { 10, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
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
	                           { 1, "ZBT", time.Date(2018,3,18,12,13,14,29,loc) },
				}

func main () {

	s:=NewSess()
	mux:=http.NewServeMux()

	//mux.Handle("/show",http.HandlerFunc(s.showTable))

	mux.HandleFunc("/show",s.showTable)
	mux.HandleFunc("/sort",s.sortTable)   					// Sort?col=1,2,3  or Sort?col=1 Sort?col=2 Sort?col=3
	
	log.Fatal(http.ListenAndServe("localhost:9080",mux))
}

func NewSess() Sess {

        return Sess{
        sorting   : false,
        colIndex  : 0,
	}
}



func (s *Sess) showTable(w http.ResponseWriter, req *http.Request) {

	var Items []Item				// Items allocated equal to nil. Add items using append which will make a slice on first execution
							//   alternatively use short declaration:   Items:=[]Item{}  with zero values
							//                                               :=[]Item{...} non-zero values
							//					    make(Items,10) 10 structs of zero values.

 	s.Lock()
	sorting:=s.sorting
	s.Unlock()
	if sorting {
		select {
	    		case <-s.sortCh:  						// wait for sort to finish
		}
	}	
	s.Lock()									// stop sorts
 	for i,r := range cs { log.Printf("%d   %#v", i, *r) }
 	for _,r := range cs { 
		Items  =append(Items, Item{Col1: r.Col1.(int), Col2: r.Col2.(string)} )
	}
	s.Unlock()
        Result_:=Result{Title:"Aazon.com", Items: Items }

	var customSortHTML = template.Must(template.New("issuelist").Parse(`
<h1>{{.Title}} issues</h1>
<customSort>
<head>
<style>	
table, th, td {
    border: 1px solid black;
}
</style>
</head>
<body>
<table>
<tr style='text-align: left'>
<th><a href="http://localhost:9080/sort?SortCol=1">Col1</a></th>
<th><a href="http://localhost:9080/sort?SortCol=2">Col2</a></th>
</tr>
{{range .Items}}
<tr>
<td>{{.Col1}}</td>
<td>{{.Col2}}</td>
</tr>
{{end}}
</table>
</body>
</customSort>
`))

	if err := customSortHTML.Execute(w, Result_); err != nil {		// render to client
		log.Fatal(err)
	}
}


	

func (s *Sess) sortTable (w http.ResponseWriter, req *http.Request) {
	// ignore request if already sorting..


	q:=req.URL.Query()
	s.Lock()
	switch s.sorting {
	case false:
 		if s.colIndex < maxSortCols_   {  							// colIndex 0,1,2 upto 3 columns can be specified for sorting
			var err error
			reqCol,err:=strconv.Atoi(q["SortCol"][0])			//..Sort?SortCol=2
			if err !=nil {
				log.Panic(err) 
			}
			if s.colIndex > 0 && reqCol == s.colSort[s.colIndex-1] {
				log.Println("Double click detected...sending close on context")
				s.cancelC()						// double tap detectd - commence sorting  
				break
			}
			for _,v := range s.colSort[:s.colIndex+1] {			// ignore input if already specified colun
				if v == reqCol {	
				    break
				}
			}
			s.colSort[s.colIndex]=reqCol
		}
		switch {
	 	case s.colIndex == 0:
			var ctxC context.Context					// use long declaration - see later. Note scope of these vars are local to case opton
			//var secs,_  = time.ParseDuration("3s") 			// long declaration, typically used for package var & when type differs from expression
			secs,_:=time.ParseDuration("3s") 				// use short declar - is neater. All local vars that are copied into asyncSort as a result 
			ctxT, _ := context.WithTimeout(context.Background(),secs)       //  of func by value on argument assignment end up taking a local copy of them
			ctxC, s.cancelC = context.WithCancel(context.Background())      // s.cancelC was declared outside this lexical block, so cannot be declared again 
											//    in short declar. So we are forced to predeclare ctxC in long form earlier..
			s.sortCh = make(chan struct{})					// previous channel will be GC'd
			// nb: anonymous func must appear after dependent func variables are declared.
			asyncSort := func () {
        					// block waiting on contexts
        					select {                                // Block and wait for
        					case <-ctxT.Done():                     //  ..select timeout
        					case <-ctxC.Done():                     //  ..supplied all columns
        					}
        					// perform sort
        					s.Lock()
        					s.sorting=true                          // all dml on s is barred when sorting=true
        					s.colIndex=0
        					s.Unlock()
        					s.performSort()
        					s.Lock()
        					s.sorting=false
        					for i,_ := range s.colSort {
                					s.colSort[i]=0
        					}
        					close(s.sortCh)
        					s.Unlock()
					      }
			go asyncSort()
		case s.colIndex > maxSortCols_-1:					
			s.cancelC()
		}
		s.colIndex+=1
	case true: 
	}
	s.Unlock()
	select {
	    	case <-s.sortCh :								// wait for sort to finish..
	}
 	s.showTable(w,req) 	
}


func (s *Sess) performSort () {
	// struct literal to define Sort argument
        sort.Sort(customSort{ tab: cs, less: func(i,j *rowT) bool { // use struct literal to initialise customSort arg in sort.Sort
            var vi,vj interface{}
	    // NB: s.colSort is being read outside of mutex. Should be safe as writes to s.colSort are barred during sort operation
            for _,v := range s.colSort {							
		if v == 0 {
			break
		}
		log.Println("in performSort.  Sort column ",s.colIndex,v)
                switch v {
                case 1 :
                        vi=i.Col1
                        vj=j.Col1
                case 2:
                        vi=i.Col2
                        vj=j.Col2
                case 3:
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
}
