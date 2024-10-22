package repositories

import( 
  "clinic-pulse-go/domain"
	)

type DoctorRepository interface {
    FindByUserId(userId string) (*doctor.Doctor, error)

    FindByAll() ([]DoctorBasicInfo, error)

    FindById(doctorId string) (*doctor.Doctor, error)

    Save(doctor *doctor.Doctor) error
}

type DoctorBasicInfo struct {
    ID       string
    FullName string
}
