// maubot - A plugin-based Matrix bot system written in Go.
// Copyright (C) 2018 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package matrix

import (
	"maubot.xyz"
	"maunium.net/go/gomatrix"
	"maunium.net/go/gomatrix/format"
)

type EventFuncsImpl struct {
	*gomatrix.Event
	Client *Client
}

func (client *Client) ParseEvent(mxEvent *gomatrix.Event) *maubot.Event {
	if mxEvent == nil {
		return nil
	}
	mxEvent.Content.RemoveReplyFallback()
	return &maubot.Event{
		EventFuncs: &EventFuncsImpl{
			Event:  mxEvent,
			Client: client,
		},
		Event: mxEvent,
	}
}

func (evt *EventFuncsImpl) MarkRead() error {
	return evt.Client.MarkRead(evt.RoomID, evt.ID)
}

func (evt *EventFuncsImpl) Reply(text string) (string, error) {
	content := format.RenderMarkdown(text)
	content.MsgType = gomatrix.MsgNotice
	return evt.ReplyContent(content)
}

func (evt *EventFuncsImpl) ReplyContent(content gomatrix.Content) (string, error) {
	content.SetReply(evt.Event)
	return evt.SendContent(content)
}

func (evt *EventFuncsImpl) SendMessage(text string) (string, error) {
	return evt.Client.SendMessage(evt.RoomID, text)
}

func (evt *EventFuncsImpl) SendMessagef(text string, args ...interface{}) (string, error) {
	return evt.Client.SendMessagef(evt.RoomID, text, args...)
}

func (evt *EventFuncsImpl) SendContent(content gomatrix.Content) (string, error) {
	return evt.Client.SendContent(evt.RoomID, content)
}

func (evt *EventFuncsImpl) SendMessageEvent(evtType gomatrix.EventType, content interface{}) (eventID string, err error) {
	return evt.Client.SendMessageEvent(evt.RoomID, evtType, content)
}