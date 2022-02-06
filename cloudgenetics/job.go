package cloudgenetics

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func listJobs(c *gin.Context, db *gorm.DB) []Job {
	var jobs []Job
	userid := userid(c, db)
	db.Find(&jobs, Dataset{OwnerID: userid})
	return jobs
}
