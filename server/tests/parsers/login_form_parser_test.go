package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jawherbou/url_analyzer/server/services/parsers"
)

func Test_ParseLoginForm(t *testing.T) {
	htmlData := "<form><input type='password'/></form>"

	mockAnalysis := &MockAnalysisResponse{}
	loginFormParser := parsers.NewLoginFormParser()

	loginFormParser.Parse(htmlData, mockAnalysis)

	expectedLoginFound := true
	assert.Equal(t, expectedLoginFound, mockAnalysis.HasLogin, "SetHasLogin should be called with true value as there is a login form")
}
