# Software Architecture

This document outlines the software architecture of Dasko, which is implemented in the Go programming language.

The leading design choice is to make keep *everything a file*, meaning all slideshows, readings, assignments, submissions, and grades are files.


## Broad Strokes

Dasko has four main modules:

1. Course Parsing
    - parsing the folder structure of a Dasko course's source to derive a useable representation by other modules of Dasko
2. Content Presentation
    - taking a representation of a Dasko course and compiling the representation into consumable formats, like PDFs or HTML
3. Activity Interaction
    - receiving and storing activity interactions 
4. Activity Grading
    - accessing student submissions and applying grades to the submissions
5. CLI Interface
    - accessing all previous modules via a command line interface using the `dasko` utility
6. Content Interaction over the Web
    - serving the course materials over the web for access by students, graders, and instructors, permitting submissions, grading, and more.