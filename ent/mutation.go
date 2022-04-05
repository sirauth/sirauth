// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sirauth/ent/predicate"
	"sirauth/ent/realm"
	"sirauth/ent/realmoauthclient"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRealm            = "Realm"
	TypeRealmOAuthClient = "RealmOAuthClient"
)

// RealmMutation represents an operation that mutates the Realm nodes in the graph.
type RealmMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Realm, error)
	predicates    []predicate.Realm
}

var _ ent.Mutation = (*RealmMutation)(nil)

// realmOption allows management of the mutation configuration using functional options.
type realmOption func(*RealmMutation)

// newRealmMutation creates new mutation for the Realm entity.
func newRealmMutation(c config, op Op, opts ...realmOption) *RealmMutation {
	m := &RealmMutation{
		config:        c,
		op:            op,
		typ:           TypeRealm,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRealmID sets the ID field of the mutation.
func withRealmID(id int) realmOption {
	return func(m *RealmMutation) {
		var (
			err   error
			once  sync.Once
			value *Realm
		)
		m.oldValue = func(ctx context.Context) (*Realm, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Realm.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRealm sets the old Realm of the mutation.
func withRealm(node *Realm) realmOption {
	return func(m *RealmMutation) {
		m.oldValue = func(context.Context) (*Realm, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RealmMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RealmMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RealmMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RealmMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Realm.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *RealmMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *RealmMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Realm entity.
// If the Realm object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RealmMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *RealmMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the RealmMutation builder.
func (m *RealmMutation) Where(ps ...predicate.Realm) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RealmMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Realm).
func (m *RealmMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RealmMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, realm.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RealmMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case realm.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RealmMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case realm.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Realm field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RealmMutation) SetField(name string, value ent.Value) error {
	switch name {
	case realm.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Realm field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RealmMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RealmMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RealmMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Realm numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RealmMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RealmMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RealmMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Realm nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RealmMutation) ResetField(name string) error {
	switch name {
	case realm.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Realm field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RealmMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RealmMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RealmMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RealmMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RealmMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RealmMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RealmMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Realm unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RealmMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Realm edge %s", name)
}

// RealmOAuthClientMutation represents an operation that mutates the RealmOAuthClient nodes in the graph.
type RealmOAuthClientMutation struct {
	config
	op            Op
	typ           string
	id            *int
	realm_id      *int64
	addrealm_id   *int64
	client_secret *string
	client_id     *string
	redirect_urls *[]string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*RealmOAuthClient, error)
	predicates    []predicate.RealmOAuthClient
}

var _ ent.Mutation = (*RealmOAuthClientMutation)(nil)

// realmoauthclientOption allows management of the mutation configuration using functional options.
type realmoauthclientOption func(*RealmOAuthClientMutation)

// newRealmOAuthClientMutation creates new mutation for the RealmOAuthClient entity.
func newRealmOAuthClientMutation(c config, op Op, opts ...realmoauthclientOption) *RealmOAuthClientMutation {
	m := &RealmOAuthClientMutation{
		config:        c,
		op:            op,
		typ:           TypeRealmOAuthClient,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRealmOAuthClientID sets the ID field of the mutation.
func withRealmOAuthClientID(id int) realmoauthclientOption {
	return func(m *RealmOAuthClientMutation) {
		var (
			err   error
			once  sync.Once
			value *RealmOAuthClient
		)
		m.oldValue = func(ctx context.Context) (*RealmOAuthClient, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().RealmOAuthClient.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRealmOAuthClient sets the old RealmOAuthClient of the mutation.
func withRealmOAuthClient(node *RealmOAuthClient) realmoauthclientOption {
	return func(m *RealmOAuthClientMutation) {
		m.oldValue = func(context.Context) (*RealmOAuthClient, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RealmOAuthClientMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RealmOAuthClientMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RealmOAuthClientMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RealmOAuthClientMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().RealmOAuthClient.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetRealmID sets the "realm_id" field.
func (m *RealmOAuthClientMutation) SetRealmID(i int64) {
	m.realm_id = &i
	m.addrealm_id = nil
}

// RealmID returns the value of the "realm_id" field in the mutation.
func (m *RealmOAuthClientMutation) RealmID() (r int64, exists bool) {
	v := m.realm_id
	if v == nil {
		return
	}
	return *v, true
}

// OldRealmID returns the old "realm_id" field's value of the RealmOAuthClient entity.
// If the RealmOAuthClient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RealmOAuthClientMutation) OldRealmID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRealmID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRealmID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRealmID: %w", err)
	}
	return oldValue.RealmID, nil
}

// AddRealmID adds i to the "realm_id" field.
func (m *RealmOAuthClientMutation) AddRealmID(i int64) {
	if m.addrealm_id != nil {
		*m.addrealm_id += i
	} else {
		m.addrealm_id = &i
	}
}

// AddedRealmID returns the value that was added to the "realm_id" field in this mutation.
func (m *RealmOAuthClientMutation) AddedRealmID() (r int64, exists bool) {
	v := m.addrealm_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetRealmID resets all changes to the "realm_id" field.
func (m *RealmOAuthClientMutation) ResetRealmID() {
	m.realm_id = nil
	m.addrealm_id = nil
}

// SetClientSecret sets the "client_secret" field.
func (m *RealmOAuthClientMutation) SetClientSecret(s string) {
	m.client_secret = &s
}

// ClientSecret returns the value of the "client_secret" field in the mutation.
func (m *RealmOAuthClientMutation) ClientSecret() (r string, exists bool) {
	v := m.client_secret
	if v == nil {
		return
	}
	return *v, true
}

// OldClientSecret returns the old "client_secret" field's value of the RealmOAuthClient entity.
// If the RealmOAuthClient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RealmOAuthClientMutation) OldClientSecret(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldClientSecret is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldClientSecret requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldClientSecret: %w", err)
	}
	return oldValue.ClientSecret, nil
}

// ResetClientSecret resets all changes to the "client_secret" field.
func (m *RealmOAuthClientMutation) ResetClientSecret() {
	m.client_secret = nil
}

// SetClientID sets the "client_id" field.
func (m *RealmOAuthClientMutation) SetClientID(s string) {
	m.client_id = &s
}

// ClientID returns the value of the "client_id" field in the mutation.
func (m *RealmOAuthClientMutation) ClientID() (r string, exists bool) {
	v := m.client_id
	if v == nil {
		return
	}
	return *v, true
}

// OldClientID returns the old "client_id" field's value of the RealmOAuthClient entity.
// If the RealmOAuthClient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RealmOAuthClientMutation) OldClientID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldClientID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldClientID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldClientID: %w", err)
	}
	return oldValue.ClientID, nil
}

// ResetClientID resets all changes to the "client_id" field.
func (m *RealmOAuthClientMutation) ResetClientID() {
	m.client_id = nil
}

// SetRedirectUrls sets the "redirect_urls" field.
func (m *RealmOAuthClientMutation) SetRedirectUrls(s []string) {
	m.redirect_urls = &s
}

// RedirectUrls returns the value of the "redirect_urls" field in the mutation.
func (m *RealmOAuthClientMutation) RedirectUrls() (r []string, exists bool) {
	v := m.redirect_urls
	if v == nil {
		return
	}
	return *v, true
}

// OldRedirectUrls returns the old "redirect_urls" field's value of the RealmOAuthClient entity.
// If the RealmOAuthClient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RealmOAuthClientMutation) OldRedirectUrls(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRedirectUrls is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRedirectUrls requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRedirectUrls: %w", err)
	}
	return oldValue.RedirectUrls, nil
}

// ResetRedirectUrls resets all changes to the "redirect_urls" field.
func (m *RealmOAuthClientMutation) ResetRedirectUrls() {
	m.redirect_urls = nil
}

// Where appends a list predicates to the RealmOAuthClientMutation builder.
func (m *RealmOAuthClientMutation) Where(ps ...predicate.RealmOAuthClient) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RealmOAuthClientMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (RealmOAuthClient).
func (m *RealmOAuthClientMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RealmOAuthClientMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.realm_id != nil {
		fields = append(fields, realmoauthclient.FieldRealmID)
	}
	if m.client_secret != nil {
		fields = append(fields, realmoauthclient.FieldClientSecret)
	}
	if m.client_id != nil {
		fields = append(fields, realmoauthclient.FieldClientID)
	}
	if m.redirect_urls != nil {
		fields = append(fields, realmoauthclient.FieldRedirectUrls)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RealmOAuthClientMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case realmoauthclient.FieldRealmID:
		return m.RealmID()
	case realmoauthclient.FieldClientSecret:
		return m.ClientSecret()
	case realmoauthclient.FieldClientID:
		return m.ClientID()
	case realmoauthclient.FieldRedirectUrls:
		return m.RedirectUrls()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RealmOAuthClientMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case realmoauthclient.FieldRealmID:
		return m.OldRealmID(ctx)
	case realmoauthclient.FieldClientSecret:
		return m.OldClientSecret(ctx)
	case realmoauthclient.FieldClientID:
		return m.OldClientID(ctx)
	case realmoauthclient.FieldRedirectUrls:
		return m.OldRedirectUrls(ctx)
	}
	return nil, fmt.Errorf("unknown RealmOAuthClient field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RealmOAuthClientMutation) SetField(name string, value ent.Value) error {
	switch name {
	case realmoauthclient.FieldRealmID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRealmID(v)
		return nil
	case realmoauthclient.FieldClientSecret:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetClientSecret(v)
		return nil
	case realmoauthclient.FieldClientID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetClientID(v)
		return nil
	case realmoauthclient.FieldRedirectUrls:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRedirectUrls(v)
		return nil
	}
	return fmt.Errorf("unknown RealmOAuthClient field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RealmOAuthClientMutation) AddedFields() []string {
	var fields []string
	if m.addrealm_id != nil {
		fields = append(fields, realmoauthclient.FieldRealmID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RealmOAuthClientMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case realmoauthclient.FieldRealmID:
		return m.AddedRealmID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RealmOAuthClientMutation) AddField(name string, value ent.Value) error {
	switch name {
	case realmoauthclient.FieldRealmID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRealmID(v)
		return nil
	}
	return fmt.Errorf("unknown RealmOAuthClient numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RealmOAuthClientMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RealmOAuthClientMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RealmOAuthClientMutation) ClearField(name string) error {
	return fmt.Errorf("unknown RealmOAuthClient nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RealmOAuthClientMutation) ResetField(name string) error {
	switch name {
	case realmoauthclient.FieldRealmID:
		m.ResetRealmID()
		return nil
	case realmoauthclient.FieldClientSecret:
		m.ResetClientSecret()
		return nil
	case realmoauthclient.FieldClientID:
		m.ResetClientID()
		return nil
	case realmoauthclient.FieldRedirectUrls:
		m.ResetRedirectUrls()
		return nil
	}
	return fmt.Errorf("unknown RealmOAuthClient field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RealmOAuthClientMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RealmOAuthClientMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RealmOAuthClientMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RealmOAuthClientMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RealmOAuthClientMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RealmOAuthClientMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RealmOAuthClientMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown RealmOAuthClient unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RealmOAuthClientMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown RealmOAuthClient edge %s", name)
}
