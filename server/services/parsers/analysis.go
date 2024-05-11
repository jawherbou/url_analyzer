package parsers

type AnalysisResponse interface {
	// setting the page title.
	SetTitle(title string)

	// setting the page html version.
	SetHtmlVersion(version string)

	// setting if page has login or not.
	SetHasLogin(hasLogin bool)

	// add heading and its level
	AddHeading(heading string, content string)

	// add link
	AddLink(link Link)
}

type AnalysisSuccessResponse struct {
	Title    string    `json:"title"`
	Version  string    `json:"htmlVersion"`
	HasLogin bool      `json:"hasLogin"`
	Headings []Heading `json:"headings"`
	Links    []Link    `json:"links"`
}

type Heading struct {
	TagName  string   `json:"tagName"`
	Contents []string `json:"contents"`
}

type Link struct {
	Url       string `json:"url"`
	LinkType  string `json:"linkType"`
	Reachable bool   `json:"reachable"`
}

func NewAnalysisResponse() AnalysisResponse {
	headings := make([]Heading, 0)
	links := make([]Link, 0)
	return &AnalysisSuccessResponse{
		Headings: headings,
		Links:    links,
	}
}

func (ap *AnalysisSuccessResponse) SetTitle(title string) {
	ap.Title = title
}

func (ap *AnalysisSuccessResponse) SetHtmlVersion(version string) {
	ap.Version = version
}

func (ap *AnalysisSuccessResponse) SetHasLogin(hasLogin bool) {
	ap.HasLogin = hasLogin
}

func (ap *AnalysisSuccessResponse) AddHeading(tag string, content string) {
	// if headings is empty
	if len(ap.Headings) == 0 {
		contents := make([]string, 0)
		contents = append(contents, content)
		ap.Headings = append(ap.Headings, Heading{
			TagName:  tag,
			Contents: contents,
		})
		return
	}

	for i, heading := range ap.Headings {
		if heading.TagName == tag {
			heading.Contents = append(heading.Contents, content)
			ap.Headings[i] = heading
			break
		}
		if i == len(ap.Headings)-1 {
			contents := make([]string, 0)
			contents = append(contents, content)
			ap.Headings = append(ap.Headings, Heading{
				TagName:  tag,
				Contents: contents,
			})
		}
	}
}

func (ap *AnalysisSuccessResponse) AddLink(link Link) {
	ap.Links = append(ap.Links, link)
}
