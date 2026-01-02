package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Struktur data untuk template
type Data struct {
	Name      string
	LowerName string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run tools/generator.go [Name]")
		return
	}

	name := os.Args[1]
	data := Data{
		Name:      name,
		LowerName: strings.ToLower(name),
	}

	// file yang akan dibuat
	files := map[string]string{
		"repositories/" + data.LowerName + "_repository.go":      repoInterfaceTemplate,
		"repositories/" + data.LowerName + "_repository_impl.go": repoImplTemplate,
		"services/" + data.LowerName + "_service.go":             serviceTemplate,
		"controllers/" + data.LowerName + "_controller.go":       controllerTemplate,
	}

	for path, temp := range files {
		t := template.Must(template.New("gen").Parse(temp))
		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("Gagal buat file %s: %v\n", path, err)
			continue
		}
		t.Execute(f, data)
		f.Close()
		fmt.Printf("Generated: %s\n", path)
	}
}

// --- TEMPLATES ---

const repoInterfaceTemplate = `package repositories
import (
	"context"
	"github.com/cardinaldeacre/OpenMusic-API-GoLang/model/domain"
)

type {{.Name}}Repository interface {
	Create(ctx context.Context, data *domain.{{.Name}}) error
	// Tambahkan method lainnya di sini
}`

const repoImplTemplate = `package repositories
import (
	"context"
	"github.com/cardinaldeacre/OpenMusic-API-GoLang/model/domain"
	"gorm.io/gorm"
)

type {{.Name}}RepositoryImpl struct { db *gorm.DB }

func New{{.Name}}Repository(db *gorm.DB) {{.Name}}Repository {
	return &{{.Name}}RepositoryImpl{db}
}

func (r *{{.Name}}RepositoryImpl) Create(ctx context.Context, data *domain.{{.Name}}) error {
	return r.db.WithContext(ctx).Create(data).Error
}`

const serviceTemplate = `package services
import "github.com/cardinaldeacre/OpenMusic-API-GoLang/repositories"

type {{.Name}}Service interface {}
type {{.LowerName}}Service struct { repo repositories.{{.Name}}Repository }

func New{{.Name}}Service(repo repositories.{{.Name}}Repository) {{.Name}}Service {
	return &{{.LowerName}}Service{repo}
}`

const controllerTemplate = `package controllers
import (
	"github.com/cardinaldeacre/OpenMusic-API-GoLang/services"
	"github.com/gofiber/fiber/v2"
)

type {{.Name}}Controller struct { Service services.{{.Name}}Service }

func New{{.Name}}Controller(s services.{{.Name}}Service) *{{.Name}}Controller {
	return &{{.Name}}Controller{Service: s}
}`
