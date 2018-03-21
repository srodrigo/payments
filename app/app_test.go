package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateApp(t *testing.T) {
	app := CreateApp()

	assert.NotNil(t, app.Router)
}
