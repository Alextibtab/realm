// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kcarretto/realm/tavern/ent/tome"
)

// Tome is the model entity for the Tome schema.
type Tome struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name of the tome
	Name string `json:"name,omitempty"`
	// Information about the tome
	Description string `json:"description,omitempty"`
	// JSON string describing what parameters are used with the tome
	Parameters string `json:"parameters,omitempty"`
	// The size of the tome in bytes
	Size int `json:"size,omitempty"`
	// A SHA3 digest of the content field
	Hash string `json:"hash,omitempty"`
	// The timestamp for when the Tome was created
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// The timestamp for when the Tome was last modified
	LastModifiedAt time.Time `json:"lastModifiedAt,omitempty"`
	// The content of the tome
	Content []byte `json:"content,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tome) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tome.FieldContent:
			values[i] = new([]byte)
		case tome.FieldID, tome.FieldSize:
			values[i] = new(sql.NullInt64)
		case tome.FieldName, tome.FieldDescription, tome.FieldParameters, tome.FieldHash:
			values[i] = new(sql.NullString)
		case tome.FieldCreatedAt, tome.FieldLastModifiedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tome", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tome fields.
func (t *Tome) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tome.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tome.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tome.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case tome.FieldParameters:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parameters", values[i])
			} else if value.Valid {
				t.Parameters = value.String
			}
		case tome.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				t.Size = int(value.Int64)
			}
		case tome.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				t.Hash = value.String
			}
		case tome.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tome.FieldLastModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastModifiedAt", values[i])
			} else if value.Valid {
				t.LastModifiedAt = value.Time
			}
		case tome.FieldContent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value != nil {
				t.Content = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Tome.
// Note that you need to call Tome.Unwrap() before calling this method if this Tome
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tome) Update() *TomeUpdateOne {
	return (&TomeClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tome entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tome) Unwrap() *Tome {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tome is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tome) String() string {
	var builder strings.Builder
	builder.WriteString("Tome(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("parameters=")
	builder.WriteString(t.Parameters)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", t.Size))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(t.Hash)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("lastModifiedAt=")
	builder.WriteString(t.LastModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(fmt.Sprintf("%v", t.Content))
	builder.WriteByte(')')
	return builder.String()
}

// Tomes is a parsable slice of Tome.
type Tomes []*Tome

func (t Tomes) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
