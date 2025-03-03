package statusx_test

import (
	"errors"
	"github.com/go-leo/leo/v3/proto/leo/status"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatusx "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/genproto/googleapis/type/phone_number"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"testing"
)

func TestCode(t *testing.T) {
	st := statusx.Internal()
	assert.Equal(t, codes.Internal, st.Code())
}

func TestGRPCStatus(t *testing.T) {
	text := "password is invalid"
	domain := "account"
	metadata := map[string]string{"username": "leo"}
	number := "13013013013"
	key := "WWW-Authenticate"
	value := "Basic realm=xxx"
	st := statusx.Unauthenticated(
		statusx.Message(text),
		statusx.ErrorInfo(text, domain, metadata),
		statusx.Headers(http.Header{key: {value}}),
		statusx.Extra(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)
	grpcStatus := st.GRPCStatus()
	assert.Equal(t, codes.Unauthenticated, grpcStatus.Code())
	assert.Equal(t, text, grpcStatus.Message())
	details := grpcStatus.Details()
	assert.Equal(t, 3, len(details))

	errorInfo := details[0].(*errdetails.ErrorInfo)
	assert.Equal(t, text, errorInfo.GetReason())
	assert.Equal(t, domain, errorInfo.GetDomain())
	assert.Equal(t, metadata, errorInfo.GetMetadata())

	phoneNumber := details[1].(*phone_number.PhoneNumber)
	assert.Equal(t, number, phoneNumber.GetE164Number())

	response := details[2].(*httpstatusx.HttpResponse)
	assert.Equal(t, http.StatusUnauthorized, int(response.GetStatus()))
	assert.Equal(t, key, response.GetHeaders()[0].GetKey())
	assert.Equal(t, value, response.GetHeaders()[0].GetValue())
}

func TestHTTPStatus(t *testing.T) {
	text := "password is invalid"
	domain := "account"
	metadata := map[string]string{"username": "leo"}
	number := "13013013013"
	key := "WWW-Authenticate"
	value := "Basic realm=xxx"
	st := statusx.Unauthenticated(
		statusx.Message(text),
		statusx.ErrorInfo(text, domain, metadata),
		statusx.Headers(http.Header{key: {value}}),
		statusx.Extra(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)

	data, err := st.MarshalJSON()
	assert.NoError(t, err)
	body := &status.HttpBody{}
	err = protojson.Unmarshal(data, body)
	assert.NoErrorf(t, err, "unmarshal http body")
	assert.Equal(t, text, body.GetError().GetMessage())
	assert.Equal(t, int(codes.Unauthenticated), int(body.GetError().GetStatus()))
	assert.Equal(t, http.StatusUnauthorized, int(body.GetError().GetCode()))

	details := body.GetError().GetDetails()
	d1, err := details[0].UnmarshalNew()
	assert.NoError(t, err, "UnmarshalNew")
	errorInfo := d1.(*errdetails.ErrorInfo)
	assert.Equal(t, text, errorInfo.GetReason())
	assert.Equal(t, domain, errorInfo.GetDomain())
	assert.Equal(t, metadata, errorInfo.GetMetadata())

	d2, err := details[1].UnmarshalNew()
	assert.NoError(t, err, "UnmarshalNew")
	phoneNumber := d2.(*phone_number.PhoneNumber)
	assert.Equal(t, number, phoneNumber.GetE164Number())

}

func TestIs(t *testing.T) {
	text := "password is invalid"
	domain := "account"
	metadata := map[string]string{"username": "leo"}
	number := "13013013013"
	key := "WWW-Authenticate"
	value := "Basic realm=xxx"
	st1 := statusx.Unauthenticated(
		statusx.Message(text),
		statusx.ErrorInfo(text, domain, metadata),
		statusx.Headers(http.Header{key: {value}}),
		statusx.Extra(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)

	st2 := statusx.Unauthenticated(
		statusx.Message(text),
		statusx.ErrorInfo(text, domain, metadata),
		statusx.Headers(http.Header{key: {value}}),
		statusx.Extra(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)

	assert.True(t, errors.Is(st1, st2))

	st3 := statusx.Internal(
		statusx.Message(text),
		statusx.ErrorInfo(text, domain, metadata),
		statusx.Headers(http.Header{key: {value}}),
		statusx.Extra(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)
	assert.False(t, errors.Is(st1, st3))
}
