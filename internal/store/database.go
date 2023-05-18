// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

// Package store defines the data storage interfaces.
package store

import (
	"context"

	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

type (
	// PrincipalStore defines the principal data storage.
	PrincipalStore interface {
		/*
		 * PRINCIPAL RELATED OPERATIONS.
		 */

		// Find finds the principal by id.
		Find(ctx context.Context, id int64) (*types.Principal, error)

		// FindByUID finds the principal by uid.
		FindByUID(ctx context.Context, uid string) (*types.Principal, error)

		// FindByEmail finds the principal by email.
		FindByEmail(ctx context.Context, email string) (*types.Principal, error)

		/*
		 * USER RELATED OPERATIONS.
		 */

		// FindUser finds the user by id.
		FindUser(ctx context.Context, id int64) (*types.User, error)

		// FindUserByUID finds the user by uid.
		FindUserByUID(ctx context.Context, uid string) (*types.User, error)

		// FindUserByEmail finds the user by email.
		FindUserByEmail(ctx context.Context, email string) (*types.User, error)

		// CreateUser saves the user details.
		CreateUser(ctx context.Context, user *types.User) error

		// UpdateUser updates an existing user.
		UpdateUser(ctx context.Context, user *types.User) error

		// DeleteUser deletes the user.
		DeleteUser(ctx context.Context, id int64) error

		// ListUsers returns a list of users.
		ListUsers(ctx context.Context, params *types.UserFilter) ([]*types.User, error)

		// CountUsers returns a count of users.
		CountUsers(ctx context.Context) (int64, error)

		/*
		 * SERVICE ACCOUNT RELATED OPERATIONS.
		 */

		// FindServiceAccount finds the service account by id.
		FindServiceAccount(ctx context.Context, id int64) (*types.ServiceAccount, error)

		// FindServiceAccountByUID finds the service account by uid.
		FindServiceAccountByUID(ctx context.Context, uid string) (*types.ServiceAccount, error)

		// CreateServiceAccount saves the service account.
		CreateServiceAccount(ctx context.Context, sa *types.ServiceAccount) error

		// UpdateServiceAccount updates the service account details.
		UpdateServiceAccount(ctx context.Context, sa *types.ServiceAccount) error

		// DeleteServiceAccount deletes the service account.
		DeleteServiceAccount(ctx context.Context, id int64) error

		// ListServiceAccounts returns a list of service accounts for a specific parent.
		ListServiceAccounts(ctx context.Context,
			parentType enum.ParentResourceType, parentID int64) ([]*types.ServiceAccount, error)

		// CountServiceAccounts returns a count of service accounts for a specific parent.
		CountServiceAccounts(ctx context.Context,
			parentType enum.ParentResourceType, parentID int64) (int64, error)

		/*
		 * SERVICE RELATED OPERATIONS.
		 */

		// FindService finds the service by id.
		FindService(ctx context.Context, id int64) (*types.Service, error)

		// FindServiceByUID finds the service by uid.
		FindServiceByUID(ctx context.Context, uid string) (*types.Service, error)

		// CreateService saves the service.
		CreateService(ctx context.Context, sa *types.Service) error

		// UpdateService updates the service.
		UpdateService(ctx context.Context, sa *types.Service) error

		// DeleteService deletes the service.
		DeleteService(ctx context.Context, id int64) error

		// ListServices returns a list of service for a specific parent.
		ListServices(ctx context.Context) ([]*types.Service, error)

		// CountServices returns a count of service for a specific parent.
		CountServices(ctx context.Context) (int64, error)
	}

	// PrincipalInfoView defines helper utility for fetching types.PrincipalInfo objects.
	// It uses the same underlying data storage as PrincipalStore.
	PrincipalInfoView interface {
		Find(ctx context.Context, id int64) (*types.PrincipalInfo, error)
		FindMany(ctx context.Context, ids []int64) ([]*types.PrincipalInfo, error)
	}

	// PathStore defines the path data storage.
	// It is used to store routing paths for repos & spaces.
	PathStore interface {
		// Create creates a new path.
		Create(ctx context.Context, path *types.Path) error

		// Find finds the path for the given id.
		Find(ctx context.Context, id int64) (*types.Path, error)

		// FindWithLock finds the path for the given id and locks the entry.
		FindWithLock(ctx context.Context, id int64) (*types.Path, error)

		// FindValue finds the path for the given value.
		FindValue(ctx context.Context, value string) (*types.Path, error)

		// FindPrimary finds the primary path for a target.
		FindPrimary(ctx context.Context, targetType enum.PathTargetType, targetID int64) (*types.Path, error)

		// FindPrimaryWithLock finds the primary path for a target and locks the db entry.
		FindPrimaryWithLock(ctx context.Context, targetType enum.PathTargetType, targetID int64) (*types.Path, error)

		// Update updates an existing path.
		Update(ctx context.Context, path *types.Path) error

		// Delete deletes a specific path.
		Delete(ctx context.Context, id int64) error

		// Count returns the count of paths for a target.
		Count(ctx context.Context, targetType enum.PathTargetType, targetID int64,
			opts *types.PathFilter) (int64, error)

		// List lists all paths for a target.
		List(ctx context.Context, targetType enum.PathTargetType, targetID int64,
			opts *types.PathFilter) ([]*types.Path, error)

		// ListPrimaryDescendantsWithLock lists all primary paths that are descendants of the given path and locks them.
		ListPrimaryDescendantsWithLock(ctx context.Context, value string) ([]*types.Path, error)
	}

	// SpaceStore defines the space data storage.
	SpaceStore interface {
		// Find the space by id.
		Find(ctx context.Context, id int64) (*types.Space, error)

		// FindByRef finds the space using the spaceRef as either the id or the space path.
		FindByRef(ctx context.Context, spaceRef string) (*types.Space, error)

		// Create creates a new space
		Create(ctx context.Context, space *types.Space) error

		// Update updates the space details.
		Update(ctx context.Context, space *types.Space) error

		// UpdateOptLock updates the space using the optimistic locking mechanism.
		UpdateOptLock(ctx context.Context, space *types.Space,
			mutateFn func(space *types.Space) error) (*types.Space, error)

		// Delete deletes the space.
		Delete(ctx context.Context, id int64) error

		// Count the child spaces of a space.
		Count(ctx context.Context, id int64, opts *types.SpaceFilter) (int64, error)

		// List returns a list of child spaces in a space.
		List(ctx context.Context, id int64, opts *types.SpaceFilter) ([]*types.Space, error)
	}

	// RepoStore defines the repository data storage.
	RepoStore interface {
		// Find the repo by id.
		Find(ctx context.Context, id int64) (*types.Repository, error)

		// FindByRef finds the repo using the repoRef as either the id or the repo path.
		FindByRef(ctx context.Context, repoRef string) (*types.Repository, error)

		// Create a new repo.
		Create(ctx context.Context, repo *types.Repository) error

		// Update the repo details.
		Update(ctx context.Context, repo *types.Repository) error

		// UpdateOptLock the repo details using the optimistic locking mechanism.
		UpdateOptLock(ctx context.Context, repo *types.Repository,
			mutateFn func(repository *types.Repository) error) (*types.Repository, error)

		// Delete the repo.
		Delete(ctx context.Context, id int64) error

		// Count of repos in a space.
		Count(ctx context.Context, parentID int64, opts *types.RepoFilter) (int64, error)

		// List returns a list of repos in a space.
		List(ctx context.Context, parentID int64, opts *types.RepoFilter) ([]*types.Repository, error)
	}

	// RepoGitInfoView defines the repository GitUID view.
	RepoGitInfoView interface {
		Find(ctx context.Context, id int64) (*types.RepositoryGitInfo, error)
	}

	// TokenStore defines the token data storage.
	TokenStore interface {
		// Find finds the token by id
		Find(ctx context.Context, id int64) (*types.Token, error)

		// FindByUID finds the token by principalId and tokenUID
		FindByUID(ctx context.Context, principalID int64, tokenUID string) (*types.Token, error)

		// Create saves the token details.
		Create(ctx context.Context, token *types.Token) error

		// Delete deletes the token with the given id.
		Delete(ctx context.Context, id int64) error

		// DeleteForPrincipal deletes all tokens for a specific principal
		DeleteForPrincipal(ctx context.Context, principalID int64) error

		// List returns a list of tokens of a specific type for a specific principal.
		List(ctx context.Context, principalID int64, tokenType enum.TokenType) ([]*types.Token, error)

		// Count returns a count of tokens of a specifc type for a specific principal.
		Count(ctx context.Context, principalID int64, tokenType enum.TokenType) (int64, error)
	}

	// PullReqStore defines the pull request data storage.
	PullReqStore interface {
		// Find the pull request by id.
		Find(ctx context.Context, id int64) (*types.PullReq, error)

		// FindByNumberWithLock finds the pull request by repo ID and the pull request number
		// and acquires an exclusive lock of the pull request database row for the duration of the transaction.
		FindByNumberWithLock(ctx context.Context, repoID, number int64) (*types.PullReq, error)

		// FindByNumber finds the pull request by repo ID and the pull request number.
		FindByNumber(ctx context.Context, repoID, number int64) (*types.PullReq, error)

		// Create a new pull request.
		Create(ctx context.Context, pullreq *types.PullReq) error

		// Update the pull request. It will set new values to the Version and Updated fields.
		Update(ctx context.Context, pr *types.PullReq) error

		// UpdateOptLock the pull request details using the optimistic locking mechanism.
		UpdateOptLock(ctx context.Context, pr *types.PullReq,
			mutateFn func(pr *types.PullReq) error) (*types.PullReq, error)

		// UpdateActivitySeq the pull request's activity sequence number.
		// It will set new values to the ActivitySeq, Version and Updated fields.
		UpdateActivitySeq(ctx context.Context, pr *types.PullReq) (*types.PullReq, error)

		// Delete the pull request.
		Delete(ctx context.Context, id int64) error

		// Count of pull requests in a space.
		Count(ctx context.Context, opts *types.PullReqFilter) (int64, error)

		// List returns a list of pull requests in a space.
		List(ctx context.Context, opts *types.PullReqFilter) ([]*types.PullReq, error)
	}

	PullReqActivityStore interface {
		// Find the pull request activity by id.
		Find(ctx context.Context, id int64) (*types.PullReqActivity, error)

		// Create a new pull request activity. Value of the Order field should be fetched with UpdateActivitySeq.
		// Value of the SubOrder field (for replies) should be the incremented ReplySeq field (non-replies have 0).
		Create(ctx context.Context, act *types.PullReqActivity) error

		// CreateWithPayload create a new system activity from the provided payload.
		CreateWithPayload(ctx context.Context,
			pr *types.PullReq, principalID int64, payload types.PullReqActivityPayload) (*types.PullReqActivity, error)

		// Update the pull request activity. It will set new values to the Version and Updated fields.
		Update(ctx context.Context, act *types.PullReqActivity) error

		// UpdateOptLock updates the pull request activity using the optimistic locking mechanism.
		UpdateOptLock(ctx context.Context,
			act *types.PullReqActivity,
			mutateFn func(act *types.PullReqActivity) error,
		) (*types.PullReqActivity, error)

		// Count returns number of pull request activities in a pull request.
		Count(ctx context.Context, prID int64, opts *types.PullReqActivityFilter) (int64, error)

		// CountUnresolved returns number of unresolved comments.
		CountUnresolved(ctx context.Context, prID int64) (int, error)

		// List returns a list of pull request activities in a pull request (a timeline).
		List(ctx context.Context, prID int64, opts *types.PullReqActivityFilter) ([]*types.PullReqActivity, error)
	}

	// CodeCommentView is to manipulate only code-comment subset of PullReqActivity.
	// It's used by internal service that migrates code comment line numbers after new commits.
	CodeCommentView interface {
		// ListNotAtSourceSHA loads code comments that need to be updated after a new commit.
		// Resulting list is ordered by the file name and the relevant line number.
		ListNotAtSourceSHA(ctx context.Context, prID int64, sourceSHA string) ([]*types.CodeComment, error)

		// ListNotAtMergeBaseSHA loads code comments that need to be updated after merge base update.
		// Resulting list is ordered by the file name and the relevant line number.
		ListNotAtMergeBaseSHA(ctx context.Context, prID int64, targetSHA string) ([]*types.CodeComment, error)

		// UpdateAll updates code comments (pull request activity of types code-comment).
		// entities coming from the input channel.
		UpdateAll(ctx context.Context, codeComments []*types.CodeComment) error
	}

	// PullReqReviewStore defines the pull request review storage.
	PullReqReviewStore interface {
		// Find returns the pull request review entity or an error if it doesn't exist.
		Find(ctx context.Context, id int64) (*types.PullReqReview, error)

		// Create creates a new pull request review.
		Create(ctx context.Context, v *types.PullReqReview) error
	}

	// PullReqReviewerStore defines the pull request reviewer storage.
	PullReqReviewerStore interface {
		// Find returns the pull request reviewer or an error if it doesn't exist.
		Find(ctx context.Context, prID, principalID int64) (*types.PullReqReviewer, error)

		// Create creates the new pull request reviewer.
		Create(ctx context.Context, v *types.PullReqReviewer) error

		// Update updates the pull request reviewer.
		Update(ctx context.Context, v *types.PullReqReviewer) error

		// List returns all pull request reviewers for the pull request.
		List(ctx context.Context, prID int64) ([]*types.PullReqReviewer, error)
	}

	// WebhookStore defines the webhook data storage.
	WebhookStore interface {
		// Find finds the webhook by id.
		Find(ctx context.Context, id int64) (*types.Webhook, error)

		// Create creates a new webhook.
		Create(ctx context.Context, hook *types.Webhook) error

		// Update updates an existing webhook.
		Update(ctx context.Context, hook *types.Webhook) error

		// UpdateOptLock updates the webhook using the optimistic locking mechanism.
		UpdateOptLock(ctx context.Context, hook *types.Webhook,
			mutateFn func(hook *types.Webhook) error) (*types.Webhook, error)

		// Delete deletes the webhook for the given id.
		Delete(ctx context.Context, id int64) error

		// Count counts the webhooks for a given parent type and id.
		Count(ctx context.Context, parentType enum.WebhookParent, parentID int64,
			opts *types.WebhookFilter) (int64, error)

		// List lists the webhooks for a given parent type and id.
		List(ctx context.Context, parentType enum.WebhookParent, parentID int64,
			opts *types.WebhookFilter) ([]*types.Webhook, error)
	}

	// WebhookExecutionStore defines the webhook execution data storage.
	WebhookExecutionStore interface {
		// Find finds the webhook execution by id.
		Find(ctx context.Context, id int64) (*types.WebhookExecution, error)

		// Create creates a new webhook execution entry.
		Create(ctx context.Context, hook *types.WebhookExecution) error

		// ListForWebhook lists the webhook executions for a given webhook id.
		ListForWebhook(ctx context.Context, webhookID int64,
			opts *types.WebhookExecutionFilter) ([]*types.WebhookExecution, error)

		// ListForTrigger lists the webhook executions for a given trigger id.
		ListForTrigger(ctx context.Context, triggerID string) ([]*types.WebhookExecution, error)
	}
)