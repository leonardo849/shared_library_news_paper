package units_tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/leonardo849/shared_library_news_paper/pkg/consts"
	"github.com/leonardo849/shared_library_news_paper/pkg/dto"
)

func TestPublishUserAuth(t *testing.T) {
	const (
		username = "batman123"
		
		role = "CUSTOMER"
	)
	authid := uuid.New().String()
	dto := dto.AuthPublishUserCreated{
		AuthId: authid,
		Username: username,
		Role: role,
	}
	if dto.AuthId != authid || dto.Username != username || dto.Role != role {
		t.Error("error dto.AuthPublishUserCreated ")
	}
}

func TestRoles(t *testing.T) {
	if !(consts.Ceo == "CEO" && consts.Customer == "CUSTOMER" && consts.Journalist == "JOURNALIST" && consts.Developer == "DEVELOPER" && len(consts.Roles) == 4) {
		t.Error("error in roles")
	}
} 

func TestAuthConsts(t *testing.T) {
	if consts.ExchangeNameAuthEvents != "auth_events" || consts.KeyUserAuthDeleted != "user.auth.deleted" || consts.KeyUserAuthVerified != "user.auth.verified" {
		t.Error("error in auth consts")
	}
}

func TestEmailConsts(t *testing.T)  {
	if consts.ExchangeNameEmail != "email_direct" || consts.KeyEmail != "email" {
		t.Error("error in email consts")
	}
}

