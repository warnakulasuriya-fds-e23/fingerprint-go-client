package requestobjects

type MatchTemplatesReqObj struct {
	Probe     []byte `json:"probe"`
	Candidate []byte `json:"candidate"`
}
