# Simple state machine example

This example contains a simple.go which is the main program and example.go will have the execution steps of a state machine. Util.go describes state machine.
* The state machine has an init event which starts from Idle state and reaches state1.
* Close event ends with closed state.
* Every transition event is between init and close. A transition from state1 to state2 is named as state1_2.
* Every intermediate state except from idle and closed are reentrants only if there is an error and these reentrant states are handled in error transition.

A small ASCII Diagram for how it works
```                                                                  
                                                                     +--------->-----+
                                     +----->-------+                 |    error      |
                                     |   error     |                 +---------<-----+
       +-----------+                 +-----<-------+                 |               |
       |           |                 |             |    state1_2     |    state2     |
       | idle      |     init        |  state1     |                 |               |
       |           +---------------->+             +---------------->+               |
       |           |                 |             |                 +---------------+
       +-----------+                 +-------------+                         |
                                                                             |
                                                                             |  state2_3
                                                                             |
                                                                             v
                                                                      +--------------------->---+
                                      +---------------+    state3_4   |               |         |
        +-----------+    close        |               |               |     state3    |  error  |
        |           |                 |  state4       +<--------------+               |         |
        |           +<----------------+               |               |               |         |
        |   closed  |                 |               |               +----------------^-----<--+
        |           |                 +-------->------+
        +-----------+                 ^               |
                                      |    error      |
                                      +--------<------+
```
* Above diagram represents a simple state machine that starts with Idle and ends with closed.
* Every event is a transition from State X to State Y. An Error transition occurs from State X to State X.
* Error transition is re-entrant and executes as per the # of retries.
* While executing a job if the state machine foes down before updating a state. During a restart the state machine pick ups from the last updated state and executes the last event transition. Even if we start FSM from beginning events.
