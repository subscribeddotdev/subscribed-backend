// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Message is an object representing the database table.
type Message struct {
	ID            string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrgID         string    `boil:"org_id" json:"org_id" toml:"org_id" yaml:"org_id"`
	ApplicationID string    `boil:"application_id" json:"application_id" toml:"application_id" yaml:"application_id"`
	EventTypeID   string    `boil:"event_type_id" json:"event_type_id" toml:"event_type_id" yaml:"event_type_id"`
	Payload       string    `boil:"payload" json:"payload" toml:"payload" yaml:"payload"`
	SentAt        time.Time `boil:"sent_at" json:"sent_at" toml:"sent_at" yaml:"sent_at"`

	R *messageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L messageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MessageColumns = struct {
	ID            string
	OrgID         string
	ApplicationID string
	EventTypeID   string
	Payload       string
	SentAt        string
}{
	ID:            "id",
	OrgID:         "org_id",
	ApplicationID: "application_id",
	EventTypeID:   "event_type_id",
	Payload:       "payload",
	SentAt:        "sent_at",
}

var MessageTableColumns = struct {
	ID            string
	OrgID         string
	ApplicationID string
	EventTypeID   string
	Payload       string
	SentAt        string
}{
	ID:            "messages.id",
	OrgID:         "messages.org_id",
	ApplicationID: "messages.application_id",
	EventTypeID:   "messages.event_type_id",
	Payload:       "messages.payload",
	SentAt:        "messages.sent_at",
}

// Generated where

var MessageWhere = struct {
	ID            whereHelperstring
	OrgID         whereHelperstring
	ApplicationID whereHelperstring
	EventTypeID   whereHelperstring
	Payload       whereHelperstring
	SentAt        whereHelpertime_Time
}{
	ID:            whereHelperstring{field: "\"messages\".\"id\""},
	OrgID:         whereHelperstring{field: "\"messages\".\"org_id\""},
	ApplicationID: whereHelperstring{field: "\"messages\".\"application_id\""},
	EventTypeID:   whereHelperstring{field: "\"messages\".\"event_type_id\""},
	Payload:       whereHelperstring{field: "\"messages\".\"payload\""},
	SentAt:        whereHelpertime_Time{field: "\"messages\".\"sent_at\""},
}

// MessageRels is where relationship names are stored.
var MessageRels = struct {
	Application string
	Org         string
	EventType   string
}{
	Application: "Application",
	Org:         "Org",
	EventType:   "EventType",
}

// messageR is where relationships are stored.
type messageR struct {
	Application *Application  `boil:"Application" json:"Application" toml:"Application" yaml:"Application"`
	Org         *Organization `boil:"Org" json:"Org" toml:"Org" yaml:"Org"`
	EventType   *EventType    `boil:"EventType" json:"EventType" toml:"EventType" yaml:"EventType"`
}

// NewStruct creates a new relationship struct
func (*messageR) NewStruct() *messageR {
	return &messageR{}
}

func (r *messageR) GetApplication() *Application {
	if r == nil {
		return nil
	}
	return r.Application
}

func (r *messageR) GetOrg() *Organization {
	if r == nil {
		return nil
	}
	return r.Org
}

func (r *messageR) GetEventType() *EventType {
	if r == nil {
		return nil
	}
	return r.EventType
}

// messageL is where Load methods for each relationship are stored.
type messageL struct{}

var (
	messageAllColumns            = []string{"id", "org_id", "application_id", "event_type_id", "payload", "sent_at"}
	messageColumnsWithoutDefault = []string{"id", "org_id", "application_id", "event_type_id", "payload", "sent_at"}
	messageColumnsWithDefault    = []string{}
	messagePrimaryKeyColumns     = []string{"id"}
	messageGeneratedColumns      = []string{}
)

type (
	// MessageSlice is an alias for a slice of pointers to Message.
	// This should almost always be used instead of []Message.
	MessageSlice []*Message
	// MessageHook is the signature for custom Message hook methods
	MessageHook func(context.Context, boil.ContextExecutor, *Message) error

	messageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	messageType                 = reflect.TypeOf(&Message{})
	messageMapping              = queries.MakeStructMapping(messageType)
	messagePrimaryKeyMapping, _ = queries.BindMapping(messageType, messageMapping, messagePrimaryKeyColumns)
	messageInsertCacheMut       sync.RWMutex
	messageInsertCache          = make(map[string]insertCache)
	messageUpdateCacheMut       sync.RWMutex
	messageUpdateCache          = make(map[string]updateCache)
	messageUpsertCacheMut       sync.RWMutex
	messageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var messageAfterSelectHooks []MessageHook

var messageBeforeInsertHooks []MessageHook
var messageAfterInsertHooks []MessageHook

var messageBeforeUpdateHooks []MessageHook
var messageAfterUpdateHooks []MessageHook

var messageBeforeDeleteHooks []MessageHook
var messageAfterDeleteHooks []MessageHook

var messageBeforeUpsertHooks []MessageHook
var messageAfterUpsertHooks []MessageHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Message) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Message) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Message) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Message) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Message) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Message) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Message) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Message) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Message) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMessageHook registers your hook function for all future operations.
func AddMessageHook(hookPoint boil.HookPoint, messageHook MessageHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		messageAfterSelectHooks = append(messageAfterSelectHooks, messageHook)
	case boil.BeforeInsertHook:
		messageBeforeInsertHooks = append(messageBeforeInsertHooks, messageHook)
	case boil.AfterInsertHook:
		messageAfterInsertHooks = append(messageAfterInsertHooks, messageHook)
	case boil.BeforeUpdateHook:
		messageBeforeUpdateHooks = append(messageBeforeUpdateHooks, messageHook)
	case boil.AfterUpdateHook:
		messageAfterUpdateHooks = append(messageAfterUpdateHooks, messageHook)
	case boil.BeforeDeleteHook:
		messageBeforeDeleteHooks = append(messageBeforeDeleteHooks, messageHook)
	case boil.AfterDeleteHook:
		messageAfterDeleteHooks = append(messageAfterDeleteHooks, messageHook)
	case boil.BeforeUpsertHook:
		messageBeforeUpsertHooks = append(messageBeforeUpsertHooks, messageHook)
	case boil.AfterUpsertHook:
		messageAfterUpsertHooks = append(messageAfterUpsertHooks, messageHook)
	}
}

