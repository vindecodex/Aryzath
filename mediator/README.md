# Mediator

This pattern lets you collaborate with every object without letting them know about each other.

The mediator will be the one to handle the communication between objects.

Single Responsibility since each objects acts on there own.

Open/Closed Principle since you can add more objects.

Common Problem is online games or chat, when someone offline that pointer reference of object will become nil, so we are going to solve this using the mediator pattern. Instead of talking to a pointer reference we're going to communicate with the mediator then the mediator will be the one sends information.
