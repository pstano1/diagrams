package gocart

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/blushft/go-diagrams/nodes/gcp"
)

func GenerateCustomerAPISchema(filename string) error {
	d, err := diagram.New(diagram.Filename(filename), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}
	user := aws.General.User(diagram.NodeLabel("User"))

	front := gcp.Api.Endpoints(diagram.NodeLabel("Gateway"))
	API := gcp.Compute.ComputeEngine(diagram.NodeLabel("API"))
	customersAPI := gcp.Compute.ComputeEngine(diagram.NodeLabel("customer's API"))

	app := diagram.NewGroup("App")

	d.Connect(API, customersAPI, diagram.Bidirectional())
	d.Connect(front, API, diagram.Bidirectional()).Group(app)
	d.Connect(user, front, diagram.Forward())

	if err := d.Render(); err != nil {
		return err
	}
	return nil
}
