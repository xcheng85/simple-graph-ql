package graph

import (
	"github.com/xcheng85/simple-graphql-go/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	playersMap map[string]*model.Player
}
