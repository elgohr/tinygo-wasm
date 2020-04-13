package webauthn

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AuthenticatorID [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#type-AuthenticatorId
type AuthenticatorID string

// String returns the AuthenticatorID as string value.
func (t AuthenticatorID) String() string {
	return string(t)
}

// AuthenticatorProtocol [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#type-AuthenticatorProtocol
type AuthenticatorProtocol string

// String returns the AuthenticatorProtocol as string value.
func (t AuthenticatorProtocol) String() string {
	return string(t)
}

// AuthenticatorProtocol values.
const (
	AuthenticatorProtocolU2f   AuthenticatorProtocol = "u2f"
	AuthenticatorProtocolCtap2 AuthenticatorProtocol = "ctap2"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AuthenticatorProtocol) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AuthenticatorProtocol) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AuthenticatorProtocol) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AuthenticatorProtocol(in.String()) {
	case AuthenticatorProtocolU2f:
		*t = AuthenticatorProtocolU2f
	case AuthenticatorProtocolCtap2:
		*t = AuthenticatorProtocolCtap2

	default:
		in.AddError(errors.New("unknown AuthenticatorProtocol value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AuthenticatorProtocol) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AuthenticatorTransport [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#type-AuthenticatorTransport
type AuthenticatorTransport string

// String returns the AuthenticatorTransport as string value.
func (t AuthenticatorTransport) String() string {
	return string(t)
}

// AuthenticatorTransport values.
const (
	AuthenticatorTransportUsb      AuthenticatorTransport = "usb"
	AuthenticatorTransportNfc      AuthenticatorTransport = "nfc"
	AuthenticatorTransportBle      AuthenticatorTransport = "ble"
	AuthenticatorTransportCable    AuthenticatorTransport = "cable"
	AuthenticatorTransportInternal AuthenticatorTransport = "internal"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AuthenticatorTransport) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AuthenticatorTransport) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AuthenticatorTransport) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AuthenticatorTransport(in.String()) {
	case AuthenticatorTransportUsb:
		*t = AuthenticatorTransportUsb
	case AuthenticatorTransportNfc:
		*t = AuthenticatorTransportNfc
	case AuthenticatorTransportBle:
		*t = AuthenticatorTransportBle
	case AuthenticatorTransportCable:
		*t = AuthenticatorTransportCable
	case AuthenticatorTransportInternal:
		*t = AuthenticatorTransportInternal

	default:
		in.AddError(errors.New("unknown AuthenticatorTransport value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AuthenticatorTransport) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// VirtualAuthenticatorOptions [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#type-VirtualAuthenticatorOptions
type VirtualAuthenticatorOptions struct {
	Protocol                    AuthenticatorProtocol  `json:"protocol"`
	Transport                   AuthenticatorTransport `json:"transport"`
	HasResidentKey              bool                   `json:"hasResidentKey"`
	HasUserVerification         bool                   `json:"hasUserVerification"`
	AutomaticPresenceSimulation bool                   `json:"automaticPresenceSimulation,omitempty"` // If set to true, tests of user presence will succeed immediately. Otherwise, they will not be resolved. Defaults to true.
}

// Credential [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#type-Credential
type Credential struct {
	CredentialID string `json:"credentialId"`
	RpIDHash     string `json:"rpIdHash"`   // SHA-256 hash of the Relying Party ID the credential is scoped to. Must be 32 bytes long. See https://w3c.github.io/webauthn/#rpidhash
	PrivateKey   string `json:"privateKey"` // The private key in PKCS#8 format.
	SignCount    int64  `json:"signCount"`  // Signature counter. This is incremented by one for each successful assertion. See https://w3c.github.io/webauthn/#signature-counter
}
