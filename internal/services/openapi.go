package services

import (
	"hgoj/internal/dao"
	"hgoj/internal/models"
)

type Host struct {
	ID             int32   `gorm:"id"    json:"id"`
	CPUType        string  `gorm:"cpu_type" json:"cpu_type"`
	CPUPrice       float64 `gorm:"cpu_price" json:"cpu_price"`
	MonitorType    string  `gorm:"monitor_type" json:"monitor_type"`
	MonitorPrice   float64 `gorm:"monitor_price" json:"monitor_price"`
	MainBoardType  string  `gorm:"main_board_type" json:"main_board_type"`
	MainBoardPrice float64 `gorm:"main_board_price" json:"main_board_price"`
	PowerType      string  `gorm:"power_type" json:"power_type"`

	PowerPrice  float64 `gorm:"power_price" json:"power_price"`
	SSDType     string  `gorm:"ssd_type" json:"ssd_type"`
	SSDPrice    float64 `gorm:"ssd_price" json:"ssd_price"`
	MemoryType  string  `gorm:"memory_type" json:"memory_type"`
	MemoryPrice float64 `gorm:"memory_price" json:"memory_price"`
	CrateType   string  `gorm:"crate_type" json:"crate_type"`
	CratePrice  float64 `gorm:"crate_price" json:"crate_price"`
	TotalPrice  float64 `gorm:"total_price" json:"total_price"`
}

// QueryHostByArea 通过价格区间获取该区间所有host
func (pr *Host) QueryHostByPriceArea(min, max int) (pro []models.Host, err error) {
	err = dao.Orm.Debug().Where("total_price <= ? AND total_price >= ?", max, min).Find(&pro).Error
	return pro, err
}
