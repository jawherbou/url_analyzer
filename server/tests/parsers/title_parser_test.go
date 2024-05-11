package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jawherbou/url_analyzer/server/services/parsers"
)

// MockAnalysisResponse is a mock implementation of the AnalysisResponse interface
type MockAnalysisResponse struct {
	Title string
}

// AddHeading implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) AddHeading(heading string, level string) {
	panic("unimplemented")
}

// AddLink implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) AddLink(link parsers.Link) {
	panic("unimplemented")
}

// SetHasLogin implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) SetHasLogin(hasLogin bool) {
	panic("unimplemented")
}

// SetHtmlVersion implements parsers.AnalysisResponse.
func (m *MockAnalysisResponse) SetHtmlVersion(version string) {
	panic("unimplemented")
}

// SetTitle is a mock implementation of the SetTitle method
func (m *MockAnalysisResponse) SetTitle(title string) {
	m.Title = title
}
func Test_ParseTitle(t *testing.T) {
	htmlData := "<title>Sample Title</title>"

	mockAnalysis := &MockAnalysisResponse{}
	titleParser := parsers.NewTitleParser()

	titleParser.Parse(htmlData, mockAnalysis)

	expectedTitle := "Sample Title"
	assert.Equal(t, expectedTitle, mockAnalysis.Title, "SetTitle should be called with the expected title")

}
