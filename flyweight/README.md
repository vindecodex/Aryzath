# Flyweight

Main purpose of the pattern is to reduce RAM consumption by sharing the resources instead of creating them again and again. This pattern reduces the redundancy of storing data.

From our example every time we add a player into our game instead of creating a new team object for each player we point to the created object in our map that was already been created.

From our example it seems like it doesn't really useful but try to imagine if the Team object was a big object holding a lot of data's and methods.
