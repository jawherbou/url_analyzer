package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jawherbou/url_analyzer/server/services/parsers"
)

func Test_ParseTitle(t *testing.T) {
	htmlData := "<title>Sample Title</title>"

	mockAnalysis := &MockAnalysisResponse{}
	titleParser := parsers.NewTitleParser()

	titleParser.Parse(htmlData, mockAnalysis)

	expectedTitle := "Sample Title"
	assert.Equal(t, expectedTitle, mockAnalysis.Title, "SetTitle should be called with the expected title")
}

func Test_ParseEmptyTitle(t *testing.T) {
	htmlData := ""

	mockAnalysis := &MockAnalysisResponse{}
	titleParser := parsers.NewTitleParser()

	titleParser.Parse(htmlData, mockAnalysis)

	expectedTitle := ""
	assert.Equal(t, expectedTitle, mockAnalysis.Title, "SetTitle should not be called if no title exists")
}
