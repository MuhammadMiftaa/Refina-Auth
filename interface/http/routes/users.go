package routes

import (
	"refina-auth/interface/http/handler"
	"refina-auth/internal/repository"
	"refina-auth/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func UserRoutes(version *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	User_repo := repository.NewUsersRepository(db)
	User_serv := service.NewUsersService(User_repo)

	OTP_repo := repository.NewOTPRepository(redis)
	OTP_serv := service.NewOTPService(OTP_repo)

	User_handler := handler.NewUsersHandler(User_serv, OTP_serv)

	auth := version.Group("/auth")
	{
		auth.POST("login", User_handler.Login)
		auth.POST("register", User_handler.Register)
		auth.POST("send/otp", User_handler.SendOTP)
		auth.POST("verify/otp", User_handler.VerifyOTP)

		auth.GET("google/oauth", User_handler.OAuthHandler("google"))
		auth.GET("callback/google", User_handler.CallbackGoogle)
		auth.GET("github/oauth", User_handler.OAuthHandler("github"))
		auth.GET("callback/github", User_handler.CallbackGithub)
		auth.GET("microsoft/oauth", User_handler.OAuthHandler("microsoft"))
		auth.GET("callback/microsoft", User_handler.CallbackMicrosoft)
	}
}
