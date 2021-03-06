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

//func NewRayCastFromPointer(ptr gdnative.Pointer) RayCast {
func newRayCastFromPointer(ptr gdnative.Pointer) RayCast {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RayCast{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type RayCast struct {
	Spatial
	owner gdnative.Object
}

func (o *RayCast) BaseClass() string {
	return "RayCast"
}

/*
        Undocumented
	Args: [{ false node Object}], Returns: void
*/
func (o *RayCast) AddException(node ObjectImplementer) {
	//log.Println("Calling RayCast.AddException()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "add_exception")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rid RID}], Returns: void
*/
func (o *RayCast) AddExceptionRid(rid gdnative.Rid) {
	//log.Println("Calling RayCast.AddExceptionRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(rid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "add_exception_rid")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *RayCast) ClearExceptions() {
	//log.Println("Calling RayCast.ClearExceptions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "clear_exceptions")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *RayCast) ForceRaycastUpdate() {
	//log.Println("Calling RayCast.ForceRaycastUpdate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "force_raycast_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *RayCast) GetCastTo() gdnative.Vector3 {
	//log.Println("Calling RayCast.GetCastTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_cast_to")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Object
*/
func (o *RayCast) GetCollider() ObjectImplementer {
	//log.Println("Calling RayCast.GetCollider()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collider")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *RayCast) GetColliderShape() gdnative.Int {
	//log.Println("Calling RayCast.GetColliderShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collider_shape")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *RayCast) GetCollisionMask() gdnative.Int {
	//log.Println("Calling RayCast.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false bit int}], Returns: bool
*/
func (o *RayCast) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling RayCast.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_mask_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *RayCast) GetCollisionNormal() gdnative.Vector3 {
	//log.Println("Calling RayCast.GetCollisionNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *RayCast) GetCollisionPoint() gdnative.Vector3 {
	//log.Println("Calling RayCast.GetCollisionPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_point")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RayCast) GetExcludeParentBody() gdnative.Bool {
	//log.Println("Calling RayCast.GetExcludeParentBody()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_exclude_parent_body")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RayCast) IsCollideWithAreasEnabled() gdnative.Bool {
	//log.Println("Calling RayCast.IsCollideWithAreasEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_collide_with_areas_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RayCast) IsCollideWithBodiesEnabled() gdnative.Bool {
	//log.Println("Calling RayCast.IsCollideWithBodiesEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_collide_with_bodies_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RayCast) IsColliding() gdnative.Bool {
	//log.Println("Calling RayCast.IsColliding()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_colliding")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RayCast) IsEnabled() gdnative.Bool {
	//log.Println("Calling RayCast.IsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false node Object}], Returns: void
*/
func (o *RayCast) RemoveException(node ObjectImplementer) {
	//log.Println("Calling RayCast.RemoveException()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "remove_exception")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rid RID}], Returns: void
*/
func (o *RayCast) RemoveExceptionRid(rid gdnative.Rid) {
	//log.Println("Calling RayCast.RemoveExceptionRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(rid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "remove_exception_rid")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false local_point Vector3}], Returns: void
*/
func (o *RayCast) SetCastTo(localPoint gdnative.Vector3) {
	//log.Println("Calling RayCast.SetCastTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(localPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_cast_to")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RayCast) SetCollideWithAreas(enable gdnative.Bool) {
	//log.Println("Calling RayCast.SetCollideWithAreas()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collide_with_areas")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RayCast) SetCollideWithBodies(enable gdnative.Bool) {
	//log.Println("Calling RayCast.SetCollideWithBodies()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collide_with_bodies")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *RayCast) SetCollisionMask(mask gdnative.Int) {
	//log.Println("Calling RayCast.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *RayCast) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling RayCast.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *RayCast) SetEnabled(enabled gdnative.Bool) {
	//log.Println("Calling RayCast.SetEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask bool}], Returns: void
*/
func (o *RayCast) SetExcludeParentBody(mask gdnative.Bool) {
	//log.Println("Calling RayCast.SetExcludeParentBody()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_exclude_parent_body")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RayCastImplementer is an interface that implements the methods
// of the RayCast class.
type RayCastImplementer interface {
	SpatialImplementer
	AddException(node ObjectImplementer)
	AddExceptionRid(rid gdnative.Rid)
	ClearExceptions()
	ForceRaycastUpdate()
	GetCastTo() gdnative.Vector3
	GetCollider() ObjectImplementer
	GetColliderShape() gdnative.Int
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	GetCollisionNormal() gdnative.Vector3
	GetCollisionPoint() gdnative.Vector3
	GetExcludeParentBody() gdnative.Bool
	IsCollideWithAreasEnabled() gdnative.Bool
	IsCollideWithBodiesEnabled() gdnative.Bool
	IsColliding() gdnative.Bool
	IsEnabled() gdnative.Bool
	RemoveException(node ObjectImplementer)
	RemoveExceptionRid(rid gdnative.Rid)
	SetCastTo(localPoint gdnative.Vector3)
	SetCollideWithAreas(enable gdnative.Bool)
	SetCollideWithBodies(enable gdnative.Bool)
	SetCollisionMask(mask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
	SetEnabled(enabled gdnative.Bool)
	SetExcludeParentBody(mask gdnative.Bool)
}
