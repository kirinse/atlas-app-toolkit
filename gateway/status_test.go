package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"testing"
)

func TestStatus(t *testing.T) {
	// test REST status from gRPC one
	stat, statName := HTTPStatus(context.Background(), status.New(codes.OK, "success message"))

	if stat != http.StatusOK {
		t.Errorf("invalid http status code %d - expected: %d", stat, http.StatusOK)
	}
	if statName != codes.OK.String() {
		t.Errorf("invalid http status codename %q - expected: %q", statName, codes.OK.String())
	}

	// test REST status from incoming context
	md := metadata.Pairs(
		runtime.MetadataPrefix+"status-code", CodeName(Created),
	)
	ctx := metadata.NewIncomingContext(context.Background(), md)
	stat, statName = HTTPStatus(ctx, nil)

	if stat != http.StatusCreated {
		t.Errorf("invalid http status code %d - expected: %d", stat, http.StatusCreated)
	}
	if statName != CodeName(Created) {
		t.Errorf("invalid http status codename %q - expected: %q", statName, codes.OK.String())
	}
}

func TestCodeName(t *testing.T) {
	// test renamed code
	if cn := CodeName(codes.Unimplemented); cn != "NOT_IMPLEMENTED" {
		t.Errorf("invalid code name: %s - expected: %s", cn, "NOT_IMPLEMENTED")
	}

	// test custom code
	if cn := CodeName(LongRunning); cn != "LONG_RUNNING_OP" {
		t.Errorf("invalid code name: %s - expected: %s", cn, "LONG_RUNNING_OP")
	}

	// test standard code
	if cn := CodeName(codes.OutOfRange); cn != code.Code_name[int32(code.Code_OUT_OF_RANGE)] {
		t.Errorf("invalid code name: %s - expected: %s", cn, code.Code_name[int32(code.Code_OUT_OF_RANGE)])
	}
}

func TestCode(t *testing.T) {
	// test renamed code
	if c := Code("NOT_IMPLEMENTED"); c != codes.Unimplemented {
		t.Errorf("invalid code: %s - expected: %s", c, codes.Unimplemented)
	}
	// test custom code
	if c := Code("LONG_RUNNING_OP"); c != LongRunning {
		t.Errorf("invalid code: %s - expected: %s", c, LongRunning)
	}
	// test standard code
	if c := Code(code.Code_name[int32(code.Code_OUT_OF_RANGE)]); c != codes.OutOfRange {
		t.Errorf("invalid code: %s - expected: %s", c, codes.OutOfRange)
	}
}

func TestHTTPStatusFromCode(t *testing.T) {
	// test overwritten code
	if sc := HTTPStatusFromCode(codes.Canceled); sc != 499 {
		t.Errorf("invalid http status: %d - expected: %d", sc, 499)
	}
	// test custom code
	if sc := HTTPStatusFromCode(Created); sc != http.StatusCreated {
		t.Errorf("invalid http status: %d - expected: %d", sc, http.StatusCreated)
	}
	// test standard code
	if sc := HTTPStatusFromCode(codes.NotFound); sc != http.StatusNotFound {
		t.Errorf("invalid http status: %d - expected: %d", sc, http.StatusNotFound)
	}
}

func TestOldStatusCreated(t *testing.T) {
	OldStatusCreatedOnUpdate = true
	s := HTTPStatusFromCode(Updated)
	if s != http.StatusCreated {
		t.Error("if OldStatusCreatedOnUpdate is false true should be StatusCreated")
	}
	OldStatusCreatedOnUpdate = false
	s = HTTPStatusFromCode(Updated)
	if s != http.StatusOK {
		t.Error("if OldStatusCreatedOnUpdate is false status should be StatusOk")
	}
}
