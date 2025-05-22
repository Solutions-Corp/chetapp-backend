package repository

import (
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	CreateCompany(company *model.Company) error
	GetCompanyByID(id uuid.UUID) (*model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	UpdateCompany(company *model.Company) error
	DeleteCompany(id uuid.UUID) error
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{
		db: db,
	}
}

func (r *companyRepository) CreateCompany(company *model.Company) error {
	return r.db.Create(company).Error
}

func (r *companyRepository) GetCompanyByID(id uuid.UUID) (*model.Company, error) {
	var company model.Company
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) GetAllCompanies() ([]model.Company, error) {
	var companies []model.Company
	if err := r.db.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *companyRepository) UpdateCompany(company *model.Company) error {
	return r.db.Save(company).Error
}

func (r *companyRepository) DeleteCompany(id uuid.UUID) error {
	return r.db.Delete(&model.Company{}, "id = ?", id).Error
}
