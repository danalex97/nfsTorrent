package cache_torrent

import (
  "github.com/danalex97/nfsTorrent/torrent"
)

type Forwarder struct {
  *Leader

  from string
  to   string
}

func NewForwarder(leader *Leader, from, to string) *Forwarder {
  return &Forwarder{
    Leader : leader,

    from : from,
    to   : to,
  }
}

func (f *Forwarder) Recv(m interface {}) {
  id := f.GetId(m)

  if id == f.from {
    // follower --> peer

    switch msg := m.(type) {
    case torrent.Request:
      _, ok := f.Storage.Have(msg.Index)
      if !ok {
        // In this momment, the connection follower - leader is unchoked
        // but we don't have the piece. We will try to favor the transfer of
        // the piece by letting the picker know.

        //[TODO] let picker know
      }
    }
  } else {
    // peer --> follower

    switch m.(type) {
    case torrent.Have:
      f.Transport.ControlSend(f.from, m)
    }
  }
}
