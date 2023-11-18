package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/kcarretto/realm/tavern/internal/auth"
	"github.com/kcarretto/realm/tavern/internal/ent"
)

// Files is the resolver for the files field.
func (r *queryResolver) Files(ctx context.Context, where *ent.FileWhereInput) ([]*ent.File, error) {
	query, err := r.client.File.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Quests is the resolver for the quests field.
func (r *queryResolver) Quests(ctx context.Context, where *ent.QuestWhereInput) ([]*ent.Quest, error) {
	query, err := r.client.Quest.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, where *ent.TaskWhereInput) ([]*ent.Task, error) {
	query, err := r.client.Task.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Beacons is the resolver for the beacons field.
func (r *queryResolver) Beacons(ctx context.Context, where *ent.BeaconWhereInput) ([]*ent.Beacon, error) {
	query, err := r.client.Beacon.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Hosts is the resolver for the hosts field.
func (r *queryResolver) Hosts(ctx context.Context, where *ent.HostWhereInput) ([]*ent.Host, error) {
	query, err := r.client.Host.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, where *ent.TagWhereInput) ([]*ent.Tag, error) {
	query, err := r.client.Tag.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Tomes is the resolver for the tomes field.
func (r *queryResolver) Tomes(ctx context.Context, where *ent.TomeWhereInput) ([]*ent.Tome, error) {
	query, err := r.client.Tome.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, where *ent.UserWhereInput) ([]*ent.User, error) {
	query, err := r.client.User.Query().CollectFields(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect fields: %w", err)
	}
	if where != nil {
		query, err := where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return query.All(ctx)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	if authUser := auth.UserFromContext(ctx); authUser != nil {
		return authUser, nil
	}
	return nil, fmt.Errorf("no authenticated user present in request context")
}
