package signer

import (
	"context"
	"net/http"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/bnb-chain/greenfield-common/go/hash"
	storagetypes "github.com/bnb-chain/greenfield/x/storage/types"

	"github.com/bnb-chain/greenfield-storage-provider/base/gfspapp"
	"github.com/bnb-chain/greenfield-storage-provider/base/types/gfsperrors"
	"github.com/bnb-chain/greenfield-storage-provider/base/types/gfspp2p"
	"github.com/bnb-chain/greenfield-storage-provider/core/module"
	"github.com/bnb-chain/greenfield-storage-provider/core/rcmgr"
	"github.com/bnb-chain/greenfield-storage-provider/core/task"
	"github.com/bnb-chain/greenfield-storage-provider/pkg/metrics"
	virtualgrouptypes "github.com/bnb-chain/greenfield/x/virtualgroup/types"
)

var (
	ErrSignMsg                      = gfsperrors.Register(module.SignModularName, http.StatusBadRequest, 120001, "sign message with private key failed")
	ErrSealObjectOnChain            = gfsperrors.Register(module.SignModularName, http.StatusBadRequest, 120002, "send sealObject msg failed")
	ErrRejectUnSealObjectOnChain    = gfsperrors.Register(module.SignModularName, http.StatusBadRequest, 120003, "send rejectUnSealObject msg failed")
	ErrDiscontinueBucketOnChain     = gfsperrors.Register(module.SignModularName, http.StatusBadRequest, 120004, "send discontinueBucket msg failed")
	ErrDanglingPointer              = gfsperrors.Register(module.SignModularName, http.StatusBadRequest, 120005, "sign or tx msg pointer dangling")
	ErrCreateGVGOnChain             = gfsperrors.Register(module.SignModularName, http.StatusBadRequest, 120006, "send create gvg msg failed")
	ErrCompleteMigrateBucketOnChain = gfsperrors.Register(module.SignModularName, http.StatusBadRequest, 120007, "send complete migrate bucket msg failed")
)

var _ module.Signer = &SignModular{}

type SignModular struct {
	baseApp *gfspapp.GfSpBaseApp
	client  *GreenfieldChainSignClient
}

func (s *SignModular) Name() string {
	return module.SignModularName
}

func (s *SignModular) Start(ctx context.Context) error {
	return nil
}

func (s *SignModular) Stop(ctx context.Context) error {
	return nil
}

func (s *SignModular) ReserveResource(ctx context.Context, state *rcmgr.ScopeStat) (
	rcmgr.ResourceScopeSpan, error) {
	return &rcmgr.NullScope{}, nil
}

func (s *SignModular) ReleaseResource(ctx context.Context, span rcmgr.ResourceScopeSpan) {
	span.Done()
}

