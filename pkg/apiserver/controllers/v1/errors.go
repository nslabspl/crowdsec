package v1

import (
	"net/http"

	"github.com/wojtekxtx/crowdsec/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (c *Controller) HandleDBErrors(gctx *gin.Context, err error) {
	switch errors.Cause(err) {
	case database.ItemNotFound:
		gctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	case database.UserExists:
		gctx.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	case database.HashError:
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case database.InsertFail:
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	case database.QueryFail:
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	case database.ParseTimeFail:
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	case database.ParseDurationFail:
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	default:
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}
