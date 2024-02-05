package types

import (
	// "fmt"
	// "reflect"
	sdkerrors "cosmossdk.io/errors"
	// errorsmod "cosmossdk.io/errors"
	// sdk "github.com/cosmos/cosmos-sdk/types"
    // sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// "github.com/pkg/errors"
	// grpccodes "google.golang.org/grpc/codes"
	// grpcstatus "google.golang.org/grpc/status"
	
)

var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
    ErrWrongLoanState = sdkerrors.Register(ModuleName, 2, "wrong loan state")
	ErrDeadline = sdkerrors.Register(ModuleName, 3, "deadline")
)
