package parsers

import "github.com/jawherbou/url_analyzer/server/services/parsers"

// MockAnalysisResponse is a mock implementation of the AnalysisResponse interface
type MockAnalysisResponse struct {
	Title    string
	Version  string
	HasLogin bool
	Headings []MockHeading
}

type MockHeading struct {
	TagName  string
	Contents []string
}

// AddHeading implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) AddHeading(heading string, content string) {
	contents := []string{content}
	m.Headings = append(m.Headings, MockHeading{
		TagName:  heading,
		Contents: contents,
	})
}

// AddLink implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) AddLink(link parsers.Link) {
	panic("unimplemented")
}

// SetHasLogin implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) SetHasLogin(hasLogin bool) {
	m.HasLogin = hasLogin
}

// SetHtmlVersion implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) SetHtmlVersion(version string) {
	m.Version = version
}

// SetTitle is a mock implementation of the SetTitle method
func (m *MockAnalysisResponse) SetTitle(title string) {
	m.Title = title
}
