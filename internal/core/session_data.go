package core

import (
	"github.com/gin-gonic/gin"
)

type ContextAccessFunc func() *gin.Context

type SessionData struct {
	AccessContext *ContextAccessFunc
}

var _sessionData *SessionData

func SessionDataInstant() SessionData {
	if _sessionData == nil {
		panic("Session data must initial first")
	}
	return *_sessionData
}

func (s SessionData) GetUserInfo() *UserInfo {
	if s.AccessContext != nil {
		ctx := (*s.AccessContext)()
		d, exist := ctx.Get("UserInfo")
		if d == nil || exist == false {
			return nil
		}
		result := d.(UserInfo)
		return &result
	} else {
		return nil
	}
}

func InitSessionData(accessor ContextAccessFunc) {
	_sessionData = new(SessionData)
	_sessionData.AccessContext = &accessor
}

func (s SessionData) Get(key string) interface{} {
	if s.AccessContext != nil {
		ctx := (*s.AccessContext)()
		d, exist := ctx.Get(key)
		if d == nil || exist == false {
			return nil
		}
		return d
	} else {
		return nil
	}
}

func (s SessionData) Set(key string, value interface{}) {
	if s.AccessContext != nil {
		ctx := (*s.AccessContext)()
		ctx.Set(key, value)
	}
}
