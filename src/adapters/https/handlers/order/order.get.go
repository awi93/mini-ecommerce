package order

import (
	"log"
	"net/http"
	"strconv"

	"github.com/awi93/mini-ecommerce/src/utils/httputil"
	"github.com/go-chi/chi"
)

func (h *Handler) Find(r *http.Request) httputil.Response {
	var idString = chi.URLParam(r, "id")
	if idString == "" {
		return httputil.Response{
			StatusCode: 400,
			Body: map[string]interface{}{
				"error": "bad request",
				"cause": "invalid request from client",
			},
		}
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		return httputil.Response{
			StatusCode: 400,
			Body: map[string]interface{}{
				"error": "bad request",
				"cause": "invalid request from client",
			},
		}
	}

	result, err := h.service.FindOrderDetailById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return httputil.Response{
				StatusCode: 404,
				Body: map[string]interface{}{
					"error": "not found",
					"cause": "data not found",
				},
			}
		}
		return httputil.Response{
			StatusCode: 500,
			Body: map[string]interface{}{
				"error": "internal server error",
				"cause": "fail to fetch data from database",
			},
		}
	}

	return httputil.Response{
		StatusCode: 200,
		Body:       result,
	}
}
