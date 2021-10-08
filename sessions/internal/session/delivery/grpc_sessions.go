package delivery

import (
	"context"

	"github.com/ferjmc/api_ddd/sessions/internal/csrf"
	"github.com/ferjmc/api_ddd/sessions/internal/models"
	"github.com/ferjmc/api_ddd/sessions/internal/session"
	"github.com/ferjmc/api_ddd/sessions/pkg/grpc_errors"
	"github.com/ferjmc/api_ddd/sessions/pkg/logger"
	sessionService "github.com/ferjmc/api_ddd/sessions/proto"
	"google.golang.org/grpc/status"

	uuid "github.com/satori/go.uuid"
)

type SessionsService struct {
	logger logger.Logger
	sessUC session.SessUseCase
	csrfUC csrf.UseCase
	sessionService.UnimplementedAuthorizationServiceServer
}

func NewSessionsService(logger logger.Logger, sessUC session.SessUseCase, csrfUC csrf.UseCase) *SessionsService {
	return &SessionsService{logger: logger, sessUC: sessUC, csrfUC: csrfUC}
}

func (s *SessionsService) CreateSession(ctx context.Context, r *sessionService.CreateSessionRequest) (*sessionService.CreateSessionResponse, error) {

	userUUID, err := uuid.FromString(r.UserID)
	if err != nil {
		s.logger.Errorf("uuid.FromString: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "uuid.FromString: %v", err)
	}
	sess, err := s.sessUC.CreateSession(ctx, userUUID)
	if err != nil {
		s.logger.Errorf("sessUC.CreateSession: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "sessUC.CreateSession: %v", err)
	}

	return &sessionService.CreateSessionResponse{Session: s.sessionJSONToProto(sess)}, nil
}

func (s *SessionsService) GetSessionByID(ctx context.Context, r *sessionService.GetSessionByIDRequest) (*sessionService.GetSessionByIDResponse, error) {

	sess, err := s.sessUC.GetSessionByID(ctx, r.SessionID)
	if err != nil {
		s.logger.Errorf("sessUC.GetSessionByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "sessUC.GetSessionByID: %v", err)
	}

	return &sessionService.GetSessionByIDResponse{Session: s.sessionJSONToProto(sess)}, nil
}

func (s *SessionsService) DeleteSession(ctx context.Context, r *sessionService.DeleteSessionRequest) (*sessionService.DeleteSessionResponse, error) {

	if err := s.sessUC.DeleteSession(ctx, r.SessionID); err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "sessUC.DeleteSession: %v", err)
	}

	return &sessionService.DeleteSessionResponse{SessionID: r.SessionID}, nil
}

func (s *SessionsService) CreateCsrfToken(ctx context.Context, r *sessionService.CreateCsrfTokenRequest) (*sessionService.CreateCsrfTokenResponse, error) {

	token, err := s.csrfUC.GetCSRFToken(ctx, r.GetCsrfTokenInput().GetSessionID())
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "csrfUC.CreateCsrfToken: %v", err)
	}

	return &sessionService.CreateCsrfTokenResponse{CsrfToken: &sessionService.CsrfToken{Token: token}}, nil
}

func (s *SessionsService) CheckCsrfToken(ctx context.Context, r *sessionService.CheckCsrfTokenRequest) (*sessionService.CheckCsrfTokenResponse, error) {
	isValid, err := s.csrfUC.ValidateCSRFToken(ctx, r.GetCsrfTokenCheck().GetSessionID(), r.GetCsrfTokenCheck().GetToken())
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "csrfUC.CheckToken: %v", err)
	}

	return &sessionService.CheckCsrfTokenResponse{CheckResult: &sessionService.CheckResult{Result: isValid}}, nil
}

func (s *SessionsService) sessionJSONToProto(sess *models.Session) *sessionService.Session {
	return &sessionService.Session{
		UserID:    sess.UserID.String(),
		SessionID: sess.SessionID,
	}
}
