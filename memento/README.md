# Memento

- It creates a snapshot for the current state to be used for the future roll back.

How it works ?

The <b>originator</b> will save it's state by creating a <b>memento</b> that holds current state of <b>originator</b>. The created <b>memento</b>
will be passed to the <b>caretaker</b>. <b>Caretaker</b> is the one holds all the <b>memento</b>, if <b>originator</b> wants to roll back from a specific
snapshot then the <b>originator</b> will request for a <b>memento</b> from <b>caretaker</b> then <b>caretaker</b> will give the <b>memento</b> back to <b>originator</b>
