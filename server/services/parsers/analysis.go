package parsers

type AnalysisResponse interface {

	// setting the page title.
	SetTitle(title string)

	// setting the page html version.
	SetHtmlVersion(version string)
}

func NewAnalysisResponse() AnalysisResponse {
	return &AnalysisSuccessResponse{}
}

type AnalysisSuccessResponse struct {
	title   string
	version string
}

func (ap *AnalysisSuccessResponse) SetTitle(title string) {
	ap.title = title
}

func (ap *AnalysisSuccessResponse) SetHtmlVersion(version string) {
	ap.version = version
}
