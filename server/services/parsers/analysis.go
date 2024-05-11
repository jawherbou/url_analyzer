package parsers

type AnalysisResponse interface {
	// setting the page title.
	SetTitle(title string)

	// setting the page html version.
	SetHtmlVersion(version string)

	// setting if page has login or not.
	SetHasLogin(hasLogin bool)

	// add heading and its level
	AddHeading(heading string, level string)
}

func NewAnalysisResponse() AnalysisResponse {
	headings := make([]Heading, 0)
	return &AnalysisSuccessResponse{
		headings: headings,
	}
}

type AnalysisSuccessResponse struct {
	title    string
	version  string
	hasLogin bool
	headings []Heading
}

type Heading struct {
	tagName string
	levels  []string
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

func (ap *AnalysisSuccessResponse) AddHeading(tag string, level string) {
	// if headings is empty
	if len(ap.headings) == 0 {
		levels := make([]string, 0)
		levels = append(levels, level)
		ap.headings = append(ap.headings, Heading{
			tagName: tag,
			levels:  levels,
		})
		return
	}

	for i, heading := range ap.headings {
		if heading.tagName == tag {
			heading.levels = append(heading.levels, level)
			ap.headings[i] = heading
			break
		}
		if i == len(ap.headings)-1 {
			levels := make([]string, 0)
			levels = append(levels, level)
			ap.headings = append(ap.headings, Heading{
				tagName: tag,
				levels:  levels,
			})
		}
	}
}
