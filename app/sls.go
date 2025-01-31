package app

import (
	"context"
	"database/sql"
	"deficonnect/sonicflare/web"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Success           bool                `json:"success"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              interface{}         `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
}

func wrapHandler(fn func(ctx context.Context, r *http.Request) (Response, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := fn(r.Context(), r)
		if err != nil {
			log.Error(err)
			web.SendErrorfJSON(w, err.Error())
			return
		}

		d, err := json.Marshal(response.Body)
		if err != nil {
			log.Errorf("Error marshalling data: %s", err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(d)
	}
}

func (m module) wrapHandlerWithTx(fn func(ctx context.Context, tx *sql.Tx, r *http.Request) (Response, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := m.db.BeginTransaction()
		if err != nil {
			log.Error(err)
			web.SendErrorfJSON(w, "Internal server DB error")
			return
		}
		response, err := fn(r.Context(), tx, r)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			web.SendErrorfJSON(w, err.Error())
			return
		}

		if response.Success {
			if err := tx.Commit(); err != nil {
				log.Error(err)
				web.SendErrorfJSON(w, "Internal error. Please try again later")
				return
			}
		} else {
			tx.Rollback()
		}

		d, err := json.Marshal(response.Body)
		if err != nil {
			log.Errorf("Error marshalling data: %s", err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(d)
	}
}

func SendErrorfJSON(errorMessage string, args ...interface{}) (Response, error) {
	data := map[string]interface{}{
		"data":    nil,
		"success": false,
		"error":   fmt.Sprintf(errorMessage, args...),
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            data,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func SendAuthErrorfJSON(errorMessage string, args ...interface{}) (Response, error) {
	data := map[string]interface{}{
		"data":           nil,
		"success":        false,
		"error":          fmt.Sprintf(errorMessage, args...),
		"_un_authorized": true,
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            data,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func SendJSON(data interface{}) (Response, error) {
	d := map[string]interface{}{
		"data":    data,
		"success": true,
		"error":   nil,
	}

	resp := Response{
		StatusCode:      200,
		Success:         true,
		IsBase64Encoded: false,
		Body:            d,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func SendPagedJSON(data interface{}, totalCount int64) (Response, error) {
	d := map[string]interface{}{
		"data":    data,
		"total":   totalCount,
		"success": true,
		"error":   nil,
	}

	resp := Response{
		StatusCode:      200,
		Success:         true,
		IsBase64Encoded: false,
		Body:            d,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}
