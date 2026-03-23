package model

import (

	"gorm.io/gorm"
)

//gorm will createb id,createdat,deletedat,updatedat
type Shortener_Model struct{
 gorm.Model
 Actual_url string
 Short_url string
}


