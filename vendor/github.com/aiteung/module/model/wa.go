package model

type MediaType int

const (
	_              = iota
	MediaTypeImage = MediaType(iota)
	MediaTypeVideo
	MediaTypeAudio
	MediaTypeDocument
)

var (
	AppInfo = map[MediaType]string{
		MediaTypeImage:    "WhatsApp Image Keys",
		MediaTypeVideo:    "WhatsApp Video Keys",
		MediaTypeAudio:    "WhatsApp Audio Keys",
		MediaTypeDocument: "WhatsApp Document Keys",
	}
)
