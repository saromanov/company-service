package main

import (
    "fmt"

    micro "github.com/micro/go-micro"
    proto "github.com/saromanov/company-service/proto/company"
    "golang.org/x/net/context"
)


func main() {
    // Create a new service. Optionally include some options here.
    service := micro.NewService(micro.Name("go.micro.srv.company"))

    // Create new greeter client
    greeter := proto.NewCompanyClient("go.micro.srv.company", service.Client())

    req := &proto.CompanyItem{Name:"big"}
    // Call the greeter
    rsp, err := greeter.Create(context.TODO(), &proto.CreateRequest{Company: req})
    if err != nil {
        fmt.Println(err)
    }

    // Print response
    fmt.Println(rsp)
}