package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func RunGenericMethodPost(t *testing.T, expectCode int, fhandler func(ctx echo.Context) error, target string, ctxParams *ContextParams, expectJson, bodyJson string) {
	request := httptest.NewRequest(
		http.MethodPost,
		target,
		strings.NewReader(bodyJson),
	)
	recorder := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(request, recorder)

	if ctxParams != nil {
		c.SetParamNames(ctxParams.Keys...)
		c.SetParamValues(ctxParams.Values...)
	}

	if assert.NoError(t, fhandler(c)) {
		assert.Equal(t, expectCode, recorder.Code)
		assert.JSONEq(t, expectJson, recorder.Body.String())
	}
}

func RunGenericMethodGet(t *testing.T, expectCode int, fhandler func(ctx echo.Context) error, target string, ctxParams *ContextParams, expectJson string) {
	request := httptest.NewRequest(http.MethodGet, target, nil)

	recorder := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(request, recorder)

	if ctxParams != nil {
		c.SetParamNames(ctxParams.Keys...)
		c.SetParamValues(ctxParams.Values...)
	}

	if assert.NoError(t, fhandler(c)) {
		assert.Equal(t, expectCode, recorder.Code)
		assert.JSONEq(t, expectJson, recorder.Body.String())
	}
}
