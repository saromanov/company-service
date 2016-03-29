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

    req := &proto.CompanyItem{Name:"big", Owner: "me"}
    // Call the greeter
    rsp, err := greeter.Create(context.TODO(), &proto.CreateRequest{Company: req})
    if err != nil {
        fmt.Println(err)
        return
    }

    // Read item

    body, err := greeter.Read(context.TODO(), &proto.ReadRequest{Id: rsp.Id})
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(body)
}