package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVisualScriptIndexGetFromPointer(ptr gdnative.Pointer) VisualScriptIndexGet {
func NewVisualScriptIndexGetFromPointer(ptr gdnative.Pointer) VisualScriptIndexGet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptIndexGet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptIndexGet struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptIndexGet) BaseClass() string {
	return "VisualScriptIndexGet"
}