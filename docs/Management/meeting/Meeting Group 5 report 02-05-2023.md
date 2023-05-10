# Meeting Group 5 report 02-05-2023

As group 8 we did a meeting with group 5, the meeting took place on the 2nd of May 2023 at 3.15pm, the meeting lasted 45 minutes. 

## Overview

During this meeting, we discuss early technical decisions and the functional specifications of the project. 

We decided to focus on QT for the GUI and C++ for the backend, as it is one of the most efficient and reliable technologies for this kind of project, 

## Why QT?

QT is one of the main tools to build C++ applications with GUI (Graphical User Interface),

We planned to use QT 6.5 LTS for long-term support and prevent possible bugs, and C++ 11 for the backend, as it is compatible with QT 6.5 and C++ 11, even if it is not the latest version of C++ it is still a very good version that fits our needs(Compatibility and stability with QT)

## Minimum version

Since C++ does not have the same standard library for each Operating System, we discussed a minimum version of Windows that the client will need to run the program, we decided to target every Windows version from Windows 7 to Windows 11, on top of that we are planning to use Bazel to cross compile the program for Linux and macOS. 

## Why Bazel? 

Bazel is a cross-compilation tool that allows us to compile the program for different operating systems and architecture, it will also be the recommended tool by Google Test which is the main testing tool that our QA suggested use. 

We plan to use Bazel 7.0 as it is the latest version of Bazel and it is compatible with QT 6.5 and C++ 11. 

## Specifications 

Both teams will push their respective specifications to the repository of the other team, we will then review them and make sure that they are compatible with our specifications. 

## Next meeting 

We decided to not plan weekly meetings, as we are in different groups and we have different schedules, we will instead plan a meeting when we need to discuss something/use Slack if the need were to arise. 