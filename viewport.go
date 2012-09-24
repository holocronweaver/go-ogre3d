package gogre3d

/* 
 #cgo LDFLAGS: -lllcoi
 #include "ogre_interface.h"
*/
import "C"

type Viewport struct {
	CPtr C.ViewportHandle
}

func (v *Viewport) SetBackgroundColour(r, g, b float32) {
	C.viewport_set_background_colour(v.CPtr, C.float(r), C.float(g), C.float(b))
}

func (v *Viewport) GetWidth() float32 {
	result := C.viewport_get_width(v.CPtr)
	return float32(result)
}

func (v *Viewport) GetHeight() float32 {
	result := C.viewport_get_height(v.CPtr)
	return float32(result)
}