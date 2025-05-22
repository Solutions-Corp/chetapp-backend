package service

import (
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/repository"
	"github.com/google/uuid"
)

type CompanyService interface {
	CreateCompany(company *model.Company) error
	GetCompanyByID(id uuid.UUID) (*model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	UpdateCompany(company *model.Company) error
	DeleteCompany(id uuid.UUID) error
}

type companyService struct {
	companyRepo repository.CompanyRepository
}

func NewCompanyService(companyRepo repository.CompanyRepository) CompanyService {
	return &companyService{
		companyRepo: companyRepo,
	}
}

func (s *companyService) CreateCompany(company *model.Company) error {
	return s.companyRepo.CreateCompany(company)
}

func (s *companyService) GetCompanyByID(id uuid.UUID) (*model.Company, error) {
	return s.companyRepo.GetCompanyByID(id)
}

func (s *companyService) GetAllCompanies() ([]model.Company, error) {
	return s.companyRepo.GetAllCompanies()
}

func (s *companyService) UpdateCompany(company *model.Company) error {
	return s.companyRepo.UpdateCompany(company)
}

func (s *companyService) DeleteCompany(id uuid.UUID) error {
	return s.companyRepo.DeleteCompany(id)
}