// One returns a single message record from the query.
func (q messageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Message, error) {
	o := &Message{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for messages")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Message records from the query.
func (q messageQuery) All(ctx context.Context, exec boil.ContextExecutor) (MessageSlice, error) {
	var o []*Message

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Message slice")
	}

	if len(messageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Message records in the query.
func (q messageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count messages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q messageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if messages exists")
	}

	return count > 0, nil
}

// Application pointed to by the foreign key.
func (o *Message) Application(mods ...qm.QueryMod) applicationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ApplicationID),
	}

	queryMods = append(queryMods, mods...)

	return Applications(queryMods...)
}

// Org pointed to by the foreign key.
func (o *Message) Org(mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrgID),
	}

	queryMods = append(queryMods, mods...)

	return Organizations(queryMods...)
}

// EventType pointed to by the foreign key.
func (o *Message) EventType(mods ...qm.QueryMod) eventTypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EventTypeID),
	}

	queryMods = append(queryMods, mods...)

	return EventTypes(queryMods...)
}

// LoadApplication allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (messageL) LoadApplication(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMessage interface{}, mods queries.Applicator) error {
	var slice []*Message
	var object *Message

	if singular {
		var ok bool
		object, ok = maybeMessage.(*Message)
		if !ok {
			object = new(Message)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMessage))
			}
		}
	} else {
		s, ok := maybeMessage.(*[]*Message)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMessage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &messageR{}
		}
		args = append(args, object.ApplicationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &messageR{}
			}

			for _, a := range args {
				if a == obj.ApplicationID {
					continue Outer
				}
			}

			args = append(args, obj.ApplicationID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`applications`),
		qm.WhereIn(`applications.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Application")
	}

	var resultSlice []*Application
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Application")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for applications")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for applications")
	}

	if len(applicationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Application = foreign
		if foreign.R == nil {
			foreign.R = &applicationR{}
		}
		foreign.R.Messages = append(foreign.R.Messages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ApplicationID == foreign.ID {
				local.R.Application = foreign
				if foreign.R == nil {
					foreign.R = &applicationR{}
				}
				foreign.R.Messages = append(foreign.R.Messages, local)
				break
			}
		}
	}

	return nil
}

// LoadOrg allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (messageL) LoadOrg(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMessage interface{}, mods queries.Applicator) error {
	var slice []*Message
	var object *Message

	if singular {
		var ok bool
		object, ok = maybeMessage.(*Message)
		if !ok {
			object = new(Message)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMessage))
			}
		}
	} else {
		s, ok := maybeMessage.(*[]*Message)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMessage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &messageR{}
		}
		args = append(args, object.OrgID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &messageR{}
			}

			for _, a := range args {
				if a == obj.OrgID {
					continue Outer
				}
			}

			args = append(args, obj.OrgID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`organizations`),
		qm.WhereIn(`organizations.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organization")
	}

	var resultSlice []*Organization
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organization")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for organizations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for organizations")
	}

	if len(organizationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Org = foreign
		if foreign.R == nil {
			foreign.R = &organizationR{}
		}
		foreign.R.OrgMessages = append(foreign.R.OrgMessages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrgID == foreign.ID {
				local.R.Org = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.OrgMessages = append(foreign.R.OrgMessages, local)
				break
			}
		}
	}

	return nil
}

