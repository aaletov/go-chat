# go-chat

## Description

The project should have two sides: server and client. For privacy purposes, server should not store any client's data and all messages must be client-side encrypted (we can use PGP for this purpose). To identify clients without passwords, the public key should be used as identifier. Message queue may be stored on client-side or not stored at all, but client-server-client message transfer is possible only when two clients are online.