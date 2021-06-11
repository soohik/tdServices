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
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.POST("/preregister", ensureNotLoggedIn(), preregister)
		userRoutes.POST("/joinlinkurl", ensureNotLoggedIn(), JoinChatByInviteLink)
		userRoutes.POST("/getallgroups", ensureNotLoggedIn(), Getallgroups)
		userRoutes.POST("/getmegroups", ensureNotLoggedIn(), Getmegroups)
		userRoutes.POST("/invatefriends", ensureNotLoggedIn(), Invategroup) //邀请
		userRoutes.POST("/sendmessage", ensureNotLoggedIn(), Sendmessage)
		userRoutes.POST("/addcontact", ensureNotLoggedIn(), AddContacts) //联系人
		userRoutes.POST("/getmecontacts", ensureNotLoggedIn(), GetmeContents)
		userRoutes.POST("/getgroupcontact", ensureNotLoggedIn(), GetmeContents)

	}

}
