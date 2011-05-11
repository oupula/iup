/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of iup.go.

	iup.go is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	iup.go is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with iup.go.  If not, see <http://www.gnu.org/licenses/>.
*/

package iup

/*
#include <stdlib.h>
#include <iup.h>

void _IupSetfAttributeId2(Ihandle *ih, const char *name, int lin, int col, const char *value) {
	IupSetfAttributeId2(ih, name, lin, col, value);
}
*/
import "C"
import "unsafe"
import "fmt"

func (ih *Ihandle) StoreAttribute(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupStoreAttribute(ih.h, cName, cValue)
}

func (ih *Ihandle) StoreAttributeId(name string, id int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupStoreAttributeId(ih.h, cName, C.int(id), cValue)
}

func (ih *Ihandle) SetAttribute(name string, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupSetAttribute(ih.h, cName, (*C.char)(value))
}

func (ih *Ihandle) SetAttributeId(name string, id int, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupSetAttributeId(ih.h, cName, C.int(id), (*C.char)(value))
}

func (ih *Ihandle) SetfAttribute(name, format string, args ...interface{}) {
	ih.StoreAttribute(name, fmt.Sprintf(format, args...))
}

func (ih *Ihandle) SetfAttributeId(name string, id int, format string, args ...interface{}) {
	ih.StoreAttributeId(name, id, fmt.Sprintf(format, args...))
}

func (ih *Ihandle) SetfAttributeId2(name string, lin int, col int, format string, args ...interface{}) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cValue))

	C._IupSetfAttributeId2(ih.h, cName, C.int(lin), C.int(col), cValue)
}

func (ih *Ihandle) SetAttributes(values string) {
	cValues := C.CString(values)
	defer C.free(unsafe.Pointer(cValues))

	C.IupSetAttributes(ih.h, cValues)
}

func (ih *Ihandle) ResetAttribute(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupResetAttribute(ih.h, cName)
}

func (ih *Ihandle) SetAttributeHandle(name string, ihNamed *Ihandle) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupSetAttributeHandle(ih.h, cName, ihNamed.h)
}

func (ih *Ihandle) GetAttributeHandle(name string) *Ihandle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return &Ihandle{h: C.IupGetAttributeHandle(ih.h, cName)}
}

func (ih *Ihandle) GetAttribute(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.GoString(C.IupGetAttribute(ih.h, cName))
}

func (ih *Ihandle) GetAttributeId(name string, id int) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.GoString(C.IupGetAttributeId(ih.h, cName, C.int(id)))
}

func (ih *Ihandle) GetIntId(name string, id int) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int(C.IupGetIntId(ih.h, cName, C.int(id)))
}

func (ih *Ihandle) GetFloatId(name string, id int) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.IupGetFloatId(ih.h, cName, C.int(id)))
}

func (ih *Ihandle) GetAttributes() string {
	return C.GoString(C.IupGetAttributes(ih.h))
}

func (ih *Ihandle) GetFloat(name string) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.IupGetFloat(ih.h, cName))
}

func (ih *Ihandle) GetInt(name string) int64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int64(C.IupGetInt(ih.h, cName))
}

func StoreGlobal(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupStoreGlobal(cName, cValue)
}

func SetGlobal(name string, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupSetGlobal(cName, (*C.char)(value))
}

func GetGlobal(name string) string {	
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return C.GoString(C.IupGetGlobal(cName))
}
