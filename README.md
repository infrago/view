# view

`view` 是 infrago 的模块包。

## 安装

```bash
go get github.com/infrago/view@latest
```

## 最小接入

```go
package main

import (
    _ "github.com/infrago/view"
    "github.com/infrago/infra"
)

func main() {
    infra.Run()
}
```

## 配置示例

```toml
[view]
driver = "default"
```

## 公开 API（摘自源码）

- `func (m *Module) Parse(body Body) (string, error)`
- `func (d *defaultDriver) Connect(inst *Instance) (Connection, error)`
- `func (c *defaultConnection) Open() error`
- `func (c *defaultConnection) Health() (Health, error)`
- `func (c *defaultConnection) Close() error`
- `func (c *defaultConnection) Parse(body Body) (string, error)`
- `func (p *defaultParser) Parse() (string, error)`
- `func Parse(body Body) (string, error)`
- `func SetFS(fsys fs.FS)`
- `func (m *Module) Register(name string, value Any)`
- `func (m *Module) RegisterDriver(name string, driver Driver)`
- `func (m *Module) RegisterHelper(name string, helper Helper)`
- `func (m *Module) RegisterConfig(config Config)`
- `func (m *Module) Config(global Map)`
- `func (m *Module) Setup()`
- `func (m *Module) Open()`
- `func (m *Module) Start()`
- `func (m *Module) Stop()`
- `func (m *Module) Close()`

## 排错

- 模块未运行：确认空导入已存在
- driver 无效：确认驱动包已引入
- 配置不生效：检查配置段名是否为 `[view]`
