package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/kwa0x2/GoLang-Postgre-API/response"
	"io/ioutil"
	"log"
	"net/http"
)

var request response.Request
var requestBody string

func RoutePost(ctx *gin.Context) {

	/*	err := ctx.BindJSON(&requestBody)
		if err != nil {
			log.Fatalf("error: %v", err)
		}*/
	/*	err := ctx.BindJSON(&request)
		if err != nil {
			log.Fatalf("error: %v", err)
		}*/

	bodyAsByteArray, _ := ioutil.ReadAll(ctx.Request.Body)
	jsonBody := string(bodyAsByteArray)

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"entity": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					//nameQuery, isOK := params.Args["name"].(string)
					//if isOK {
					//	nameQuery
					//}

					return nil, nil

				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: rootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	params := graphql.Params{Schema: schema, RequestString: jsonBody}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   &r,
	})

}
