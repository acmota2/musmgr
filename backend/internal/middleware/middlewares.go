package middleware

import (
	"errors"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const ScopeKey = "scope"
const ClassKey = "classification"
const ReqPermKey = "request_perm"
const CurrentFile = "current_file"

func setRouterScope(scope policies.Scope) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ScopeKey, scope)
		c.Next()
	}
}

func setRouterClassification(class policies.FileClassification) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ClassKey, class)
		c.Next()
	}
}

func RequirePerm(need policies.Perm) gin.HandlerFunc {
	return func(c *gin.Context) {
		scope, ok := GetContextValue[policies.Scope](c, ScopeKey)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		class, ok := GetContextValue[policies.FileClassification](c, ClassKey)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !policies.IsAllowed(scope, class, need) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func FileClassificationBlocking(q *model.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileID, err := uuid.Parse(c.Param("file_id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx := c.Request.Context()

		class, ok := GetContextValue[policies.FileClassification](c, ClassKey)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		queryArgs := model.GetFileParams{ID: fileID, Classification: int16(class)}
		file, err := q.GetFile(ctx, queryArgs)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Set(CurrentFile, file)
		c.Next()
	}
}

func SetAdminRouterScope() gin.HandlerFunc {
	return setRouterScope(policies.ScopeAdmin)
}

func SetAdminRouterClass() gin.HandlerFunc {
	return setRouterClassification(policies.ClassProtected)
}

func SetPublicRouterScope() gin.HandlerFunc {
	return setRouterScope(policies.ScopePublic)
}

func SetPublicRouterClass() gin.HandlerFunc {
	return setRouterClassification(policies.ClassPublic)
}
