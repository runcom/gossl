package gossl

const (
	AD_REASON_OFFSET = 1000 /* offset to get R_... value from AD_... */

	ERROR_NONE             = 0
	ERROR_SSL              = 1
	ERROR_WANT_READ        = 2
	ERROR_WANT_WRITE       = 3
	ERROR_WANT_X509_LOOKUP = 4
	ERROR_SYSCALL          = 5 /* look at error stack/return value/errno */
	ERROR_ZERO_RETURN      = 6
	ERROR_WANT_CONNECT     = 7
	ERROR_WANT_ACCEPT      = 8

	CTRL_NEED_TMP_RSA    = 1
	CTRL_SET_TMP_RSA     = 2
	CTRL_SET_TMP_DH      = 3
	CTRL_SET_TMP_ECDH    = 4
	CTRL_SET_TMP_RSA_CB  = 5
	CTRL_SET_TMP_DH_CB   = 6
	CTRL_SET_TMP_ECDH_CB = 7

	CTRL_GET_SESSION_REUSED       = 8
	CTRL_GET_CLIENT_CERT_REQUEST  = 9
	CTRL_GET_NUM_RENEGOTIATIONS   = 10
	CTRL_CLEAR_NUM_RENEGOTIATIONS = 11
	CTRL_GET_TOTAL_RENEGOTIATIONS = 12
	CTRL_GET_FLAGS                = 13
	CTRL_EXTRA_CHAIN_CERT         = 14

	CTRL_SET_MSG_CALLBACK     = 15
	CTRL_SET_MSG_CALLBACK_ARG = 16

	/* only applies to datagram connections */
	CTRL_SET_MTU = 17
	/* Stats */
	CTRL_SESS_NUMBER              = 20
	CTRL_SESS_CONNECT             = 21
	CTRL_SESS_CONNECT_GOOD        = 22
	CTRL_SESS_CONNECT_RENEGOTIATE = 23
	CTRL_SESS_ACCEPT              = 24
	CTRL_SESS_ACCEPT_GOOD         = 25
	CTRL_SESS_ACCEPT_RENEGOTIATE  = 26
	CTRL_SESS_HIT                 = 27
	CTRL_SESS_CB_HIT              = 28
	CTRL_SESS_MISSES              = 29
	CTRL_SESS_TIMEOUTS            = 30
	CTRL_SESS_CACHE_FULL          = 31
	CTRL_OPTIONS                  = 32
	CTRL_MODE                     = 33

	CTRL_GET_READ_AHEAD      = 40
	CTRL_SET_READ_AHEAD      = 41
	CTRL_SET_SESS_CACHE_SIZE = 42
	CTRL_GET_SESS_CACHE_SIZE = 43
	CTRL_SET_SESS_CACHE_MODE = 44
	CTRL_GET_SESS_CACHE_MODE = 45

	CTRL_GET_MAX_CERT_LIST = 50
	CTRL_SET_MAX_CERT_LIST = 51

	CTRL_SET_MAX_SEND_FRAGMENT = 52

	/* see tls1.h for macros based on these */
	CTRL_SET_TLSEXT_SERVERNAME_CB           = 53
	CTRL_SET_TLSEXT_SERVERNAME_ARG          = 54
	CTRL_SET_TLSEXT_HOSTNAME                = 55
	CTRL_SET_TLSEXT_DEBUG_CB                = 56
	CTRL_SET_TLSEXT_DEBUG_ARG               = 57
	CTRL_GET_TLSEXT_TICKET_KEYS             = 58
	CTRL_SET_TLSEXT_TICKET_KEYS             = 59
	CTRL_SET_TLSEXT_OPAQUE_PRF_INPUT        = 60
	CTRL_SET_TLSEXT_OPAQUE_PRF_INPUT_CB     = 61
	CTRL_SET_TLSEXT_OPAQUE_PRF_INPUT_CB_ARG = 62
	CTRL_SET_TLSEXT_STATUS_REQ_CB           = 63
	CTRL_SET_TLSEXT_STATUS_REQ_CB_ARG       = 64
	CTRL_SET_TLSEXT_STATUS_REQ_TYPE         = 65
	CTRL_GET_TLSEXT_STATUS_REQ_EXTS         = 66
	CTRL_SET_TLSEXT_STATUS_REQ_EXTS         = 67
	CTRL_GET_TLSEXT_STATUS_REQ_IDS          = 68
	CTRL_SET_TLSEXT_STATUS_REQ_IDS          = 69
	CTRL_GET_TLSEXT_STATUS_REQ_OCSP_RESP    = 70
	CTRL_SET_TLSEXT_STATUS_REQ_OCSP_RESP    = 71

	CTRL_SET_TLSEXT_TICKET_KEY_CB = 72

	CTRL_SET_TLS_EXT_SRP_USERNAME_CB = 75
	CTRL_SET_SRP_VERIFY_PARAM_CB     = 76
	CTRL_SET_SRP_GIVE_CLIENT_PWD_CB  = 77

	CTRL_SET_SRP_ARG                       = 78
	CTRL_SET_TLS_EXT_SRP_USERNAME          = 79
	CTRL_SET_TLS_EXT_SRP_STRENGTH          = 80
	CTRL_SET_TLS_EXT_SRP_PASSWORD          = 81
	CTRL_TLS_EXT_SEND_HEARTBEAT            = 85
	CTRL_GET_TLS_EXT_HEARTBEAT_PENDING     = 86
	CTRL_SET_TLS_EXT_HEARTBEAT_NO_REQUESTS = 87

	DTLS_CTRL_GET_TIMEOUT    = 73
	DTLS_CTRL_HANDLE_TIMEOUT = 74
	DTLS_CTRL_LISTEN         = 75

	CTRL_GET_RI_SUPPORT = 76
	CTRL_CLEAR_OPTIONS  = 77
	CTRL_CLEAR_MODE     = 78

	CTRL_GET_EXTRA_CHAIN_CERTS   = 82
	CTRL_CLEAR_EXTRA_CHAIN_CERTS = 83
	OP_NO_COMPRESSION            = 0x00020000
	FILETYPE_ASN1                = 2
	FILETYPE_PEM                 = 1
)