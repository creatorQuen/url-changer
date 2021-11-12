package app

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"url-changer/infrastructure/localservices"
)

func TestUrlCutterService_GetURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	receiveUrl := "youtube.com"
	keyStr := "31567sdf78"

	testTable := []struct {
		description      string
		repoMock         *MockUrlSaver
		returnedVariable string
		setVariable      string
		wantError        error
	}{
		{
			description: "should successful get url",
			repoMock: func(repo *MockUrlSaver) *MockUrlSaver {
				repo.EXPECT().GetFullString(gomock.Any()).Return(receiveUrl, nil).Times(1)
				return repo
			}(NewMockUrlSaver(ctrl)),
			returnedVariable: "youtube.com",
			setVariable:      "31567sdf78",
			wantError:        nil,
		},
		{
			description: "should successful get url",
			repoMock: func(repo *MockUrlSaver) *MockUrlSaver {
				repo.EXPECT().GetFullString(gomock.Any()).Return("", errors.New(receiveUrl)).Times(1)
				return repo
			}(NewMockUrlSaver(ctrl)),
			returnedVariable: "",
			setVariable:      "31567sdf78",
			wantError:        errors.New(receiveUrl),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {
			handle := NewUrlCutterService(nil, testCase.repoMock)
			str, err := handle.GetURL(keyStr)

			if err == nil {
				assert.Equal(t, testCase.returnedVariable, str)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})

	}
}

func TestMockKeyGenerator_MakeKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setUrl := "youtube.com"
	returnedKey := "31567sdf78"

	testTable := []struct {
		description  string
		service1     *MockUrlSaver
		service2     *localservices.MockICutter
		messageUrl   string
		wantResponse string
		wantError    error
	}{
		{
			description: "should succsesful",
			service1: func(mock *MockUrlSaver) *MockUrlSaver {
				mock.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				return mock
			}(NewMockUrlSaver(ctrl)),
			service2: func(mock *localservices.MockICutter) *localservices.MockICutter {
				mock.EXPECT().Generate().Return(returnedKey).Times(1)
				return mock
			}(localservices.NewMockICutter(ctrl)),
			messageUrl:   "youtube.com",
			wantResponse: returnedKey,
			wantError:    nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {
			handle := NewUrlCutterService(testCase.service2, testCase.service1)
			str, err := handle.MakeKey(setUrl)

			if err == nil {
				assert.Equal(t, testCase.wantResponse, str)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})

	}
}
