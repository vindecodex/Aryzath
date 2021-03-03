# Proxy

Proxy is used to wrap the original object from the client. The client will talk to the proxy then the proxy will decide either it will give an access to the original object.

Proxy will contain similar Function of the Original Object. So it is good to have an interface for this.

Proxy also helps to validate the clients request before it can access the original object.

It helpls to implement:
- Logging
- Authorization before accessing the Real Object
- Lazy Initialization of Real Object

