# Memory Layout

When you define a struct, Go has to decide where each field lives in memory. Fields are placed one after the other, in the order you wrote them. But there’s a catch: the CPU likes it when data is aligned — integers on 4-byte or 8-byte boundaries, for example. So Go may insert padding bytes to keep everything aligned.

## Assignment

Our over-engineering boss is at it again. He's heard about memory layout and wants to squeeze every last byte out of our structs.

Run the tests to see the current size of the structs, then update the struct definitions to minimize memory usage.
