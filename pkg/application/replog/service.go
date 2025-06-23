package replog

import (
	"context"
	"sync"

	"github.com/rotisserie/eris"
	"go.uber.org/zap"

	"github.com/sashajdn/replog/pkg/application/replog/messaging"
	"github.com/sashajdn/replog/pkg/log"
)

func NewService(
	logger *log.SugaredLogger,
	messagingClient messaging.Client,
	entryRepository ReplogRepository,
	userRepository UserRepository,
	messagingProviderRepository MessagingProviderRepository,
) *Service {
	slogger := logger.With("service", `replog`)

	return &Service{
		messagingClient:             messagingClient,
		logger:                      slogger,
		stateMachine:                NewStateMachine(slogger),
		replogRepository:            entryRepository,
		userRepository:              userRepository,
		messagingProviderRepository: messagingProviderRepository,
	}
}

type Service struct {
	wg                          sync.WaitGroup
	messagingClient             messaging.Client
	logger                      *log.SugaredLogger
	stateMachine                *stateMachine
	replogRepository            ReplogRepository
	userRepository              UserRepository
	messagingProviderRepository MessagingProviderRepository
}

func (s *Service) Run(ctx context.Context) error {
	receiver, err := s.messagingClient.Receive(ctx)
	if err != nil {
		return eris.Wrap(err, "receive messages: %w")
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		for {
			select {
			case message := <-receiver:
				if err := s.handleMessage(ctx, message); err != nil {
					s.logger.With(zap.Error(err)).Error("Failed to handle message")
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (s *Service) Close() error {
	return s.messagingClient.Close()
}

func (s *Service) handleMessage(ctx context.Context, message *messaging.Message) error {
	if message == nil {
		return eris.New(`nil message`)
	}

	if err := s.messagingProviderRepository.UpsertLastSeenMessage(ctx, "", "", ``); err != nil {
		// TODO: should we log or not?
	}

	state, err := s.tryParseStateFromContent(message.Content)
	if err != nil {
		return eris.Wrap(err, "try parse state from content")
	}

	switch state {
	case cancelled:
		if err := s.handleCancelledStateTransition(ctx, message); err != nil {
			return eris.Wrap(err, "handle cancelled state transition")
		}
	case collecting:
		if err := s.handleCollectingStateTransition(ctx, message); err != nil {
			return eris.Wrap(err, "handle collecting state transition")
		}
	case finished:
		if err := s.handleFinishedStateTransition(ctx, message); err != nil {
			return eris.Wrap(err, "handle finished state transition")
		}
	case ready:
		if err := s.handleReadyStateTransition(ctx, message); err != nil {
			return eris.Wrap(err, `handle ready state transition`)
		}
	default:
		s.logger.Fatalf(`Invalid state: %v`, state)
	}

	return nil
}

func (s *Service) handleReadyStateTransition(ctx context.Context, message *messaging.Message) error {
	if err := s.stateMachine.SetReady(); err != nil {
		return eris.Wrap(err, "set ready")
	}

	var entry = &Entry{}
	if err := s.replogRepository.CreateEntry(ctx, entry); err != nil {
	}

	return nil
}

func (s *Service) handleCollectingStateTransition(ctx context.Context, message *messaging.Message) error {
	if err := s.stateMachine.SetCollecting(); err != nil {
		return eris.Wrap(err, "set collecting")
	}

	s.replogRepository.AppendLinesToEntry()

	return nil
}

func (s *Service) handleCancelledStateTransition(ctx context.Context, _ *messaging.Message) error {
	if err := s.stateMachine.SetCancelled(); err != nil {
		return eris.Wrap(err, "set cancelled")
	}

	return nil
}

func (s *Service) handleFinishedStateTransition(ctx context.Context, _ *messaging.Message) error {
	if err := s.stateMachine.SetFinished(); err != nil {
		return eris.Wrap(err, "set finished")
	}

	return nil
}

func (s *Service) tryParseStateFromContent(_ string) (state, error) {
	return ready, eris.New("not implemented")
}
