package surveyInformation

import (
	// "fmt"
	"net/http"
	"survayData/api/model"

	// "fmt"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
)

func Init(o *echo.Group) {
	// o.POST("/getStudentData", getStudentDataRoute)
	o.POST("/setStudentData", setStudentDataRoute)
	o.POST("/getLogin", loginRoute)
}

// func getStudentDataRoute(c echo.Context) error{
// 	logginghelper.LogDebug("IN : getStudentDataRoute")

// 	var studentData model.StudentDataStructure
// 	bindErr := c.Bind(&studentData)
// 	if bindErr != nil {
// 		logginghelper.LogInfo("Error While Binding Data", bindErr)
// 		return c.JSON(http.StatusExpectationFailed, "MAHAVAN_ERRORCODE_ERROR_WHILE_BINDING_DATA")
// 	}

// 	value, err := getStudentDataService()
// 	if err != nil {
// 		logginghelper.LogInfo("Error While Binding Data", bindErr)
// 	}
// 	defer logginghelper.LogDebug("OUT : getStudentDataRoute")
// 	return c.JSON(http.StatusOK, value)
// }
func setStudentDataRoute(c echo.Context) error {
	logginghelper.LogDebug("IN : setStudentDataRoute")

	var studentData model.StudentDataStructure
	// fmt.Println(&StudentData )
	bindErr := c.Bind(&studentData)
	if bindErr != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
		return c.JSON(http.StatusExpectationFailed, "MAHAVAN_ERRORCODE_ERROR_WHILE_BINDING_DATA")
	}
	value, err := setStudentDataService(studentData)
	if err != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
	}
	defer logginghelper.LogDebug("OUT :setStudentDatakRoute")
	return c.JSON(http.StatusOK, value)
}

func loginRoute(c echo.Context) error {
	logginghelper.LogDebug("IN : loginRoute")

	var loginStructure model.LoginStructure
	bindErr := c.Bind(&loginStructure)
	if bindErr != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
		return c.JSON(http.StatusExpectationFailed, "MAHAVAN_ERRORCODE_ERROR_WHILE_BINDING_DATA")
	}

	value, err := loginService(loginStructure)
	if err != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
	}

	defer logginghelper.LogDebug("OUT : loginRoute")
	return c.JSON(http.StatusOK, value)
}
