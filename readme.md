# Concurrent Bogosort

<img align="right" width="500" src="./docs/_media/golang-what.webp" alt="Golang What"/>

## What is Bogosort?

Bogosort, also known as permutation sort, stupid sort, slow sort, shotgun sort or monkey sort is a particularly ineffective algorithm one person can ever imagine.

It is based on generate and test paradigm. The algorithm successively generates permutations of its input until it finds one that is sorted.

## How does Concurrent Bogosort works?

This code opens a new goroutine sequentially, applying the bogosort technique until one matches a sorted array.
At the end of the execution, you can see the number of opened goroutines.

## WHY?

This is just a study of Golang's Goroutines. Just for fun. Yeah, nothing more.

## Contributing

This project is an open-source project, and contributions from other developers are welcome. If you encounter any issues or have suggestions for improvement, please submit them on the project's GitHub page.

Just kidding, please do not contribute.
