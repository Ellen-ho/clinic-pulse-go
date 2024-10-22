package treatment

import (
	"time"
	"gorm.io/gorm"
)

type MedicineTreatmentEntity struct {
	ID            string     `gorm:"type:uuid;primaryKey"`
	GetMedicineAt *time.Time `gorm:"type:timestamp;nullable"`
}
