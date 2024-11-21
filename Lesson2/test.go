package main

import "fmt"

type DigitalProductInfo struct {
	Edition  string
	UserName string
	State    string
}

type Mp3Info struct {
	Digital   DigitalProductInfo
	AudioChip string
}

type BasicPhoneInfo struct {
	CameraName string
	Soc        string
	Battery    int
}
type PhoneInfo struct {
	Digital DigitalProductInfo
	Basic   BasicPhoneInfo
}

// 这是一个函数
func WhosePhone(p PhoneInfo) {
	fmt.Println("函数WhosePhone", p.Digital.UserName)
}

// 这是一个函数
func openState(p *PhoneInfo) {
	p.Digital.State = "open"
}

// 这是一个方法
func (p *PhoneInfo) Open() {
	p.Digital.State = "open"
}

func (p *PhoneInfo) close() {
	p.Digital.State = "close"
}

// 几种简单的声明结构体的方法
func main() {
	var a PhoneInfo
	a.Basic.Battery = 4000
	a.Digital = DigitalProductInfo{
		UserName: "a的UserName",
	}
	//WhosePhone(a)
	openState(&a)
	fmt.Println(a.Digital.State)
	b := PhoneInfo{
		Digital: DigitalProductInfo{
			Edition:  "v100",
			UserName: "xiaoA",
			State:    "new",
		},
		Basic: BasicPhoneInfo{
			CameraName: "pix10000000",
			Soc:        "kirin999",
			Battery:    4000,
		},
	}
	fmt.Println(a, b)
	//fmt.Println("BasicPhoneInfo", b.Basic)
	//fmt.Println("DigitalProductInfo", b.Digital)
	//fmt.Println("Battery", b.Basic.Battery)

}
