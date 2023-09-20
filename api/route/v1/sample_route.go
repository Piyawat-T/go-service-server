package route

import (
	"time"

	"github.com/Piyawat-T/go-service-server/api/controller"
	"github.com/Piyawat-T/go-service-server/bootstrap"
	"github.com/Piyawat-T/go-service-server/usecase"
	"github.com/gin-gonic/gin"
)

func SampleRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	sc := controller.SampleController{
		Env:           env,
		SampleUsecase: usecase.NewSampleUsecase(timeout),
	}
	group.GET("/sample", sc.Sample)
}
