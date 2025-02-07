package servicestorage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	models "github.com/ruziba3vich/prosphere/internal/items/models/service"
)

type ServiceStorage struct {
	db           *sql.DB
	redisClient  *redis.Client
	queryBuilder squirrel.StatementBuilderType
}

func NewServiceStorage(db *sql.DB, redisClient *redis.Client) *ServiceStorage {
	return &ServiceStorage{
		db:           db,
		redisClient:  redisClient,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (s *ServiceStorage) CreateService(ctx context.Context, req *models.CreateServiceRequest) (*models.Service, error) {
	serviceId := uuid.New().String()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query, args, err := s.queryBuilder.
		Insert("services").
		Columns("service_id", "service_name", "service_link").
		Values(serviceId, req.ServiceName, req.ServiceLink).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	service := &models.Service{
		ServiceId:   serviceId,
		ServiceName: req.ServiceName,
		ServiceLink: req.ServiceLink,
	}

	if err := s.cacheService(ctx, service); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *ServiceStorage) UpdateService(ctx context.Context, req *models.UpdateServiceRequest) (*models.Service, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query, args, err := s.queryBuilder.
		Update("services").
		Set("service_name", req.ServiceName).
		Set("service_link", req.ServiceLink).
		Where(squirrel.Eq{"service_id": req.ServiceId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	service := &models.Service{
		ServiceId:   req.ServiceId,
		ServiceName: req.ServiceName,
		ServiceLink: req.ServiceLink,
	}

	if err := s.cacheService(ctx, service); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *ServiceStorage) GetServiceById(ctx context.Context, req *models.GetServiceByIdRequest) (*models.Service, error) {
	service, err := s.getServiceFromCache(ctx, req.ServiceId)
	if err == nil {
		return service, nil
	}

	query, args, err := s.queryBuilder.
		Select("service_id", "service_name", "service_link").
		From("services").
		Where(squirrel.Eq{"service_id": req.ServiceId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	service = &models.Service{}
	if err := row.Scan(&service.ServiceId, &service.ServiceName, &service.ServiceLink); err != nil {
		return nil, err
	}

	if err := s.cacheService(ctx, service); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *ServiceStorage) GetAllServices(ctx context.Context, req *models.GetAllServicesRequest) (*models.GetAllServicesResponse, error) {

	query, args, err := s.queryBuilder.
		Select("service_id", "service_name", "service_link").
		From("services").
		Limit(uint64(req.Limit)).
		Offset(uint64((req.Page - 1) * req.Limit)).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []*models.Service
	for rows.Next() {
		service := &models.Service{}
		if err := rows.Scan(&service.ServiceId, &service.ServiceName, &service.ServiceLink); err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return &models.GetAllServicesResponse{Response: services}, nil
}

func (s *ServiceStorage) DeleteService(ctx context.Context, req *models.DeleteServiceRequest) (*models.Service, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	service, err := s.GetServiceById(ctx, &models.GetServiceByIdRequest{ServiceId: req.ServiceId})
	if err != nil {
		return nil, err
	}

	query, args, err := s.queryBuilder.
		Delete("services").
		Where(squirrel.Eq{"service_id": req.ServiceId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if err := s.redisClient.Del(ctx, req.ServiceId).Err(); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *ServiceStorage) cacheService(ctx context.Context, service *models.Service) error {
	data, err := json.Marshal(service)
	if err != nil {
		fmt.Println("Error marshaling service:", err)
		return err
	}
	return s.redisClient.Set(ctx, service.ServiceId, data, time.Hour*24).Err()
}

func (s *ServiceStorage) getServiceFromCache(ctx context.Context, serviceId string) (*models.Service, error) {
	data, err := s.redisClient.Get(ctx, serviceId).Result()
	if err != nil {
		return nil, err
	}

	var service models.Service
	if err := json.Unmarshal([]byte(data), &service); err != nil {
		return nil, err
	}

	return &service, nil
}
