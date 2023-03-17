package view

import (
	"sync"
	"time"

	. "github.com/infrago/base"
	"github.com/infrago/infra"
)

func init() {
	infra.Register(module)
}

var (
	module = &Module{
		config: Config{
			Driver: infra.DEFAULT, Root: "asset/views",
			Shared: "shared", Left: "{%", Right: "%}",
		},
		drivers: make(map[string]Driver, 0),
		helpers: make(map[string]Helper, 0),
	}
)

type (
	// 日志模块定义
	Module struct {
		connected, initialized, launched bool

		mutex   sync.Mutex
		drivers map[string]Driver
		helpers map[string]Helper

		helperActions Map

		config  Config
		connect Connect
	}

	// LogConfig 日志模块配置
	Config struct {
		Driver  string
		Root    string
		Shared  string
		Left    string
		Right   string
		Setting Map
	}

	Body struct {
		View     string
		Site     string
		Language string
		Timezone *time.Location
		Data     Map
		Model    Map
		Helpers  Map
	}
)

// Driver 为view模块注册驱动
func (this *Module) Driver(name string, driver Driver) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if driver == nil {
		panic("Invalid view driver: " + name)
	}

	if infra.Override() {
		this.drivers[name] = driver
	} else {
		if this.drivers[name] == nil {
			this.drivers[name] = driver
		}
	}
}

func (this *Module) Config(config Config) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.config = config
}

func (this *Module) Helper(name string, config Helper) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	alias := make([]string, 0)
	if name != "" {
		alias = append(alias, name)
	}
	if config.Alias != nil {
		alias = append(alias, config.Alias...)
	}

	for _, key := range alias {
		if infra.Override() {
			this.helpers[key] = config
		} else {
			if _, ok := this.helpers[key]; ok == false {
				this.helpers[key] = config
			}
		}

	}
}
