# DISTIL examples

Here are a few examples for how to use go-distil.

## Noop

This is the first example, it does the bare minimum required to interface with
BTrDB and the DISTIL engine, and copy one stream to another. If this program
is left running, any changes in the given stream will manifest in the output
stream. It hardcodes the address of the server and the streams so you will need
to modify it before you run it yourself.

## Frequency

This is a more realistic example, it gets the server addresses from environment
variables (`$DISTIL_BTRDB_ADDR` and `$DISTIL_MONGO_ADDR`). It then looks for
all uPMUs and creates an instance of the frequency distillate for each L1 channel.
By following this method, it is easy to make a program that "automatically" locates
new uPMUs and computes distillates on them, removing the need to modify the
distillate code or maintain a configuration file. It also makes use of the
Rebaser and Lead time.
