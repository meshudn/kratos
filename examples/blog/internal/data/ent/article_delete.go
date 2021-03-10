// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-kratos/kratos/examples/blog/internal/data/ent/article"
	"github.com/go-kratos/kratos/examples/blog/internal/data/ent/predicate"
)

// ArticleDelete is the builder for deleting a Article entity.
type ArticleDelete struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where adds a new predicate to the ArticleDelete builder.
func (ad *ArticleDelete) Where(ps ...predicate.Article) *ArticleDelete {
	ad.mutation.predicates = append(ad.mutation.predicates, ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *ArticleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ad.hooks) == 0 {
		affected, err = ad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ad.mutation = mutation
			affected, err = ad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ad.hooks) - 1; i >= 0; i-- {
			mut = ad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *ArticleDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *ArticleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: article.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: article.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
}

// ArticleDeleteOne is the builder for deleting a single Article entity.
type ArticleDeleteOne struct {
	ad *ArticleDelete
}

// Exec executes the deletion query.
func (ado *ArticleDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{article.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *ArticleDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
