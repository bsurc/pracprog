# Proposal: Raptor Research Practical Programming Seminar

Author(s): Kyle Shannon <kyle@pobox.com>

Last updated: 2018-10-17

## Abstract

Programming is becoming a part of every day research in nearly every field.
This proposes bridging a gap between introduction programming, and scientific
programming.  The term researcher(s) and learner(s) in this document refers to
students, graduate students, faculty, and staff.

## Background

Boise State University is striving to become an R2 institution.
Cyberinfrastructure is required to reach this goal.  Part of
cyberinfrastructure is human interaction with machines, and one specific form
of that interaction is programming.  While the University provides introductory
workshops and computer science courses, there is little to no training or
support in between.

## Proposal

We propose to develop a general purpose, practical programming short course to
introduce researchers to programming.  The course would _not_ be highly
scientific, instead focusing on common software engineering skills and tasks.
The purpose is to familiarize researchers with concepts to help them understand
what they can do.

## Rationale

Boise State currently offers researchers a few options for learning
programming.  [Software Carpentry](https://software-carpentry.org/) workshops
are put on several times a year, and Computer Science courses.  There
are also a handful of department level courses that teach programming in the
context of the department.  There are benefits and draw backs to each,
discussed below

### Software Carpentry

Software Carpentry (SC) was created 20 years ago to help researchers save time using
computing skills.  The goal of SC is highly aligned with this proposal, and
provides a nice introduction to programming.

#### Benefits
Many researchers have never been exposed to computer programming to solve
research related tasks.  SC shows concrete examples of using Python or R to
read, analyze, and plot scientific data.  This provides exposure to tools that
is necessary for the curious to continue learning programming.

It also introduces researchers to two very common tools, `bash` and `git`.
These are fundamental tools that everyone must have to move forward.

The pedagogy behind SC is intriguing, very learner-centric, and a
learn-by-doing(and making mistakes) curriculum.

#### Drawbacks
Software Carpentry is short.  In one or two days, researchers are exposed to
`bash`, `git`, and Python, R or Matlab (sometimes others).  The sessions are
compact, following along with the instructor for a bit, then working alone on
tasks.  There is little to no discussion.  Helpers are readily available to
help stuck learners, but act more as a triage unit to keep the lessons moving.

Several core concepts in programming are glossed over, or hidden by package
'magic' (numpy, etc.), and thus the new programmer has a hard time learning
from examples that contain rudimentary, lower-level code (e.g. open a file).

SC also lacks a small section on finding solutions to your problems.  This is
crucial for new users, and learning how to search the web for results can go a
long way.

The computer itself is completely ignored, and why or how it does things isn't
mentioned.

### Computer Science Courses
Computer Science (CS) courses fall on the other end of the spectrum.  The courses
are typically related to one topic, too long, and too rigid.  If the researcher
needs in-depth knowledge on a specific topic, such as databases, then the
courses can provide valuable knowledge.

There are some newer classes (Introduction to Data Science comes to mind) that
don't follow these benefits/drawbacks, and perhaps more closely align with the
departmental level courses mentioned below.

The Hatchery Unit courses also fall somewhere else on the spectrum, but still
have rigid content and arbitrary data.

#### Benefits
CS courses typically have knowledgable instructors, in-depth lessons, and
structure.

#### Drawbacks
CS courses are rigid.  One topic is exhausted thoroughly.  This is not
typically what researchers need.  Learning about how disk paging works in a
database class is not valuable to a researcher who just wants to query a
database.

With exhaustion comes length.  The classes are long, and meticulous, which may
cause a loss of interest.

Content in CS classes can be arbitrary.  Sometimes real world examples are
used, other times they are not.  It is likely that the content does not fall
within the research area of interest for the learner.

### Departmental Level Courses
These courses may be the best current approach, but can fall on the other side
of SC workshops.  These courses vary, so these ideas are generalized here.

#### Benefits
Context matters.  If the content is alien to the learner, it is just another
ambiguous step in a complex process.  Ideally the course content is familiar to
the learner.

The programming is also likely familiar to the learner, possibly reimplementing
processes done using other mediums.

#### Drawbacks
The tasks may be to specialized to the content.  The programming language is
likely what the instructor knows best, and the analysis is related to the
instructors research.  This would be fine for a specific group of learners, but
fails to teach generalized programming.

## Implementation
We propose designing, implementing, and executing an eight week short-course to
teach practical programming to a small group of students.  The course will be
centered around general software engineering, not scientific programming.  To
be clear, the course will _support_ scientific programming, both directly and
indirectly.  Directly by providing a tool to make accessing a general purpose
dataset, and indirectly by giving researchers the skills to replicate the
project with other data.

Research Computing will:
- Build a server for the purpose of the class
- Provide a github repository and access for the researchers
- Provide an instructor for eight weeks, 90 - 120 minutes for class time.
- Provide 'office hours' for ~60 minutes a week

The course will consist of X main processes:
- Obtain a moderately large dataset related to the general research field
- Extract a workable subset of data for testing and programming
- Create and populate a well-formed database with the data, checking the
	validity of the data as well.
- Create a web service for accessing the dataset over a network
- All students write tools to use the service (any language)
- All students write a service on the server (any language)
For the purpose of Raptor Research/Biology, the dataset will be the entire (up
to September 2018) eBird database.  The data is individual species sightings
around the world since 2002.  This data may not directly apply to research, but
obviously may have the potential to.

The main webserver will be written interactively, together.  The pace is
governed by the class.

After the data service is written, each student will write a tool in a
programming language of their choice.  The tool may or may not directly apply
to the researchers current work.

The course will cover the following topics:
- Learn a simple mental model of a computer
- Learn a simple mental model of data on a computer
- Using git for practical purposes
- Data 'wrangling'
- Importing data into a relational database
- Software testing
- Writing simple SQL queries
- Developing web methods to provide data 'on demand'
- Learn about different web encodings (csv, json, xml, binary, etc.)
- Develop a client for web services
- Develop a server for web services

