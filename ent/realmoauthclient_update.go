// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sirauth/ent/predicate"
	"sirauth/ent/realmoauthclient"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RealmOAuthClientUpdate is the builder for updating RealmOAuthClient entities.
type RealmOAuthClientUpdate struct {
	config
	hooks    []Hook
	mutation *RealmOAuthClientMutation
}

// Where appends a list predicates to the RealmOAuthClientUpdate builder.
func (rocu *RealmOAuthClientUpdate) Where(ps ...predicate.RealmOAuthClient) *RealmOAuthClientUpdate {
	rocu.mutation.Where(ps...)
	return rocu
}

// SetRealmID sets the "realm_id" field.
func (rocu *RealmOAuthClientUpdate) SetRealmID(i int64) *RealmOAuthClientUpdate {
	rocu.mutation.ResetRealmID()
	rocu.mutation.SetRealmID(i)
	return rocu
}

// AddRealmID adds i to the "realm_id" field.
func (rocu *RealmOAuthClientUpdate) AddRealmID(i int64) *RealmOAuthClientUpdate {
	rocu.mutation.AddRealmID(i)
	return rocu
}

// SetClientSecret sets the "client_secret" field.
func (rocu *RealmOAuthClientUpdate) SetClientSecret(s string) *RealmOAuthClientUpdate {
	rocu.mutation.SetClientSecret(s)
	return rocu
}

// SetClientID sets the "client_id" field.
func (rocu *RealmOAuthClientUpdate) SetClientID(s string) *RealmOAuthClientUpdate {
	rocu.mutation.SetClientID(s)
	return rocu
}

// SetRedirectUrls sets the "redirect_urls" field.
func (rocu *RealmOAuthClientUpdate) SetRedirectUrls(s []string) *RealmOAuthClientUpdate {
	rocu.mutation.SetRedirectUrls(s)
	return rocu
}

// Mutation returns the RealmOAuthClientMutation object of the builder.
func (rocu *RealmOAuthClientUpdate) Mutation() *RealmOAuthClientMutation {
	return rocu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rocu *RealmOAuthClientUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rocu.hooks) == 0 {
		affected, err = rocu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RealmOAuthClientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rocu.mutation = mutation
			affected, err = rocu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rocu.hooks) - 1; i >= 0; i-- {
			if rocu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rocu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rocu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rocu *RealmOAuthClientUpdate) SaveX(ctx context.Context) int {
	affected, err := rocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rocu *RealmOAuthClientUpdate) Exec(ctx context.Context) error {
	_, err := rocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rocu *RealmOAuthClientUpdate) ExecX(ctx context.Context) {
	if err := rocu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rocu *RealmOAuthClientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   realmoauthclient.Table,
			Columns: realmoauthclient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: realmoauthclient.FieldID,
			},
		},
	}
	if ps := rocu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rocu.mutation.RealmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: realmoauthclient.FieldRealmID,
		})
	}
	if value, ok := rocu.mutation.AddedRealmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: realmoauthclient.FieldRealmID,
		})
	}
	if value, ok := rocu.mutation.ClientSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: realmoauthclient.FieldClientSecret,
		})
	}
	if value, ok := rocu.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: realmoauthclient.FieldClientID,
		})
	}
	if value, ok := rocu.mutation.RedirectUrls(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: realmoauthclient.FieldRedirectUrls,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rocu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{realmoauthclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RealmOAuthClientUpdateOne is the builder for updating a single RealmOAuthClient entity.
type RealmOAuthClientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RealmOAuthClientMutation
}

// SetRealmID sets the "realm_id" field.
func (rocuo *RealmOAuthClientUpdateOne) SetRealmID(i int64) *RealmOAuthClientUpdateOne {
	rocuo.mutation.ResetRealmID()
	rocuo.mutation.SetRealmID(i)
	return rocuo
}

// AddRealmID adds i to the "realm_id" field.
func (rocuo *RealmOAuthClientUpdateOne) AddRealmID(i int64) *RealmOAuthClientUpdateOne {
	rocuo.mutation.AddRealmID(i)
	return rocuo
}

// SetClientSecret sets the "client_secret" field.
func (rocuo *RealmOAuthClientUpdateOne) SetClientSecret(s string) *RealmOAuthClientUpdateOne {
	rocuo.mutation.SetClientSecret(s)
	return rocuo
}

// SetClientID sets the "client_id" field.
func (rocuo *RealmOAuthClientUpdateOne) SetClientID(s string) *RealmOAuthClientUpdateOne {
	rocuo.mutation.SetClientID(s)
	return rocuo
}

// SetRedirectUrls sets the "redirect_urls" field.
func (rocuo *RealmOAuthClientUpdateOne) SetRedirectUrls(s []string) *RealmOAuthClientUpdateOne {
	rocuo.mutation.SetRedirectUrls(s)
	return rocuo
}

// Mutation returns the RealmOAuthClientMutation object of the builder.
func (rocuo *RealmOAuthClientUpdateOne) Mutation() *RealmOAuthClientMutation {
	return rocuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rocuo *RealmOAuthClientUpdateOne) Select(field string, fields ...string) *RealmOAuthClientUpdateOne {
	rocuo.fields = append([]string{field}, fields...)
	return rocuo
}

// Save executes the query and returns the updated RealmOAuthClient entity.
func (rocuo *RealmOAuthClientUpdateOne) Save(ctx context.Context) (*RealmOAuthClient, error) {
	var (
		err  error
		node *RealmOAuthClient
	)
	if len(rocuo.hooks) == 0 {
		node, err = rocuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RealmOAuthClientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rocuo.mutation = mutation
			node, err = rocuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rocuo.hooks) - 1; i >= 0; i-- {
			if rocuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rocuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rocuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rocuo *RealmOAuthClientUpdateOne) SaveX(ctx context.Context) *RealmOAuthClient {
	node, err := rocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rocuo *RealmOAuthClientUpdateOne) Exec(ctx context.Context) error {
	_, err := rocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rocuo *RealmOAuthClientUpdateOne) ExecX(ctx context.Context) {
	if err := rocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rocuo *RealmOAuthClientUpdateOne) sqlSave(ctx context.Context) (_node *RealmOAuthClient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   realmoauthclient.Table,
			Columns: realmoauthclient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: realmoauthclient.FieldID,
			},
		},
	}
	id, ok := rocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RealmOAuthClient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, realmoauthclient.FieldID)
		for _, f := range fields {
			if !realmoauthclient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != realmoauthclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rocuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rocuo.mutation.RealmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: realmoauthclient.FieldRealmID,
		})
	}
	if value, ok := rocuo.mutation.AddedRealmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: realmoauthclient.FieldRealmID,
		})
	}
	if value, ok := rocuo.mutation.ClientSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: realmoauthclient.FieldClientSecret,
		})
	}
	if value, ok := rocuo.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: realmoauthclient.FieldClientID,
		})
	}
	if value, ok := rocuo.mutation.RedirectUrls(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: realmoauthclient.FieldRedirectUrls,
		})
	}
	_node = &RealmOAuthClient{config: rocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{realmoauthclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}