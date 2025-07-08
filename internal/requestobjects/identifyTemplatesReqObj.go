package requestobjects

type IdentifyTemplateReqObj struct {
	ProbeCbor []byte `json:"probecbor"`
	DeviceId  string `json:"deviceid"`
}
