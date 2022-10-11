package repository

import "get_a_discount/internal/core/business/domain"

type BusinessRepository interface {
	Upsert(business domain.Business) error
}
