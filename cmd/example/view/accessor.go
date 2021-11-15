package view

import (
	"log"

	"fyne.io/fyne/v2"
	"github.com/drognisep/fynehelpers/generation"
)

func NewTypeBaseNode() fyne.CanvasObject {
	node := newTypeBaseNode()
	return node
}

func Update(obj fyne.CanvasObject, data generation.TreeModel) {
	node, ok := obj.(*typeBaseNode)
	if !ok {
		log.Println("Not a typeBaseNode")
	}
	node.update("123", data)
}
