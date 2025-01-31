package app

import (
	"deficonnect/sonicflare/web"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	NOTIFICATION_STATUS_NEW  = 0
	NOTIFICATION_STATUS_READ = 1

	NOTIFICATION_TYPE_TOOLBAR   = 0
	NOTIFICATION_TYPE_DASHBOARD = 1
)

type sendNotificationInput struct {
	Titile     string `json:"title"`
	Content    string `json:"content"`
	ActionLink string `json:"action_link"`
	ActionText string `json:"action_text"`
	Type       int    `josn:"type"`
}

func (m module) sendNotification(w http.ResponseWriter, r *http.Request) {
	var input sendNotificationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Error("sendNotification", "json::Decode", err)
		web.SendErrorfJSON(w, "cannot decode request")
		return
	}
	if err := m.db.NotifyAll(r.Context(), input.Titile, input.Content, input.ActionText, input.ActionLink, input.Type); err != nil {
		m.sendSomethingWentWrong(w, "sendNotification", err)
		return
	}
	web.SendJSON(w, "Sent to all users")
}

func (m module) getUnReadNotificationCount(w http.ResponseWriter, r *http.Request) {
	notificationType, _ := strconv.ParseInt(r.FormValue("type"), 10, 64)
	count, err := m.db.UnReadNotificationCount(r.Context(), m.server.GetUserIDTokenCtx(r), int(notificationType))
	if err != nil {
		log.Error("getNotificationCount", "UnReadNotificationCount", err)
		web.SendErrorfJSON(w, "Something went wrong. Please try again later")
	}
	web.SendJSON(w, count)
}

func (m module) getNewNotifications(w http.ResponseWriter, r *http.Request) {
	pagedReq := web.GetPaginationInfo(r)
	notificationType, _ := strconv.ParseInt(r.FormValue("type"), 10, 64)
	notification, count, err := m.db.GetNewNotifications(r.Context(), m.server.GetUserIDTokenCtx(r), int(notificationType), pagedReq.Offset, pagedReq.Limit)
	if err != nil {
		m.sendSomethingWentWrong(w, "GetNewNotifications", err)
		return
	}
	web.SendPagedJSON(w, notification, count)
}

func (m module) getNotifications(w http.ResponseWriter, r *http.Request) {
	pagedReq := web.GetPaginationInfo(r)
	notificationType, _ := strconv.ParseInt(r.FormValue("type"), 10, 64)
	notification, count, err := m.db.GetNotifications(r.Context(), m.server.GetUserIDTokenCtx(r), int(notificationType), pagedReq.Offset, pagedReq.Limit)
	if err != nil {
		m.sendSomethingWentWrong(w, "GetNotifications", err)
		return
	}
	web.SendPagedJSON(w, notification, count)
}

func (m module) getNotification(w http.ResponseWriter, r *http.Request) {
	notification, err := m.db.GetNotification(r.Context(), r.FormValue("id"))
	if err != nil {
		m.sendSomethingWentWrong(w, "GetNotification", err)
		return
	}
	if notification.AccountID != m.server.GetUserIDTokenCtx(r) {
		web.SendErrorfJSON(w, "Access denied")
		return
	}
	web.SendJSON(w, notification)
}

func (m module) sendSomethingWentWrong(w http.ResponseWriter, fn string, err error) {
	log.Error(fn, err)
	web.SendErrorfJSON(w, "Something went wrong. Please try again later")
}

func (m module) sendSomethingWentWrongSls(fn string, err error) (Response, error) {
	log.Error(fn, err)
	return m.handleError(err)
}
