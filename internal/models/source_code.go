package models

type SourceCode struct {
	SolutionId int32  `gorm:"solution_id" json:"solution_id"`
	Source     string `gorm:"source"`
}

func (SourceCode) TableName() string {

	return "source_code"
}
