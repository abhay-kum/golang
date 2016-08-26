# golang
This is a small library of Go programs:

##1] WordCountSequential.go

   This program as the name suggests, takes a text file (now a directory) as input and after processing the file/directory it will create a processed.csv file with WordCount.

##2] WordCountParallel.go

   This program uses golang's parallelism features. The way it works is you pass a directory with multiple text files to be processed, depending on the parallelism set in the program, it would create the channels and would assign one file to each channel and wait for the channel to be available for assigning the next file.
   This is initial version with some assumptions on file names (processed.csv etc)
