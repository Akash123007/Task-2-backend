package surveyInformation

import (
	"fmt"
	"survayData/api/helpers"
	"survayData/api/model"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

// func getStudentDataDAO()([]model.StudentDataStructure, error){
// 	logginghelper.LogDebug("IN : getStudentDataDAO")
// 	var studentData []model.StudentDataStructure
// 	db, dberr := helpers.GetSQLConnection()
// 	if dberr != nil {
// 		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
// 		return studentData, dberr
// 	}
// 	session := db.NewSession(nil)
// 	rows, readErr := session.Query("SELECT NAME,SIRNAME FROM student_info ;")
// 	if readErr != nil {
// 		// logginghelper.LogError(readErr)
// 		fmt.Println("Error while executing query", readErr)
// 		return studentData, dberr
// 	}
// 	for rows.Next() {
// 		var NAME 					string
// 		var SIRNAME 				string
// 		scanErr := rows.Scan ( &NAME, &SIRNAME)
// 		if scanErr != nil {
// 			fmt.Println("Error in Scanning the rows", scanErr)
// 			// logginghelper.LogError(scanErr)
// 			return nil, scanErr
// 		}
// 		studentData = append(studentData, model.StudentDataStructure{
// 			NAME: 					NAME,
// 			SIRNAME: 				SIRNAME,
// 		})
// 		fmt.Println("studentData: ",studentData)
// 	}
// 	logginghelper.LogDebug("OUT : getStudentDataDAO")
// 	return studentData, nil
// }
func setStudentDataDAO(studentData model.StudentDataStructure) (bool, error) {
	logginghelper.LogDebug("IN :setStudentDataDAO :")
	db, dberr := helpers.GetSQLConnection()
	if dberr != nil {

		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
		return false, dberr
	}
	fmt.Println(studentData)
	session := db.NewSession(nil)
	if studentData.EMAIL == "" || studentData.USERNAME == "" || studentData.PASSWORD == "" {
		return false, nil
	}
	result, err := session.InsertBySql(
		`INSERT INTO registration(EMAIL,USERNAME,PASSWORD)
	  VALUES (?,?,?)`,
		studentData.EMAIL,
		studentData.USERNAME,
		studentData.PASSWORD,
	).Exec()
	if err != nil {
		// logginghelper.LogError(readErr)
		fmt.Println("Error while executing query", err)
		return false, dberr
	}
	fmt.Println(result)
	logginghelper.LogDebug("OUT :setStudentDataDAO")
	return true, nil
}

func loginDAO(loginData model.LoginStructure) (bool, error) {
	logginghelper.LogDebug("IN : loginDAO")
	db, dberr := helpers.GetSQLConnection()
	if dberr != nil {
		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
		return false, dberr
	}
	session := db.NewSession(nil)
	var result model.LoginStructure
	fmt.Println("username", loginData.USERNAME)
	if loginData.USERNAME == " " {
		fmt.Println("Userid and password required")
		return false, nil
	}
	if loginData.PASSWORD == "" {
		return false, nil
	}
	rows, readErr := session.SelectBySql(`
	SELECT USERNAME,PASSWORD
	FROM registration 
	WHERE USERNAME =? AND PASSWORD=?`,
		loginData.USERNAME,
		loginData.PASSWORD,
	).Load(&result)
	if rows == 0 {
		return false, nil
	}
	if readErr != nil {
		logginghelper.LogError("ERROR in EXCUTING QUERY", dberr)
		return false, dberr
	}
	fmt.Println("rowsdata", rows)

	fmt.Printf("%T, %v\n", result, result)
	logginghelper.LogDebug("OUT : loginDAO")
	return true, nil
}

// delete
// func deleteUserDataDAO(userData model.RegistrationDataStructure) (bool, error) {
// 	logginghelper.LogDebug("IN : deleteUserDataDAO")

// 	session, dberr := helpers.GetSQLConnection()
// 	if dberr != nil {
// 		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
// 		return false, dberr
// 	}

// 	rows, readErr := session.Query("UPDATE registration SET EMAIL=? , USERNAME = ? , PASSWORD=? WHERE Id = ?", userData.EMAIL, userData.USERNAME, userData.PASSWORD, strconv.Itoa(userData.Id))
// 	if readErr != nil {
// 		// logginghelper.LogError(readErr)
// 		fmt.Println("Error while executing query", readErr)
// 		return false, dberr
// 	}
// 	for rows.Next() {
// 		var id int
// 		var EMAIL string
// 		var USERNAME string
// 		var PASSWORD string

// 		scanErr := rows.Scan(&id, &EMAIL, &USERNAME, &PASSWORD)
// 		if scanErr != nil {
// 			fmt.Println("Error in Scanning the rows", scanErr)
// 			// logginghelper.LogError(scanErr)
// 			return false, scanErr
// 		}

// 		userData = model.RegistrationDataStructure{
// 			Id:        id,
// 			EMAIL:     EMAIL,
// 			USERNAME:  USERNAME,
// 			PASSWORD: PASSWORD,
// 		}
// 		fmt.Println("userData: ", userData)
// 	}
// 	logginghelper.LogDebug("OUT : deleteUserDataDAO")
// 	return true, nil
// }
