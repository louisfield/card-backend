package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/louisfield/go-app-backend/cmd/user"
)

type Session struct {
	ID         uuid.UUID   `json:"id"`
	MaxPlayers int         `json:"max_players"`
	LastActive time.Time   `json:"last_active"`
	Users      []user.User `json:"users"`
}

var sessions []Session

func CreateSession(c echo.Context) error {
	session := new(Session)

	session.ID = uuid.New()
	session.LastActive = time.Now()

	if err := c.Bind(session); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sessions = append(sessions, *session)

	return c.JSON(http.StatusOK, session)
}

func AddUserToSession(c echo.Context) error {
	sessionId := c.Param("id")

	if err := c.Bind(session); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	for _, s := range sessions {
		if s.ID.String() == sessionId {
			if err := user.MaybeAddUser(c.FormValue("name"), &s.Users); err != nil {

				return c.String(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusOK, s)
		}
	}
	return c.String(http.StatusBadRequest, "Session does not exist")
}
