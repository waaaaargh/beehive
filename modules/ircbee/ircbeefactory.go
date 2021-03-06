/*
 *    Copyright (C) 2014 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package ircbee

import (
	"github.com/muesli/beehive/modules"
)

type IrcBeeFactory struct {
	modules.ModuleFactory
}

// Interface impl

func (factory *IrcBeeFactory) New(name, description string, options modules.BeeOptions) modules.ModuleInterface {
	bee := IrcBee{
		server:      options.GetValue("server").(string),
		nick:        options.GetValue("nick").(string),
	}

	for _, channel := range options.GetValue("channels").([]interface{}) {
		bee.channels = append(bee.channels, channel.(string))
	}

	// optional parameters
	if options.GetValue("password") != nil {
		bee.password = options.GetValue("password").(string)
	}
	if options.GetValue("ssl") != nil {
		bee.ssl = options.GetValue("ssl").(bool)
	}

	bee.Module = modules.Module{name, factory.Name(), description}
	return &bee
}

func (factory *IrcBeeFactory) Name() string {
	return "ircbee"
}

func (factory *IrcBeeFactory) Description() string {
	return "An IRC module for beehive"
}

func (factory *IrcBeeFactory) Options() []modules.BeeOptionDescriptor {
	opts := []modules.BeeOptionDescriptor{
		modules.BeeOptionDescriptor{
			Name:        "server",
			Description: "Hostname of IRC server, eg: irc.example.org:6667",
			Type:        "string",
			Mandatory:   true,
		},
		modules.BeeOptionDescriptor{
			Name:        "nick",
			Description: "Nickname to use for IRC",
			Type:        "string",
			Mandatory:   true,
		},
		modules.BeeOptionDescriptor{
			Name:        "password",
			Description: "Password to use to connect to IRC server",
			Type:        "string",
		},
		modules.BeeOptionDescriptor{
			Name:        "channels",
			Description: "Which channels to join",
			Type:        "[]string",
			Mandatory:   true,
		},
		modules.BeeOptionDescriptor{
			Name:        "ssl",
			Description: "Use SSL for IRC connection",
			Type:        "bool",
		},
	}
	return opts
}

func (factory *IrcBeeFactory) Events() []modules.EventDescriptor {
	events := []modules.EventDescriptor{
		modules.EventDescriptor{
			Namespace:   factory.Name(),
			Name:        "message",
			Description: "A message was received over IRC, either in a channel or a private query",
			Options: []modules.PlaceholderDescriptor{
				modules.PlaceholderDescriptor{
					Name:        "text",
					Description: "The message that was received",
					Type:        "string",
				},
				modules.PlaceholderDescriptor{
					Name:        "channel",
					Description: "The channel the message was received in",
					Type:        "string",
				},
				modules.PlaceholderDescriptor{
					Name:        "user",
					Description: "The user that sent the message",
					Type:        "string",
				},
			},
		},
	}
	return events
}

func (factory *IrcBeeFactory) Actions() []modules.ActionDescriptor {
	actions := []modules.ActionDescriptor{
		modules.ActionDescriptor{
			Namespace:   factory.Name(),
			Name:        "send",
			Description: "Sends a message to a channel or a private query",
			Options: []modules.PlaceholderDescriptor{
				modules.PlaceholderDescriptor{
					Name:        "channel",
					Description: "Which channel to send the message to",
					Type:        "string",
				},
				modules.PlaceholderDescriptor{
					Name:        "text",
					Description: "Content of the message",
					Type:        "string",
				},
			},
		},
		modules.ActionDescriptor{
			Namespace:   factory.Name(),
			Name:        "join",
			Description: "Joins a channel",
			Options: []modules.PlaceholderDescriptor{
				modules.PlaceholderDescriptor{
					Name:        "channel",
					Description: "Channel to join",
					Type:        "string",
				},
			},
		},
		modules.ActionDescriptor{
			Namespace:   factory.Name(),
			Name:        "part",
			Description: "Parts a channel",
			Options: []modules.PlaceholderDescriptor{
				modules.PlaceholderDescriptor{
					Name:        "channel",
					Description: "Channel to part",
					Type:        "string",
				},
			},
		},
	}
	return actions
}

func init() {
	f := IrcBeeFactory{}
	modules.RegisterFactory(&f)
}
