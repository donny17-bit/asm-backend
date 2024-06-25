package claim

type InputData struct {
	Page        string `json:"page"`
	Page_size   string `json:"page_size"`
	Begin_date  string `json:"begin_date"`
	End_date    string `json:"end_date"`
	No_polis    string `json:"no_polis"`
	No_cif      string `json:"no_cif"`
	Client_name string `json:"client_name"`
	Branch      string `json:"branch"`
	Business    string `json:"business"`
	Sumbis      string `json:"sumbis"`
}

type Data struct {
	Rn              string   `json:"Rn"`
	AcceptedDateOri string   `json:"AcceptedDateOri"`
	Kanwil          string   `json:"Kanwil"`
	Cabang          string   `json:"Cabang"`
	Perwakilan      string   `json:"Perwakilan"`
	SubPerwakilan   string   `json:"SubPerwakilan"`
	Namaleader0     string   `json:"Namaleader0"`
	Namaleader1     string   `json:"Namaleader1"`
	Namaleader2     string   `json:"Namaleader2"`
	Namaleader3     string   `json:"Namaleader3"`
	GroupBusiness   string   `json:"GroupBusiness"`
	Business        string   `json:"Business"`
	TahunPolis      string   `json:"TahunPolis"`
	AcceptedNo      string   `json:"AcceptedNo"`
	NoKlaim         *string  `json:"NoKlaim"`
	NoPolis         string   `json:"NoPolis"`
	NoCif           string   `json:"NoCif"`
	ClientName      string   `json:"ClientName"`
	Mo              string   `json:"Mo"`
	PrepareDate     *string  `json:"PrepareDate"`
	DateOfLoss      *string  `json:"DateOfLoss"`
	AcceptedDate    *string  `json:"AcceptedDate"`
	BeginDate       *string  `json:"BeginDate"`
	EndDate         *string  `json:"EndDate"`
	JenisPaket      *string  `json:"JenisPaket"`
	Workshop        *string  `json:"Workshop"`
	NamaDealer      *string  `json:"NamaDealer"`
	ColDesk         *string  `json:"ColDesk"`
	RiskLoc         *string  `json:"RiskLoc"`
	Tsi             float32  `json:"Tsi"`
	AcceptedClaim   float32  `json:"AcceptedClaim"`
	AcceptedClaimRp float32  `json:"AcceptedClaimRp"`
	AccKlaimGrossRp *float32 `json:"AccKlaimGrossRp"`
	OwnRetention    float32  `json:"OwnRetention"`
	CoIns           float32  `json:"CoIns"`
	Psrspl          float32  `json:"Psrspl"`
	Qsri            float32  `json:"Qsri"`
	Er1             float32  `json:"Er1"`
	Surplus1        float32  `json:"Surplus1"`
	Surplus2        float32  `json:"Surplus2"`
	Er2             float32  `json:"Er2"`
	PsrqsRi         float32  `json:"PsrqsRi"`
	PsrqsOr         float32  `json:"PsrqsOr"`
	Ors             float32  `json:"Ors"`
	Facultative     float32  `json:"Facultative"`
	Facobl          float32  `json:"Facobl"`
	Bppdan          float32  `json:"Bppdan"`
	Xl              float32  `json:"Xl"`
	Pss             float32  `json:"Pss"`
	Prgbi           float32  `json:"Prgbi"`
	Pfra            float32  `json:"Pfra"`
	Fsplnsri        *float32 `json:"Fsplnsri"`
	Psplnsri        *float32 `json:"Psplnsri"`
	Fsplnsor        *float32 `json:"Fsplnsor"`
	Psplnsor        *float32 `json:"Psplnsor"`
	Facobsrb        *float32 `json:"Facobsrb"`
	Facobindt       *float32 `json:"Facobindt"`
}