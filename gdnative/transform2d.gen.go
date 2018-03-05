package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include "gdnative.gen.h"
#include <gdnative/transform2d.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

// NewEmptyTransform2D will return a pointer to an empty
// initialized Transform2D. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyTransform2D() Pointer {
	var obj C.godot_transform2d
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromTransform2D will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromTransform2D(obj Transform2D) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewTransform2DFromPointer will return a Transform2D from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewTransform2DFromPointer(ptr Pointer) Transform2D {

	return Transform2D{base: (*C.godot_transform2d)(ptr.getBase())}
}

type Transform2D struct {
	base *C.godot_transform2d
}

func (gdt Transform2D) getBase() *C.godot_transform2d {
	return gdt.base
}

// AsString godot_transform2d_as_string [[const godot_transform2d * p_self]] godot_string
func (gdt *Transform2D) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform2d_as_string(GDNative.api, arg0)

	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharT(wchar)
	return String(goWchar.AsString())

}

// Inverse godot_transform2d_inverse [[const godot_transform2d * p_self]] godot_transform2d
func (gdt *Transform2D) Inverse() Transform2D {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform2d_inverse(GDNative.api, arg0)

	return Transform2D{base: &ret}

}

// AffineInverse godot_transform2d_affine_inverse [[const godot_transform2d * p_self]] godot_transform2d
func (gdt *Transform2D) AffineInverse() Transform2D {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform2d_affine_inverse(GDNative.api, arg0)

	return Transform2D{base: &ret}

}

// GetRotation godot_transform2d_get_rotation [[const godot_transform2d * p_self]] godot_real
func (gdt *Transform2D) GetRotation() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform2d_get_rotation(GDNative.api, arg0)

	return Real(ret)
}

// GetOrigin godot_transform2d_get_origin [[const godot_transform2d * p_self]] godot_vector2
func (gdt *Transform2D) GetOrigin() Vector2 {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform2d_get_origin(GDNative.api, arg0)

	return Vector2{base: &ret}

}

// GetScale godot_transform2d_get_scale [[const godot_transform2d * p_self]] godot_vector2
func (gdt *Transform2D) GetScale() Vector2 {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform2d_get_scale(GDNative.api, arg0)

	return Vector2{base: &ret}

}

// Orthonormalized godot_transform2d_orthonormalized [[const godot_transform2d * p_self]] godot_transform2d
func (gdt *Transform2D) Orthonormalized() Transform2D {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform2d_orthonormalized(GDNative.api, arg0)

	return Transform2D{base: &ret}

}

// Rotated godot_transform2d_rotated [[const godot_transform2d * p_self] [const godot_real p_phi]] godot_transform2d
func (gdt *Transform2D) Rotated(phi Real) Transform2D {
	arg0 := gdt.getBase()
	arg1 := phi.getBase()

	ret := C.go_godot_transform2d_rotated(GDNative.api, arg0, arg1)

	return Transform2D{base: &ret}

}

// Scaled godot_transform2d_scaled [[const godot_transform2d * p_self] [const godot_vector2 * p_scale]] godot_transform2d
func (gdt *Transform2D) Scaled(scale Vector2) Transform2D {
	arg0 := gdt.getBase()
	arg1 := scale.getBase()

	ret := C.go_godot_transform2d_scaled(GDNative.api, arg0, arg1)

	return Transform2D{base: &ret}

}

// Translated godot_transform2d_translated [[const godot_transform2d * p_self] [const godot_vector2 * p_offset]] godot_transform2d
func (gdt *Transform2D) Translated(offset Vector2) Transform2D {
	arg0 := gdt.getBase()
	arg1 := offset.getBase()

	ret := C.go_godot_transform2d_translated(GDNative.api, arg0, arg1)

	return Transform2D{base: &ret}

}

// XformVector2 godot_transform2d_xform_vector2 [[const godot_transform2d * p_self] [const godot_vector2 * p_v]] godot_vector2
func (gdt *Transform2D) XformVector2(v Vector2) Vector2 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform2d_xform_vector2(GDNative.api, arg0, arg1)

	return Vector2{base: &ret}

}

// XformInvVector2 godot_transform2d_xform_inv_vector2 [[const godot_transform2d * p_self] [const godot_vector2 * p_v]] godot_vector2
func (gdt *Transform2D) XformInvVector2(v Vector2) Vector2 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform2d_xform_inv_vector2(GDNative.api, arg0, arg1)

	return Vector2{base: &ret}

}

// BasisXformVector2 godot_transform2d_basis_xform_vector2 [[const godot_transform2d * p_self] [const godot_vector2 * p_v]] godot_vector2
func (gdt *Transform2D) BasisXformVector2(v Vector2) Vector2 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform2d_basis_xform_vector2(GDNative.api, arg0, arg1)

	return Vector2{base: &ret}

}

// BasisXformInvVector2 godot_transform2d_basis_xform_inv_vector2 [[const godot_transform2d * p_self] [const godot_vector2 * p_v]] godot_vector2
func (gdt *Transform2D) BasisXformInvVector2(v Vector2) Vector2 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform2d_basis_xform_inv_vector2(GDNative.api, arg0, arg1)

	return Vector2{base: &ret}

}

// InterpolateWith godot_transform2d_interpolate_with [[const godot_transform2d * p_self] [const godot_transform2d * p_m] [const godot_real p_c]] godot_transform2d
func (gdt *Transform2D) InterpolateWith(m Transform2D, c Real) Transform2D {
	arg0 := gdt.getBase()
	arg1 := m.getBase()
	arg2 := c.getBase()

	ret := C.go_godot_transform2d_interpolate_with(GDNative.api, arg0, arg1, arg2)

	return Transform2D{base: &ret}

}

// OperatorEqual godot_transform2d_operator_equal [[const godot_transform2d * p_self] [const godot_transform2d * p_b]] godot_bool
func (gdt *Transform2D) OperatorEqual(b Transform2D) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_transform2d_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorMultiply godot_transform2d_operator_multiply [[const godot_transform2d * p_self] [const godot_transform2d * p_b]] godot_transform2d
func (gdt *Transform2D) OperatorMultiply(b Transform2D) Transform2D {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_transform2d_operator_multiply(GDNative.api, arg0, arg1)

	return Transform2D{base: &ret}

}

// XformRect2 godot_transform2d_xform_rect2 [[const godot_transform2d * p_self] [const godot_rect2 * p_v]] godot_rect2
func (gdt *Transform2D) XformRect2(v Rect2) Rect2 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform2d_xform_rect2(GDNative.api, arg0, arg1)

	return Rect2{base: &ret}

}

// XformInvRect2 godot_transform2d_xform_inv_rect2 [[const godot_transform2d * p_self] [const godot_rect2 * p_v]] godot_rect2
func (gdt *Transform2D) XformInvRect2(v Rect2) Rect2 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform2d_xform_inv_rect2(GDNative.api, arg0, arg1)

	return Rect2{base: &ret}

}