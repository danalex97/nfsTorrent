package multi_torrent

import (
  . "github.com/danalex97/Speer/interfaces"

  "github.com/danalex97/nfsTorrent/cache_torrent"
  "strconv"
)

// A MultiElection component is made out of multiple Election components. When
// an election takes place, the nodes that were elected get banned for the
// purpose of them not getting reelected. The banning process is done by
// eliminating all the banned candidates from all the elections that did not
// take place yet. For this algorithm to be correct, we require the number of
// sub-Torrents to be smaller than 100/leaderPercent.
type MultiElection struct {
  limit int
  nodes int

  elections map[string]*cache_torrent.Election
}

func NewMultiElection(elections int, limit int, transport Transport) *MultiElection {
  e := &MultiElection{
    limit : limit,
    nodes : 0,

    elections : make(map[string]*cache_torrent.Election),
  }

  proxy := NewStripProxy(transport)
  for i := 0; i < elections; i++ {
    e.elections[strconv.Itoa(i)] = cache_torrent.NewElection(limit, proxy)
  }
  return e
}

func (e *MultiElection) Run() {
}

func (e *MultiElection) NewJoin(id string) {
  for i, election := range e.elections {
    election.NewJoin(FullId(id, i))
  }
}

func (e *MultiElection) Recv(m interface {}) {
  switch candidate := m.(type) {
  case cache_torrent.Candidate:
    e.RegisterCandidate(candidate)
  }
}

func (e *MultiElection) RegisterCandidate(candidate cache_torrent.Candidate) {
  e.nodes++
  for i, election := range e.elections {
    election.RegisterCandidate(cache_torrent.Candidate{
      Id   : FullId(candidate.Id, i),
      Up   : candidate.Up,
      Down : candidate.Down,
    })

    if e.nodes == e.limit {
      // Eliminate the already elected candidates
      toRemove := election.GetElected()
      for j, curr := range e.elections {
        if i != j {
          for _, id := range toRemove {
            curr.RemoveCandidate(FullId(ExternId(id), j))
          }
        }
      }
    }
  }
}
