package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"

	"github.com/go-leo/leo/common/stringx"
	"github.com/go-leo/leo/runner/net/http/internal/util"
)

const _MetadataHeaderPrefix = "grpc-metadata-"
const _MetadataTrailerPrefix = "grpc-trailer-"
const _XForwardedHostKey = "x-forwarded-host"
const _XForwardedForKey = "x-forwarded-for"
const _TimeoutKey = "grpc-timeout"

func newOutgoingContext(ctx context.Context, c *gin.Context) (context.Context, error) {
	var mdPairs []string
	for k, vv := range c.Request.Header {
		if len(vv) <= 0 {
			continue
		}
		k = strings.ToLower(k)
		if isRejectHeader(k) {
			continue
		}
		if isSpecialHandleHeader(k) {
			continue
		}
		for _, v := range vv {
			v, err := decodeMetadataHeader(k, v)
			if err != nil {
				return ctx, err
			}
			if isPermanentHTTPHeader(k) {
				k = "leo-" + k
			}
			k = strings.TrimPrefix(k, _MetadataHeaderPrefix)
			mdPairs = append(mdPairs, k, v)
		}
	}
	if k, v, ok := _XForwardedHost(c); ok {
		mdPairs = append(mdPairs, k, v)
	}
	if k, v, ok := _XForwardedFor(c); ok {
		mdPairs = append(mdPairs, k, v)
	}
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(mdPairs...))
	return ctx, nil
}

func _XForwardedHost(c *gin.Context) (string, string, bool) {
	host := c.GetHeader(_XForwardedHostKey)
	if stringx.IsNotBlank(host) {
		return strings.ToLower(_XForwardedHostKey), host, true
	} else if c.Request.Host != "" {
		return strings.ToLower(_XForwardedHostKey), c.Request.Host, true
	}
	return "", "", false
}

func _XForwardedFor(c *gin.Context) (string, string, bool) {
	remoteIP, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return "", "", false
	}
	forward := c.GetHeader(_XForwardedForKey)
	if stringx.IsBlank(forward) {
		return strings.ToLower(_XForwardedForKey), remoteIP, true
	}
	return strings.ToLower(_XForwardedForKey), fmt.Sprintf("%s, %s", forward, remoteIP), true
}

func _GRPCTimeout(c *gin.Context) (context.Context, context.CancelFunc, error) {
	v := c.GetHeader(_TimeoutKey)
	if stringx.IsBlank(v) {
		ctx, cancel := context.WithCancel(c.Request.Context())
		return ctx, cancel, nil
	}
	to, err := util.DecodeTimeout(v)
	if err != nil {
		ctx, cancel := context.WithCancel(c.Request.Context())
		return ctx, cancel, err
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), to)
	return ctx, cancel, nil
}

func isPermanentHTTPHeader(hdr string) bool {
	switch hdr {
	case
		"accept",
		"accept-charset",
		"accept-language",
		"accept-ranges",
		"authorization",
		"cache-control",
		"content-type",
		"cookie",
		"date",
		"expect",
		"from",
		"host",
		"if-match",
		"if-modified-since",
		"if-none-match",
		"if-schedule-tag-match",
		"if-unmodified-since",
		"max-forwards",
		"origin",
		"pragma",
		"referer",
		"user-agent",
		"via",
		"warning":
		return true
	}
	return false
}

func isRejectHeader(hdr string) bool {
	switch hdr {
	case "connection":
		return true
	default:
		return false
	}
}

func isSpecialHandleHeader(hdr string) bool {
	switch hdr {
	case _XForwardedHostKey, _XForwardedForKey, _TimeoutKey:
		return true
	default:
		return false
	}
}

func decodeMetadataHeader(k, v string) (string, error) {
	if strings.HasSuffix(k, "-bin") {
		b, err := decodeBinHeader(v)
		return string(b), err
	}
	return v, nil
}

func decodeBinHeader(v string) ([]byte, error) {
	if len(v)%4 == 0 {
		// Input was padded, or padding was not necessary.
		return base64.StdEncoding.DecodeString(v)
	}
	return base64.RawStdEncoding.DecodeString(v)
}

type Metadata struct {
	HeaderMD  metadata.MD
	TrailerMD metadata.MD
}

type metadataKey struct{}

func NewContextWithMetadata(ctx context.Context, md *Metadata) context.Context {
	return context.WithValue(ctx, metadataKey{}, md)
}

func MetadataFromContext(ctx context.Context) (md *Metadata, ok bool) {
	md, ok = ctx.Value(metadataKey{}).(*Metadata)
	return
}

func handleMetadata(c *gin.Context, md metadata.MD, prefix string) {
	for k, vs := range md {
		h := fmt.Sprintf("%s%s", prefix, k)
		for _, v := range vs {
			c.Writer.Header().Add(h, v)
		}
	}
}

func handleHeaderMetadata(c *gin.Context, md *Metadata) {
	handleMetadata(c, md.HeaderMD, _MetadataHeaderPrefix)
}

func handleTrailerMetadata(c *gin.Context, md *Metadata) {
	handleMetadata(c, md.TrailerMD, _MetadataTrailerPrefix)
}

func requestAcceptsTrailers(c *gin.Context) bool {
	te := c.Request.Header.Get("TE")
	return strings.Contains(strings.ToLower(te), "trailers")
}

func handleForwardResponseTrailerHeader(c *gin.Context, md *Metadata) {
	for k := range md.TrailerMD {
		c.Writer.Header().Add("Trailer", fmt.Sprintf("%s%s", _MetadataTrailerPrefix, k))
	}
}
