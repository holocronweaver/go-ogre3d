package gogre3d

/* 
 #cgo LDFLAGS: -lllcoi
 #include "ogre_interface.h"
*/
import "C"

type Entity struct {
	CPtr C.EntityHandle
}