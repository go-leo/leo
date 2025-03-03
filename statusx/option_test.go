package statusx_test

import (
	"github.com/go-leo/leo/v3/statusx"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"net/http"
	"testing"
	"time"
)

func TestMessage(t *testing.T) {
	text := "this is ok"
	st := statusx.OK(statusx.Message(text))
	assert.Equal(t, st.Message(), text)
}

func TestHeaders(t *testing.T) {
	header := http.Header{"X-Request-Id": {"1234567890"}}
	st := statusx.Canceled(statusx.Headers(header))
	assert.Equal(t, st.Headers(), header)
}

func TestErrorInfo(t *testing.T) {
	reason := "username is empty"
	domain := "login"
	metadata := map[string]string{"id": "1234567890"}
	st := statusx.InvalidArgument(statusx.ErrorInfo(reason, domain, metadata))
	assert.Equal(t, reason, st.ErrorInfo().GetReason())
	assert.Equal(t, domain, st.ErrorInfo().GetDomain())
	assert.Equal(t, metadata, st.ErrorInfo().GetMetadata())
}

func TestRetryInfo(t *testing.T) {
	retryDelay := time.Second
	st := statusx.Unknown(statusx.RetryInfo(retryDelay))
	assert.Equal(t, retryDelay, st.RetryInfo().GetRetryDelay().AsDuration())
}

func TestDebugInfo(t *testing.T) {
	stackEntries := []string{
		"line 1",
		"line 2",
		"line 3",
	}
	detail := "user exists"
	st := statusx.AlreadyExists(statusx.DebugInfo(stackEntries, detail))
	assert.Equal(t, stackEntries, st.DebugInfo().GetStackEntries())
	assert.Equal(t, detail, st.DebugInfo().GetDetail())
}

func TestQuotaFailure(t *testing.T) {
	violations := []*errdetails.QuotaFailure_Violation{
		{
			Subject:     "123.111.222.444",
			Description: "Service disabled",
		},
	}
	st := statusx.ResourceExhausted(statusx.QuotaFailure(violations))
	assert.Equal(t, len(violations), len(st.QuotaFailure().GetViolations()))
	assert.Equal(t, violations[0].GetSubject(), st.QuotaFailure().GetViolations()[0].GetSubject())
	assert.Equal(t, violations[0].GetDescription(), st.QuotaFailure().GetViolations()[0].GetDescription())
}

func TestPreconditionFailure(t *testing.T) {
	violations := []*errdetails.PreconditionFailure_Violation{
		{
			Type:        "TOS",
			Subject:     "123.111.222.444",
			Description: "Service disabled",
		},
	}
	st := statusx.FailedPrecondition(statusx.PreconditionFailure(violations))
	assert.Equal(t, len(violations), len(st.PreconditionFailure().GetViolations()))
	assert.Equal(t, violations[0].GetType(), st.PreconditionFailure().GetViolations()[0].GetType())
	assert.Equal(t, violations[0].GetSubject(), st.PreconditionFailure().GetViolations()[0].GetSubject())
	assert.Equal(t, violations[0].GetDescription(), st.PreconditionFailure().GetViolations()[0].GetDescription())
}

func TestBadRequest(t *testing.T) {
	violations := []*errdetails.BadRequest_FieldViolation{
		{
			Field:       "username",
			Description: "Service disabled",
			Reason:      "is empty",
			LocalizedMessage: &errdetails.LocalizedMessage{
				Locale:  "zh-CN",
				Message: "用户名不能为空",
			},
		},
	}
	st := statusx.InvalidArgument(statusx.BadRequest(violations))
	assert.Equal(t, len(violations), len(st.BadRequest().GetFieldViolations()))
	assert.Equal(t, violations[0].GetField(), st.BadRequest().GetFieldViolations()[0].GetField())
	assert.Equal(t, violations[0].GetDescription(), st.BadRequest().GetFieldViolations()[0].GetDescription())
	assert.Equal(t, violations[0].GetReason(), st.BadRequest().GetFieldViolations()[0].GetReason())
	assert.Equal(t, violations[0].GetLocalizedMessage(), st.BadRequest().GetFieldViolations()[0].GetLocalizedMessage())
}

func TestRequestInfo(t *testing.T) {
	requestId := ""
	servingData := ""
	st := statusx.PermissionDenied(statusx.RequestInfo(requestId, servingData))
	assert.Equal(t, requestId, st.RequestInfo().GetRequestId())
	assert.Equal(t, servingData, st.RequestInfo().GetServingData())
}

func TestResourceInfo(t *testing.T) {
	resourceType := "resourceType"
	resourceName := "resourceName"
	owner := "owner"
	description := "description"
	st := statusx.Aborted(statusx.ResourceInfo(resourceType, resourceName, owner, description))
	assert.Equal(t, resourceType, st.ResourceInfo().GetResourceType())
	assert.Equal(t, resourceName, st.ResourceInfo().GetResourceName())
	assert.Equal(t, owner, st.ResourceInfo().GetOwner())
	assert.Equal(t, description, st.ResourceInfo().GetDescription())
}

func TestHelp(t *testing.T) {
	links := []*errdetails.Help_Link{
		{
			Description: "Description",
			Url:         "Url",
		},
	}
	st := statusx.OutOfRange(statusx.Help(links))
	assert.Equal(t, len(links), len(st.Help().GetLinks()))
	assert.Equal(t, links[0].GetDescription(), st.Help().GetLinks()[0].GetDescription())
	assert.Equal(t, links[0].GetUrl(), st.Help().GetLinks()[0].GetUrl())
}

func TestLocalizedMessage(t *testing.T) {
	locale := "locale"
	message := "message"
	st := statusx.DataLoss(statusx.LocalizedMessage(locale, message))
	assert.Equal(t, locale, st.LocalizedMessage().GetLocale())
	assert.Equal(t, message, st.LocalizedMessage().GetMessage())
}

func TestDetail(t *testing.T) {
	now := time.Now()
	detail := &timeofday.TimeOfDay{
		Hours:   int32(now.Hour()),
		Minutes: int32(now.Minute()),
		Seconds: int32(now.Second()),
		Nanos:   int32(now.Nanosecond()),
	}
	st := statusx.Unimplemented(statusx.Extra(detail))
	assert.Equal(t, detail.GetHours(), st.Extra().(*timeofday.TimeOfDay).GetHours())
	assert.Equal(t, detail.GetMinutes(), st.Extra().(*timeofday.TimeOfDay).GetMinutes())
	assert.Equal(t, detail.GetSeconds(), st.Extra().(*timeofday.TimeOfDay).GetSeconds())
	assert.Equal(t, detail.GetNanos(), st.Extra().(*timeofday.TimeOfDay).GetNanos())
}
