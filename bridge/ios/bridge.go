//go:build ios
// +build ios

package ios

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>
void Log(const char *text) {
  NSString *nss = [NSString stringWithUTF8String:text];
  NSLog(@"%@", nss);
}
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"

	librsyncdartwrapper "github.com/filemingo/librsync-dart-wrapper"
	"github.com/filemingo/librsync-dart-wrapper/bridge"
	"github.com/filemingo/librsync-dart-wrapper/logging"
)

type CallbackWriter = librsyncdartwrapper.CallbackWriter

type Stream interface {
  bridge.Stream
}

func Test(signatureStr string, targetFilePath string, callback func([]byte)) {
}

func ComputeDelta(signatureStr string, targetFilePath string, stream Stream) {
	librsyncdartwrapper.ComputeDelta(signatureStr, targetFilePath, func(b []byte, i int) {
		stream.Send(b, i)
	})
}

type iosLogger struct {
}

func (l iosLogger) D(TAG, message string) {
	l.Write([]byte(fmt.Sprintf("D/ [%v]: %v", TAG, message)))
}
func (l iosLogger) E(TAG, message string) {
	l.Write([]byte(fmt.Sprintf("E/ [%v]: %v", TAG, message)))
}
func (l iosLogger) V(TAG, message string) {
	l.Write([]byte(fmt.Sprintf("V/ [%v]: %v", TAG, message)))
}
func (l iosLogger) I(TAG, message string) {
	l.Write([]byte(fmt.Sprintf("I/ [%v]: %v", TAG, message)))
}

func (l iosLogger) W(TAG, message string) {
	l.Write([]byte(fmt.Sprintf("W/ [%v]: %v", TAG, message)))
}

func (l iosLogger) Write(p []byte) (n int, err error) {
	p = append(p, 0)
	cstr := (*C.char)(unsafe.Pointer(&p[0]))
	C.Log(cstr)
	return len(p), nil
}

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		bridge.Init()
		logging.AddLogger(iosLogger{})
	})
}
