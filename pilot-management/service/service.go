package service

import (
	"../domain/entity"
	"../repository"
	"../repository/impl/postgresql"
)

type ServiceImpl struct {
	pilotRepo repository.PilotRepo
}

func MakeServiceImpl() ServiceImpl {
	pilotRepo := postgresql.MakePostgresPilotRepo()
	return ServiceImpl{pilotRepo: &pilotRepo}
}

func (s ServiceImpl) ListPilots() ([]entity.Pilot, error) {
	return s.pilotRepo.ListPilots()
}

//
//func (s ServiceImpl) GetPilot(id string) (Pilot, error) {
//	return s.pilotRepo.GetPilot(id)
//}
//
//func (s ServiceImpl) CreatePilot(params CreatePilotParams) (Pilot, error) {
//	return s.pilotRepo.CreatePilot(params)
//}
//
//func (s ServiceImpl) UpdatePilot(params UpdatePilotParams) (Pilot, error) {
//	return s.pilotRepo.UpdatePilot(params)
//}
//
//func (s ServiceImpl) DeletePilot(id string) error {
//	return s.pilotRepo.DeletePilot(id)
//}
