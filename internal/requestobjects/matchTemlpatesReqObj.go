package requestobjects

type MatchTemplatesReqObj struct {
	ProbeCbor     []byte `json:"probecbor"`
	CandidateCbor []byte `json:"candidatecbor"`
	DeviceId      string `json:"deviceid"`
}
