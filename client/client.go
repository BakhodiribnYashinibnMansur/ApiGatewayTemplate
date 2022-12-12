package client

import (
	"api_getaway_web/config"
	"fmt"
	"sync"

	"api_getaway_web/genproto/user_service"

	"google.golang.org/grpc"
)

var cfg = config.Config()
var (
	onceUserService sync.Once

	instanceUserService user_service.UserServiceClient
)

// User ...
func UserService() user_service.UserServiceClient {
	onceUserService.Do(func() {
		connUser, err := grpc.Dial(
			fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
			grpc.WithInsecure())
		if err != nil {
			panic(fmt.Errorf("user service dial host: %s port:%d err: %s",
				cfg.UserServiceHost, cfg.UserServicePort, err))
		}
		instanceUserService = user_service.NewUserServiceClient(connUser)
	})
	return instanceUserService
}
