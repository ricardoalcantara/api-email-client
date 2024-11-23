package types

type DashboardDto struct {
	Templates int64 `json:"templates"`
	Emails    int64 `json:"emails"`
	Smtps     int64 `json:"smtps"`
	ApiKeys   int64 `json:"api_keys"`
}
