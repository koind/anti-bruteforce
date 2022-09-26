package list

import (
	"github.com/go-redis/redis"
	"github.com/koind/anti-bruteforce/internal/service"
	"net"
)

const (
	white = "white"
	black = "black"
)

type Storage struct {
	r *redis.Client
}

type Config interface {
	GetRedisAddr() string
	GetPassword() string
	GetDBNumber() int
}

func NewStorage(config Config) *Storage {
	return &Storage{
		r: redis.NewClient(&redis.Options{
			Addr:     config.GetRedisAddr(),
			Password: config.GetPassword(),
			DB:       config.GetDBNumber(),
		}),
	}
}

func (s *Storage) Check(ip net.IP) (service.IPStatus, error) {
	whiteList, err := s.r.SMembers(white).Result()

	if err != nil {
		return service.Undefined, err
	}
	for _, whiteNet := range whiteList {
		_, ipNet, err := net.ParseCIDR(whiteNet)
		if err != nil {
			return service.Undefined, err
		}

		if ipNet.Contains(ip) {
			return service.Allowed, nil
		}
	}

	blackList, err := s.r.SMembers(black).Result()

	if err != nil {
		return service.Undefined, err
	}
	for _, blackNet := range blackList {
		_, ipNet, err := net.ParseCIDR(blackNet)
		if err != nil {
			return service.Undefined, err
		}

		if ipNet.Contains(ip) {
			return service.Rejected, nil
		}
	}

	return service.Undefined, nil
}

func (s *Storage) ShowList() map[string][]string {
	list := make(map[string][]string)

	whiteList, _ := s.r.SMembersMap(white).Result()
	blackList, _ := s.r.SMembersMap(black).Result()

	for whiteNet := range whiteList {
		list[white] = append(list[white], whiteNet)
	}
	for blackNet := range blackList {
		list[black] = append(list[black], blackNet)
	}

	return list
}

func (s *Storage) AddWhiteNet(net string) error {
	if err := s.r.SAdd(white, net).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddBlackNet(net string) error {
	if err := s.r.SAdd(black, net).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) RemoveWhiteNet(net string) error {
	if err := s.r.SRem(white, net).Err(); err != nil {
		return err
	}

	return nil
}
func (s *Storage) RemoveBlackNet(net string) error {
	if err := s.r.SRem(black, net).Err(); err != nil {
		return err
	}

	return nil
}
