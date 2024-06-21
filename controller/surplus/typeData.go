package surplus

type InputData struct {
	Page        string `json:"page"`
	Page_size   string `json:"page_size"`
	Periode     string `json:"periode"`
	No_polis    string `json:"no_polis"`
	No_cif      string `json:"no_cif"`
	Client_name string `json:"Client_name"`
	Branch      string `json:"Branch"`
	Business    string `json:"business"`
	Sumbis      string `json:"sumbis"`
}

type Data struct {
	Rn               string  `json:"Rn"`
	Mthname          string  `json:"Mthname"`
	Periode          string  `json:"Periode"`
	Kanwil           string  `json:"Kanwil"`
	Cabang           string  `json:"Cabang"`
	Perwakilan       string  `json:"Perwakilan"`
	SubPerwakilan    string  `json:"SubPerwakilan"`
	Namaleader0      string  `json:"Namaleader0"`
	Namaleader1      string  `json:"Namaleader1"`
	Namaleader2      string  `json:"Namaleader2"`
	Namaleader3      string  `json:"Namaleader3"`
	Mo               string  `json:"Mo"`
	GroupBusiness    string  `json:"GroupBusiness"`
	Business         string  `json:"Business"`
	ClientName       *string `json:"ClientName"`
	NoPolis          string  `json:"NoPolis"`
	NoCif            string  `json:"NoCif"`
	JenisPaket       *string `json:"JenisPaket"`
	Keterangan       *string `json:"Keterangan"`
	NamaCeding       *string `json:"NamaCeding"`
	NamaDealer       *string `json:"NamaDealer"`
	Tsi              float32 `json:"Tsi"`
	Gpw              float32 `json:"Gpw"`
	Disc             float32 `json:"Disc"`
	Disc2            float32 `json:"Disc2"`
	Comm             float32 `json:"Comm"`
	Oc               float32 `json:"Oc"`
	Bkp              float32 `json:"Bkp"`
	Ngpw             float32 `json:"Ngpw"`
	Ri               float32 `json:"Ri"`
	Ricom            float32 `json:"Ricom"`
	Npw              float32 `json:"Npw"`
	CadPremi         float32 `json:"CadPremi"`
	CadPremi1        float32 `json:"CadPremi1"`
	PremiumReserve   float32 `json:"PremiumReserve"`
	Npe              float32 `json:"Npe"`
	AcceptedClaim    float32 `json:"AcceptedClaim"`
	RejectedClaim    float32 `json:"RejectedClaim"`
	OutstandingClaim float32 `json:"OutstandingClaim"`
	ReversedClaim    float32 `json:"ReversedClaim"`
	SurplusUw        float32 `json:"SurplusUw"`
}