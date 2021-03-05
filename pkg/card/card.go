package card

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

type Card struct {
	Id int64
	Issuer string
	Number string
	Currency string
	Virtual  bool
}

type Service struct {
	mu sync.RWMutex
	cards []*Card
}

func NewService() *Service {
	return &Service{mu: sync.RWMutex{}, cards: []*Card{}}
}

func (s *Service) All() []*Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cards
}

func (s *Service) StoreCard(id int64,issuer string, currency string, virtual bool) error {
	err := s.checkUserId(id)
	if virtual && err != nil {
		log.Println(err)
		return errors.New("user dont have base card")
	}

	cardNumber := strconv.Itoa(rand.Intn(999)) + strconv.Itoa(int(id))

	card := &Card{
		Id: id,
		Issuer: issuer,
		Number: cardNumber,
		Currency: currency,
		Virtual: virtual,
	}
	s.cards = append(s.cards, card)
	return nil
}

func (s *Service) checkUserId(userId int64) error {
	for _, c := range s.cards {
		log.Println(c)
		if c.Id == userId {
			return nil
		}
	}
	return errors.New("user not found")
}

func (s *Service) GetCardsByUserId(userId int64) []*Card {
	userCards := make([]*Card, 0)
	for _, c := range s.cards {
		if c.Id == userId {
			userCards = append(userCards, c)
		}
	}

	return userCards
}