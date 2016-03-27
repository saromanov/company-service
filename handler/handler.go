package handler

import (

	//"github.com/micro/go-micro/errors"
	"github.com/saromanov/company-service/db"
	//"github.com/micro/go-platform/kv"
	"golang.org/x/net/context"

	company "github.com/saromanov/company-service/proto/company"
)

type Company struct{}

func (c*Company) Create(ctx context.Context, req *company.CreateRequest, rsp *company.CreateResponse) error {
	return db.Create(req.Company)
}

func (c*Company) Read(ctx context.Context, req *company.ReadRequest, rsp *company.ReadResponse) error {
	user, err := db.Read(req.Id)
	if err != nil {
		return err
	}
	rsp.Company = user
	return nil
}
