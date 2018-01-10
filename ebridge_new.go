package ebridge_new

/*
#cgo CFLAGS: -I/opt/egcs/epics/extensions/include -I/opt/egcs/epics/base/include -I/opt/egcs/epics/base/include/os/Linux
#cgo LDFLAGS: -L/opt/egcs/epics/extensions/lib/linux-x86_64 -lezca
#include <stdio.h>
#include <tsDefs.h>
#include <cadef.h>
#include "ezca.h"
*/
import "C"
import (
	"errors"
	"unsafe"
	"fmt"

		
)

func ezcaInit() {
	C.ezcaSetTimeout(0.2)
	C.ezcaSetRetryCount(18)
}

func LongGet(pv string) (int, error) {
	ezcaInit()
	result := new(int)
	
	ezcaReturn := C.ezcaGet(C.CString(pv), C.ezcaLong, 1, unsafe.Pointer(result))

	fmt.Println(ezcaReturn)

	//

	if ezcaReturn != C.EZCA_OK {
		
	
			
		return -1, errors.New("long PV获取失败")

		
	}
	
	return *result, nil
}

func StringGet(pv string) (string, error) {
	ezcaInit()
	rawResult := make([]byte, 100)
	src := C.CBytes(rawResult)

	ezcaReturn := C.ezcaGet(C.CString(pv), C.ezcaString, 1, src)

	if ezcaReturn != C.EZCA_OK {
		return "", errors.New("string PV获取失败")
	}
	return string(C.GoBytes(src, 100)), nil
}

func DoubleGet(pv string) (float64, error) {
	ezcaInit()
	result := new(float64)
	ezcaReturn := C.ezcaGet(C.CString(pv), C.ezcaDouble, 1, unsafe.Pointer(result))
	if ezcaReturn != C.EZCA_OK {
		return -1, errors.New("double PV获取失败")
	}
	return *result, nil
}

func BoolGet(pv string) (int16, error) {
	ezcaInit()
	result := new(int16)
	ezcaReturn := C.ezcaGet(C.CString(pv), C.ezcaByte, 1, unsafe.Pointer(result))
	if ezcaReturn != C.EZCA_OK {
		return -1, errors.New("bool PV获取失败")
	}
	return *result, nil
}
