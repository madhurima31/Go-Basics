package main

import ("fmt" 
	    "net/http"
	     )



func indexPage_handler( w http.ResponseWriter , r *http.Request){
	fmt.Fprintf( w , "Home page")
}

func main() {
	http.HandleFunc("/", indexPage_handler)
	http.ListenAndServe(":8080", nil)
}