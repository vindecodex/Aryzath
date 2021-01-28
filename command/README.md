# Command

It follows the single responsibility principle. Each Command Class performs specific actions.

It follows Open/Closed Principle since it is easy to add new commands.

It also allows undo features easily introduced by uncle bob. [Undo](/undo)

It also allows queing of commands.

It executes any methods in single command by setting up what command it will handle.
