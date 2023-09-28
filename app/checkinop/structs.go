package checkinop

// {"retcode":0,"message":"OK","data":{"total_sign_day":11,"today":"2023-09-13","is_sign":true,"first_bind":false,"is_sub":true,"region":"os_asia","month_last_day":false}}
type CheckInResponse struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    CheckInData `json:"data"`
}

type CheckInData struct {
	TotalSignDay int    `json:"total_sign_day"`
	Today        string `json:"today"`
	IsSign       bool   `json:"is_sign"`
	FirstBind    bool   `json:"first_bind"`
	IsSub        bool   `json:"is_sub"`
	Region       string `json:"region"`
	MonthLastDay bool   `json:"month_last_day"`
}

// {"retcode":0,"message":"OK","data":{"total_sign_day":24,"today":"2023-09-27","is_sign":true,"first_bind":false,"is_sub":true,"region":"os_asia","month_last_day":false}}
type ClaimResult struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		TotalSignDay int    `json:"total_sign_day"`
		Today        string `json:"today"`
		IsSign       bool   `json:"is_sign"`
		FirstBind    bool   `json:"first_bind"`
		IsSub        bool   `json:"is_sub"`
		Region       string `json:"region"`
		MonthLastDay bool   `json:"month_last_day"`
	} `json:"data"`
	AppMessage string `json:"-"`
}
