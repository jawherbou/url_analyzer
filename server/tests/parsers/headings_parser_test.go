package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jawherbou/url_analyzer/server/services/parsers"
)

func Test_ParseHeadings(t *testing.T) {
	htmlData := "<h1>first heading</h1>"

	mockAnalysis := &MockAnalysisResponse{}
	headingsParser := parsers.NewHeadingsParser()

	headingsParser.Parse(htmlData, mockAnalysis)

	contents := []string{"first heading"}
	expectedHeadings := []MockHeading{{
		TagName:  "h1",
		Contents: contents,
	}}
	assert.Equal(t, expectedHeadings, mockAnalysis.Headings, "AddHeading should be called with the expected headings")
}
