package v201705

type TrialService struct {
	Auth
}

func NewTrialService(auth *Auth) *TrialService {
	return &TrialService{Auth: *auth}
}
