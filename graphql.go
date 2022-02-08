package main 


type Tutorial struct{
	ID string
	Title string
	Author Author
	Comments []Comment
}

type Author struct{
	Name string
	Toturials []int
}
type Comment struct{
	Body string 
}

func Populate() []Tutorial{
	author := Author{Name : "Chinonso", Toturials: []int{1}}
	tutorial := Tutorial{
		Author: author, 
		Title: "Learning Graphql",
		ID: "1",
		Comments: []Comment{
			{Body :"First Comment"},
		}, 
	}
	var t []Tutorial
	t = append(t, tutorial)
	return t
}
