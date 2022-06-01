# Authentication Component
Source code: [comp_auth.go](../comp_auth.go)

Handles packets necessary to complete authentication and responds with according packets. Sends the `init` packet unpon connection open.
Invalid packets related to auth received from server or detected incompabilities may result in the client being closed. In this case, an error state is set that can be read using the `self:state()` method.

**Important: ** the auth component does not automatically disconnect if authentication fails due to an invalid password or already being logged in; it is up to the API user to handle these cases by subscribing to the `kick` and `legacy_kick` packets.

Handles the `hello`, `srp_bytes_salt_b`, `accept_auth` and `time_of_day` packets (the last one is only handled when received the first time and sets the state to active).
May send `init`, `first_srp`, `srp_bytes_a`, `srp_bytes_m`, `init2` and `ready` packets.

## Functions

`self:username([username])`: Sets or gets the username (string). Setting may not occur after having connected the client. A username must be set before connecting.
`self:password([password])`: Sets or gets the password (string). Setting may not occur after having connected the client. By default, an empty password is used.
`self:language([language])`: Sets or gets the language sent to server. Setting may not occur after having connected the client. By default, "en_US" is used.
`self:version([version])`: Sets or gets the version string sent to server. Setting may not occur after having connected the client. By default, "hydra-dragonfire" is used.
`self:state()`: Returns `state, error`. State is one of "init", "requested", "verified", "active", "error". If state is "error", error is a string containing a description of the problem that occured. Otherwise, error is nil.
