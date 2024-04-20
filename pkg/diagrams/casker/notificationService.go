package casker

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/blushft/go-diagrams/nodes/gcp"
	"github.com/blushft/go-diagrams/nodes/saas"
)

func GenerateNotificationService(filename string) error {
	d, err := diagram.New(diagram.Filename(filename), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}
	user := aws.General.User(diagram.NodeLabel("User"))
	db := gcp.Database.Sql(diagram.NodeLabel("Database"))
	front := gcp.Api.Endpoints(diagram.NodeLabel("Gateway"))
	server := gcp.Compute.ComputeEngine(diagram.NodeLabel("API"))
	tasks := gcp.Devtools.Tasks(diagram.NodeLabel("Tasks"))
	telegram := saas.Chat.Telegram(diagram.NodeLabel("Telegram BOT"))

	app := diagram.NewGroup("App")
	app.NewGroup("services").
		Add(server, tasks)

	app.Add(db)

	d.Connect(server, tasks, diagram.Bidirectional())
	d.Connect(front, server, diagram.Forward()).Group(app)
	d.Connect(server, db, diagram.Bidirectional())
	d.Connect(tasks, db, diagram.Reverse())
	d.Connect(server, telegram, diagram.Forward())
	d.Connect(user, front, diagram.Forward())

	if err := d.Render(); err != nil {
		return err
	}
	return nil
}
