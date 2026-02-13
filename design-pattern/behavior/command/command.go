package command

import "fmt"

// 1. Command 命令接口：定义执行操作的方法
type Command interface {
	Execute() // 执行命令
	Undo()    // 撤销命令（可选，增强功能）
}

// 2. Receiver 接收者：实际执行命令的对象（灯）
type Light struct {
	IsOn bool // 灯的状态：true=开，false=关
}

// 开灯操作（接收者的具体行为）
func (l *Light) TurnOn() {
	l.IsOn = true
	fmt.Println("灯已打开 ✨")
}

// 关灯操作（接收者的具体行为）
func (l *Light) TurnOff() {
	l.IsOn = false
	fmt.Println("灯已关闭 ❌")
}

// 3. ConcreteCommand 具体命令：开灯命令
type LightOnCommand struct {
	light *Light // 绑定接收者（灯）
}

// 实现 Command 接口的 Execute 方法
func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

// 实现 Command 接口的 Undo 方法（撤销=关灯）
func (c *LightOnCommand) Undo() {
	c.light.TurnOff()
}

// 4. ConcreteCommand 具体命令：关灯命令
type LightOffCommand struct {
	light *Light // 绑定接收者（灯）
}

// 实现 Command 接口的 Execute 方法
func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

// 实现 Command 接口的 Undo 方法（撤销=开灯）
func (c *LightOffCommand) Undo() {
	c.light.TurnOn()
}

// 5. Invoker 调用者：遥控器（负责调用命令，不关心具体操作）
type RemoteControl struct {
	currentCommand Command   // 当前执行的命令
	history        []Command // 命令历史（用于撤销/重做）
}

// 设置要执行的命令
func (r *RemoteControl) SetCommand(cmd Command) {
	r.currentCommand = cmd
}

// 执行命令（并记录历史）
func (r *RemoteControl) PressButton() {
	if r.currentCommand != nil {
		r.currentCommand.Execute()
		r.history = append(r.history, r.currentCommand)
	} else {
		fmt.Println("未设置任何命令 🚫")
	}
}

// 撤销上一个命令
func (r *RemoteControl) PressUndo() {
	if len(r.history) == 0 {
		fmt.Println("无命令可撤销 📜")
		return
	}
	// 取出最后一个命令并撤销
	lastCmd := r.history[len(r.history)-1]
	lastCmd.Undo()
	// 移除已撤销的命令
	r.history = r.history[:len(r.history)-1]
}
