package main

type Extension struct {
	Element
}

type Element struct {
	Id        string      `json:"id"`
	Extension []Extension `json:"extension,omitempty"`
}

type DomainResource struct {
	Element
	Extension []Extension `json:"extension,omitempty"`
}

type Resource struct {
	DomainResource
}

// Resource godoc
// @Summary get
// @Description get resource
//
// @Accept json
// @Param Authorization header string true "access token sent using Bearer prefix"
// @Param ID path string true "ID"
//
// @Success 200 {object} Resource
// @Failure 400
// @Tags resource
// @Router /Resource/{ID} [get]
func _() {}

// @title Medsensio Demo
// @version 0.0.1
// @description Swaggo Crash Demo
// @license.name Apache 2, doesn't really matter
//
// @contact.name medsensio
// @contact.uti https://medsens.io
// @contact.email someome@medsens.io
//
// @BasePath /
func main() {}
