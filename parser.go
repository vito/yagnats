package yagnats

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Parser func(*bufio.Reader) (Packet, error)

var PARSERS = map[string]Parser{
	// PING\s*\r\n
	"PING": func(io *bufio.Reader) (Packet, error) {
		io.ReadBytes('\n')
		return &PingPacket{}, nil
	},

	// PONG\s*\r\n
	"PONG": func(io *bufio.Reader) (Packet, error) {
		io.ReadBytes('\n')
		return &PongPacket{}, nil
	},

	// +OK\s*\r\n
	"+OK": func(io *bufio.Reader) (Packet, error) {
		io.ReadBytes('\n')
		return &OKPacket{}, nil
	},

	// -ERR '(message)'\r\n
	"-ERR": func(io *bufio.Reader) (Packet, error) {
		bytes, _ := io.ReadBytes('\n')
		re := regexp.MustCompile(`\s*'(.*)'\r\n`)
		match := re.FindSubmatchIndex(bytes)
		return &ERRPacket{Message: string(bytes[match[2]:match[3]])}, nil
	},

	// INFO (payload)\r\n
	"INFO": func(io *bufio.Reader) (Packet, error) {
		bytes, _ := io.ReadBytes('\n')
		re := regexp.MustCompile(`\s*([^\s]+)\r\n`)
		match := re.FindSubmatchIndex(bytes)
		return &InfoPacket{Payload: string(bytes[match[2]:match[3]])}, nil
	},

	// MSG (subject) (subscriber-id) (reply)? (length)\r\n(byte * length)\r\n
	"MSG": func(io *bufio.Reader) (Packet, error) {
		bytes, _ := io.ReadBytes('\n')
		re := regexp.MustCompile(`\s*([^\s]+)\s+(\d+)\s+(([^\s]+)\s+)?(\d+)\r\n`)
		matches := re.FindStringSubmatch(string(bytes))

		subID, _ := strconv.Atoi(matches[2])
		payloadLen, _ := strconv.Atoi(matches[5])
		payload := make([]byte, payloadLen)
		io.Read(payload)
		io.ReadBytes('\n')

		return &MsgPacket{
			Subject: matches[1],
			SubID:   subID,
			ReplyTo: matches[4],
			Payload: string(payload),
		}, nil
	},
}

func Parse(io *bufio.Reader) (val Packet, err error) {
	header, err := readWord(io)
	if err != nil {
		return nil, err
	}

	parser := PARSERS[string(header)]
	if parser == nil {
		return nil, errors.New(fmt.Sprintf("Unknown header: %s", header))
	}

	return parser(io)
}

func readWord(io *bufio.Reader) ([]byte, error) {
	word := new(bytes.Buffer)

	for {
		next, err := io.ReadByte()
		if err != nil {
			return nil, err
		}

		switch next {
		case ' ', '\t', '\r', '\n':
			return word.Bytes(), nil

		default:
			word.WriteByte(next)
		}
	}
}
