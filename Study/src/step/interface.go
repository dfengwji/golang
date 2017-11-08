// interface
package step

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type phoneConnecter struct {
	name string
}

type tvConnecter struct {
	name string
}

func (this phoneConnecter) Name() string {
	return this.name
}

func (this phoneConnecter) Connect() {
	fmt.Println("Connected:", this.name)
}

func Disconnect1(usb USB) {
	//类型断言
	if pc, ok := usb.(phoneConnecter); ok {
		fmt.Println("disconnected:", pc.name)
		return
	}
	fmt.Println("unknow decive!!!")
}

func Disconnect(usb interface{}) {
	//类型断言
	//	if pc, ok := usb.(phoneConnecter); ok {
	//		fmt.Println("disconnected:", pc.name)
	//		return
	//	}
	//	fmt.Println("unknow decive!!!")
	switch v := usb.(type) {
	case phoneConnecter:
		fmt.Println("disconnected:", v.name)
	default:
		fmt.Println("unknow decive!!!")
	}
}

func StudyInterface() {
	pc := phoneConnecter{"phoneConn"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	Disconnect(a)

	var b interface{}
	fmt.Println(b == nil)
	var p *int = nil
	b = p
	fmt.Println(b == nil)
}
