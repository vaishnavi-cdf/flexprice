package types

import (
	"time"

	ierr "github.com/flexprice/flexprice/internal/errors"
	"github.com/samber/lo"
)

// SubscriptionStatus is the status of a subscription
// For now taking inspiration from Stripe's subscription statuses
// https://stripe.com/docs/api/subscriptions/object#subscription_object-status
type SubscriptionStatus string

const (
	SubscriptionStatusActive            SubscriptionStatus = "active"
	SubscriptionStatusPaused            SubscriptionStatus = "paused"
	SubscriptionStatusCancelled         SubscriptionStatus = "cancelled"
	SubscriptionStatusIncomplete        SubscriptionStatus = "incomplete"
	SubscriptionStatusIncompleteExpired SubscriptionStatus = "incomplete_expired"
	SubscriptionStatusPastDue           SubscriptionStatus = "past_due"
	SubscriptionStatusTrialing          SubscriptionStatus = "trialing"
	SubscriptionStatusUnpaid            SubscriptionStatus = "unpaid"
)

func (s SubscriptionStatus) String() string {
	return string(s)
}

func (s SubscriptionStatus) Validate() error {
	allowed := []SubscriptionStatus{
		SubscriptionStatusActive,
		SubscriptionStatusPaused,
		SubscriptionStatusCancelled,
		SubscriptionStatusIncomplete,
		SubscriptionStatusIncompleteExpired,
		SubscriptionStatusPastDue,
		SubscriptionStatusTrialing,
		SubscriptionStatusUnpaid,
	}
	if !lo.Contains(allowed, s) {
		return ierr.NewError("invalid subscription status").
			WithHint("Invalid subscription status").
			WithReportableDetails(map[string]any{
				"status":         s,
				"allowed_status": allowed,
			}).
			Mark(ierr.ErrValidation)
	}
	return nil
}

// PauseStatus represents the pause state of a subscription
type PauseStatus string

const (
	// PauseStatusNone indicates the subscription is not paused
	PauseStatusNone PauseStatus = "none"

	// PauseStatusActive indicates the subscription is currently paused
	PauseStatusActive PauseStatus = "active"

	// PauseStatusScheduled indicates the subscription is scheduled to be paused
	PauseStatusScheduled PauseStatus = "scheduled"

	// PauseStatusCompleted indicates the pause has been completed (subscription resumed)
	PauseStatusCompleted PauseStatus = "completed"

	// PauseStatusCancelled indicates the pause was cancelled
	PauseStatusCancelled PauseStatus = "cancelled"
)

func (s PauseStatus) String() string {
	return string(s)
}

func (s PauseStatus) Validate() error {
	allowed := []PauseStatus{
		PauseStatusNone,
		PauseStatusActive,
		PauseStatusScheduled,
		PauseStatusCompleted,
		PauseStatusCancelled,
	}

	if !lo.Contains(allowed, s) {
		return ierr.NewError("invalid pause status").
			WithHint("Invalid pause status").
			WithReportableDetails(map[string]any{
				"status":         s,
				"allowed_status": allowed,
			}).
			Mark(ierr.ErrValidation)
	}

	return nil
}

// SubscriptionFilter represents filters for subscription queries
type SubscriptionFilter struct {
	*QueryFilter
	*TimeRangeFilter

	// CustomerID filters by customer ID
	CustomerID string `json:"customer_id,omitempty" form:"customer_id"`
	// PlanID filters by plan ID
	PlanID string `json:"plan_id,omitempty" form:"plan_id"`
	// SubscriptionStatus filters by subscription status
	SubscriptionStatus []SubscriptionStatus `json:"subscription_status,omitempty" form:"subscription_status"`
	// BillingCadence filters by billing cadence
	BillingCadence []BillingCadence `json:"billing_cadence,omitempty" form:"billing_cadence"`
	// BillingPeriod filters by billing period
	BillingPeriod []BillingPeriod `json:"billing_period,omitempty" form:"billing_period"`
	// IncludeCanceled includes canceled subscriptions if true
	IncludeCanceled bool `json:"include_canceled,omitempty" form:"include_canceled"`
	// ActiveAt filters subscriptions that are active at the given time
	ActiveAt *time.Time `json:"active_at,omitempty" form:"active_at"`

	// WithLineItems includes line items in the response
	WithLineItems bool `json:"with_line_items,omitempty" form:"with_line_items"`
}

// NewSubscriptionFilter creates a new SubscriptionFilter with default values
func NewSubscriptionFilter() *SubscriptionFilter {
	return &SubscriptionFilter{
		QueryFilter: NewDefaultQueryFilter(),
	}
}

// NewNoLimitSubscriptionFilter creates a new SubscriptionFilter with no pagination limits
func NewNoLimitSubscriptionFilter() *SubscriptionFilter {
	return &SubscriptionFilter{
		QueryFilter: NewNoLimitQueryFilter(),
	}
}

// Validate validates the subscription filter
func (f SubscriptionFilter) Validate() error {
	if f.QueryFilter != nil {
		if err := f.QueryFilter.Validate(); err != nil {
			return err
		}
	}

	if f.TimeRangeFilter != nil {
		if err := f.TimeRangeFilter.Validate(); err != nil {
			return err
		}
	}

	// Validate subscription status values
	for _, status := range f.SubscriptionStatus {
		if err := status.Validate(); err != nil {
			return err
		}
	}

	// Validate billing cadence values
	for _, cadence := range f.BillingCadence {
		if err := cadence.Validate(); err != nil {
			return err
		}
	}

	// Validate billing period values
	for _, period := range f.BillingPeriod {
		if err := period.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetLimit implements BaseFilter interface
func (f *SubscriptionFilter) GetLimit() int {
	if f.QueryFilter == nil {
		return NewDefaultQueryFilter().GetLimit()
	}
	return f.QueryFilter.GetLimit()
}

// GetOffset implements BaseFilter interface
func (f *SubscriptionFilter) GetOffset() int {
	if f.QueryFilter == nil {
		return NewDefaultQueryFilter().GetOffset()
	}
	return f.QueryFilter.GetOffset()
}

// GetSort implements BaseFilter interface
func (f *SubscriptionFilter) GetSort() string {
	if f.QueryFilter == nil {
		return NewDefaultQueryFilter().GetSort()
	}
	return f.QueryFilter.GetSort()
}

// GetOrder implements BaseFilter interface
func (f *SubscriptionFilter) GetOrder() string {
	if f.QueryFilter == nil {
		return NewDefaultQueryFilter().GetOrder()
	}
	return f.QueryFilter.GetOrder()
}

// GetStatus implements BaseFilter interface
func (f *SubscriptionFilter) GetStatus() string {
	if f.QueryFilter == nil {
		return NewDefaultQueryFilter().GetStatus()
	}
	return f.QueryFilter.GetStatus()
}

// GetExpand implements BaseFilter interface
func (f *SubscriptionFilter) GetExpand() Expand {
	if f.QueryFilter == nil {
		return NewDefaultQueryFilter().GetExpand()
	}
	return f.QueryFilter.GetExpand()
}

func (f *SubscriptionFilter) IsUnlimited() bool {
	if f.QueryFilter == nil {
		return NewDefaultQueryFilter().IsUnlimited()
	}
	return f.QueryFilter.IsUnlimited()
}
