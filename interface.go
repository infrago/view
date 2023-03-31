package view

import (
	"fmt"
	"log"

	. "github.com/infrago/base"
	"github.com/infrago/infra"
)

func (this *Module) Register(name string, value Any) {
	switch obj := value.(type) {
	case Driver:
		this.Driver(name, obj)
	case Helper:
		this.Helper(name, obj)
	}
}
func (this *Module) Configure(global Map) {
	var config Map
	if vvv, ok := global["view"].(Map); ok {
		config = vvv
	}
	if config == nil {
		return
	}

	//设置值
	if driver, ok := config["driver"].(string); ok {
		this.config.Driver = driver
	}
	if vv, ok := config["root"].(string); ok {
		this.config.Root = vv
	}
	if vv, ok := config["shared"].(string); ok {
		this.config.Shared = vv
	}
	if vv, ok := config["left"].(string); ok {
		this.config.Left = vv
	}
	if vv, ok := config["right"].(string); ok {
		this.config.Right = vv
	}

	if setting, ok := config["setting"].(Map); ok {
		this.config.Setting = setting
	}
}
func (this *Module) Initialize() {
	if this.initialized {
		return
	}

	if this.config.Root == "" {
		this.config.Root = "asset/views"
	}
	if this.config.Shared == "" {
		this.config.Shared = "shared"
	}
	if this.config.Left == "" {
		this.config.Left = "{%"
	}
	if this.config.Right == "" {
		this.config.Right = "%}"
	}

	this.helperActions = Map{}
	for key, helper := range this.helpers {
		if helper.Action != nil {
			this.helperActions[key] = helper.Action
		}
	}

	this.initialized = true
}
func (this *Module) Connect() {
	if this.connected {
		return
	}

	driver, ok := this.drivers[this.config.Driver]
	if ok == false {
		panic("Invalid view driver: " + this.config.Driver)
	}

	inst := &Instance{
		nil, this.config, this.config.Setting,
	}

	// 建立连接
	connect, err := driver.Connect(inst)
	if err != nil {
		panic("Failed to connect to view: " + err.Error())
	}

	// 打开连接
	err = connect.Open()
	if err != nil {
		panic("Failed to open view connect: " + err.Error())
	}

	inst.connect = connect

	// 保存连接
	// this.connect = connect
	this.instance = inst

	this.connected = true
}
func (this *Module) Launch() {
	if this.launched {
		return
	}

	log.Println(fmt.Sprintf("%s VIEW is running with %d helpers.", infra.INFRAGO, len(this.helpers)))

	this.launched = true
}
func (this *Module) Terminate() {
	if this.instance != nil {
		this.instance.connect.Close()
	}

	this.launched = false
	this.connected = false
	this.initialized = false
}
