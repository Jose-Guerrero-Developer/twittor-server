package routes

import (
	"github.com/devJGuerrero/twittor-server/controllers"
	"github.com/devJGuerrero/twittor-server/galex/router"
	"github.com/devJGuerrero/twittor-server/middlewares"
)

var subscribe = router.Subscribe

/*API subscribe the routes to the *mux.Router instance */
func API() {
	/* Route Start */
	var Auth controllers.Auth
	subscribe("POST", "/api/sign", Auth.Sign)
	/* Routes Users */
	var User controllers.User
	subscribe("POST", "/api/users", middlewares.ValidateTokenAccess(User.Store))
	/* Routes Profile */
	var Profile controllers.Profile
	subscribe("GET", "/api/profiles", middlewares.ValidateTokenAccess(Profile.Get))
	subscribe("GET", "/api/profiles/{id}/avatar", middlewares.ValidateTokenAccess(Profile.GetAvatar))
	subscribe("GET", "/api/profiles/{id}/banner", middlewares.ValidateTokenAccess(Profile.GetBanner))
	subscribe("PUT", "/api/profiles/{id}", middlewares.ValidateTokenAccess(Profile.Update))
	/* Routes Tweets */
	var Tweet controllers.Tweet
	subscribe("GET", "/api/tweets", middlewares.ValidateTokenAccess(Tweet.Get))
	subscribe("GET", "/api/tweets/{id}", middlewares.ValidateTokenAccess(Tweet.GetID))
	subscribe("GET", "/api/tweets/profile/{id}", middlewares.ValidateTokenAccess(Tweet.GetProfile))
	subscribe("POST", "/api/tweets", middlewares.ValidateTokenAccess(Tweet.Store))
	subscribe("PUT", "/api/tweets/{id}", middlewares.ValidateTokenAccess(Tweet.Update))
	subscribe("DELETE", "/api/tweets/{id}", middlewares.ValidateTokenAccess(Tweet.Delete))
	/* Routes Followers */
	var Follow controllers.Follow
	subscribe("GET", "/api/followers/tweets", middlewares.ValidateTokenAccess(Tweet.GetFollow))
	subscribe("GET", "/api/followers", middlewares.ValidateTokenAccess(Follow.GetFollowed))
	subscribe("GET", "/api/followers/profile/{id}", middlewares.ValidateTokenAccess(Follow.GetProfile))
	subscribe("GET", "/api/followers/profile/{idProfile}/follow/{idFollow}", middlewares.ValidateTokenAccess(Follow.Exists))
	subscribe("POST", "/api/followers", middlewares.ValidateTokenAccess(Follow.Store))
	subscribe("DELETE", "/api/followers", middlewares.ValidateTokenAccess(Follow.Delete))
	/* Routes Uploads */
	var Upload controllers.Upload
	subscribe("POST", "/api/profiles/{id}/upload/avatar", middlewares.ValidateTokenAccess(Upload.Avatar))
	subscribe("POST", "/api/profiles/{id}/upload/banner", middlewares.ValidateTokenAccess(Upload.Banner))
}
