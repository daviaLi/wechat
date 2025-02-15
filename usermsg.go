package wechat

import (
	"encoding/xml"
)

const (
	//MsgTypeText 表示文本消息
	MsgTypeText MsgType = "text"
	//MsgTypeImage 表示图片消息
	MsgTypeImage = "image"
	//MsgTypeVoice 表示语音消息
	MsgTypeVoice = "voice"
	//MsgTypeVideo 表示视频消息
	MsgTypeVideo = "video"
	//MsgTypeShortVideo 表示短视频消息[限接收]
	MsgTypeShortVideo = "shortvideo"
	//MsgTypeLocation 表示坐标消息[限接收]
	MsgTypeLocation = "location"
	//MsgTypeLink 表示链接消息[限接收]
	MsgTypeLink = "link"
	//MsgTypeMusic 表示音乐消息[限回复]
	MsgTypeMusic = "music"
	//MsgTypeNews 表示图文消息[限回复]
	MsgTypeNews = "news"
	//MsgTypeTransfer 表示消息消息转发到客服
	MsgTypeTransfer = "transfer_customer_service"
	//MsgTypeEvent 表示事件推送消息
	MsgTypeEvent = "event"
)

const (
	//EventSubscribe 订阅
	EventSubscribe EventType = "subscribe"
	//EventUnsubscribe 取消订阅
	EventUnsubscribe = "unsubscribe"
	//EventScan 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventScan = "SCAN"
	//EventLocation 上报地理位置事件
	EventLocation = "LOCATION"
	//EventClick 点击菜单拉取消息时的事件推送
	EventClick = "CLICK"
	//EventView 点击菜单跳转链接时的事件推送
	EventView = "VIEW"
	//EventScancodePush 扫码推事件的事件推送
	EventScancodePush = "scancode_push"
	//EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventScancodeWaitmsg = "scancode_waitmsg"
	//EventPicSysphoto 弹出系统拍照发图的事件推送
	EventPicSysphoto = "pic_sysphoto"
	//EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
	EventPicPhotoOrAlbum = "pic_photo_or_album"
	//EventPicWeixin 弹出微信相册发图器的事件推送
	EventPicWeixin = "pic_weixin"
	//EventLocationSelect 弹出地理位置选择器的事件推送
	EventLocationSelect = "location_select"
	//EventTemplateSendJobFinish 发送模板消息推送通知
	EventTemplateSendJobFinish = "TEMPLATESENDJOBFINISH"
	//EventWxaMediaCheck 异步校验图片/音频是否含有违法违规内容推送事件
	EventWxaMediaCheck = "wxa_media_check"
)

type (
	// WxMsg 混合用户消息，业务判断的主体
	WxMsg struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string
		FromUserName string
		CreateTime   int64
		MsgId        int64
		MsgType      string
		Content      string  // text
		AgentID      int     // corp
		PicUrl       string  // image
		MediaId      string  // image/voice/video/shortvideo
		Format       string  // voice
		Recognition  string  // voice
		ThumbMediaId string  // video
		LocationX    float32 `xml:"Latitude"`  // location
		LocationY    float32 `xml:"Longitude"` // location
		Precision    float32 // LOCATION
		Scale        int     // location
		Label        string  // location
		Title        string  // link
		Description  string  // link
		Url          string  // link
		Event        string  // event
		EventKey     string  // event
		SessionFrom  string  // event|user_enter_tempsession
		Ticket       string
		FileKey      string
		FileMd5      string
		FileTotalLen string
		TaskId       string

		ScanCodeInfo struct {
			ScanType   string
			ScanResult string
		}
	}

	// WxMsgEnc 加密的用户消息
	WxMsgEnc struct {
		XMLName    xml.Name `xml:"xml"`
		ToUserName string
		AgentID    int
		Encrypt    string
		AgentType  string
	}
)
