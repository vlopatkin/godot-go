package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVisualScriptIndexSetFromPointer(ptr gdnative.Pointer) VisualScriptIndexSet {
func newVisualScriptIndexSetFromPointer(ptr gdnative.Pointer) VisualScriptIndexSet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptIndexSet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptIndexSet struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptIndexSet) BaseClass() string {
	return "VisualScriptIndexSet"
}

// VisualScriptIndexSetImplementer is an interface that implements the methods
// of the VisualScriptIndexSet class.
type VisualScriptIndexSetImplementer interface {
	VisualScriptNodeImplementer
}
