package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	JWT "github.com/ezzycreative1/majoo-test/app/user/handler"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthMiddleware for Authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := c.Request.Header.Get("secret-key")
		env := os.Getenv("ENV")
		if env != "local" && govalidator.IsNull(secretKey) {
			if secretKey != os.Getenv("SECRET_KEY") {
				c.AbortWithStatus(401)
			}
		}
		c.Next()
	}
}

// TokenAuthMiddleware ..
// func TokenAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := auth.TokenValid(c)
// 		if err != nil {
// 			BaseHandler.RespondUnauthorized(c, "")
// 			return
// 		}
// 		c.Next()
// 	}
// }

//AuthUserMiddlerWare ...
func AuthorizeUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})
			return

		}

		data := strings.Split(authHeader, " ")
		if len(data) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization bearer header found"})
			return
		}
		log.Println(len(data))
		tokenString := authHeader[len(BearerSchema):]

		if token, err := JWT.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token or Token Expired"})

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("userID", claims["userID"])
					fmt.Println("during authorization", claims["userID"])
					isExp := getTokenRemainingValidity(claims["exp"])
					log.Println(isExp)
					if !isExp {
						ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
							"error": "Token Expired",
						})
						return
					}
					// if claims["roleID"] != 0 {
					// 	ctx.AbortWithStatus(http.StatusUnauthorized)
					// }
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}

			}
		}
	}
}

// MaxAllowed specify max allowed connections
func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}

func getTokenRemainingValidity(timestamp interface{}) bool {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return true
		}
	}
	return false
}

// LoggerToFile ..
func LoggerToFile() gin.HandlerFunc {
	// log file
	fileName := path.Join("logging", "lintasarta.log")
	// write file
	//src, err := os.OpenFile("lintasarta.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	// instantiation
	logger := logrus.New()
	// Set output
	logger.Out = src
	// Set log level
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.AddHook(lfHook)
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// Stop time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Log format
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
