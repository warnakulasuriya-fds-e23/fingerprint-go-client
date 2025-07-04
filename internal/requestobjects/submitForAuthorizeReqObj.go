package requestobjects

type SubmitForAuthorizeReqObj struct {
	ProbeCbor []byte `json:"probecbor"`
	ClientId  string `json:"clientid"`
}
