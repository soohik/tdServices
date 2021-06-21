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
		userRoutes.POST("/Invatefriends", ensureNotLoggedIn(), InviteFriends)
		userRoutes.POST("/getallgroups", ensureNotLoggedIn(), Getaddgroups)         //获取自己加入的组
		userRoutes.POST("/getallchats", ensureNotLoggedIn(), GetallChats)           //获取自己加入的组
		userRoutes.POST("/getmegroups", ensureNotLoggedIn(), Getmegroups)           //获取自己创建的组
		userRoutes.POST("/createsuperchats", ensureNotLoggedIn(), Createsupergroup) //创建组
		userRoutes.POST("/sendmessage", ensureNotLoggedIn(), Sendmessage)
		userRoutes.POST("/addcontact", ensureNotLoggedIn(), AddContacts) //联系人
		userRoutes.POST("/getmecontacts", ensureNotLoggedIn(), GetmeContents)
		userRoutes.POST("/savegroupcontact", ensureNotLoggedIn(), SavegroupContents)
		userRoutes.POST("/savechatcontacts", ensureNotLoggedIn(), Savechatcontacts) //保存聊天记录人
		userRoutes.POST("/addtask", ensureNotLoggedIn(), Addtask)

	}

}
