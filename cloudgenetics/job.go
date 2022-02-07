package cloudgenetics

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func listJobs(c *gin.Context, db *gorm.DB) []Job {
	var jobs []Job
	userid := userid(c, db)
	db.Find(&jobs, Dataset{OwnerID: userid})
	return jobs
}

func getJobInfo(c *gin.Context, db *gorm.DB) Job {
	var job Job
	jobid, err := uuid.Parse(c.Param("uuid"))
	owner_id := userid(c, db)
	if err != nil {
		log.Print("Get files dataset: ", err)
	}
	db.Find(&job, Job{ID: jobid, OwnerID: owner_id})
	return job
}
