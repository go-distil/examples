# DISTIL examples

This repo contains a number of example distillates written in Go for the DISTIL processing framework designed for use with the BTrDB time series database. The examples are well commented and increase in complexity. Each contains two files:

1. main.go - configure and run the distillate including the input/output stream names
2. algorithm.go - contains the function that defines how to process the time series data

Distillates must be compiled and can be run either from a remote system (such as your laptop) or on the server with BTrDB. Running a distillate on a remote system requires the data to stream over the network to the running distillate and then the output data must return to the server. Thus, running distillates remotely is intended for testing only. 

The distillate executable expects an environmental variable, `$BTRDB_ENDPOINTS`, to be set to either an IP address with port (xxx.xxx.xxx.xxx:4410) or a server name such as "btrdb.myserver.local." Note that distillates connecting to the older version 3.x of BTrDB required different environmental variables: `$DISTIL_BTRDB_ADDR` and `$DISTIL_MONGO_ADDR`.



## 1. Noop

This is the first example, it does the bare minimum required to interface with BTrDB and the DISTIL engine, and copies one stream to another. If this program is left running, any changes in the given stream will manifest in the output stream. It hardcodes the input and output streams so you will need to modify it before you run it yourself.

## 2. Frequency

This is a more realistic example. It gets the server addresses from the environment variable `$BTRDB_ENDPOINTS`. It then looks for all uPMUs and creates an instance of the frequency distillate for each L1 channel. By following this method, it is easy to make a program that "automatically" locates new uPMUs and computes distillates on them, removing the need to modify the
distillate code or maintain a configuration file. It also makes use of the Rebaser and Lead time.



# The DISTIL System

For an indepth overview of the design of the DISTL system, please see [DISTIL: Design and Implemention of a Scalable Synchrophasor Data Processing System](http://ieeexplore.ieee.org/document/7436312/) by Michael P Andersen, Sam Kumar,Connor Brooks, Alexandra von Meier, and David E. Culler . This paper appeared in the 2015 IEEE International Conference on Smart Grid Communications.

## Abstract
*The introduction and deployment of cheap, high precision, high sample rate next-generation synchrophasors en masse in both the transmission and distribution tier – while invaluable for fault diagnosis, situational awareness and capacity planning – poses a problem for existing methods of phasor data analysis and storage. Addressing this, we present the design and implementation of a novel architecture for synchrophasor data analysis on distributed commodity hardware. At the core is a new feature-rich timeseries store, BTrDB. Capable of sustained writes and reads in excess of 16 million points per second per cluster node, advanced query functionality and highly efficient storage, this database enables novel analysis and visualization techniques. Leveraging this, a distillate framework has been developed that enables agile development of scalable analysis pipelines with strict guarantees on result integrity despite asynchronous changes in data or out of order arrival. Finally, the system is evaluated in a pilot deployment, archiving more than 216 billion raw datapoints and 515 billion derived datapoints from 13 devices in just 3.9TB. We show that the system is capable of scaling to handle complex analytics and storage for tens of thousands of next-generation synchrophasors on off-the-shelf servers.*
