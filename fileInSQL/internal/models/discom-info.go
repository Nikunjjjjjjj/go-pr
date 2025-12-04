package models

type OfficeData struct {
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	OfficeName        string `gorm:"column:officename" json:"officename"`
	Pincode           string `gorm:"column:pincode" json:"pincode"`
	OfficeType        string `gorm:"column:officetype" json:"officetype"`
	DeliveryStatus    string `gorm:"column:deliverystatus" json:"deliverystatus"`
	DivisionName      string `gorm:"column:divisionname" json:"divisionname"`
	RegionName        string `gorm:"column:regionname" json:"regionname"`
	CircleName        string `gorm:"column:circlename" json:"circlename"`
	Taluk             string `gorm:"column:taluk" json:"taluk"`
	DistrictName      string `gorm:"column:districtname" json:"districtname"`
	StateName         string `gorm:"column:statename" json:"statename"`
	Telephone         string `gorm:"column:telephone" json:"telephone"`
	RelatedSubOffice  string `gorm:"column:relatedsuboffice" json:"relatedsuboffice"`
	RelatedHeadOffice string `gorm:"column:relatedheadoffice" json:"relatedheadoffice"`
	Discom            string `gorm:"column:discom" json:"discom"`
}

func (OfficeData) TableName() string {
	return "office_data"
}
