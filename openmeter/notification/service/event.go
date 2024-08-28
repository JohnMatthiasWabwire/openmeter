package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/openmeterio/openmeter/openmeter/notification"
	"github.com/openmeterio/openmeter/pkg/models"
)

func (s Service) ListEvents(ctx context.Context, params notification.ListEventsInput) (notification.ListEventsResult, error) {
	if err := params.Validate(ctx, s); err != nil {
		return notification.ListEventsResult{}, fmt.Errorf("invalid params: %w", err)
	}

	return s.repo.ListEvents(ctx, params)
}

func (s Service) GetEvent(ctx context.Context, params notification.GetEventInput) (*notification.Event, error) {
	if err := params.Validate(ctx, s); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	return s.repo.GetEvent(ctx, params)
}

func (s Service) CreateEvent(ctx context.Context, params notification.CreateEventInput) (*notification.Event, error) {
	if err := params.Validate(ctx, s); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	logger := s.logger.WithGroup("event").With(
		"operation", "create",
		"namespace", params.Namespace,
	)

	logger.Debug("creating event")

	rule, err := s.repo.GetRule(ctx, notification.GetRuleInput{
		Namespace: params.Namespace,
		ID:        params.RuleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get rule: %w", err)
	}

	if rule.DeletedAt != nil {
		return nil, notification.NotFoundError{
			NamespacedID: models.NamespacedID{
				Namespace: params.Namespace,
				ID:        params.RuleID,
			},
		}
	}

	if rule.Disabled {
		return nil, notification.ValidationError{
			Err: errors.New("failed to send event: rule is disabled"),
		}
	}

	event, err := s.repo.CreateEvent(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create event: %w", err)
	}

	if err = s.eventHandler.Dispatch(event); err != nil {
		return nil, fmt.Errorf("failed to dispatch event: %w", err)
	}

	return event, nil
}

func (s Service) UpdateEventDeliveryStatus(ctx context.Context, params notification.UpdateEventDeliveryStatusInput) (*notification.EventDeliveryStatus, error) {
	if err := params.Validate(ctx, s); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	return s.repo.UpdateEventDeliveryStatus(ctx, params)
}

func (s Service) ListEventsDeliveryStatus(ctx context.Context, params notification.ListEventsDeliveryStatusInput) (notification.ListEventsDeliveryStatusResult, error) {
	if err := params.Validate(ctx, s); err != nil {
		return notification.ListEventsDeliveryStatusResult{}, fmt.Errorf("invalid params: %w", err)
	}

	return s.repo.ListEventsDeliveryStatus(ctx, params)
}

func (s Service) GetEventDeliveryStatus(ctx context.Context, params notification.GetEventDeliveryStatusInput) (*notification.EventDeliveryStatus, error) {
	if err := params.Validate(ctx, s); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	return s.repo.GetEventDeliveryStatus(ctx, params)
}
