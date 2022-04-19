# chat
chat


Please write a client and a server that talk to each other via UDP. The server has to cache messages in Redis. It is not expect guaranteed delivery.
Task details:
• Redis backed chat room (Go language)
• Create a chat room based on UDP protocol
• Clients send messages via UDP to Server. Then it is broadcast to all other clients
• Server pushes message to Redis for temporary history. History is limited to 20 messages
• When new client connects to chat server it receives last 20 messages (in correct order)
• Client may delete any message he/she has previously written (but not messages from others).
• When client deletes message, it is removed in chat screen for all clients. And if it was saved in Redis, it will also be removed there as well, new clients will see history without it.
• When all clients disconnect, the DB is flushed.
• Tests are (very) welcome