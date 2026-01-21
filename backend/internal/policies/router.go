package policies

import "github.com/gin-gonic/gin"

var ScopeKey = "scope"
var ClassKey = "classification"
var ReqPermKey = "request_perm"

func setRouterScope(scope Scope) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ScopeKey, scope)
		c.Next()
	}
}

func RequirePerm(perm Perm) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ReqPermKey, perm)
		c.Next()
	}
}

func SetAdminRouterScope() gin.HandlerFunc {
	return setRouterScope(ScopeAdmin)
}

func SetPublicRouterScope() gin.HandlerFunc {
	return setRouterScope(ScopePublic)
}
