package handler

import (
	"api_getaway_web/config"
	"api_getaway_web/docs"
	"api_getaway_web/package/response"
	"api_getaway_web/tools/middleware"
	"api_getaway_web/util/logrus_log"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	logrus *logrus_log.Logger
	config *config.Configuration
}

func NewHandler(logrus *logrus_log.Logger, config *config.Configuration) *Handler {
	return &Handler{logrus: logrus, config: config}
}

func (handler *Handler) InitRoutes() (route *gin.Engine) {
	// logrus := handler.logrus
	config := handler.config
	route = gin.New()
	middleware.GinMiddleware(route)
	//swagger settings
	docs.SwaggerInfo.Title = config.AppName
	docs.SwaggerInfo.Version = config.AppVersion
	docs.SwaggerInfo.Host = config.AppURL
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//routers
	Ping(route, handler)
	UserRouter(route, handler)
	return
}

// handler Funcs
func (handler *Handler) handleResponse(ctx *gin.Context, status response.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		handler.logrus.Info(
			"---Response--->",
			"  code  :", status.Code,
			"  status  :", status.Status,
			"  description  :", status.Description,
			"  data  :", data,
		)
	case code < 400:
		handler.logrus.Warn(
			"!!!Response--->",
			"  code  :", status.Code,
			"  status  :", status.Status,
			"  description  :", status.Description,
			"  data  :", data,
		)
	default:
		handler.logrus.Error(
			"!!!Response--->  ",
			"  code  :", status.Code,
			"  status  :", status.Status,
			"  description  :", status.Description,
			"  data  :", data,
		)
	}
	ctx.JSON(status.Code, response.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (handler *Handler) GetOffsetParam(ctx *gin.Context) (offset int, err error) {
	offsetStr := ctx.DefaultQuery("offset", handler.config.DefaultOffset)

	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		handler.handleResponse(ctx, response.BadEnvironment, ErrorNotANumberOffset)
		return
	}

	return offset, nil
}

func (handler *Handler) GetLimitParam(ctx *gin.Context) (limit int, err error) {
	limitStr := ctx.DefaultQuery("limit", handler.config.DefaultLimit)

	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handler.handleResponse(ctx, response.BadEnvironment, ErrorNotANumberLimit)
		return
	}

	return limit, nil
}

func (handler *Handler) GetStringParam(ctx *gin.Context, query string) (param string) {
	param = ctx.Query(query)
	if param == "" {
		handler.handleResponse(ctx, response.BadEnvironment, ErrorParamIsEmpty)
		return
	}

	return param
}

func (handler *Handler) GrpcErrorConvert(ctx *gin.Context, serviceError error) {
	st, ok := status.FromError(serviceError)
	if !ok || st.Code() == codes.Internal {
		handler.handleResponse(ctx, response.InternalServerError, st.Message())
	} else if st.Code() == codes.NotFound {
		handler.handleResponse(ctx, response.NotFound, st.Message())
	} else if st.Code() == codes.InvalidArgument {
		handler.handleResponse(ctx, response.InvalidArgument, st.Message())
	} else if st.Code() == codes.Unavailable {
		handler.handleResponse(ctx, response.Unavailable, st.Message())
	} else if st.Code() == codes.OK {
		handler.handleResponse(ctx, response.OK, st.Message())
	} else if st.Code() == codes.AlreadyExists {
		handler.handleResponse(ctx, response.AlreadyExists, st.Message())
	} else if st.Code() == codes.Canceled {
		handler.handleResponse(ctx, response.Canceled, st.Message)
	} else if st.Code() == codes.Unknown {
		handler.handleResponse(ctx, response.Unknown, st.Message())
	} else if st.Code() == codes.DeadlineExceeded {
		handler.handleResponse(ctx, response.DeadlineExceeded, st.Message())
	} else if st.Code() == codes.PermissionDenied {
		handler.handleResponse(ctx, response.PermissionDenied, st.Message())
	} else if st.Code() == codes.ResourceExhausted {
		handler.handleResponse(ctx, response.ResourceExhausted, st.Message())
	} else if st.Code() == codes.FailedPrecondition {
		handler.handleResponse(ctx, response.FailedPrecondition, st.Message())
	} else if st.Code() == codes.Aborted {
		handler.handleResponse(ctx, response.Aborted, st.Message())
	} else if st.Code() == codes.OutOfRange {
		handler.handleResponse(ctx, response.OutOfRange, st.Message())
	} else if st.Code() == codes.Unimplemented {
		handler.handleResponse(ctx, response.Unimplemented, st.Message())
	} else if st.Code() == codes.DataLoss {
		handler.handleResponse(ctx, response.DataLos, st.Message())
	} else if st.Code() == codes.Unauthenticated {
		handler.handleResponse(ctx, response.Unauthorized, st.Message())
	} else {
		handler.handleResponse(ctx, response.Unknown, st.Message())
	}

}

// Can be added as many as need like belows examples
// 400	BAD_CONTINUATION_TOKEN	Invalid continuation token passed.
// 400	BAD_PAGE	Page number does not exist or is an invalid format (e.g. negative).
// 400	BAD_REQUEST	The resource you’re creating already exists.
// 400	INVALID_ARGUMENT	Invalid argument value passed.
// 400	INVALID_AUTH	Authentication/OAuth token is invalid.
// 400	INVALID_AUTH_HEADER	Authentication header is invalid.
// 400	INVALID_BATCH	Batched request is missing or invalid.
// 400	INVALID_BODY	A request body that was not in JSON format was passed.
// 400	UNSUPPORTED_OPERATION	Requested operation not supported.
// 401	ACCESS_DENIED	Authentication unsuccessful.
// 401	NO_AUTH	Authentication not provided.
// 403	NOT_AUTHORIZED	User has not been authorized to perform that action.
// 404	NOT_FOUND	Invalid URL.
// 405	METHOD_NOT_ALLOWED	Method is not allowed for this endpoint.
// 409	REQUEST_CONFLICT	Requested operation resulted in conflict.
// 429	HIT_RATE_LIMIT	Hourly rate limit has been reached for this token. Default rate limits are 2,000 calls per hour.
// 500	EXPANSION_FAILED	Unhandled error occurred during expansion; the request is likely to succeed if you don’t ask for expansions, but contact Eventbrite support if this problem persists.
// 500	INTERNAL_ERROR	Unhandled error occurred in Eventbrite. contact Eventbrite support if this problem persists.
