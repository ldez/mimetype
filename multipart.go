package mimetype

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

const (
	// MultipartAlternative alternative mime type.
	MultipartAlternative = "multipart/alternative"
	// MultipartAppledouble appledouble mime type.
	MultipartAppledouble = "multipart/appledouble"
	// MultipartByteranges byteranges mime type.
	MultipartByteranges = "multipart/byteranges"
	// MultipartDigest digest mime type.
	MultipartDigest = "multipart/digest"
	// MultipartEncrypted encrypted mime type.
	MultipartEncrypted = "multipart/encrypted"
	// MultipartExample example mime type.
	MultipartExample = "multipart/example"
	// MultipartFormData form-data mime type.
	MultipartFormData = "multipart/form-data"
	// MultipartHeaderSet header-set mime type.
	MultipartHeaderSet = "multipart/header-set"
	// MultipartMixed mixed mime type.
	MultipartMixed = "multipart/mixed"
	// MultipartMultilingual multilingual mime type.
	MultipartMultilingual = "multipart/multilingual"
	// MultipartParallel parallel mime type.
	MultipartParallel = "multipart/parallel"
	// MultipartRelated related mime type.
	MultipartRelated = "multipart/related"
	// MultipartReport report mime type.
	MultipartReport = "multipart/report"
	// MultipartSigned signed mime type.
	MultipartSigned = "multipart/signed"
	// MultipartVndBintMedPlus vnd.bint.med-plus mime type.
	MultipartVndBintMedPlus = "multipart/vnd.bint.med-plus"
	// MultipartVoiceMessage voice-message mime type.
	MultipartVoiceMessage = "multipart/voice-message"
	// MultipartXMixedReplace x-mixed-replace mime type.
	MultipartXMixedReplace = "multipart/x-mixed-replace"
)

// IsMultipart checks if the mime types is multipart.
func IsMultipart(mt string) bool {
	switch mt {
	case "multipart/alternative":
		return true
	case "multipart/appledouble":
		return true
	case "multipart/byteranges":
		return true
	case "multipart/digest":
		return true
	case "multipart/encrypted":
		return true
	case "multipart/example":
		return true
	case "multipart/form-data":
		return true
	case "multipart/header-set":
		return true
	case "multipart/mixed":
		return true
	case "multipart/multilingual":
		return true
	case "multipart/parallel":
		return true
	case "multipart/related":
		return true
	case "multipart/report":
		return true
	case "multipart/signed":
		return true
	case "multipart/vnd.bint.med-plus":
		return true
	case "multipart/voice-message":
		return true
	case "multipart/x-mixed-replace":
		return true
	default:
		return false
	}
}
