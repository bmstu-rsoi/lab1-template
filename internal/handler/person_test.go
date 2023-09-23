package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Astemirdum/person-service/internal/errs"

	"github.com/Astemirdum/person-service/internal/model"

	"github.com/Astemirdum/person-service/internal/handler"
	service_mocks "github.com/Astemirdum/person-service/internal/handler/mocks"
	"github.com/Astemirdum/person-service/pkg/validate"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestHandler_CreatePerson(t *testing.T) {
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

func ptr[T any](a T) *T {
	return &a
}

func TestHandler_ListPerson(t *testing.T) {
	t.Parallel()
	type response struct {
		expectedCode int
		expectedBody string
	}
	type mockBehavior func(r *service_mocks.MockPersonService)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		response     response
		wantErr      bool
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockPersonService) {
				r.EXPECT().
					List(context.Background()).
					Return([]model.Person{{
						ID:      1,
						Name:    "lol",
						Age:     ptr(30),
						Address: ptr("kek"),
						Work:    ptr("omm"),
					},
						{
							ID:      2,
							Name:    "lol1",
							Age:     ptr(31),
							Address: ptr("kek"),
							Work:    ptr("omm"),
						},
					}, nil)
			},
			response: response{
				expectedCode: http.StatusOK,
				expectedBody: `[{"id":1,"name":"lol","age":30,"address":"kek","work":"omm"},{"id":2,"name":"lol1","age":31,"address":"kek","work":"omm"}]`,
			},
			wantErr: false,
		},
		{
			name: "err. internal",
			mockBehavior: func(r *service_mocks.MockPersonService) {
				r.EXPECT().List(context.Background()).
					Return(nil, errors.New("db internal"))
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
				http.MethodGet, "/persons", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			ctx := e.NewContext(r, w)

			tt.mockBehavior(svc)
			err := h.ListPerson(ctx)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.response.expectedCode, w.Code)
				require.Equal(t, tt.response.expectedBody, strings.Trim(w.Body.String(), "\n"))
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

func TestHandler_GetPerson(t *testing.T) {
	t.Parallel()
	type input struct {
		id int
	}
	type response struct {
		expectedCode int
		expectedBody string
	}
	type mockBehavior func(r *service_mocks.MockPersonService, id int)

	tests := []struct {
		name         string
		input        input
		mockBehavior mockBehavior
		response     response
		wantErr      bool
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockPersonService, id int) {
				r.EXPECT().
					Get(context.Background(), id).
					Return(model.Person{
						ID:      id,
						Name:    "lol1",
						Age:     ptr(31),
						Address: ptr("kek"),
						Work:    ptr("omm"),
					}, nil)
			},
			input: input{id: 2},
			response: response{
				expectedCode: http.StatusOK,
				expectedBody: `{"id":2,"name":"lol1","age":31,"address":"kek","work":"omm"}`,
			},
			wantErr: false,
		},
		{
			name: "err. not found",
			mockBehavior: func(r *service_mocks.MockPersonService, id int) {
				r.EXPECT().Get(context.Background(), id).
					Return(model.Person{}, errs.ErrNotFound)
			},
			input: input{id: 2},
			response: response{
				expectedCode: http.StatusNotFound,
				expectedBody: `{"message":"not found"}`,
			},
			wantErr: true,
		},
		{
			name: "err. internal",
			mockBehavior: func(r *service_mocks.MockPersonService, id int) {
				r.EXPECT().Get(context.Background(), id).
					Return(model.Person{}, errors.New("db internal"))
			},
			input: input{id: 3},
			response: response{
				expectedCode: http.StatusInternalServerError,
				expectedBody: `{"message":"db internal"}`,
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

			e.GET("/api/v1/persons/:id", h.GetPerson)

			r := httptest.NewRequest(
				http.MethodGet, fmt.Sprintf("/api/v1/persons/%d", tt.input.id), nil)
			w := httptest.NewRecorder()
			tt.mockBehavior(svc, tt.input.id)
			e.ServeHTTP(w, r)
			require.Equal(t, tt.response.expectedCode, w.Code)
			require.Equal(t, tt.response.expectedBody, strings.Trim(w.Body.String(), "\n"))
		})
	}
}

func TestHandler_DeletePerson(t *testing.T) {
	t.Parallel()
	type input struct {
		id int
	}
	type response struct {
		expectedCode int
		expectedBody string
	}
	type mockBehavior func(r *service_mocks.MockPersonService, id int)

	tests := []struct {
		name         string
		input        input
		mockBehavior mockBehavior
		response     response
		wantErr      bool
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockPersonService, id int) {
				r.EXPECT().
					Delete(context.Background(), id).
					Return(nil)
			},
			input: input{id: 2},
			response: response{
				expectedCode: http.StatusNoContent,
				expectedBody: `ok`,
			},
			wantErr: false,
		},
		{
			name: "ok not found",
			mockBehavior: func(r *service_mocks.MockPersonService, id int) {
				r.EXPECT().Delete(context.Background(), id).
					Return(errs.ErrNotFound)
			},
			input: input{id: 2},
			response: response{
				expectedCode: http.StatusNoContent,
				expectedBody: `ok`,
			},
			wantErr: true,
		},
		{
			name: "err. internal",
			mockBehavior: func(r *service_mocks.MockPersonService, id int) {
				r.EXPECT().Delete(context.Background(), id).
					Return(errors.New("db internal"))
			},
			input: input{id: 3},
			response: response{
				expectedCode: http.StatusInternalServerError,
				expectedBody: `{"message":"db internal"}`,
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

			e.DELETE("/api/v1/persons/:id", h.DeletePerson)

			r := httptest.NewRequest(
				http.MethodDelete, fmt.Sprintf("/api/v1/persons/%d", tt.input.id), nil)
			w := httptest.NewRecorder()
			tt.mockBehavior(svc, tt.input.id)
			e.ServeHTTP(w, r)
			require.Equal(t, tt.response.expectedCode, w.Code)
			require.Equal(t, tt.response.expectedBody, strings.Trim(w.Body.String(), "\n"))
		})
	}
}
