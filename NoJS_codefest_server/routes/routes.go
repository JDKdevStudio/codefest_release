package routes

import (
	"NoJS_codefest_server/controllers"
	"NoJS_codefest_server/models"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func DefineRoutes(server *echo.Echo) {
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣤⣶⣾⣿⣿⣷⢰⣆⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⣾⣿⣿⣿⣿⣿⣿⣿⡏⠀⢿⡄⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠀⢸⡇⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⢸⡇⠀
⠀⠀⠀⠀⠀⢀⣠⣾⣿⣿⡿⠿⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠀⠀⠀⣾⠃⠀
⠀⠀⠀⢀⣴⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣿⣿⠃⠀⠀⠀⢠⡟⠀⠀
⠀⠀⢠⣿⣿⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⣿⡿⠃⠀⠀⠀⠀⣼⠃⠀⠀
⠀⢠⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⡟⠁⠀⠀⠀⠀⣼⠇⠀⠀⠀
 ⣾⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣿⣿⡿⠋⠀⠀⠀⠀⢀⣼⠏⠀⠀⠀⠀
⢸⣿⣿⣧⠀⠀⠀⠀⢀⣠⣴⣿⣿⣿⠿⠋⠀⠀⠀⠀⠀⣠⡾⠃⠀⠀⠀⠀⠀
⢸⣿⣿⣿⣷⣶⣶⣿⣿⣿⡿⠟⠋⠁⠀⠀⠀⠀⠀⣠⡾⠋⠀⠀⠀⠀⠀⠀⠀
⠀⠙⠿⠿⠿⠿⠟⠛⠉⠁⠀⠀⠀⠀⠀⠀⢀⣤⡾⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠒⠶⢶⣤⣤⣤⣤⣤⣤⣴⠶⠞⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀

 _______           _______                        ______       _______  _______  _______           _______  _______ 
(  ____ \|\     /|(  ____ \    |\     /||\     /|(  ___ \     (  ____ \(  ____ \(  ____ )|\     /|(  ____ \(  ____ )
| (    \/( \   / )| (    \/    | )   ( || )   ( || (   ) )    | (    \/| (    \/| (    )|| )   ( || (    \/| (    )|
| (_____  \ (_) / | (_____     | (___) || |   | || (__/ /     | (_____ | (__    | (____)|| |   | || (__    | (____)|
(_____  )  \   /  (_____  )    |  ___  || |   | ||  __ (      (_____  )|  __)   |     __)( (   ) )|  __)   |     __)
      ) |   ) (         ) |    | (   ) || |   | || (  \ \           ) || (      | (\ (    \ \_/ / | (      | (\ (   
/\____) |   | |   /\____) |    | )   ( || (___) || )___) )    /\____) || (____/\| ) \ \__  \   /  | (____/\| ) \ \__
\_______)   \_/   \_______)    |/     \|(_______)|/ \___/     \_______)(_______/|/   \__/   \_/   (_______/|/   \__/
                                                                                                                    
$SYS-Hub Server $v1.0
Powered by golang 1.21.3
		`)
	})
	//Get path of server instance
	executablePath, err := os.Executable()
	if err != nil {
		log.Panic(err)
	}
	//middleware for jwt endpoints
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{NewClaimsFunc: func(c echo.Context) jwt.Claims { return new(models.User) }, SigningKey: []byte(os.Getenv("JWT_SECRET"))})
	server.GET("/tests/", controllers.TestController, jwtMiddleware)

	server.Static("/assets/", filepath.Join(filepath.Dir(executablePath), "public"))
	server.GET("/users/login/", controllers.UserLoginController)
	server.POST("/users/register/", controllers.UserRegisterController)
	server.POST("/projects/", controllers.ProjectRegisterController, jwtMiddleware)
	server.GET("/projects/", controllers.ProyectListController)
}
