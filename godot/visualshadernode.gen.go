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

// VisualShaderNodePortType is an enum for PortType values.
type VisualShaderNodePortType int

const (
	VisualShaderNodePortTypeBoolean   VisualShaderNodePortType = 2
	VisualShaderNodePortTypeMax       VisualShaderNodePortType = 5
	VisualShaderNodePortTypeSampler   VisualShaderNodePortType = 4
	VisualShaderNodePortTypeScalar    VisualShaderNodePortType = 0
	VisualShaderNodePortTypeTransform VisualShaderNodePortType = 3
	VisualShaderNodePortTypeVector    VisualShaderNodePortType = 1
)

//func NewVisualShaderNodeFromPointer(ptr gdnative.Pointer) VisualShaderNode {
func newVisualShaderNodeFromPointer(ptr gdnative.Pointer) VisualShaderNode {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNode{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNode struct {
	Resource
	owner gdnative.Object
}

func (o *VisualShaderNode) BaseClass() string {
	return "VisualShaderNode"
}

/*
        Returns an [Array] containing default values for all of the input ports of the node in the form [code][index0, value0, index1, value1, ...][/code].
	Args: [], Returns: Array
*/
func (o *VisualShaderNode) GetDefaultInputValues() gdnative.Array {
	//log.Println("Calling VisualShaderNode.GetDefaultInputValues()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNode", "get_default_input_values")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the default value of the input [code]port[/code].
	Args: [{ false port int}], Returns: Variant
*/
func (o *VisualShaderNode) GetInputPortDefaultValue(port gdnative.Int) gdnative.Variant {
	//log.Println("Calling VisualShaderNode.GetInputPortDefaultValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNode", "get_input_port_default_value")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *VisualShaderNode) GetOutputPortForPreview() gdnative.Int {
	//log.Println("Calling VisualShaderNode.GetOutputPortForPreview()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNode", "get_output_port_for_preview")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Sets the default input ports values using an [Array] of the form [code][index0, value0, index1, value1, ...][/code]. For example: [code][0, Vector3(0, 0, 0), 1, Vector3(0, 0, 0)][/code].
	Args: [{ false values Array}], Returns: void
*/
func (o *VisualShaderNode) SetDefaultInputValues(values gdnative.Array) {
	//log.Println("Calling VisualShaderNode.SetDefaultInputValues()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(values)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNode", "set_default_input_values")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the default value for the selected input [code]port[/code].
	Args: [{ false port int} { false value Variant}], Returns: void
*/
func (o *VisualShaderNode) SetInputPortDefaultValue(port gdnative.Int, value gdnative.Variant) {
	//log.Println("Calling VisualShaderNode.SetInputPortDefaultValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNode", "set_input_port_default_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false port int}], Returns: void
*/
func (o *VisualShaderNode) SetOutputPortForPreview(port gdnative.Int) {
	//log.Println("Calling VisualShaderNode.SetOutputPortForPreview()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNode", "set_output_port_for_preview")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeImplementer is an interface that implements the methods
// of the VisualShaderNode class.
type VisualShaderNodeImplementer interface {
	ResourceImplementer
	GetDefaultInputValues() gdnative.Array
	GetInputPortDefaultValue(port gdnative.Int) gdnative.Variant
	GetOutputPortForPreview() gdnative.Int
	SetDefaultInputValues(values gdnative.Array)
	SetInputPortDefaultValue(port gdnative.Int, value gdnative.Variant)
	SetOutputPortForPreview(port gdnative.Int)
}
