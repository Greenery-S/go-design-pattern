package command

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBoxes(t *testing.T) {
	mb := &MotherBoard{}
	start := NewStartCommand(mb)
	reboot := NewRebootCommand(mb)

	// 读取stdout, 检查输出
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	oldStdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
		err = w.Close()
		assert.NoError(t, err)
		err = r.Close()
		assert.NoError(t, err)
	}()

	// 两个按钮, 一个是start, 一个是reboot
	box1 := NewBox(start, reboot)
	box1.PressButton1() // start
	box1.PressButton2() // reboot
	t.Log("\n这句话不会出现在output里,而是出现在consul里.\n" +
		"	- 在测试通过的情况下：默认情况下，如果测试通过了（没有调用t.Fail、t.Error等函数），t.Log输出的信息不会显示在标准输出中。这是为了保持测试输出的简洁，只显示通过或失败的测试结果。\n" +
		"	- 在测试失败的情况下：如果测试失败（例如调用了t.Fail、t.Error等函数），那么t.Log输出的信息会被打印到标准输出中，帮助你调试和定位问题。\n" +
		"	- 使用`-v`或`-test.v`标志：如果你运行测试时使用了-v（verbose）标志，例如go test -v，那么所有t.Log输出的信息都会打印到标准输出，无论测试是否通过。这在调试时非常有用，可以查看所有日志信息。")
	// 两个按钮, 一个是reboot, 一个是start
	box2 := NewBox(reboot, start)
	box2.PressButton1() // reboot
	box2.PressButton2() // start

	// 读取管道中的print数据
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	assert.NoError(t, err)
	output := string(buf[:n])

	expected :=
		"system starting\n" +
			"system rebooting\n" +
			"system rebooting\n" +
			"system starting\n"

	assert.Equal(t, expected, output)

}
