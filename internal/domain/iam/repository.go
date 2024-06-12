package iam

import "context"

type OrganizationRepository interface {
	Insert(ctx context.Context, org *Organization) error
}

type MemberRepository interface {
	Insert(ctx context.Context, member *Member) error

	ExistsByOr(ctx context.Context, email Email, loginProviderID LoginProviderID) (bool, error)
}
