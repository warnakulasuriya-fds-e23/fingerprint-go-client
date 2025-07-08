package requestobjects

type SubmitForAuthorizeReqObj struct {
	ProbeCbor []byte `json:"probecbor"`
	DeviceId  string `json:"deviceid"`
}
