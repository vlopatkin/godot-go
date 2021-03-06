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

// VisualShaderNodeTextureUniformColorDefault is an enum for ColorDefault values.
type VisualShaderNodeTextureUniformColorDefault int

const (
	VisualShaderNodeTextureUniformColorDefaultBlack VisualShaderNodeTextureUniformColorDefault = 1
	VisualShaderNodeTextureUniformColorDefaultWhite VisualShaderNodeTextureUniformColorDefault = 0
)

// VisualShaderNodeTextureUniformTextureType is an enum for TextureType values.
type VisualShaderNodeTextureUniformTextureType int

const (
	VisualShaderNodeTextureUniformTypeAniso     VisualShaderNodeTextureUniformTextureType = 3
	VisualShaderNodeTextureUniformTypeColor     VisualShaderNodeTextureUniformTextureType = 1
	VisualShaderNodeTextureUniformTypeData      VisualShaderNodeTextureUniformTextureType = 0
	VisualShaderNodeTextureUniformTypeNormalmap VisualShaderNodeTextureUniformTextureType = 2
)

//func NewVisualShaderNodeTextureUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeTextureUniform {
func newVisualShaderNodeTextureUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeTextureUniform {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeTextureUniform{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeTextureUniform struct {
	VisualShaderNodeUniform
	owner gdnative.Object
}

func (o *VisualShaderNodeTextureUniform) BaseClass() string {
	return "VisualShaderNodeTextureUniform"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeTextureUniform::ColorDefault
*/
func (o *VisualShaderNodeTextureUniform) GetColorDefault() VisualShaderNodeTextureUniformColorDefault {
	//log.Println("Calling VisualShaderNodeTextureUniform.GetColorDefault()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeTextureUniform", "get_color_default")

	// Call the parent method.
	// enum.VisualShaderNodeTextureUniform::ColorDefault
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeTextureUniformColorDefault(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeTextureUniform::TextureType
*/
func (o *VisualShaderNodeTextureUniform) GetTextureType() VisualShaderNodeTextureUniformTextureType {
	//log.Println("Calling VisualShaderNodeTextureUniform.GetTextureType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeTextureUniform", "get_texture_type")

	// Call the parent method.
	// enum.VisualShaderNodeTextureUniform::TextureType
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeTextureUniformTextureType(ret)
}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/
func (o *VisualShaderNodeTextureUniform) SetColorDefault(aType gdnative.Int) {
	//log.Println("Calling VisualShaderNodeTextureUniform.SetColorDefault()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeTextureUniform", "set_color_default")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/
func (o *VisualShaderNodeTextureUniform) SetTextureType(aType gdnative.Int) {
	//log.Println("Calling VisualShaderNodeTextureUniform.SetTextureType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeTextureUniform", "set_texture_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeTextureUniformImplementer is an interface that implements the methods
// of the VisualShaderNodeTextureUniform class.
type VisualShaderNodeTextureUniformImplementer interface {
	VisualShaderNodeUniformImplementer
	SetColorDefault(aType gdnative.Int)
	SetTextureType(aType gdnative.Int)
}
