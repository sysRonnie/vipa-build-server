package network

import "github.com/labstack/echo/v4"


func ExtractIPandAgent(c echo.Context) (string, string) {
	
	ip := c.RealIP()
	ua := c.Request().UserAgent()
	return ip, ua
}
