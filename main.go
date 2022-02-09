package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/graphql-go/graphql"
)

func main() {
	tutorials := Populate()
	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Tutorial": &graphql.Field{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	var tutorialType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"author": &graphql.Field{
					Type: authorType,
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(commentType),
				},
			},
		},
	)

	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type: tutorialType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, tutorial := range tutorials {
						i,_:= strconv.Atoi(tutorial.ID)
						if i == id{
							return tutorial,nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type: graphql.NewList(tutorialType),
			Description: "Get full tutorial list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return tutorials,nil 
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemeConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemeConfig)
	// err != nil {
	// 	log.Fatalf("Error creating schema",err)
	// }
	query := `{ 
		tutorial(:id)
		list{
			id
			title
			author
		}
		}`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("fail", r.Errors)
	}
	rJson, _ := json.Marshal(r)
	fmt.Printf("%s ", rJson)
}
