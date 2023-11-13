package service

import (
	"github.com/graphql-go/graphql"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"github.com/kwa0x2/GoLang-Postgre-API/repo"
	"log"
)

var requestBody string

var user = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var users []models.User

func RoutePost() graphql.Schema {

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: user,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					nameQuery, isOK := params.Args["id"].(int)
					if isOK {

						user := repo.GetUserByID(nameQuery)
						return user, nil
					}

					return nil, nil

				},
			},
			"users": &graphql.Field{
				Type: graphql.NewList(user),
				Args: graphql.FieldConfigArgument{
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"size": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					page, isOK := params.Args["page"].(int)
					size, isOK := params.Args["size"].(int)
					if isOK {
						user := repo.GetUsersPageable(page, size)
						return user, nil
					}

					return nil, nil

				},
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: user, // the return type for this field
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					// marshall and cast the argument value
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)

					newUser := &models.User{
						Username: username,
						Password: password,
					}

					savedUser := repo.Create(newUser)

					return savedUser, nil
				},
			},
			"updateUser": &graphql.Field{
				Type: user, // the return type for this field
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					// marshall and cast the argument value
					id, _ := params.Args["id"].(int)
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)

					newUser := &models.User{
						ID:       id,
						Username: username,
						Password: password,
					}

					savedUser := repo.Update(newUser)

					return savedUser, nil
				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: rootQuery, Mutation: rootMutation}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema

}
