package gonet

type Cluster struct {
	Addrs []string
}

// TODO implement some basic functions around cluster communication
// e.g. broadcast
//
// Also, possibly define some functions for building protocols,
// e.g. having separate data and control planes, or defining some standard
// for creating protocol messages.
