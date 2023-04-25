# gonet

This is a library to speed development of distributed applications.
The purpose is to encapsulate all of the lower-level networking calls
and some basic fault tolerance.

# Simulator

One aspect of this library is the network simulator.
The goal is to implement the interface in such a way that programs using this networking layer
can accept a drop-in replacement which would allow unit testing to be done from a single OS process.

You could also do your integration testing on a single host, where every address points to a different
`localhost` port. However, the purpose of the simulator is to also be extensible and offer features
such as jitter, message dropping, message corruption, etc. In general, this Simulator layer can
be used to test against CAP theorem problems.