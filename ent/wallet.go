// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/wallet"
	"github.com/flexprice/flexprice/internal/types"
	"github.com/shopspring/decimal"
)

// Wallet is the model entity for the Wallet schema.
type Wallet struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// EnvironmentID holds the value of the "environment_id" field.
	EnvironmentID string `json:"environment_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID string `json:"customer_id,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance decimal.Decimal `json:"balance,omitempty"`
	// CreditBalance holds the value of the "credit_balance" field.
	CreditBalance decimal.Decimal `json:"credit_balance,omitempty"`
	// WalletStatus holds the value of the "wallet_status" field.
	WalletStatus string `json:"wallet_status,omitempty"`
	// AutoTopupTrigger holds the value of the "auto_topup_trigger" field.
	AutoTopupTrigger *string `json:"auto_topup_trigger,omitempty"`
	// AutoTopupMinBalance holds the value of the "auto_topup_min_balance" field.
	AutoTopupMinBalance *decimal.Decimal `json:"auto_topup_min_balance,omitempty"`
	// AutoTopupAmount holds the value of the "auto_topup_amount" field.
	AutoTopupAmount *decimal.Decimal `json:"auto_topup_amount,omitempty"`
	// WalletType holds the value of the "wallet_type" field.
	WalletType string `json:"wallet_type,omitempty"`
	// ConversionRate holds the value of the "conversion_rate" field.
	ConversionRate decimal.Decimal `json:"conversion_rate,omitempty"`
	// Config holds the value of the "config" field.
	Config       types.WalletConfig `json:"config,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wallet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wallet.FieldAutoTopupMinBalance, wallet.FieldAutoTopupAmount:
			values[i] = &sql.NullScanner{S: new(decimal.Decimal)}
		case wallet.FieldMetadata, wallet.FieldConfig:
			values[i] = new([]byte)
		case wallet.FieldBalance, wallet.FieldCreditBalance, wallet.FieldConversionRate:
			values[i] = new(decimal.Decimal)
		case wallet.FieldID, wallet.FieldTenantID, wallet.FieldStatus, wallet.FieldCreatedBy, wallet.FieldUpdatedBy, wallet.FieldEnvironmentID, wallet.FieldName, wallet.FieldCustomerID, wallet.FieldCurrency, wallet.FieldDescription, wallet.FieldWalletStatus, wallet.FieldAutoTopupTrigger, wallet.FieldWalletType:
			values[i] = new(sql.NullString)
		case wallet.FieldCreatedAt, wallet.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wallet fields.
func (w *Wallet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				w.ID = value.String
			}
		case wallet.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				w.TenantID = value.String
			}
		case wallet.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				w.Status = value.String
			}
		case wallet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case wallet.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		case wallet.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				w.CreatedBy = value.String
			}
		case wallet.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				w.UpdatedBy = value.String
			}
		case wallet.FieldEnvironmentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value.Valid {
				w.EnvironmentID = value.String
			}
		case wallet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case wallet.FieldCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				w.CustomerID = value.String
			}
		case wallet.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				w.Currency = value.String
			}
		case wallet.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case wallet.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case wallet.FieldBalance:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value != nil {
				w.Balance = *value
			}
		case wallet.FieldCreditBalance:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field credit_balance", values[i])
			} else if value != nil {
				w.CreditBalance = *value
			}
		case wallet.FieldWalletStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_status", values[i])
			} else if value.Valid {
				w.WalletStatus = value.String
			}
		case wallet.FieldAutoTopupTrigger:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auto_topup_trigger", values[i])
			} else if value.Valid {
				w.AutoTopupTrigger = new(string)
				*w.AutoTopupTrigger = value.String
			}
		case wallet.FieldAutoTopupMinBalance:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field auto_topup_min_balance", values[i])
			} else if value.Valid {
				w.AutoTopupMinBalance = new(decimal.Decimal)
				*w.AutoTopupMinBalance = *value.S.(*decimal.Decimal)
			}
		case wallet.FieldAutoTopupAmount:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field auto_topup_amount", values[i])
			} else if value.Valid {
				w.AutoTopupAmount = new(decimal.Decimal)
				*w.AutoTopupAmount = *value.S.(*decimal.Decimal)
			}
		case wallet.FieldWalletType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_type", values[i])
			} else if value.Valid {
				w.WalletType = value.String
			}
		case wallet.FieldConversionRate:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field conversion_rate", values[i])
			} else if value != nil {
				w.ConversionRate = *value
			}
		case wallet.FieldConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.Config); err != nil {
					return fmt.Errorf("unmarshal field config: %w", err)
				}
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Wallet.
// This includes values selected through modifiers, order, etc.
func (w *Wallet) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// Update returns a builder for updating this Wallet.
// Note that you need to call Wallet.Unwrap() before calling this method if this Wallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wallet) Update() *WalletUpdateOne {
	return NewWalletClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Wallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wallet) Unwrap() *Wallet {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wallet is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wallet) String() string {
	var builder strings.Builder
	builder.WriteString("Wallet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(w.TenantID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(w.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(w.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(w.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(w.EnvironmentID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(w.Name)
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(w.CustomerID)
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(w.Currency)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(w.Description)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", w.Metadata))
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", w.Balance))
	builder.WriteString(", ")
	builder.WriteString("credit_balance=")
	builder.WriteString(fmt.Sprintf("%v", w.CreditBalance))
	builder.WriteString(", ")
	builder.WriteString("wallet_status=")
	builder.WriteString(w.WalletStatus)
	builder.WriteString(", ")
	if v := w.AutoTopupTrigger; v != nil {
		builder.WriteString("auto_topup_trigger=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := w.AutoTopupMinBalance; v != nil {
		builder.WriteString("auto_topup_min_balance=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := w.AutoTopupAmount; v != nil {
		builder.WriteString("auto_topup_amount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("wallet_type=")
	builder.WriteString(w.WalletType)
	builder.WriteString(", ")
	builder.WriteString("conversion_rate=")
	builder.WriteString(fmt.Sprintf("%v", w.ConversionRate))
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(fmt.Sprintf("%v", w.Config))
	builder.WriteByte(')')
	return builder.String()
}

// Wallets is a parsable slice of Wallet.
type Wallets []*Wallet
