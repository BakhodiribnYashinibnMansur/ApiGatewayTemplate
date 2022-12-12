package handler_test

import (
	"api_getaway_web/config"
	"api_getaway_web/util/logrus_log"
	"encoding/json"

	"api_getaway_web/package/handler"
	"api_getaway_web/package/response"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

var (
	Host   = "http://localhost:9090"
	Client *resty.Client
	Logrus *logrus_log.Logger
	Token  string
	Router *gin.Engine
)

func TestInitMain(t *testing.T) {

	logrus := logrus_log.GetLogger()
	cfg := config.Config()
	handlers := handler.NewHandler(logrus, cfg)
	Router = handlers.InitRoutes()
	Logrus = logrus_log.GetLogger()
	Client = resty.New()

}

func TestUserSignIn(t *testing.T) {

	t.Parallel()

	type UserSignInStruct struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type TestCase struct {
		Title               string            `json:"title"`
		URL                 string            `json:"url"`
		Method              string            `json:"method"`
		Param               map[string]string `json:"query"`
		Body                UserSignInStruct  `json:"body"`
		BodyInCorrectStatus bool              `json:"body_in_correct_status"`
		BodyInCorrect       map[string]string `json:"body_in_correct"`
		WantStatusCode      int               `json:"wantStatusCode"`
		WantResult          string            `json:"wantResult"`
		WantError           string            `json:"wantError"`
	}

	tests := []TestCase{
		{
			Title:  "200 Correct Request User Sign In",
			URL:    Host + "/api/v1/account/sign-in",
			Method: "POST",
			Param:  map[string]string{},
			Body: UserSignInStruct{
				Username: "username",
				Password: "password",
			},
			BodyInCorrectStatus: false,
			WantStatusCode:      200,
			WantResult:          "",
			WantError:           "",
		},
		{
			Title:  "Not Found Request User Sign In",
			URL:    Host + "/api/v1/account/sign-in",
			Method: "POST",
			Param:  map[string]string{},
			Body: UserSignInStruct{
				Username: "username2",
				Password: "password2",
			},
			BodyInCorrectStatus: false,
			WantStatusCode:      404,
			WantResult:          "",
			WantError:           "The requested resource not found.",
		},
		{
			Title:               "Body InCorrectRequest User Sign In",
			URL:                 Host + "/api/v1/account/sign-in",
			Method:              "POST",
			Param:               map[string]string{},
			BodyInCorrectStatus: true,
			BodyInCorrect: map[string]string{
				"user": "username",
				"pas":  "pass",
			},
			WantStatusCode: 400,
			WantResult:     "",
			WantError:      "The requested resource not found.",
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			logrus.Info("---------- TEST CASE TITLE ", test.Title)

			var body interface{}

			if test.BodyInCorrectStatus {
				body = test.BodyInCorrect
			}
			body = test.Body

			resp, err := HttpClientRequest(Router, test.URL, test.Method, test.Param, Token, body)
			if err != nil {
				t.Error(err)
			}

			if statusCode := resp.Code; statusCode != test.WantStatusCode {
				t.Errorf("Status code %v want %v", statusCode, test.WantStatusCode)
			}
			var data response.Response
			err = json.Unmarshal(resp.Body.Bytes(), &data)
			if err != nil {
				t.Error(err)
			}

			if statusCode := resp.Code; statusCode == test.WantStatusCode {
				if data.Data == nil {
					t.Error("Data should be nil")
				}
			}
		})
	}
}
