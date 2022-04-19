package surveyInformation

import (
	"survayData/api/model"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

// func getStudentDataService()([]model.StudentDataStructure,error){
// 	logginghelper.LogDebug("IN : getStudentDataService")

// 	value,err := setStudentDataDAO()
// 	if err != nil {
// 		return false, err
// 	}

// 	logginghelper.LogDebug("OUT : getStudentDataService")
// 	return value, err
// }
func setStudentDataService(studentData model.StudentDataStructure) (bool, error) {
	logginghelper.LogDebug("IN : setStudentDataService")

	value, err := setStudentDataDAO(studentData)
	if err != nil {
		return false, err
	}
	logginghelper.LogDebug("OUT : setStudentDataService")
	return value, err
}

func loginService(loginData model.LoginStructure) (bool, error) {
	logginghelper.LogDebug("IN : loginService")

	value, err := loginDAO(loginData)
	if err != nil {
		return false, err
	}

	logginghelper.LogDebug("OUT : loginService")
	return value, err
}
