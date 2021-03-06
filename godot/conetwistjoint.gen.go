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

// ConeTwistJointParam is an enum for Param values.
type ConeTwistJointParam int

const (
	ConeTwistJointParamBias       ConeTwistJointParam = 2
	ConeTwistJointParamMax        ConeTwistJointParam = 5
	ConeTwistJointParamRelaxation ConeTwistJointParam = 4
	ConeTwistJointParamSoftness   ConeTwistJointParam = 3
	ConeTwistJointParamSwingSpan  ConeTwistJointParam = 0
	ConeTwistJointParamTwistSpan  ConeTwistJointParam = 1
)

//func NewConeTwistJointFromPointer(ptr gdnative.Pointer) ConeTwistJoint {
func newConeTwistJointFromPointer(ptr gdnative.Pointer) ConeTwistJoint {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ConeTwistJoint{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type ConeTwistJoint struct {
	Joint
	owner gdnative.Object
}

func (o *ConeTwistJoint) BaseClass() string {
	return "ConeTwistJoint"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *ConeTwistJoint) X_GetSwingSpan() gdnative.Real {
	//log.Println("Calling ConeTwistJoint.X_GetSwingSpan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConeTwistJoint", "_get_swing_span")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *ConeTwistJoint) X_GetTwistSpan() gdnative.Real {
	//log.Println("Calling ConeTwistJoint.X_GetTwistSpan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConeTwistJoint", "_get_twist_span")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false swing_span float}], Returns: void
*/
func (o *ConeTwistJoint) X_SetSwingSpan(swingSpan gdnative.Real) {
	//log.Println("Calling ConeTwistJoint.X_SetSwingSpan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(swingSpan)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConeTwistJoint", "_set_swing_span")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false twist_span float}], Returns: void
*/
func (o *ConeTwistJoint) X_SetTwistSpan(twistSpan gdnative.Real) {
	//log.Println("Calling ConeTwistJoint.X_SetTwistSpan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(twistSpan)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConeTwistJoint", "_set_twist_span")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int}], Returns: float
*/
func (o *ConeTwistJoint) GetParam(param gdnative.Int) gdnative.Real {
	//log.Println("Calling ConeTwistJoint.GetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConeTwistJoint", "get_param")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *ConeTwistJoint) SetParam(param gdnative.Int, value gdnative.Real) {
	//log.Println("Calling ConeTwistJoint.SetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConeTwistJoint", "set_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ConeTwistJointImplementer is an interface that implements the methods
// of the ConeTwistJoint class.
type ConeTwistJointImplementer interface {
	JointImplementer
	X_GetSwingSpan() gdnative.Real
	X_GetTwistSpan() gdnative.Real
	X_SetSwingSpan(swingSpan gdnative.Real)
	X_SetTwistSpan(twistSpan gdnative.Real)
	GetParam(param gdnative.Int) gdnative.Real
	SetParam(param gdnative.Int, value gdnative.Real)
}
