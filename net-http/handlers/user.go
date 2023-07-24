package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/learn-go/net-http/database"
)

const (
	// 64 为编码秘钥
	cookieStoreAuthKey = "fd0ae3c9fcf92c705a0dba902d6982009e16642ac28dfefdb63b6965e280dc12"
	// 32 位 AES 加密秘钥
	cookieStoreEncryptKey = "94d7992c78e1a18a696542d2706d1099"

	sessionCookieName = "user-session"
)

var sessionStore *sessions.CookieStore

func init() {
	sessionStore = sessions.NewCookieStore(
		[]byte(cookieStoreAuthKey),
		[]byte(cookieStoreEncryptKey),
	)

	sessionStore.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   60 * 15,
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, sessionCookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("username")
	pass := r.FormValue("password")
	_, err = database.AuthenticateUser(name, pass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session.Values["authenticated"] = true
	err = session.Save(r, w)
	_, _ = fmt.Fprintln(w, "Login success", err)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, sessionCookieName)
	session.Values["authenticated"] = false
	err := session.Save(r, w)
	_, _ = fmt.Fprintln(w, "Logout success", err)
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, sessionCookieName)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	_, _ = fmt.Fprintln(w, "User Profile")
}
