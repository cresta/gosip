package sip

import (
	"bytes"
	"testing"
)

// This repo is a public one, please remove any sensitive information before pushing to github
var (
	SIPInviteFromAudioCodesSBC = "INVITE sip:siprec-ip@siprec-ip;user=phone SIP/2.0\r\n" +
		"Via: SIP/2.0/TCP 8.8.8.8:5060;alias;branch=z9hG4bKac586296494\r\n" +
		"Max-Forwards: 70\r\n" +
		"From: <sip:10.0.0.31;user=phone>;tag=1c1789648591\r\n" +
		"To: <sip:siprec-ip@siprec-ip;user=phone>\r\n" +
		"Call-ID: 1974564438218202316425@8.8.8.8\r\n" +
		"CSeq: 1 INVITE\r\n" +
		"Contact: <sip:8.8.8.8:5060;transport=tcp>;+sip.src\r\n" +
		"Supported: replaces,resource-priority,sdp-anat\r\n" +
		"Allow: REGISTER,OPTIONS,INVITE,ACK,CANCEL,BYE,NOTIFY,PRACK,REFER,INFO,SUBSCRIBE,UPDATE\r\n" +
		"Require: siprec\r\n" +
		"User-Agent: Mediant VE SBC/v.7.20A.256.399\r\n" +
		"Content-Type: multipart/mixed;boundary=boundary_ac15d5\r\n" +
		"Content-Length: 1986\r\n" +
		"\r\n" +
		"--boundary_ac15d5\r\n" +
		"Content-Type: application/sdp\r\n" +
		"\r\n" +
		"v=0\r\n" +
		"o=AudiocodesGW 1194408899 488170781 IN IP4 8.8.8.8\r\n" +
		"s=SBC-Call\r\n" +
		"c=IN IP4 8.8.8.8\r\n" +
		"t=0 0\r\n" +
		"m=audio 7232 RTP/AVP 0 101\r\n" +
		"c=IN IP4 8.8.8.8\r\n" +
		"a=ptime:20\r\n" +
		"a=sendonly\r\n" +
		"a=label:1\r\n" +
		"a=rtpmap:0 PCMU/8000\r\n" +
		"a=rtpmap:101 telephone-event/8000\r\n" +
		"a=fmtp:101 0-15,16\r\n" +
		"m=audio 7236 RTP/AVP 0 101\r\n" +
		"c=IN IP4 8.8.8.8\r\n" +
		"a=ptime:20\r\n" +
		"a=sendonly\r\n" +
		"a=label:2\r\n" +
		"a=rtpmap:0 PCMU/8000\r\n" +
		"a=rtpmap:101 telephone-event/8000\r\n" +
		"a=fmtp:101 0-15,16\r\n" +
		"\r\n" +
		"--boundary_ac15d5\r\n" +
		"Content-Type: application/rs-metadata\r\n" +
		"Content-Disposition: recording-session\r\n" +
		"\r\n" +
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n" +
		"<recording xmlns=\"urn:ietf:params:xml:ns:recording\" xmlns:ac=\"http://AudioCodes\">\r\n" +
		"  <datamode>complete</datamode>\r\n" +
		"  <group id=\"00000000-0000-00b6-d16f-0000000004e4\">\r\n" +
		"    <associate-time>2023-08-21T16:42:05Z</associate-time>\r\n" +
		"  </group>\r\n" +
		"  <session id=\"0000-0000-0000-0000-bf41fa8f111f3e43\">\r\n" +
		"    <group-ref>00000000-0000-00b6-d16f-0000000004e4</group-ref>\r\n" +
		"    <associate-time>2023-08-21T16:42:05Z</associate-time>\r\n" +
		"  </session>\r\n" +
		"  <participant id=\"+18888888888\" session=\"0000-0000-0000-0000-bf41fa8f111f3e43\">\r\n" +
		"    <nameID aor=\"+18888888888@fakeip.com\"></nameID>\r\n" +
		"    <associate-time>2023-08-21T16:42:05Z</associate-time>\r\n" +
		"    <send>00000000-1c3c-00b6-d16f-0000000004e4</send>\r\n" +
		"    <recv>00000001-2d1c-00b6-d16f-0000000004e4</recv>\r\n" +
		"  </participant>\r\n" +
		"  <participant id=\"+17777777777\" session=\"0000-0000-0000-0000-bf41fa8f111f3e43\">\r\n" +
		"    <nameID aor=\"+17777777777@fakeip.com\"></nameID>\r\n" +
		"    <associate-time>2023-08-21T16:42:05Z</associate-time>\r\n" +
		"    <send>00000001-2d1c-00b6-d16f-0000000004e4</send>\r\n" +
		"    <recv>00000000-1c3c-00b6-d16f-0000000004e4</recv>\r\n" +
		"  </participant>\r\n" +
		"  <stream id=\"00000000-1c3c-00b6-d16f-0000000004e4\" session=\"0000-0000-0000-0000-bf41fa8f111f3e43\">\r\n" +
		"    <label>1</label>\r\n" +
		"  </stream>\r\n" +
		"  <stream id=\"00000001-2d1c-00b6-d16f-0000000004e4\" session=\"0000-0000-0000-0000-bf41fa8f111f3e43\">\r\n" +
		"    <label>2</label>\r\n" +
		"  </stream>\r\n" +
		"</recording>\r\n" +
		"--boundary_ac15d5--"
)

func TestParse(t *testing.T) {
	buf := bytes.Buffer{}
	buf.WriteString(SIPInviteFromAudioCodesSBC)
	bytes := buf.Bytes()
	msg, err := ParseMsg(bytes)
	if err != nil {
		t.Error(err)
	}
	if msg == nil || msg.Payload == nil {
		t.Error("msg is nil")
	}
	if msg.Payload.ContentType() != "multipart/mixed;boundary=boundary_ac15d5" {
		t.Error("content type is wrong")
	}
}
