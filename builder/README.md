# Builder

Build only applies to creating a product with common process of building but with different version of it.

The director is the one will generate the final product using the methods created by the concrete classes.

The builder (concrete class) contains the methods to build the entire product but it's method should only be used by the director.

The builder interface will be the bridge to interact with the concrete builder and the director.

The Product is the object which will be the result of build.
