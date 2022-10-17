// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package check

import (
	"github.com/harness/gitness/types"
)

var (
	ErrRepositoryRequiresSpaceID = &ValidationError{
		"SpaceID required - Repositories don't exist outside of a space.",
	}
)

// Repo checks the provided repository and returns an error in it isn't valid.
func Repo(repo *types.Repository) error {
	// validate name
	if err := PathName(repo.PathName); err != nil {
		return err
	}

	// validate display name
	if err := Name(repo.Name); err != nil {
		return err
	}

	// validate repo within a space
	if repo.SpaceID <= 0 {
		return ErrRepositoryRequiresSpaceID
	}

	// TODO: validate defaultBranch, ...

	return nil
}