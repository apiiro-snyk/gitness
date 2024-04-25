// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repo

import (
	"context"
	"fmt"

	apiauth "github.com/harness/gitness/app/api/auth"
	"github.com/harness/gitness/app/api/usererror"
	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/app/auth/authz"
	"github.com/harness/gitness/app/services/publicaccess"
	"github.com/harness/gitness/app/store"
	"github.com/harness/gitness/types/enum"
)

// GetRepo fetches an active repo (not one that is currently being imported).
func GetRepo(
	ctx context.Context,
	publicAccess publicaccess.PublicAccess,
	repoStore store.RepoStore,
	repoRef string,
) (*Repository, error) {
	if repoRef == "" {
		return nil, usererror.BadRequest("A valid repository reference must be provided.")
	}

	repo, err := repoStore.FindByRef(ctx, repoRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find repository: %w", err)
	}

	if repo.Importing {
		return nil, usererror.BadRequest("Repository import is in progress.")
	}

	isPublic, err := apiauth.CheckRepoIsPublic(ctx, publicAccess, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to check if repo is public: %w", err)
	}

	return &Repository{
		Repository: *repo,
		IsPublic:   isPublic,
	}, nil
}

// GetRepoCheckAccess fetches an active repo (not one that is currently being imported)
// and checks if the current user has permission to access it.
func GetRepoCheckAccess(
	ctx context.Context,
	repoStore store.RepoStore,
	authorizer authz.Authorizer,
	publicAccess publicaccess.PublicAccess,
	session *auth.Session,
	repoRef string,
	reqPermission enum.Permission,
) (*Repository, error) {
	repo, err := GetRepo(ctx, publicAccess, repoStore, repoRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find repo: %w", err)
	}

	if err = apiauth.CheckRepo(ctx, authorizer, session, &repo.Repository, reqPermission); err != nil {
		return nil, fmt.Errorf("access check failed: %w", err)
	}

	return repo, nil
}
