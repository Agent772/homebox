// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/location"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// LocationQuery is the builder for querying Location entities.
type LocationQuery struct {
	config
	ctx          *QueryContext
	order        []location.OrderOption
	inters       []Interceptor
	predicates   []predicate.Location
	withGroup    *GroupQuery
	withParent   *LocationQuery
	withChildren *LocationQuery
	withItems    *ItemQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LocationQuery builder.
func (lq *LocationQuery) Where(ps ...predicate.Location) *LocationQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LocationQuery) Limit(limit int) *LocationQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LocationQuery) Offset(offset int) *LocationQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LocationQuery) Unique(unique bool) *LocationQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LocationQuery) Order(o ...location.OrderOption) *LocationQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryGroup chains the current query on the "group" edge.
func (lq *LocationQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, location.GroupTable, location.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (lq *LocationQuery) QueryParent() *LocationQuery {
	query := (&LocationClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, location.ParentTable, location.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (lq *LocationQuery) QueryChildren() *LocationQuery {
	query := (&LocationClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, location.ChildrenTable, location.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryItems chains the current query on the "items" edge.
func (lq *LocationQuery) QueryItems() *ItemQuery {
	query := (&ItemClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, location.ItemsTable, location.ItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Location entity from the query.
// Returns a *NotFoundError when no Location was found.
func (lq *LocationQuery) First(ctx context.Context) (*Location, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{location.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LocationQuery) FirstX(ctx context.Context) *Location {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Location ID from the query.
// Returns a *NotFoundError when no Location ID was found.
func (lq *LocationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{location.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LocationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Location entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Location entity is found.
// Returns a *NotFoundError when no Location entities are found.
func (lq *LocationQuery) Only(ctx context.Context) (*Location, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{location.Label}
	default:
		return nil, &NotSingularError{location.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LocationQuery) OnlyX(ctx context.Context) *Location {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Location ID in the query.
// Returns a *NotSingularError when more than one Location ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LocationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{location.Label}
	default:
		err = &NotSingularError{location.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LocationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Locations.
func (lq *LocationQuery) All(ctx context.Context) ([]*Location, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Location, *LocationQuery]()
	return withInterceptors[[]*Location](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LocationQuery) AllX(ctx context.Context) []*Location {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Location IDs.
func (lq *LocationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(location.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LocationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LocationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LocationQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LocationQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LocationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LocationQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LocationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LocationQuery) Clone() *LocationQuery {
	if lq == nil {
		return nil
	}
	return &LocationQuery{
		config:       lq.config,
		ctx:          lq.ctx.Clone(),
		order:        append([]location.OrderOption{}, lq.order...),
		inters:       append([]Interceptor{}, lq.inters...),
		predicates:   append([]predicate.Location{}, lq.predicates...),
		withGroup:    lq.withGroup.Clone(),
		withParent:   lq.withParent.Clone(),
		withChildren: lq.withChildren.Clone(),
		withItems:    lq.withItems.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LocationQuery) WithGroup(opts ...func(*GroupQuery)) *LocationQuery {
	query := (&GroupClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withGroup = query
	return lq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LocationQuery) WithParent(opts ...func(*LocationQuery)) *LocationQuery {
	query := (&LocationClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withParent = query
	return lq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LocationQuery) WithChildren(opts ...func(*LocationQuery)) *LocationQuery {
	query := (&LocationClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withChildren = query
	return lq
}

// WithItems tells the query-builder to eager-load the nodes that are connected to
// the "items" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LocationQuery) WithItems(opts ...func(*ItemQuery)) *LocationQuery {
	query := (&ItemClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withItems = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Location.Query().
//		GroupBy(location.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LocationQuery) GroupBy(field string, fields ...string) *LocationGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LocationGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = location.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Location.Query().
//		Select(location.FieldCreatedAt).
//		Scan(ctx, &v)
func (lq *LocationQuery) Select(fields ...string) *LocationSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LocationSelect{LocationQuery: lq}
	sbuild.label = location.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LocationSelect configured with the given aggregations.
func (lq *LocationQuery) Aggregate(fns ...AggregateFunc) *LocationSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LocationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !location.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LocationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Location, error) {
	var (
		nodes       = []*Location{}
		withFKs     = lq.withFKs
		_spec       = lq.querySpec()
		loadedTypes = [4]bool{
			lq.withGroup != nil,
			lq.withParent != nil,
			lq.withChildren != nil,
			lq.withItems != nil,
		}
	)
	if lq.withGroup != nil || lq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, location.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Location).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Location{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withGroup; query != nil {
		if err := lq.loadGroup(ctx, query, nodes, nil,
			func(n *Location, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	if query := lq.withParent; query != nil {
		if err := lq.loadParent(ctx, query, nodes, nil,
			func(n *Location, e *Location) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := lq.withChildren; query != nil {
		if err := lq.loadChildren(ctx, query, nodes,
			func(n *Location) { n.Edges.Children = []*Location{} },
			func(n *Location, e *Location) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withItems; query != nil {
		if err := lq.loadItems(ctx, query, nodes,
			func(n *Location) { n.Edges.Items = []*Item{} },
			func(n *Location, e *Item) { n.Edges.Items = append(n.Edges.Items, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LocationQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*Location, init func(*Location), assign func(*Location, *Group)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Location)
	for i := range nodes {
		if nodes[i].group_locations == nil {
			continue
		}
		fk := *nodes[i].group_locations
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_locations" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lq *LocationQuery) loadParent(ctx context.Context, query *LocationQuery, nodes []*Location, init func(*Location), assign func(*Location, *Location)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Location)
	for i := range nodes {
		if nodes[i].location_children == nil {
			continue
		}
		fk := *nodes[i].location_children
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(location.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "location_children" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lq *LocationQuery) loadChildren(ctx context.Context, query *LocationQuery, nodes []*Location, init func(*Location), assign func(*Location, *Location)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Location)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Location(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(location.ChildrenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.location_children
		if fk == nil {
			return fmt.Errorf(`foreign-key "location_children" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "location_children" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (lq *LocationQuery) loadItems(ctx context.Context, query *ItemQuery, nodes []*Location, init func(*Location), assign func(*Location, *Item)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Location)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Item(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(location.ItemsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.location_items
		if fk == nil {
			return fmt.Errorf(`foreign-key "location_items" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "location_items" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (lq *LocationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LocationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for i := range fields {
			if fields[i] != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LocationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(location.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = location.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LocationGroupBy is the group-by builder for Location entities.
type LocationGroupBy struct {
	selector
	build *LocationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LocationGroupBy) Aggregate(fns ...AggregateFunc) *LocationGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LocationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LocationQuery, *LocationGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LocationGroupBy) sqlScan(ctx context.Context, root *LocationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LocationSelect is the builder for selecting fields of Location entities.
type LocationSelect struct {
	*LocationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LocationSelect) Aggregate(fns ...AggregateFunc) *LocationSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LocationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LocationQuery, *LocationSelect](ctx, ls.LocationQuery, ls, ls.inters, v)
}

func (ls *LocationSelect) sqlScan(ctx context.Context, root *LocationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
