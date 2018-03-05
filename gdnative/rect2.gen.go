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
#include <gdnative/rect2.h>
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

// NewEmptyRect2 will return a pointer to an empty
// initialized Rect2. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyRect2() Pointer {
	var obj C.godot_rect2
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromRect2 will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromRect2(obj Rect2) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewRect2FromPointer will return a Rect2 from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewRect2FromPointer(ptr Pointer) Rect2 {

	return Rect2{base: (*C.godot_rect2)(ptr.getBase())}
}

type Rect2 struct {
	base *C.godot_rect2
}

func (gdt Rect2) getBase() *C.godot_rect2 {
	return gdt.base
}

// NewRect2WithPositionAndSize godot_rect2_new_with_position_and_size [[godot_rect2 * r_dest] [const godot_vector2 * p_pos] [const godot_vector2 * p_size]] void
func NewRect2WithPositionAndSize(pos Vector2, size Vector2) Rect2 {
	var dest C.godot_rect2
	arg1 := pos.getBase()
	arg2 := size.getBase()
	C.go_godot_rect2_new_with_position_and_size(GDNative.api, &dest, arg1, arg2)
	return Rect2{base: &dest}
}

// NewRect2 godot_rect2_new [[godot_rect2 * r_dest] [const godot_real p_x] [const godot_real p_y] [const godot_real p_width] [const godot_real p_height]] void
func NewRect2(x Real, y Real, width Real, height Real) Rect2 {
	var dest C.godot_rect2
	arg1 := x.getBase()
	arg2 := y.getBase()
	arg3 := width.getBase()
	arg4 := height.getBase()
	C.go_godot_rect2_new(GDNative.api, &dest, arg1, arg2, arg3, arg4)
	return Rect2{base: &dest}
}

// AsString godot_rect2_as_string [[const godot_rect2 * p_self]] godot_string
func (gdt *Rect2) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_rect2_as_string(GDNative.api, arg0)

	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharT(wchar)
	return String(goWchar.AsString())

}

// GetArea godot_rect2_get_area [[const godot_rect2 * p_self]] godot_real
func (gdt *Rect2) GetArea() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_rect2_get_area(GDNative.api, arg0)

	return Real(ret)
}

// Intersects godot_rect2_intersects [[const godot_rect2 * p_self] [const godot_rect2 * p_b]] godot_bool
func (gdt *Rect2) Intersects(b Rect2) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_rect2_intersects(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// Encloses godot_rect2_encloses [[const godot_rect2 * p_self] [const godot_rect2 * p_b]] godot_bool
func (gdt *Rect2) Encloses(b Rect2) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_rect2_encloses(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// HasNoArea godot_rect2_has_no_area [[const godot_rect2 * p_self]] godot_bool
func (gdt *Rect2) HasNoArea() Bool {
	arg0 := gdt.getBase()

	ret := C.go_godot_rect2_has_no_area(GDNative.api, arg0)

	return Bool(ret)
}

// Clip godot_rect2_clip [[const godot_rect2 * p_self] [const godot_rect2 * p_b]] godot_rect2
func (gdt *Rect2) Clip(b Rect2) Rect2 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_rect2_clip(GDNative.api, arg0, arg1)

	return Rect2{base: &ret}

}

// Merge godot_rect2_merge [[const godot_rect2 * p_self] [const godot_rect2 * p_b]] godot_rect2
func (gdt *Rect2) Merge(b Rect2) Rect2 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_rect2_merge(GDNative.api, arg0, arg1)

	return Rect2{base: &ret}

}

// HasPoint godot_rect2_has_point [[const godot_rect2 * p_self] [const godot_vector2 * p_point]] godot_bool
func (gdt *Rect2) HasPoint(point Vector2) Bool {
	arg0 := gdt.getBase()
	arg1 := point.getBase()

	ret := C.go_godot_rect2_has_point(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// Grow godot_rect2_grow [[const godot_rect2 * p_self] [const godot_real p_by]] godot_rect2
func (gdt *Rect2) Grow(by Real) Rect2 {
	arg0 := gdt.getBase()
	arg1 := by.getBase()

	ret := C.go_godot_rect2_grow(GDNative.api, arg0, arg1)

	return Rect2{base: &ret}

}

// Expand godot_rect2_expand [[const godot_rect2 * p_self] [const godot_vector2 * p_to]] godot_rect2
func (gdt *Rect2) Expand(to Vector2) Rect2 {
	arg0 := gdt.getBase()
	arg1 := to.getBase()

	ret := C.go_godot_rect2_expand(GDNative.api, arg0, arg1)

	return Rect2{base: &ret}

}

// OperatorEqual godot_rect2_operator_equal [[const godot_rect2 * p_self] [const godot_rect2 * p_b]] godot_bool
func (gdt *Rect2) OperatorEqual(b Rect2) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_rect2_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// GetPosition godot_rect2_get_position [[const godot_rect2 * p_self]] godot_vector2
func (gdt *Rect2) GetPosition() Vector2 {
	arg0 := gdt.getBase()

	ret := C.go_godot_rect2_get_position(GDNative.api, arg0)

	return Vector2{base: &ret}

}

// GetSize godot_rect2_get_size [[const godot_rect2 * p_self]] godot_vector2
func (gdt *Rect2) GetSize() Vector2 {
	arg0 := gdt.getBase()

	ret := C.go_godot_rect2_get_size(GDNative.api, arg0)

	return Vector2{base: &ret}

}

// SetPosition godot_rect2_set_position [[godot_rect2 * p_self] [const godot_vector2 * p_pos]] void
func (gdt *Rect2) SetPosition(pos Vector2) {
	arg0 := gdt.getBase()
	arg1 := pos.getBase()

	C.go_godot_rect2_set_position(GDNative.api, arg0, arg1)
}

// SetSize godot_rect2_set_size [[godot_rect2 * p_self] [const godot_vector2 * p_size]] void
func (gdt *Rect2) SetSize(size Vector2) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_rect2_set_size(GDNative.api, arg0, arg1)
}