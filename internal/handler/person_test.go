package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Astemirdum/person-service/internal/model"

	"github.com/Astemirdum/person-service/internal/handler"
	service_mocks "github.com/Astemirdum/person-service/internal/handler/mocks"
	"github.com/Astemirdum/person-service/pkg/validate"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestHandler_Create(t *testing.T) {
	t.Parallel()
	type input struct {
		body string
		req  model.Person
	}
	type response struct {
		expectedCode int
		expectedBody string
		header       string
	}
	type mockBehavior func(r *service_mocks.MockPersonService, req model.Person)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		input        input
		response     response
		wantErr      bool
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockPersonService, req model.Person) {
				r.EXPECT().
					Create(context.Background(), req).
					Return(1, nil)
			},
			input: input{
				body: `{
    "name": "lol",
    "address": "kek",
    "work": "lolkek",
    "age": 31
}`,
			},
			response: response{
				expectedCode: http.StatusOK,
				expectedBody: `OK`,
				header:       `/api/v1/persons/1`, //
			},
			wantErr: false,
		},
		{
			name:         "err. name required",
			mockBehavior: func(r *service_mocks.MockPersonService, inp model.Person) {},
			input: input{
				body: `{
    "address": "kek",
    "work": "lolkek",
    "age": 31
}`,
			},
			response: response{
				expectedCode: http.StatusBadRequest,
				expectedBody: `{"message":"Key: 'Person.Name' Error:Field validation for 'Name' failed on the 'required' tag","errors":{"additionalProperties":""}}`,
			},
			wantErr: true,
		},
		{
			name: "err. internal",
			mockBehavior: func(r *service_mocks.MockPersonService, inp model.Person) {
				r.EXPECT().Create(context.Background(), inp).
					Return(0, errors.New("db internal"))
			},
			input: input{
				body: `{
    "name": "lol",
    "address": "kek",
    "work": "lolkek",
    "age": 31
}`,
			},
			response: response{
				expectedCode: http.StatusInternalServerError,
				expectedBody: `code=500, message=db internal`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()
			svc := service_mocks.NewMockPersonService(c)
			log := zap.NewExample().Named("test")
			h := handler.New(svc, log)

			e := echo.New()
			e.Validator = validate.NewCustomValidator()

			r := httptest.NewRequest(
				http.MethodPost, "/persons", strings.NewReader(tt.input.body))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			ctx := e.NewContext(r, w)
			require.NoError(t, json.NewDecoder(strings.NewReader(tt.input.body)).Decode(&tt.input.req))

			tt.mockBehavior(svc, tt.input.req)
			err := h.CreatePerson(ctx)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.response.expectedCode, w.Code)
				require.Equal(t, tt.response.expectedBody, strings.Trim(w.Body.String(), "\n"))
				require.Equal(t, tt.response.header, w.Header().Get("Location"))
			} else {
				require.Error(t, err)
				var er *echo.HTTPError
				if errors.As(err, &er) {
					require.Equal(t, tt.response.expectedCode, er.Code)
					require.Equal(t, tt.response.expectedBody, er.Error())
				} else {
					require.Equal(t, tt.response.expectedCode, w.Code)
					require.Equal(t, tt.response.expectedBody, strings.Trim(w.Body.String(), "\n"))
				}
			}
		})
	}
}
