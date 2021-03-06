package httphandlers

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-changer/app"
)

func Test_UrlCutter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	returnKey := "52fdfc072182"
	errorServer := echo.NewHTTPError(http.StatusInternalServerError)

	testTable := []struct {
		description  string
		service      *app.MockKeyGenerator
		messageJSON  string
		wantHttpCode int
		wantError    error
	}{
		{
			description: "should return statusOK",
			service: func(mock *app.MockKeyGenerator) *app.MockKeyGenerator {
				mock.EXPECT().MakeKey(gomock.Any()).Return(returnKey, nil).Times(1)
				return mock
			}(app.NewMockKeyGenerator(ctrl)),
			messageJSON:  `{"long_url": "https://www.youtube.com"}`,
			wantHttpCode: http.StatusOK,
			wantError:    nil,
		},
		{
			description:  "should return http.StatusBadRequest when request body data incorrect",
			service:      app.NewMockKeyGenerator(ctrl),
			messageJSON:  `{"long_url"}: "https://www.youtube.com"}`,
			wantHttpCode: http.StatusBadRequest,
			wantError:    echo.NewHTTPError(http.StatusBadRequest),
		},
		{
			description: "should return statusOK",
			service: func(mock *app.MockKeyGenerator) *app.MockKeyGenerator {
				mock.EXPECT().MakeKey(gomock.Any()).Return("", errorServer).Times(1)
				return mock
			}(app.NewMockKeyGenerator(ctrl)),
			messageJSON:  `{"long_url": "https://www.youtube.com"}`,
			wantHttpCode: 500,
			wantError:    echo.NewHTTPError(http.StatusInternalServerError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.messageJSON))
			request.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			recorder := httptest.NewRecorder()
			eh := echo.New()
			context := eh.NewContext(request, recorder)
			context.SetPath("/urlCut")
			handle := NewUrlGenerator(testCase.service)

			if err := handle.UrlCutter(context); err == nil {
				assert.Equal(t, testCase.wantHttpCode, recorder.Code)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})
	}
}

func Test_GetUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	returnUrl := "https://www.youtube.com"
	//inputParam := "key"

	testTable := []struct {
		description  string
		service      *app.MockKeyGenerator
		messageJSON  string
		wantHttpCode int
		wantError    error
		key          string
	}{
		{
			description:  "should return statusOK",
			wantHttpCode: http.StatusMovedPermanently,
			service: func(mock *app.MockKeyGenerator) *app.MockKeyGenerator {
				mock.EXPECT().GetURL(gomock.Any()).Return(returnUrl, nil).Times(1)
				return mock
			}(app.NewMockKeyGenerator(ctrl)),
			messageJSON: `{"long_url": "https://www.youtube.com"}`,
			wantError:   nil,
			key:         "123",
		},
		{
			description:  "should return StatusInternalServerError",
			wantHttpCode: http.StatusInternalServerError,
			service: func(mock *app.MockKeyGenerator) *app.MockKeyGenerator {
				mock.EXPECT().GetURL(gomock.Any()).Return("", echo.NewHTTPError(500)).Times(1)
				return mock
			}(app.NewMockKeyGenerator(ctrl)),
			messageJSON: `{"long_url": "https://www.youtube.com"}`,
			wantError:   echo.NewHTTPError(http.StatusInternalServerError),
			key:         "123",
		},
		{
			description:  "should return StatusBadRequest",
			wantHttpCode: http.StatusBadRequest,
			service:      app.NewMockKeyGenerator(ctrl),
			messageJSON:  `{"long_url": "https://www.youtube.com"}`,
			wantError:    echo.NewHTTPError(http.StatusBadRequest),
			key:          "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/:key", nil)
			request.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			recorder := httptest.NewRecorder()
			eh := echo.New()
			context := eh.NewContext(request, recorder)
			context.SetPath("/:key")
			context.SetParamNames("key")
			context.SetParamValues(testCase.key)
			handle := NewUrlGenerator(testCase.service)

			if err := handle.GetUrl(context); err == nil {
				assert.Equal(t, testCase.wantHttpCode, recorder.Code)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})
	}
}
