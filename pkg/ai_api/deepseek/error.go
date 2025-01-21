package deepseek

import "errors"

var ErrNoMessage = errors.New("completions request no message")
var ErrMessageNoContent = errors.New("message no content")
var ErrMessageRole = errors.New("message role not in (system, user, assistant, tool)")
var ErrFrequencyPenalty = errors.New("frequency penalty not in range [-2.0, 2.0]")
var ErrMaxTokens = errors.New("max tokens not in range [1, 8192]")
var ErrPresencePenalty = errors.New("presence penalty not in range [-2.0, 2.0]")

var ErrRespNotOK = errors.New("response is not 200OK")
