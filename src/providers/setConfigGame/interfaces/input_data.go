package setConfigGameinterfaces

type InputData struct {
	XUserWorldToken string `json:"xUserWorldToken" validate:"required"`
	XProxyAuth      string `json:"xProxyAuth"`
	Animations      *bool  `json:"animations"`
	Tutorial        *bool  `json:"tutorial"`
}
