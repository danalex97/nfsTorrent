package torrent

// This file follows the 'Upload' BitTorrent 5.3.0 release

type Connector struct {
  from string
  to   string

  interested bool 
  choked     bool

  upload    *Upload
  download  *Download
  handshake *Handshake

  components *Components
}

func NewConnector(from, to string, components *Components) *Connector {
  connector := new(Connector)

  connector.from  = from
  connector.to    = to

  connector.components = components
  connector.handshake  = NewHandshake(connector)
  connector.upload     = NewUpload(connector)
  connector.download   = NewDownload(connector)

  connector.interested   = false
  connector.choked     = true  // I am chocked by all peers

  return connector
}

func (c *Connector) Run() {
  go c.handshake.Run()
  go c.upload.Run()
  go c.download.Run()
}

func (c *Connector) Recv(m interface {}) {
  c.handshake.Recv(m)
  c.upload.Recv(m)
  c.download.Recv(m)
}

/*
 * Returns the downoad rate of the connection.
 */
func (c *Connector) Rate() float64 {
  // [TODO]
  return 0
}

func (c *Connector) Choke() {
  c.upload.Choke()
}

func (c *Connector) Unchoke() {
  c.upload.Unchoke()
}

func (c *Connector) RequestMore() {
  c.download.RequestMore()
}