func (s *SignModular) SignCreateBucketApproval(ctx context.Context, bucket *storagetypes.MsgCreateBucket) ([]byte, error) {
	msg := bucket.GetApprovalBytes()
	sig, err := s.client.Sign(SignApproval, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *SignModular) SignCreateObjectApproval(ctx context.Context, object *storagetypes.MsgCreateObject) ([]byte, error) {
	msg := object.GetApprovalBytes()
	sig, err := s.client.Sign(SignApproval, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *SignModular) SignReplicatePieceApproval(ctx context.Context, task task.ApprovalReplicatePieceTask) ([]byte, error) {
	msg := task.GetSignBytes()
	sig, err := s.client.Sign(SignOperator, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *SignModular) SignReceivePieceTask(ctx context.Context, task task.ReceivePieceTask) ([]byte, error) {
	msg := task.GetSignBytes()
	sig, err := s.client.Sign(SignOperator, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *SignModular) SignSecondarySealBls(ctx context.Context, objectID uint64, gvgId uint32, checksums [][]byte) ([]byte, error) {
	msg := storagetypes.NewSecondarySpSealObjectSignDoc(sdkmath.NewUint(objectID), gvgId, hash.GenerateIntegrityHash(checksums)).GetBlsSignHash()
	sig, err := s.client.blsKm.Sign(msg[:])
	if err != nil {
		return nil, err
	}
	// log.Debugw("bls signature length", "len", len(sig), "object_id", objectID, "gvg_id", gvgId, "checksums", checksums)
	return sig, nil
}

func (s *SignModular) SignSecondaryCompleteMigrationBls(_ context.Context, bucketId uint64, spId, srcGvgId, destGvgId uint32) ([]byte, error) {
	msg := storagetypes.NewSecondarySpMigrationBucketSignDoc(sdkmath.NewUint(bucketId), spId, srcGvgId, destGvgId).GetBlsSignHash()
	sig, err := s.client.blsKm.Sign(msg[:])
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *SignModular) SignP2PPingMsg(ctx context.Context, ping *gfspp2p.GfSpPing) ([]byte, error) {
	msg := ping.GetSignBytes()
	sig, err := s.client.Sign(SignOperator, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *SignModular) SignP2PPongMsg(ctx context.Context, pong *gfspp2p.GfSpPong) ([]byte, error) {
	msg := pong.GetSignBytes()
	sig, err := s.client.Sign(SignOperator, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *SignModular) SealObject(ctx context.Context, object *storagetypes.MsgSealObject) error {
	var (
		err       error
		startTime = time.Now()
	)
	defer func() {
		metrics.SealObjectTimeHistogram.WithLabelValues(s.Name()).Observe(time.Since(startTime).Seconds())
	}()
	_, err = s.client.SealObject(ctx, SignSeal, object)
	return err
}

func (s *SignModular) RejectUnSealObject(ctx context.Context, rejectObject *storagetypes.MsgRejectSealObject) error {
	var (
		err       error
		startTime = time.Now()
	)
	defer func() {
		metrics.RejectUnSealObjectTimeHistogram.WithLabelValues(s.Name()).Observe(time.Since(startTime).Seconds())
		if err == nil {
			metrics.RejectUnSealObjectSucceedCounter.WithLabelValues(s.Name()).Inc()
		} else {
			metrics.RejectUnSealObjectFailedCounter.WithLabelValues(s.Name()).Inc()
		}
	}()
	_, err = s.client.RejectUnSealObject(ctx, SignSeal, rejectObject)
	return err
}

func (s *SignModular) DiscontinueBucket(ctx context.Context, bucket *storagetypes.MsgDiscontinueBucket) error {
	var (
		err       error
		startTime = time.Now()
	)
	defer func() {
		metrics.DiscontinueBucketTimeHistogram.WithLabelValues(s.Name()).Observe(time.Since(startTime).Seconds())
		if err == nil {
			metrics.DiscontinueBucketSucceedCounter.WithLabelValues(s.Name()).Inc()
		} else {
			metrics.DiscontinueBucketFailedCounter.WithLabelValues(s.Name()).Inc()
		}
	}()
	_, err = s.client.DiscontinueBucket(ctx, SignGc, bucket)
	return err
}

func (s *SignModular) CreateGlobalVirtualGroup(ctx context.Context, gvg *virtualgrouptypes.MsgCreateGlobalVirtualGroup) error {
	var (
		err       error
		startTime = time.Now()
	)
	defer func() {
		metrics.CreateGlobalVirtualGroupTimeHistogram.WithLabelValues(s.Name()).Observe(time.Since(startTime).Seconds())
		if err == nil {
			metrics.CreateGlobalVirtualGroupSucceedCounter.WithLabelValues(s.Name()).Inc()
		} else {
			metrics.CreateGlobalVirtualGroupFailedCounter.WithLabelValues(s.Name()).Inc()
		}
	}()
	_, err = s.client.CreateGlobalVirtualGroup(ctx, SignOperator, gvg)
	return err
}

func (s *SignModular) CompleteMigrateBucket(ctx context.Context, completeMigrateBucket *storagetypes.MsgCompleteMigrateBucket) error {
	var (
		err       error
		startTime = time.Now()
	)
	defer func() {
		metrics.CompleteMigrateBucketTimeHistogram.WithLabelValues(s.Name()).Observe(time.Since(startTime).Seconds())
		if err == nil {
			metrics.CompleteMigrateBucketSucceedCounter.WithLabelValues(s.Name()).Inc()
		} else {
			metrics.CompleteMigrateBucketFailedCounter.WithLabelValues(s.Name()).Inc()
		}
	}()
	_, err = s.client.CompleteMigrateBucket(ctx, SignOperator, completeMigrateBucket)
	return err
}
