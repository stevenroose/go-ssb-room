// Package members implements helpers for accessing the currently logged in admin or moderator of an active request.
package members

import (
	"context"
	"net/http"

	"github.com/ssb-ngi-pointer/go-ssb-room/roomdb"
	"go.mindeco.de/http/auth"
)

type roomMemberContextKeyType string

var roomMemberContextKey roomMemberContextKeyType = "ssb:room:httpcontext:member"

// FromContext returns the member or nil if not logged in
func FromContext(ctx context.Context) *roomdb.Member {
	v := ctx.Value(roomMemberContextKey)

	m, ok := v.(roomdb.Member)
	if !ok {
		return nil
	}

	return &m
}

// ContextInjecter returns middleware for injecting a member into the context of the request.
// Retreive it using FromContext(ctx)
func ContextInjecter(mdb roomdb.MembersService, a *auth.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			v, err := a.AuthenticateRequest(req)
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			mid, ok := v.(int64)
			if !ok {
				next.ServeHTTP(w, req)
				return
			}

			member, err := mdb.GetByID(req.Context(), mid)
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			ctx := context.WithValue(req.Context(), roomMemberContextKey, member)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}

// TemplateHelper returns a function to be used with the http/render package.
// It has to return a function twice because the first is evaluated with the request before it gets passed onto html/template's FuncMap.
func TemplateHelper() func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		no := func() *roomdb.Member { return nil }

		member := FromContext(r.Context())
		if member == nil {
			return no
		}

		yes := func() *roomdb.Member { return member }
		return yes
	}
}