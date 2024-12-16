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
	return ro.RequestPayload.Build()
}

func (ro *RequestObject) InvalidCardData() Payload{
	return ro.RequestPayload.InvalidCardData().Build()
}

func (ro *RequestObject) InvalidVendor() Payload{
	return ro.RequestPayload.InvalidVendor().Build()
}