// LoadEventType allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (messageL) LoadEventType(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMessage interface{}, mods queries.Applicator) error {
	var slice []*Message
	var object *Message

	if singular {
		var ok bool
		object, ok = maybeMessage.(*Message)
		if !ok {
			object = new(Message)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMessage))
			}
		}
	} else {
		s, ok := maybeMessage.(*[]*Message)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMessage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &messageR{}
		}
		args = append(args, object.EventTypeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &messageR{}
			}

			for _, a := range args {
				if a == obj.EventTypeID {
					continue Outer
				}
			}

			args = append(args, obj.EventTypeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`event_types`),
		qm.WhereIn(`event_types.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load EventType")
	}

	var resultSlice []*EventType
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice EventType")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for event_types")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for event_types")
	}

	if len(eventTypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.EventType = foreign
		if foreign.R == nil {
			foreign.R = &eventTypeR{}
		}
		foreign.R.Messages = append(foreign.R.Messages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EventTypeID == foreign.ID {
				local.R.EventType = foreign
				if foreign.R == nil {
					foreign.R = &eventTypeR{}
				}
				foreign.R.Messages = append(foreign.R.Messages, local)
				break
			}
		}
	}

	return nil
}

// SetApplication of the message to the related item.
// Sets o.R.Application to related.
// Adds o to related.R.Messages.
func (o *Message) SetApplication(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Application) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"application_id"}),
		strmangle.WhereClause("\"", "\"", 2, messagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ApplicationID = related.ID
	if o.R == nil {
		o.R = &messageR{
			Application: related,
		}
	} else {
		o.R.Application = related
	}

	if related.R == nil {
		related.R = &applicationR{
			Messages: MessageSlice{o},
		}
	} else {
		related.R.Messages = append(related.R.Messages, o)
	}

	return nil
}

// SetOrg of the message to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.OrgMessages.
func (o *Message) SetOrg(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"org_id"}),
		strmangle.WhereClause("\"", "\"", 2, messagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrgID = related.ID
	if o.R == nil {
		o.R = &messageR{
			Org: related,
		}
	} else {
		o.R.Org = related
	}

	if related.R == nil {
		related.R = &organizationR{
			OrgMessages: MessageSlice{o},
		}
	} else {
		related.R.OrgMessages = append(related.R.OrgMessages, o)
	}

	return nil
}

// SetEventType of the message to the related item.
// Sets o.R.EventType to related.
// Adds o to related.R.Messages.
func (o *Message) SetEventType(ctx context.Context, exec boil.ContextExecutor, insert bool, related *EventType) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"event_type_id"}),
		strmangle.WhereClause("\"", "\"", 2, messagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EventTypeID = related.ID
	if o.R == nil {
		o.R = &messageR{
			EventType: related,
		}
	} else {
		o.R.EventType = related
	}

	if related.R == nil {
		related.R = &eventTypeR{
			Messages: MessageSlice{o},
		}
	} else {
		related.R.Messages = append(related.R.Messages, o)
	}

	return nil
}

// Messages retrieves all the records using an executor.
func Messages(mods ...qm.QueryMod) messageQuery {
	mods = append(mods, qm.From("\"messages\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"messages\".*"})
	}

	return messageQuery{q}
}

// FindMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMessage(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Message, error) {
	messageObj := &Message{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"messages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, messageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from messages")
	}

	if err = messageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return messageObj, err
	}

	return messageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Message) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no messages provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(messageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	messageInsertCacheMut.RLock()
	cache, cached := messageInsertCache[key]
	messageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			messageAllColumns,
			messageColumnsWithDefault,
			messageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(messageType, messageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"messages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"messages\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into messages")
	}

	if !cached {
		messageInsertCacheMut.Lock()
		messageInsertCache[key] = cache
		messageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Message.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Message) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	messageUpdateCacheMut.RLock()
	cache, cached := messageUpdateCache[key]
	messageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update messages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"messages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, messagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, append(wl, messagePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update messages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for messages")
	}

	if !cached {
		messageUpdateCacheMut.Lock()
		messageUpdateCache[key] = cache
		messageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q messageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for messages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MessageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, messagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all message")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Message) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no messages provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(messageColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	messageUpsertCacheMut.RLock()
	cache, cached := messageUpsertCache[key]
	messageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			messageAllColumns,
			messageColumnsWithDefault,
			messageColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert messages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(messagePrimaryKeyColumns))
			copy(conflict, messagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"messages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(messageType, messageMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert messages")
	}

	if !cached {
		messageUpsertCacheMut.Lock()
		messageUpsertCache[key] = cache
		messageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Message record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Message) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Message provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), messagePrimaryKeyMapping)
	sql := "DELETE FROM \"messages\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for messages")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q messageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no messageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for messages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MessageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(messageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for messages")
	}

	if len(messageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Message) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMessage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"messages\".* FROM \"messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MessageSlice")
	}

	*o = slice

	return nil
}

// MessageExists checks if the Message row exists.
func MessageExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"messages\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if messages exists")
	}

	return exists, nil
}

// Exists checks if the Message row exists.
func (o *Message) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MessageExists(ctx, exec, o.ID)
}
