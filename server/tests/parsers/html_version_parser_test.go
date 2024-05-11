package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jawherbou/url_analyzer/server/services/parsers"
)

func Test_ParseHtmlVersion(t *testing.T) {
	htmlData := "<!DOCTYPE html>"

	mockAnalysis := &MockAnalysisResponse{}
	htmlVersionParser := parsers.NewHtmlVersionParser()

	htmlVersionParser.Parse(htmlData, mockAnalysis)

	expectedVersion := "HTML 5"
	assert.Equal(t, expectedVersion, mockAnalysis.Version, "SetHtmlVersion should be called with the expected version")
}
