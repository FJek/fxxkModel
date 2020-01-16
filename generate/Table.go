package generate

type Table struct {
	Name    string `gorm:"column:name"`
	Comment string `gorm:"column:comment"`
}
