`cleanupca` util deletes certain certificates in pem file
`[XDG_CONFIGDIR]/ca.pem`.  In UNIX like operation systems
`[XDG_CONFIGDIR]` is `~/.config/` most likely.

`cleanupca` deletes copies of certificates so that each certificate
appears at most once.  `cleanupca` deletes certificates whose expiration
date lies in the past.
