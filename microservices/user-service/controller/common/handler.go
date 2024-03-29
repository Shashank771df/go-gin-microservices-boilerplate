package common

import (
	"app/core"
	"app/core/database"
	c "app/microservices/user-service/config"
	l "app/microservices/user-service/lib"
	"net/http"
)

type Handler struct {
}

func (o Handler) Sync(ctx core.AppContext) {
	ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"Sync": true,
	})
}

func (o Handler) Health(ctx core.AppContext) {
	ctx.JSON(http.StatusOK, l.DB.CheckDBHealth(database.HealthIn{
		Micro: c.AppInfo.Name,
		Trace: "",
	}))
}
