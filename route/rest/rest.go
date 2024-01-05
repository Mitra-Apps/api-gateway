package rest

import (
	"net/http"

	storePb "github.com/Mitra-Apps/be-store-service/domain/proto/store"
	userPb "github.com/Mitra-Apps/be-user-service/domain/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/labstack/echo/v4"
)

type Rest struct {
	userService  userPb.UserServiceClient
	storeService storePb.StoreServiceClient
}

func New(userService userPb.UserServiceClient, storeService storePb.StoreServiceClient) *Rest {
	return &Rest{userService, storeService}
}

func (r *Rest) Register(e *echo.Echo) {
	e.GET("/ping", r.ping)
	user := e.Group("/api/v1/user")
	user.GET("", r.getUsers)
	user.POST("/login", r.login)

	store := e.Group("/api/v1/stores")
	r.registerStoreService(store)
}

func (r *Rest) ping(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}

func convertGrpcToHttpErrorResponse(err error) *echo.HTTPError {
	// Handling gRPC errors
	statusCode, ok := status.FromError(err)
	if ok {
		httpStatusCode := grpcStatusToHTTPStatus(statusCode.Code())
		return echo.NewHTTPError(httpStatusCode, statusCode.Message())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, "Failed getting error message")
}

func grpcStatusToHTTPStatus(grpcCode codes.Code) int {
	switch grpcCode {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusRequestedRangeNotSatisfiable
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
