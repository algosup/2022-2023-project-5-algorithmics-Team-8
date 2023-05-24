# Meeting group 5 - 12-05-2023

As group 8 we did a meeting with group 5, the meeting took place on the 12th of May 2023 at 9.25, the meeting lasted 25 minutes.
 
## Overview

During this meeting, we mostly talked about the technical specifications and the algorithm itself, with some comments on the functional specifications.

Group 5 has taken into account our comments on their first version of the technical specifications, and we have also taken into account their comments on our functional specifications. 

We agreed on making 2 algorithms, one main and another one as an "alternative" algorithm, in case the main one is not working as intended. 

We discussed the LTS agreement and versioning in general, we chose to use the latest LTS version of Bazel, with Google test 1.12.0 as it is compatible with both CPP 11 and macOS. 

After discussing with the client, we decided to completely remove a GUI and replace it with a CLI, that will interact with I/O files.

Remove every mention of .csv files and replace it with .config files in our functional specifications, we also sent an example of an input and output .config file to group 5.

We also talked about a possible API, it will be out-of-scope and only be implemented if the project is finished before the deadline to allow us and the client to make a GUI for the project.

We discussed the algorithm, we agreed on using a struct for the tank implementation, and we also agreed on providing a prototype of the algorithm in pseudo-code.

And finally, we agreed on having one comment per function in the code, to ensure a better readability of the code and being able to use Doxygen to generate documentation of the code.

## Why LTS and not the latest version?

We decided to use the latest LTS version of Bazel, it is the most stable version of Bazel, and it is also a targeted version that is going to be supported for a long time, otherwise, we could have used the latest version of Bazel, but it is not as stable as the LTS version and it is subject to change in a near future.

## Why an alternative algorithm?

Since we are going to make only one technical specification for 2 groups, we decided to make an alternative algorithm that will be used by the other group and vice-versa.

## Why an API?

Implementing an API allows us to first focus on the core functionalities and ensure a working algorithm before implementing a GUI.

## Why Doxygen?

Doxygen is a tool that allows us to generate documentation of the code, it is going to be useful both for us and the client in case they want to make changes to the code in the future with other developers.

## Glossary

- **CLI**: Command Line Interface
- **GUI**: Graphical User Interface
- **LTS**: Long Term Support
- **I/O**: Input/Output
- **API**: Application Programming Interface