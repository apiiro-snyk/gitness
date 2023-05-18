// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package space

import (
	"context"
	"fmt"
	"math"

	apiauth "github.com/harness/gitness/internal/api/auth"
	"github.com/harness/gitness/internal/auth"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

// Delete deletes a space.
func (c *Controller) Delete(ctx context.Context, session *auth.Session, spaceRef string) error {
	space, err := c.spaceStore.FindByRef(ctx, spaceRef)
	if err != nil {
		return err
	}
	if err = apiauth.CheckSpace(ctx, c.authorizer, session, space, enum.PermissionSpaceDelete, false); err != nil {
		return err
	}
	sfilter := &types.SpaceFilter{
		Page:  1,
		Size:  int(math.MaxInt),
		Query: "",
		Order: enum.OrderAsc,
		Sort:  enum.SpaceAttrNone,
	}
	return c.DeleteNoAuth(ctx, session, space.ID, sfilter)
}

// DeleteNoAuth bypasses these permission
// PermissionSpaceDelete, PermissionSpaceView, PermissionRepoView, PermissionRepoDelete.
func (c *Controller) DeleteNoAuth(
	ctx context.Context,
	session *auth.Session,
	spaceID int64,
	filter *types.SpaceFilter,
) error {
	subSpaces, _, err := c.ListSpacesNoAuth(ctx, spaceID, filter)
	if err != nil {
		return fmt.Errorf("failed to list space %d sub spaces: %w", spaceID, err)
	}
	for _, space := range subSpaces {
		err = c.DeleteNoAuth(ctx, session, space.ID, filter)
		if err != nil {
			return fmt.Errorf("failed to delete space %d: %w", space.ID, err)
		}
	}
	err = c.deleteRepositoriesNoAuth(ctx, session, spaceID)
	if err != nil {
		return fmt.Errorf("failed to delete repositories of space %d: %w", spaceID, err)
	}
	err = c.spaceStore.Delete(ctx, spaceID)
	if err != nil {
		return fmt.Errorf("spaceStore failed to delete space %d: %w", spaceID, err)
	}
	return nil
}