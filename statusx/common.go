package statusx

import (
	"github.com/go-leo/leo/v3/proto/leo/status"
	"github.com/go-leo/leo/v3/statusx/internal/statuspb"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	kKey       = "X-Leo-Status-Key"
	kSeparator = ", "
)

func marshalHttpBody(st *sampleStatus) ([]byte, error) {
	body := &status.HttpBody{
		Error: &status.HttpBody_Status{
			Status:     code.Code(st.Code()),
			Message:    st.Message(),
			Code:       int32(st.StatusCode()),
			Identifier: st.Identifier(),
			Details:    statuspb.ToHttpDetails(st.err.GetDetailInfo()),
		},
	}
	return protojson.MarshalOptions{}.Marshal(body)
}

func unmarshalHttpBody(data []byte) (*status.HttpBody, error) {
	body := &status.HttpBody{}
	err := protojson.UnmarshalOptions{}.Unmarshal(data, body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
