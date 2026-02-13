package command

import (
	"fmt"
	"testing"
)

// 客户端代码
func TestLight_TurnOff(t *testing.T) {
	// 1. 创建接收者（灯）
	bedroomLight := &Light{IsOn: false}

	// 2. 创建具体命令（绑定接收者）
	lightOnCmd := &LightOnCommand{light: bedroomLight}
	lightOffCmd := &LightOffCommand{light: bedroomLight}

	// 3. 创建调用者（遥控器）
	remote := &RemoteControl{}

	// 4. 执行命令：开灯
	fmt.Println("=== 第一次操作 ===")
	remote.SetCommand(lightOnCmd)
	remote.PressButton() // 输出：灯已打开 ✨

	// 5. 执行命令：关灯
	fmt.Println("\n=== 第二次操作 ===")
	remote.SetCommand(lightOffCmd)
	remote.PressButton() // 输出：灯已关闭 ❌

	// 6. 撤销操作（撤销关灯 → 开灯）
	fmt.Println("\n=== 撤销操作 ===")
	remote.PressUndo() // 输出：灯已打开 ✨

	// 7. 再次撤销（撤销开灯 → 关灯）
	fmt.Println("\n=== 再次撤销 ===")
	remote.PressUndo() // 输出：灯已关闭 ❌

	// 8. 无命令时撤销
	fmt.Println("\n=== 无命令撤销 ===")
	remote.PressUndo() // 输出：无命令可撤销 📜
}
