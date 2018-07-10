package service

import (
	"quoter/src/api/model/domain"
	"quoter/src/api/model/repository"
	"quoter/src/api/service/dto"
	"strconv"
)

func GetQuotes(author string, genre string, random string) []dto.QuoteDto {
	var entities []domain.QuoteEntity

	if len(random) > 0 {
		fetchRandom, _ := strconv.ParseBool(random)

		if fetchRandom {
			entities = repository.GetRandomQuote(author, genre)
			return fromEntitiesToDtos(&entities)
		}
	}

	entities = repository.GetQuotes(author, genre)
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

func SaveQuote(requestDto dto.QuoteDto) (int64, error) {
	entity := domain.QuoteEntity{
		Quote:  requestDto.Quote,
		Genre:  requestDto.Genre,
		Author: requestDto.Author,
	}

	return repository.InsertQuote(&entity)
}

func DeleteQuoteById(id int64) {
	repository.DeleteQuoteById(id)
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
		Quote:  entity.Quote,
		Author: entity.Author,
		Genre:  entity.Genre,
		Id:     entity.Id,
	}
}
