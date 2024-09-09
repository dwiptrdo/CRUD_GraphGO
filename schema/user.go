package schema

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"

	database "graph-go/database"
	models "graph-go/model"
	util "graph-go/utils"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"created_by": &graphql.Field{
				Type: graphql.String,
			},
			"updated_by": &graphql.Field{
				Type: graphql.String,
			},
			"nama": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var UserQuery = graphql.Fields{
	"user": &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(int)
			var user models.User
			result := database.DB.Where("status = ?", "active").First(&user, id)
			if result.RecordNotFound() {
				return nil, fmt.Errorf("Data Tidak Ditemukan!")
			}
			return user, nil
		},
	},
	"users": &graphql.Field{
		Type: graphql.NewList(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var users []models.User
			database.DB.Where("status = ?", "active").Find(&users)
			return users, nil
		},
	},
}

var UserMutation = graphql.Fields{
	"createUser": &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"created_by": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"nama": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"kata_sandi": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"username": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var user models.User

			timeNow := time.Now()
			user.CreatedAt = timeNow
			user.UpdatedAt = timeNow

			fields := map[string]interface{}{
				"created_by": &user.CreatedBy,
				"nama":       &user.Nama,
				"kata_sandi": &user.Kata_Sandi,
				"username":   &user.Username,
				"email":      &user.Email,
			}

			err := util.Create(&user, p, fields)
			if err != nil {
				return nil, err
			}
			return user, nil
		},
	},
	"updateUser": &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"updated_by": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"nama": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"kata_sandi": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"username": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(int)
			var user models.User
			database.DB.First(&user, id)

			timeNow := time.Now()
			user.UpdatedAt = timeNow

			fields := map[string]interface{}{
				"updated_by": &user.UpdatedBy,
				"nama":       &user.Nama,
				"kata_sandi": &user.Kata_Sandi,
				"username":   &user.Username,
				"email":      &user.Email,
			}
			err := util.Update(&user, p, fields)
			if err != nil {
				return nil, err
			}
			return user, nil
		},
	},
	"deleteUser": &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(int)
			var user models.User
			database.DB.First(&user, id)
			user.Status = "inactive"

			err := util.Delete(&user, id)
			if err != nil {
				return false, err
			}
			return fmt.Sprintf("User dengan username %s sukses dihapus!", user.Username), nil
		},
	},
}
