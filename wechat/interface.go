package wechat

type MsgHandler interface {
	MsgHandle() (*ResponseMessage, error)
}

type Archiver interface {
	Archive() error
}

type TextReader interface {
	String() string
}
