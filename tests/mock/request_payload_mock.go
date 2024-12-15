package mock

type RequestObject struct {
	RequestPayload
}

func NewRequestObject() *RequestObject {
	RequestPayload := NewRequestPayload()

	return &RequestObject{
		RequestPayload: *RequestPayload,
	}
}

func (ro *RequestObject) ValidPayload() Payload {
	return ro.RequestPayload.Payload
}
