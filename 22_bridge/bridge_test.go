package bridge

import "testing"

func TestMessageSMS_Send(t *testing.T) {
	// 2种抽象类 和 2种实现类 可以随意组合
	commonMsgViaSMS := NewCommonMessage(ViaSMS())
	commonMsgViaEmail := NewCommonMessage(ViaEmail())
	urgencyMsgViaSMS := NewUrgencyMessage(ViaSMS())
	urgencyMsgViaEmail := NewUrgencyMessage(ViaEmail())

	commonMsgViaSMS.SendMessage("hi,今天五点有个会议你参加一下", "正式员工小马")
	commonMsgViaEmail.SendMessage("您好,今天五点将召开会议,希望您能参加", "主管老王")
	urgencyMsgViaSMS.SendMessage("在吗在吗,今天五点有个会议你必须参加!! 不然扣钱", "实习生小张")
	urgencyMsgViaEmail.SendMessage("您好,今天五点将召开会议,您务必参加,否则后果自负", "外包狗小李")
}
