// +build cgo darwin

/*
macserial is a simple package to get an Apple Computer's Serial Number
It's not very well tested, so use it with care.
*/
package macserial

/*
#cgo LDFLAGS: -framework CoreFoundation -framework IOKit
#include <stdio.h>
#include <CoreFoundation/CoreFoundation.h>
#include <IOKit/IOKitLib.h>
char * getserial() {
    io_service_t platformExpert = IOServiceGetMatchingService(kIOMasterPortDefault,
            IOServiceMatching("IOPlatformExpertDevice"));

    if (platformExpert) {
        CFTypeRef serialNumberAsCFString =
            IORegistryEntryCreateCFProperty(platformExpert,
                    CFSTR(kIOPlatformSerialNumberKey),
                    kCFAllocatorDefault, 0);
        if (serialNumberAsCFString) {
            CFIndex bufsize = CFStringGetLength(serialNumberAsCFString) + 1;
            char *buf;
            buf = malloc(bufsize);
            if (buf != NULL) {
                Boolean result = CFStringGetCString(serialNumberAsCFString, buf, bufsize, kCFStringEncodingMacRoman);
                if (result) {
                    free((void*)serialNumberAsCFString);
                    IOObjectRelease(platformExpert);
                    return buf;
                }
            }
        }
        free((void *)serialNumberAsCFString);
        IOObjectRelease(platformExpert);
    }
    return NULL;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//Get returns an Apple Computer's Serial Number or an error if one occurred
func Get() (string, error) {
	serial := C.getserial()
	defer C.free(unsafe.Pointer(serial))
	if serial == nil {
		return "", fmt.Errorf("unable to get serial")
	}
	return C.GoString(serial), nil
}
