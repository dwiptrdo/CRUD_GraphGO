package utils

import (
	"graph-go/database"

	"github.com/graphql-go/graphql"
)

func Create(model interface{}, p graphql.ResolveParams, fields map[string]interface{}) error {
	for key, value := range fields {
		if val, ok := p.Args[key]; ok {
			switch field := value.(type) {
			case *string:
				*field = val.(string)
			case *int:
				*field = val.(int)
			}
		}
	}
	return database.DB.Create(model).Error
}

func Update(model interface{}, p graphql.ResolveParams, fields map[string]interface{}) error {
	for key, value := range fields {
		if val, ok := p.Args[key]; ok {
			switch field := value.(type) {
			case *string:
				*field = val.(string)
			case *int:
				*field = val.(int)
			}
		}
	}
	database.DB.Save(model)
	return nil
}

func Delete(model interface{}, id int) error {

	return database.DB.Where(id).Save(model).Error
}
