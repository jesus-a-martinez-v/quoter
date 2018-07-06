package service

import (
	"quoter/src/api/model/repository"
	"quoter/src/api/model/domain"
	"quoter/src/api/service/dto"
)

func GetQuotes(author string, genre string) []dto.QuoteDto {
	entities := repository.GetQuotes(author, genre)
	return fromEntitiesToDtos(&entities)
}

func GetQuoteById(id int64) *dto.QuoteDto {
	entity := repository.GetQuoteById(id)
	if entity == nil {
		return nil
	}

	quoteDto := fromEntityToDto(*entity)
	return &quoteDto
}

func fromEntitiesToDtos(entities *[]domain.QuoteEntity) []dto.QuoteDto {
	dtos := make([]dto.QuoteDto, len(*entities))

	for index, entity := range *entities {
		dtos[index] = fromEntityToDto(entity)
	}

	return dtos
}
func fromEntityToDto(entity domain.QuoteEntity) dto.QuoteDto {
	return dto.QuoteDto{
		Quote: entity.Quote,
		Author: entity.Author,
		Genre: entity.Genre,
		Id: entity.Id,
	}
}