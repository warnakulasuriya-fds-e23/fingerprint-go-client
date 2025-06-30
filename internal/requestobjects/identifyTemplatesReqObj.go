package requestobjects

type IdentifyTemplateReqObj struct {
	ProbeCbor []byte `json:"probecbor"`
	ClientId  string `json:"clientid"`
}
