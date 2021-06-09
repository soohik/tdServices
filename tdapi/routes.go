// routes.go

package main

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// JoinChatByInviteLink
	// Group user related routes together
	userRoutes := router.Group("/phone")
	{
		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.POST("/preregister", ensureNotLoggedIn(), preregister)
		userRoutes.POST("/joinlinkurl", ensureNotLoggedIn(), JoinChatByInviteLink)
		userRoutes.POST("/getallgroups", ensureNotLoggedIn(), Getallgroups)
		userRoutes.POST("/getmegroups", ensureNotLoggedIn(), Getmegroups)

	}

}
