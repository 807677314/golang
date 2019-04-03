package wechat

import "github.com/esap/wechat"

func (c *WechatControllers) WxHandler() {
	wv := wechat.VerifyURL(c.Ctx.ResponseWriter, c.Ctx.Request)

	content := wv.Msg.Content
	//toUserName := wv.Msg.ToUserName
	//fromUserName := wv.Msg.FromUserName
	msgType := wv.Msg.MsgType
	//msgId := wv.Msg.MsgId

	//log.Println(toUserName, fromUserName, msgType, msgId)
	switch msgType {
	case "event":
		wv.NewText("您好，请问有什么可以帮助你的吗？").Reply()
	case "text":
		wv.NewText(content).Send()
	}
	c.StopRun()

}
