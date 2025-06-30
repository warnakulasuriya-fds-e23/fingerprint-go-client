package requestobjects

type MatchTemplatesReqObj struct {
	ProbeCbor     []byte `json:"probecbor"`
	CandidateCbor []byte `json:"candidatecbor"`
	ClientId      string `json:"clientid"`
}
