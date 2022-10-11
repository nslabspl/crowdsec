// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wojtekxtx/crowdsec/pkg/database/ent/alert"
	"github.com/wojtekxtx/crowdsec/pkg/database/ent/decision"
)

// DecisionCreate is the builder for creating a Decision entity.
type DecisionCreate struct {
	config
	mutation *DecisionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DecisionCreate) SetCreatedAt(t time.Time) *DecisionCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableCreatedAt(t *time.Time) *DecisionCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DecisionCreate) SetUpdatedAt(t time.Time) *DecisionCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableUpdatedAt(t *time.Time) *DecisionCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetUntil sets the "until" field.
func (dc *DecisionCreate) SetUntil(t time.Time) *DecisionCreate {
	dc.mutation.SetUntil(t)
	return dc
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableUntil(t *time.Time) *DecisionCreate {
	if t != nil {
		dc.SetUntil(*t)
	}
	return dc
}

// SetScenario sets the "scenario" field.
func (dc *DecisionCreate) SetScenario(s string) *DecisionCreate {
	dc.mutation.SetScenario(s)
	return dc
}

// SetType sets the "type" field.
func (dc *DecisionCreate) SetType(s string) *DecisionCreate {
	dc.mutation.SetType(s)
	return dc
}

// SetStartIP sets the "start_ip" field.
func (dc *DecisionCreate) SetStartIP(i int64) *DecisionCreate {
	dc.mutation.SetStartIP(i)
	return dc
}

// SetNillableStartIP sets the "start_ip" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableStartIP(i *int64) *DecisionCreate {
	if i != nil {
		dc.SetStartIP(*i)
	}
	return dc
}

// SetEndIP sets the "end_ip" field.
func (dc *DecisionCreate) SetEndIP(i int64) *DecisionCreate {
	dc.mutation.SetEndIP(i)
	return dc
}

// SetNillableEndIP sets the "end_ip" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableEndIP(i *int64) *DecisionCreate {
	if i != nil {
		dc.SetEndIP(*i)
	}
	return dc
}

// SetStartSuffix sets the "start_suffix" field.
func (dc *DecisionCreate) SetStartSuffix(i int64) *DecisionCreate {
	dc.mutation.SetStartSuffix(i)
	return dc
}

// SetNillableStartSuffix sets the "start_suffix" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableStartSuffix(i *int64) *DecisionCreate {
	if i != nil {
		dc.SetStartSuffix(*i)
	}
	return dc
}

// SetEndSuffix sets the "end_suffix" field.
func (dc *DecisionCreate) SetEndSuffix(i int64) *DecisionCreate {
	dc.mutation.SetEndSuffix(i)
	return dc
}

// SetNillableEndSuffix sets the "end_suffix" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableEndSuffix(i *int64) *DecisionCreate {
	if i != nil {
		dc.SetEndSuffix(*i)
	}
	return dc
}

// SetIPSize sets the "ip_size" field.
func (dc *DecisionCreate) SetIPSize(i int64) *DecisionCreate {
	dc.mutation.SetIPSize(i)
	return dc
}

// SetNillableIPSize sets the "ip_size" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableIPSize(i *int64) *DecisionCreate {
	if i != nil {
		dc.SetIPSize(*i)
	}
	return dc
}

// SetScope sets the "scope" field.
func (dc *DecisionCreate) SetScope(s string) *DecisionCreate {
	dc.mutation.SetScope(s)
	return dc
}

// SetValue sets the "value" field.
func (dc *DecisionCreate) SetValue(s string) *DecisionCreate {
	dc.mutation.SetValue(s)
	return dc
}

// SetOrigin sets the "origin" field.
func (dc *DecisionCreate) SetOrigin(s string) *DecisionCreate {
	dc.mutation.SetOrigin(s)
	return dc
}

// SetSimulated sets the "simulated" field.
func (dc *DecisionCreate) SetSimulated(b bool) *DecisionCreate {
	dc.mutation.SetSimulated(b)
	return dc
}

// SetNillableSimulated sets the "simulated" field if the given value is not nil.
func (dc *DecisionCreate) SetNillableSimulated(b *bool) *DecisionCreate {
	if b != nil {
		dc.SetSimulated(*b)
	}
	return dc
}

