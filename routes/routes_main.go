package routes

import (
	"github.com/Pugpaprika21/go-gin/controller"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		userRouter := apiRouter.Group("/user")
		{
			userRepository := repository.NewUserRepository()
			userController := controller.NewUserController(userRepository)
			userRouter.POST("/create", userController.Create)
			userRouter.GET("/detail/:id", userController.GetUserDetailByID)
			userRouter.PUT("/update/:id", userController.UpdateUser)
			userRouter.DELETE("/delete/:id", userController.DeleteByID)
		}

		masterRouter := apiRouter.Group("/master")
		{
			masterRolesRouter := masterRouter.Group("/roles")
			{
				masterRolesRepository := repository.NewMasterRoleRepository()
				masterRolesController := controller.NewMasterRolesController(masterRolesRepository)
				masterRolesRouter.POST("/create", masterRolesController.Create)
				masterRolesRouter.GET("/show", masterRolesController.FindAll)
				masterRolesRouter.GET("/show/:id", masterRolesController.FindByID)
				masterRolesRouter.PUT("/update/:id", masterRolesController.UpdateByID)
				masterRolesRouter.DELETE("/delete/:id", masterRolesController.DeleteByID)
			}

			masterSettingRolesRouter := masterRouter.Group("/setting_roles")
			{
				masterSettingRolesRepository := repository.NewMasterSettingRolesRepository()
				masterSettingRolesController := controller.NewMasterSettingRolesController(masterSettingRolesRepository)
				masterSettingRolesRouter.POST("/create", masterSettingRolesController.Create)
				masterSettingRolesRouter.POST("/delete", masterSettingRolesController.DeleteRoleByIDS)
			}

			meetingRoomRouter := masterRouter.Group("/meeting_room")
			{
				meetingRoomRepository := repository.NewMasterMeetingRoomRepository()
				meetingRoomController := controller.NewMasterMeetingRoomController(meetingRoomRepository)
				meetingRoomRouter.POST("/create", meetingRoomController.Create)
				meetingRoomRouter.GET("/show", meetingRoomController.FindAll)
				meetingRoomRouter.GET("/show/:id", meetingRoomController.FindByID)
				meetingRoomRouter.PUT("/update/:id", meetingRoomController.UpdateByID)
				meetingRoomRouter.DELETE("/delete/:id", meetingRoomController.DeleteByID)
			}

			meetingRoomBasicEquipmentsRouter := masterRouter.Group("/meeting_room_basic_equipments")
			{
				meetingRoomBasicEquipmentsRepository := repository.NewMasterMeetingRoomBasicEquipmentsRepository()
				meetingRoomBasicEquipmentsController := controller.NewMasterMeetingRoomBasicEquipmentsController(meetingRoomBasicEquipmentsRepository)
				meetingRoomBasicEquipmentsRouter.POST("/create", meetingRoomBasicEquipmentsController.Create)
				meetingRoomBasicEquipmentsRouter.GET("/show", meetingRoomBasicEquipmentsController.FindAll)
				meetingRoomBasicEquipmentsRouter.GET("/show/:id", meetingRoomBasicEquipmentsController.FindByID)
				meetingRoomBasicEquipmentsRouter.PUT("/update/:id", meetingRoomBasicEquipmentsController.UpdateByID)
				meetingRoomBasicEquipmentsRouter.DELETE("/delete/:id", meetingRoomBasicEquipmentsController.DeleteByID)
			}

			meetingRoomDetailsRouter := masterRouter.Group("/meeting_room_details")
			{
				meetingRoomDetailsRepository := repository.NewMasterMeetingRoomDetailsRepository()
				meetingRoomDetailsController := controller.NewMasterMeetingRoomDetails(meetingRoomDetailsRepository)
				meetingRoomDetailsRouter.POST("/create", meetingRoomDetailsController.Create)
				meetingRoomDetailsRouter.GET("/show", meetingRoomDetailsController.FindAll)
				meetingRoomDetailsRouter.GET("/show/:id", meetingRoomDetailsController.FindByID)
				meetingRoomDetailsRouter.PUT("/update/:id", meetingRoomDetailsController.UpdateByID)
				meetingRoomDetailsRouter.DELETE("/delete/:id", meetingRoomDetailsController.DeleteByID)
			}
		}

		processRouter := apiRouter.Group("/process")
		{
			meetingRoomEquipmentsRouter := processRouter.Group("/meeting_room_equipments")
			{
				meetingRoomEquipmentsRepository := repository.NewMeetingRoomEquipmentsRepository()
				meetingRoomEquipmentsController := controller.NewMeetingRoomEquipmentsController(meetingRoomEquipmentsRepository)
				meetingRoomEquipmentsRouter.POST("/create", meetingRoomEquipmentsController.CreateEquipment)
				meetingRoomEquipmentsRouter.GET("/show/:id", meetingRoomEquipmentsController.CreateEquipment)
				meetingRoomEquipmentsRouter.PUT("/update/:id", meetingRoomEquipmentsController.UpdateEquipment)
				meetingRoomEquipmentsRouter.DELETE("/delete/:id", meetingRoomEquipmentsController.DeletetEquipmentByID)
			}

			meetingRoomFormRequestsRouter := processRouter.Group("/meeting_room_form_requests")
			{
				meetingRoomFormRequestsRepository := repository.NewMeetingRoomFormRequestsRepository()
				meetingRoomFormRequestsController := controller.NewMeetingRoomFormRequestsController(meetingRoomFormRequestsRepository)
				meetingRoomFormRequestsRouter.POST("/create", meetingRoomFormRequestsController.CreateMeeting)
				meetingRoomFormRequestsRouter.GET("/show_meeting/:id", meetingRoomFormRequestsController.GetMeetingRoomByID)
				meetingRoomFormRequestsRouter.GET("/show/:userId", meetingRoomFormRequestsController.GetMeetingByUser)
				meetingRoomFormRequestsRouter.PUT("/update/id", meetingRoomFormRequestsController.UpdateByID)
				meetingRoomFormRequestsRouter.DELETE("/delete/:id", meetingRoomFormRequestsController.DeleteMeetingRoom)
			}
		}
	}
}
