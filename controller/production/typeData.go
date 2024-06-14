package production

type InputData struct {
	Page        string `json:"page"`
	Page_size   string `json:"page_size"`
	Begin_date  string `json:"begin_date"`
	End_date    string `json:"end_date"`
	No_polis    string `json:"no_polis"`
	No_cif      string `json:"no_cif"`
	Client_name string `json:"Client_name"`
	Branch      string `json:"Branch"`
	Business    string `json:"business"`
	Sumbis      string `json:"sumbis"`
}

type Data struct {
	Rn            string  `json:"Rn"`
	TglProd       string  `json:"TglProd"`
	ThnBln        string  `json:"ThnBln"`
	ProdDate      string  `json:"ProdDate"`
	BeginDate     string  `json:"BeginDate"`
	EndDate       string  `json:"EndDate"`
	Mo            string  `json:"Mo"`
	ClientName    string  `json:"ClientName"`
	Kanwil        string  `json:"Kanwil"`
	Cabang        string  `json:"Cabang"`
	Perwakilan    string  `json:"Perwakilan"`
	SubPerwakilan string  `json:"SubPerwakilan"`
	Jnner         string  `json:"Jnner"`
	JenisProd     string  `json:"JenisProd"`
	JenisPaket    *string `json:"JenisPaket"`
	// JenisPaket    NullableString `json:"JenisPaket"`
	Ket *string `json:"Ket"`
	// NamaCeding    NullableString `json:"NamaCeding"`
	NamaCeding    *string `json:"NamaCeding"`
	Namaleader0   string  `json:"Namaleader0"`
	Namaleader1   string  `json:"Namaleader1"`
	Namaleader2   string  `json:"Namaleader2"`
	Namaleader3   string  `json:"Namaleader3"`
	GroupBusiness string  `json:"GroupBusiness"`
	Business      string  `json:"Business"`
	NoKontrak     *string `json:"NoKontrak"`
	NoPolis       string  `json:"NoPolis"`
	NoCif         string  `json:"NoCif"`
	ProdKe        string  `json:"ProdKe"`
	NamaDealer    *string `json:"NamaDealer"`
	Tsi           float32 `json:"Tsi"`
	Gpw           float32 `json:"Gpw"`
	Disc          float32 `json:"Disc"`
	Disc2         float32 `json:"Disc2"`
	Comm          float32 `json:"Comm"`
	Oc            float32 `json:"Oc"`
	Bkp           float32 `json:"Bkp"`
	Ngpw          float32 `json:"Ngpw"`
	Ri            float32 `json:"Ri"`
	Ricom         float32 `json:"Ricom"`
	Npw           float32 `json:"Npw"`
}