// SetOwnerID sets the "owner" edge to the Alert entity by ID.
func (dc *DecisionCreate) SetOwnerID(id int) *DecisionCreate {
	dc.mutation.SetOwnerID(id)
	return dc
}

// SetNillableOwnerID sets the "owner" edge to the Alert entity by ID if the given value is not nil.
func (dc *DecisionCreate) SetNillableOwnerID(id *int) *DecisionCreate {
	if id != nil {
		dc = dc.SetOwnerID(*id)
	}
	return dc
}

// SetOwner sets the "owner" edge to the Alert entity.
func (dc *DecisionCreate) SetOwner(a *Alert) *DecisionCreate {
	return dc.SetOwnerID(a.ID)
}

// Mutation returns the DecisionMutation object of the builder.
func (dc *DecisionCreate) Mutation() *DecisionMutation {
	return dc.mutation
}

// Save creates the Decision in the database.
func (dc *DecisionCreate) Save(ctx context.Context) (*Decision, error) {
	var (
		err  error
		node *Decision
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DecisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DecisionCreate) SaveX(ctx context.Context) *Decision {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DecisionCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DecisionCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DecisionCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := decision.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := decision.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.Simulated(); !ok {
		v := decision.DefaultSimulated
		dc.mutation.SetSimulated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DecisionCreate) check() error {
	if _, ok := dc.mutation.Scenario(); !ok {
		return &ValidationError{Name: "scenario", err: errors.New(`ent: missing required field "Decision.scenario"`)}
	}
	if _, ok := dc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Decision.type"`)}
	}
	if _, ok := dc.mutation.Scope(); !ok {
		return &ValidationError{Name: "scope", err: errors.New(`ent: missing required field "Decision.scope"`)}
	}
	if _, ok := dc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Decision.value"`)}
	}
	if _, ok := dc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "Decision.origin"`)}
	}
	if _, ok := dc.mutation.Simulated(); !ok {
		return &ValidationError{Name: "simulated", err: errors.New(`ent: missing required field "Decision.simulated"`)}
	}
	return nil
}

func (dc *DecisionCreate) sqlSave(ctx context.Context) (*Decision, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DecisionCreate) createSpec() (*Decision, *sqlgraph.CreateSpec) {
	var (
		_node = &Decision{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: decision.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: decision.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: decision.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: decision.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := dc.mutation.Until(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: decision.FieldUntil,
		})
		_node.Until = &value
	}
	if value, ok := dc.mutation.Scenario(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: decision.FieldScenario,
		})
		_node.Scenario = value
	}
	if value, ok := dc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: decision.FieldType,
		})
		_node.Type = value
	}
	if value, ok := dc.mutation.StartIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: decision.FieldStartIP,
		})
		_node.StartIP = value
	}
	if value, ok := dc.mutation.EndIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: decision.FieldEndIP,
		})
		_node.EndIP = value
	}
	if value, ok := dc.mutation.StartSuffix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: decision.FieldStartSuffix,
		})
		_node.StartSuffix = value
	}
	if value, ok := dc.mutation.EndSuffix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: decision.FieldEndSuffix,
		})
		_node.EndSuffix = value
	}
	if value, ok := dc.mutation.IPSize(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: decision.FieldIPSize,
		})
		_node.IPSize = value
	}
	if value, ok := dc.mutation.Scope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: decision.FieldScope,
		})
		_node.Scope = value
	}
	if value, ok := dc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: decision.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := dc.mutation.Origin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: decision.FieldOrigin,
		})
		_node.Origin = value
	}
	if value, ok := dc.mutation.Simulated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: decision.FieldSimulated,
		})
		_node.Simulated = value
	}
	if nodes := dc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   decision.OwnerTable,
			Columns: []string{decision.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.alert_decisions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DecisionCreateBulk is the builder for creating many Decision entities in bulk.
type DecisionCreateBulk struct {
	config
	builders []*DecisionCreate
}

// Save creates the Decision entities in the database.
func (dcb *DecisionCreateBulk) Save(ctx context.Context) ([]*Decision, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Decision, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DecisionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DecisionCreateBulk) SaveX(ctx context.Context) []*Decision {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DecisionCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DecisionCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
