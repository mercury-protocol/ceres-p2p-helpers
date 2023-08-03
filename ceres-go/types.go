package ceresgo

type PendingTransfer struct {
	Cid    string
	Sender string
}

type DataCollectorMsg struct {
	Cid       string `json:"cid"`
	Signature []byte `json:"signature"`
}

type PeerInfo struct {
	Address      string `json:"address"`
	TCP          string `json:"tcp"`
	Webtransport string `json:"webtransport"`
	Sig1         string `json:"sig1"`
	Sig2         string `json:"sig2"`
}
