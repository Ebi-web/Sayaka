package messages

const TextType = "text"

type TextReply struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewTextReply(text string) *TextReply {
	return &TextReply{
		Type: TextType,
		Text: text,
	}
}

func (t *TextReply) Struct() TextReply {
	return *t
}
