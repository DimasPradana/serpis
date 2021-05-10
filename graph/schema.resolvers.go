package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/DimasPradana/kantor/serpis/config"
	"github.com/DimasPradana/kantor/serpis/graph/generated"
	"github.com/DimasPradana/kantor/serpis/graph/model"
	"github.com/sirupsen/logrus"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Datlogin(ctx context.Context, nmLogin string) (*model.DatLogin, error) {

	config.GetEnvfile()
	pbbuser := os.Getenv("pbbuser")
	pbbpassword := os.Getenv("pbbpassword")
	pbbhost := os.Getenv("pbbhost")
	pbbport := os.Getenv("pbbport")
	pbbservicename := os.Getenv("pbbservicename")

	var namalogin, nip, password string
	var datlogin *model.DatLogin

	// query
	// qry := fmt.Sprintf("select * from dat_login where NM_LOGIN like '%%%v%%'", strings.ToUpper(nmLogin))
	qry := fmt.Sprintf("select * from dat_login where NM_LOGIN = '%v'", strings.ToUpper(nmLogin))
	logrus.Infof("print query : %v", qry)

	// build koneksi
	kon, _ := config.KonekOracle(pbbuser, pbbpassword, pbbhost, pbbport,
		pbbservicename)

	defer kon.Close()

	rows, err := kon.Query(qry)

	if err != nil {
		logrus.Infof("errornya di baris 50: %v\n%v", err.Error(), qry)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&namalogin,
			&nip,
			&password); err != nil {
			logrus.Infof("NM_LOGIN: %v, NIP: %v, PASSWORD: %v\n", namalogin, nip, password)
		}
		datlogin = &model.DatLogin{
			NmLogin:  namalogin,
			Nip:      nip,
			Password: password,
		}
	}

	return datlogin, nil
}

func (r *queryResolver) Datlogins(ctx context.Context) ([]*model.DatLogin, error) {

	config.GetEnvfile()
	pbbuser := os.Getenv("pbbuser")
	pbbpassword := os.Getenv("pbbpassword")
	pbbhost := os.Getenv("pbbhost")
	pbbport := os.Getenv("pbbport")
	pbbservicename := os.Getenv("pbbservicename")

	var namalogin, nip, password string
	var datlogins []*model.DatLogin

	// query
	qry := fmt.Sprintf("select * from dat_login")

	// build koneksi
	kon, _ := config.KonekOracle(pbbuser, pbbpassword, pbbhost, pbbport,
		pbbservicename)
	defer kon.Close()
	rows, err := kon.Query(qry)
	if err != nil {
		logrus.Infof("errornya di baris 93: %v\n%v", err.Error(), qry)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(
			&namalogin,
			&nip,
			&password); err != nil {
			logrus.Infof("NM_LOGIN: %v, NIP: %v, PASSWORD: %v\n", namalogin, nip, password)
		}
		datlogins = append(datlogins, &model.DatLogin{NmLogin: namalogin, Nip: nip, Password: password})

	}

	return datlogins, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
