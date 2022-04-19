package model

type StudentDataStructure struct {
	EMAIL    string `json:"email" bson:"email"`
	USERNAME string `json:"username" bson:"username"`
	PASSWORD string `json:"password" bson:"password"`
}

type LoginStructure struct {
	USERNAME string `json:"username" bson:"username"`
	PASSWORD string `json:"password" bson:"password"`
}

// type DataData struct {
// 	USERNAME string `json:"username" bson:"username"`
// 	PASSWORD string `json:"password" bson:"password"`
// }
// Username        string      `json:"username,omitempty" bson:"USERNAME" validate:"required"`
