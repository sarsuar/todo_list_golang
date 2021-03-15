package main

import (
	"log"
	"html/template"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string
	Done bool
}

func IsNotDone(todo Todo) bool {
	return !todo.Done
}

func main() {
	tmpl,err := template.New("template.html").Funcs(template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("template.html")
	if err != nil {
		log.Fatal("Cannot expand template",err);
		return
	}
	todos := []Todo{
		{"Ssad",false},
		{"asdasd",false},
		{"sadasdasd",false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
		if r.Method == http.MethodPost{
			param := r.FormValue("id")
			index,_ := strconv.ParseInt(param,10,0)
			todos[index].Done = true
		}

		err := tmpl.Execute(w,todos)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}