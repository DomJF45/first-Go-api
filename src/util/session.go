package util

import (
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
)

var SessionManager *scs.SessionManager

func InitSession() {
	SessionManager = scs.New()
	SessionManager.Store = mysqlstore.New(DB)
}
