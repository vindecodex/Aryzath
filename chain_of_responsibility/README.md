# Chain of Responsibility

With this pattern you can execute chain of functionality upon recieving the request. Does
Chain of functionality also called as handler.

Each handler can terminate the chain or can continue the chain, depends upon the condition of the request.

It follows the single responsibility principle since you can do specific things on each handler.

It follows the Open/Closed Principle since you can introduce more handler without breaking existing code.
