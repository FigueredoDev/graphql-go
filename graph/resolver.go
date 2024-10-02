package graph

import "github.com/FigueredoDev/graphql-go.git/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDb *database.Category
}
