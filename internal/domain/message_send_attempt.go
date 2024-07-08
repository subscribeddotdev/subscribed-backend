package domain

import (
	"errors"
	"time"
)

const httpLastSuccessStatusCode = 299

type StatusCode uint

type SendAttemptStatus string

func (s SendAttemptStatus) String() string {
	return string(s)
}

/*func (s SendAttemptStatus) validate() error {
	if s != "succeeded" && s != "failed" {
		return fmt.Errorf("%s is not a valid message send attempt status", s)
	}

	return nil
}*/

var (
	sendAttemptStatusSucceeded = SendAttemptStatus("succeeded")
	sendAttemptStatusFailed    = SendAttemptStatus("failed")
)

type MsgSendAttemptID string

func (i MsgSendAttemptID) String() string {
	return string(i)
}

func NewMsgSendAttemptID() MsgSendAttemptID {
	return MsgSendAttemptID(NewID().WithPrefix("msa"))
}

type MessageSendAttempt struct {
	id             MsgSendAttemptID
	messageID      MessageID
	endpointID     EndpointID
	timestamp      time.Time
	status         SendAttemptStatus
	response       string // TODO: make this value-object
	statusCode     StatusCode
	requestHeaders Headers
}

func NewMessageSendAttempt(
	endpointID EndpointID,
	messageID MessageID,
	response string,
	statusCode StatusCode,
	requestHeaders Headers,
) (*MessageSendAttempt, error) {
	if endpointID.String() == "" {
		return nil, errors.New("endpointURL cannot be empty")
	}

	if messageID.String() == "" {
		return nil, errors.New("messageID cannot be empty")
	}

	status := sendAttemptStatusFailed
	if statusCode <= httpLastSuccessStatusCode {
		status = sendAttemptStatusSucceeded
	}

	return &MessageSendAttempt{
		id:             NewMsgSendAttemptID(),
		messageID:      messageID,
		endpointID:     endpointID,
		timestamp:      time.Now().UTC(),
		status:         status,
		response:       response,
		statusCode:     statusCode,
		requestHeaders: requestHeaders,
	}, nil
}

func (m *MessageSendAttempt) ID() MsgSendAttemptID {
	return m.id
}

func (m *MessageSendAttempt) MessageID() MessageID {
	return m.messageID
}

func (m *MessageSendAttempt) EndpointID() EndpointID {
	return m.endpointID
}

func (m *MessageSendAttempt) Timestamp() time.Time {
	return m.timestamp
}

func (m *MessageSendAttempt) Status() SendAttemptStatus {
	return m.status
}

func (m *MessageSendAttempt) Response() string {
	return m.response
}

func (m *MessageSendAttempt) StatusCode() StatusCode {
	return m.statusCode
}

func (m *MessageSendAttempt) RequestHeaders() Headers {
	return m.requestHeaders
}
