#gobase
gobase is a database that I am building up from scratch.  I am using a B+ tree index for the indexing data structure.  B+ trees are fast (not the fastest), easy to persist to disk (and therefore cache), and easy to select all items (as it has a link list at the bottom of the tree).  

All data is kept at the bottom of the tree not intermingled in the middle.

#Roadmap
+ Deletes
+ Persistance
+ Caching
+ Schema
+ Command line utility with multiple tables


+ (don't know how yet but want)
+ RA engine
+ Data Definition Language
+ Data Manipulation Language
+ Modularity (swap out all parts like caching and persistance with different modules etc)
+ Network interface
+ SQL Engine? maybe (or maybe new language)
+ Database Drivers

#Progress

####06/03/2014
Inserts are working!

Clocked in on my computer at 78,000ish inserts per second!

Ran a test that would insert one item and then check to make sure each item was still in the database one at a time.  This clocked in at about 5,000 for the first second making that about 13,000,000 reads per second on top of 5,000 inserts during the same second.
