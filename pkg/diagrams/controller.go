package diagrams

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pstano1/diagrams.git/pkg/diagrams/casker"
	gocart "github.com/pstano1/diagrams.git/pkg/diagrams/go-cart"
	"go.uber.org/zap"
)

const (
	dotDirectoryPath = "./go-diagrams/"
	assetsBasePath   = "./pkg/diagrams/"
)

type IDiagrams interface {
	GenerateDiagrams()

	generateDiagram(filePath string, generator func(string) error, getPath func() string, name string)
}

type Diagrams struct {
	log *zap.Logger
}

func New(logger *zap.Logger) IDiagrams {
	return &Diagrams{
		log: logger,
	}
}

func (d *Diagrams) GenerateDiagrams() {
	d.generateDiagram("notificationService.dot", casker.GenerateNotificationService, casker.GetAssetsPath, "notificationService")

	d.generateDiagram("customerAPISchema.dot", gocart.GenerateCustomerAPISchema, gocart.GetAssetsPath, "customerAPISchema")
}

func (d *Diagrams) generateDiagram(filename string, generator func(string) error, getPath func() string, name string) {
	d.log.Info("generating dot file",
		zap.String("file", name),
	)
	err := generator(name)
	if err != nil {
		d.log.Error("error while generating dot file",
			zap.Error(err),
		)
		return
	}
	generatePNG := exec.Command("sh", "-c", fmt.Sprintf("cd %s && dot -Tpng %s > %s.png", dotDirectoryPath, filename, name))
	if err := generatePNG.Run(); err != nil {
		d.log.Error("error while generating png file",
			zap.Error(err),
		)
		return
	}
	wd, _ := os.Getwd()
	src := filepath.Join(wd, dotDirectoryPath, fmt.Sprintf("%s.png", name))
	dst := filepath.Join(wd, assetsBasePath, getPath(), fmt.Sprintf("%s.png", name))
	if err := os.Rename(src, dst); err != nil {
		return
	}
}
