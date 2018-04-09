package torrent

import (
  . "github.com/danalex97/Speer/interfaces"
)

/* Original BitTorrent protocol messages. */
type choke struct { id string }
type unchoke struct { id string }
type interested struct { id string }
type notInterested struct { id string }

// A have message is sent to neighbours whenever a peer
// obtains a new piece.
type have struct {
  id    string
  index int // Index of the piece that I have
}

type request struct {
  index  int
  begin  int
  length int
}

// The actual piece as a response to a `request` message
type piece struct {
  index  int
  begin  int
  piece  Data
}

// We do not model endgame mode, so we have no `cancel` message

/* Tracker control messages. */
type join struct {
  id string
}

type neighbours struct {
  ids []string
}

type trackerReq struct {
  from string
}

type trackerRes struct {
  id string
}

/* Used to model seeds:
    - each peers sends a seed request
    - in seed response finds if its a seed how many pieces it has
 */
type seedReq struct {
  from string
}

type seedRes struct {
  pieces []request
}
