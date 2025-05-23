// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flexprice/flexprice/ent/customer"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
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
	// ExternalID holds the value of the "external_id" field.
	ExternalID string `json:"external_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// AddressLine1 holds the value of the "address_line1" field.
	AddressLine1 string `json:"address_line1,omitempty"`
	// AddressLine2 holds the value of the "address_line2" field.
	AddressLine2 string `json:"address_line2,omitempty"`
	// AddressCity holds the value of the "address_city" field.
	AddressCity string `json:"address_city,omitempty"`
	// AddressState holds the value of the "address_state" field.
	AddressState string `json:"address_state,omitempty"`
	// AddressPostalCode holds the value of the "address_postal_code" field.
	AddressPostalCode string `json:"address_postal_code,omitempty"`
	// AddressCountry holds the value of the "address_country" field.
	AddressCountry string `json:"address_country,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata     map[string]string `json:"metadata,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldMetadata:
			values[i] = new([]byte)
		case customer.FieldID, customer.FieldTenantID, customer.FieldStatus, customer.FieldCreatedBy, customer.FieldUpdatedBy, customer.FieldEnvironmentID, customer.FieldExternalID, customer.FieldName, customer.FieldEmail, customer.FieldAddressLine1, customer.FieldAddressLine2, customer.FieldAddressCity, customer.FieldAddressState, customer.FieldAddressPostalCode, customer.FieldAddressCountry:
			values[i] = new(sql.NullString)
		case customer.FieldCreatedAt, customer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case customer.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				c.TenantID = value.String
			}
		case customer.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = value.String
			}
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case customer.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = value.String
			}
		case customer.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = value.String
			}
		case customer.FieldEnvironmentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value.Valid {
				c.EnvironmentID = value.String
			}
		case customer.FieldExternalID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				c.ExternalID = value.String
			}
		case customer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case customer.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case customer.FieldAddressLine1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line1", values[i])
			} else if value.Valid {
				c.AddressLine1 = value.String
			}
		case customer.FieldAddressLine2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line2", values[i])
			} else if value.Valid {
				c.AddressLine2 = value.String
			}
		case customer.FieldAddressCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_city", values[i])
			} else if value.Valid {
				c.AddressCity = value.String
			}
		case customer.FieldAddressState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_state", values[i])
			} else if value.Valid {
				c.AddressState = value.String
			}
		case customer.FieldAddressPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_postal_code", values[i])
			} else if value.Valid {
				c.AddressPostalCode = value.String
			}
		case customer.FieldAddressCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_country", values[i])
			} else if value.Valid {
				c.AddressCountry = value.String
			}
		case customer.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Customer.
// This includes values selected through modifiers, order, etc.
func (c *Customer) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return NewCustomerClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(c.TenantID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(c.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(c.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(c.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(c.EnvironmentID)
	builder.WriteString(", ")
	builder.WriteString("external_id=")
	builder.WriteString(c.ExternalID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("address_line1=")
	builder.WriteString(c.AddressLine1)
	builder.WriteString(", ")
	builder.WriteString("address_line2=")
	builder.WriteString(c.AddressLine2)
	builder.WriteString(", ")
	builder.WriteString("address_city=")
	builder.WriteString(c.AddressCity)
	builder.WriteString(", ")
	builder.WriteString("address_state=")
	builder.WriteString(c.AddressState)
	builder.WriteString(", ")
	builder.WriteString("address_postal_code=")
	builder.WriteString(c.AddressPostalCode)
	builder.WriteString(", ")
	builder.WriteString("address_country=")
	builder.WriteString(c.AddressCountry)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", c.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer
