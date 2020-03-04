package gql

import (
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/kumarankeerthi/go-learning/graphql-go/domain"
)

var CustomerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Customer",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"firstname": &graphql.Field{
			Type: graphql.String,
		},
		"lastname": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var TicketType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Ticket",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"traveldate": &graphql.Field{
			Type: graphql.String,
		},
		"customer": &graphql.Field{
			Type: CustomerType,
		},
		"route": &graphql.Field{
			Type: RouteType,
		},
	},
})

var RouteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Route",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"source": &graphql.Field{
			Type: graphql.String,
		},
		"destination": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ticket": &graphql.Field{
			Type: graphql.NewList(TicketType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				tid := params.Args["id"].(string)
				filteredList := Filter(domain.Tickets, func(t domain.Ticket) bool {
					return strings.EqualFold(t.ID, tid)
				})
				return filteredList, nil
			},
		},
	},
})

func Filter(ts []domain.Ticket, f func(domain.Ticket) bool) []domain.Ticket {
	filtered := make([]domain.Ticket, 0)
	for _, t := range ts {
		if f(t) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
