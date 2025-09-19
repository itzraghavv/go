# Assignment

1. Implement the `importance` methods for each message type. They should return the importance score for each message type.
   1. For a `directMessage` the importance score is based on if the message `isUrgent` or not. If it is urgent, the importance score is `50` otherwise the importance score is equal to the DM's `priorityLevel`.
   2. For a `groupMessage` the importance score is based on the message's `priorityLevel`.
   3. All `systemAlert` messages should return a `100` importance score.

2. Complete the `processNotification` function. It should identify the type and return different values for each type
   1. For a `directMessage`, return the sender's username and importance score.
   2. For a `groupMessage`, return the group's name and the importance score.
   3. For a `systemAlert`, return the alert code and the importance score.
   4. If the notification does not match any known type, return an empty string and a score of 0.
