package parsers

type AnalysisResponse interface {
	// setting the page title.
	SetTitle(title string)

	// setting the page html version.
	SetHtmlVersion(version string)

	// setting if page has login or not.
	SetHasLogin(hasLogin bool)
}

func NewAnalysisResponse() AnalysisResponse {
	return &AnalysisSuccessResponse{}
}

type AnalysisSuccessResponse struct {
	title    string
	version  string
	hasLogin bool
}

func (ap *AnalysisSuccessResponse) SetTitle(title string) {
	ap.title = title
}

func (ap *AnalysisSuccessResponse) SetHtmlVersion(version string) {
	ap.version = version
}

func (ap *AnalysisSuccessResponse) SetHasLogin(hasLogin bool) {
	ap.hasLogin = hasLogin
}
