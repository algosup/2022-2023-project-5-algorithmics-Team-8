# File Structure

## Overview

The file structure of the project is as follows:

For the time being, we do not know how many .cpp and .h files we will have, for the time being, we will replace them with *. and the extension.

```
.
├── docs
│  ├── Management 
│  │  ├── File Structure.md
│  │  ├── Project Management.md
│  │  ├── Weekly Reports
│  │  │  ├── Week 1.md
│  │  │  ├── Week 2.md
│  │  │  ├── Week 3.md
│  │  │  ├── Week 4.md
│  │  │  ├── Week 5.md
│  │  │  ├── Week 6.md
│  │  │  ├── Week 7.md
│  │  │  └── Week 8.md
│  │  ├── Client
│  │  │  └── Messages.md
│  │  └── Team.md
│  ├── Specifications 
│  | ├── Functional Specification.md 
│  | └── Technical Specification.md
│  ├── User Guide
│  │  ├── Installation.md
│  │  ├── User Guide.md
│  │  └── User Guide.pdf
| ├── QA
│  │  ├── Test Plan.md
│  │  └── Test Report.md
| ├── Research
│  │  ├── 0.Project Research.md
│  │  ├── 1.Client.md
│  │  └── 1.Product.md
│  └── README.md
├── src
│  ├── Include
│  │  ├── BUILD
│  │  └── *.h
├── test
│  ├── *Test.cpp
├── bin
│  ├── *.o
├── CI
│  ├── *.yaml
├── .gitignore
├── WORKSPACE
├── DOCKERBUILD
├── .dockerignore 
└── sebastien.config
```

## Description

### Root

The root folder contains all the files that are not part of the project itself but are used to manage it. It contains the documentation, the configuration files, the build files, etc.

We will put files used to manage DOCKER, Bazel, git, etc. 

### docs

This folder contains all the documentation of the project such as the specifications, management, user guide, etc.

### src

This folder contains all the source code of the project. It contains some .cpp files and a BUILD file that is used by Bazel to build the project.

### test

This folder contains all the tests of the project. It contains .cpp files

### CI

This folder will be used to store .yaml files that will be used by Bazel to build the project.

### bin 

This folder will contain the output of the build. It will contain .o files. 