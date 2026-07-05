package handler

import (
	"net/http"
	"pipeline/metrics"

	"api/internal/api/middleware"
	"api/internal/dto"
	"api/internal/event"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service event.Service
}

func NewEventHandler(
	service event.Service,
) *EventHandler {
	return &EventHandler{
		service: service,
	}
}

func (h *EventHandler) Track(
	c *gin.Context,
) {
	var request dto.TrackEventsRequest

	if err := c.ShouldBindJSON(
		&request,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	user := middleware.User(c)

	if user == nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "unauthorized",
			},
		)

		return
	}

	events := make([]event.Event, 0, len(request.Events))

	for _, dtoEvent := range request.Events {
		dtoEvent.UserID = user.ID
		dtoEvent.Name = user.Name
		e, err := event.FromDTO(dtoEvent)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		events = append(events, e)
	}

	ctx := event.WithRequestMetadata(

		c.Request.Context(),

		event.RequestMetadata{
			IPAddress: c.ClientIP(),

			UserAgent: c.Request.UserAgent(),

			Referrer: c.Request.Referer(),

			Language: c.GetHeader(
				"Accept-Language",
			),

			RequestID: c.GetString(
				"requestID",
			),

			SDKVersion: c.GetHeader(
				"X-SDK-Version",
			),
		},
	)

	metrics.EventsReceived.Add(
		float64(len(events)),
	)

	if err := h.service.TrackBatch(
		ctx,
		events,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusAccepted,
		gin.H{
			"status": "accepted",
		},
	)
}
