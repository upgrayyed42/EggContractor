package api

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/protobuf/proto"
)

const (
	AppVersion     = "1.21.2"
	AppBuild       = "1.21.2.1"
	ClientVersion  = 34
	PlatformString = "IOS"
)

var _apiPrefix = "https://www.auxbrain.com"

var _client *http.Client

func init() {
	if runtime.GOOS == "js" && runtime.GOARCH == "wasm" {
		// Use CORS proxy in the browser setting.
		_apiPrefix = "https://wasmegg.zw.workers.dev/?url=https://www.auxbrain.com"
	}
	_client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func Request(endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	return RequestWithContext(context.Background(), endpoint, reqMsg, respMsg)
}

func RequestWithContext(ctx context.Context, endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	return doRequestWithContext(ctx, endpoint, reqMsg, respMsg, false)
}

func RequestAuthenticated(endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	return RequestAuthenticatedWithContext(context.Background(), endpoint, reqMsg, respMsg)
}

func RequestAuthenticatedWithContext(ctx context.Context, endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	return doRequestWithContext(ctx, endpoint, reqMsg, respMsg, true)
}

func RequestFirstContact(payload *FirstContactRequestPayload) (*FirstContact, error) {
	return RequestFirstContactWithContext(context.Background(), payload)
}

func RequestFirstContactWithContext(ctx context.Context, payload *FirstContactRequestPayload) (*FirstContact, error) {
	if payload.ClientVersion == 0 {
		payload.ClientVersion = ClientVersion
	}
	resp := &FirstContact{}
	err := RequestAuthenticatedWithContext(ctx, "/ei/first_contact", payload, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func RequestCoopStatus(payload *CoopStatusRequestPayload) (*CoopStatus, error) {
	return RequestCoopStatusWithContext(context.Background(), payload)
}

func RequestCoopStatusWithContext(ctx context.Context, payload *CoopStatusRequestPayload) (*CoopStatus, error) {
	resp := &CoopStatus{}
	err := RequestAuthenticatedWithContext(ctx, "/ei/coop_status", payload, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func RequestPeriodicals(payload *GetPeriodicalsRequestPayload) (*Periodicals, error) {
	return RequestPeriodicalsWithContext(context.Background(), payload)
}

func RequestPeriodicalsWithContext(ctx context.Context, payload *GetPeriodicalsRequestPayload) (*Periodicals, error) {
	if payload.CurrentClientVersion == 0 {
		payload.CurrentClientVersion = ClientVersion
	}
	resp := &Periodicals{}
	err := RequestAuthenticatedWithContext(ctx, "/ei/get_periodicals", payload, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewBasicRequestInfo(userId string) *BasicRequestInfo {
	return &BasicRequestInfo{
		EiUserId:      userId,
		ClientVersion: ClientVersion,
		Version:       AppVersion,
		Build:         AppBuild,
		Platform:      PlatformString,
	}
}

func doRequestWithContext(ctx context.Context, endpoint string, reqMsg proto.Message, respMsg proto.Message, authenticated bool) error {
	apiUrl := _apiPrefix + endpoint
	reqBin, err := proto.Marshal(reqMsg)
	if err != nil {
		return errors.Wrapf(err, "marshaling payload %+v for %s", reqMsg, apiUrl)
	}
	enc := base64.StdEncoding
	reqDataEncoded := enc.EncodeToString(reqBin)
	log.Infof("POST %s: %+v", apiUrl, reqMsg)
	log.Debugf("POST %s data=%s", apiUrl, reqDataEncoded)
	resp, err := ctxhttp.PostForm(ctx, _client, apiUrl, url.Values{"data": {reqDataEncoded}})
	if err != nil {
		return errors.Wrapf(err, "POST %s", apiUrl)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "POST %s", apiUrl)
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return errors.Errorf("POST %s: HTTP %d: %#v", apiUrl, resp.StatusCode, string(body))
	}
	return DecodeAPIResponse(apiUrl, body, respMsg, authenticated)
}

func DecodeAPIResponse(apiUrl string, payload []byte, msg proto.Message, authenticated bool) error {
	enc := base64.StdEncoding
	buf := make([]byte, enc.DecodedLen(len(payload)))
	n, err := enc.Decode(buf, payload)
	if err != nil {
		return errors.Wrapf(err, "base64 decoding %s reponse (%#v)", apiUrl, string(payload))
	}
	if authenticated {
		authMsg := &AuthenticatedMessage{}
		err = proto.Unmarshal(buf[:n], authMsg)
		if err != nil {
			err = errors.Wrapf(err, "unmarshaling %s response as AuthenticatedMessage (%#v)", apiUrl, string(payload))
			return interpretUnmarshalError(err)
		}
		err = proto.Unmarshal(authMsg.Message, msg)
		if err != nil {
			err = errors.Wrapf(err, "unmarshaling AuthenticatedMessage payload in %s response (%#v)", apiUrl, string(payload))
			return interpretUnmarshalError(err)
		}
	} else {
		err = proto.Unmarshal(buf[:n], msg)
		if err != nil {
			err = errors.Wrapf(err, "unmarshaling %s response (%#v)", apiUrl, string(payload))
			return interpretUnmarshalError(err)
		}
	}
	return nil
}

func interpretUnmarshalError(err error) error {
	if strings.Contains(err.Error(), "contains invalid UTF-8") {
		return errors.Wrap(err, "API returned corrupted data (invalid UTF-8 in one or more string fields); "+
			"this is a known issue affecting some players, and it can only be resolved when Auxbrain fixes their server bug")
	}
	return err
}
