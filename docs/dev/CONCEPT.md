
Dasko seeks to meet the different needs of the pedagogical ecosystem.
The needs in view are

- Creating educational content
    - writing assignments, drafting readings, setting course policies, coding autograders
- Serving educational content
    - REST API, enforcing policies (who can read/write to what)
- Consuming educational content
    - (typically students) reading materials, completing assignments, viewing grades
- Grading submissions
    - manual grading, autograding, splitting grading among a larger team (like of TAs or Graders)
- Administration of an ongoing course
    - Adding/removing students from a course, granting extensions



# Proposed Structure


The file structure of a Dasko course may look as follows:

```
my-course/
├── config.yml           # Global course identity & permissions
├── chronology.yml       # Defining due date policies
├── groups.yml           # Defining user groups for authorization
├── modules/
│   └── 01-concurrency/
│       ├── mod.yml      # Module manifest
│       ├── readings/
│       │   └── intro.md # Interactive reading with quizzes
│       ├── slides/
│       │   └── lec1.md  # Markdown slides
│       └── labs/
│           └── hashmap/
│               ├── activity.yml # Assignment spec
│               ├── index.md     # Lab instructions
│               └── harness/     # Instructor tests & assets
└── .dasko/              # Local state (gitignored)
    └── work/            # Staging areas for grading
```


And here are some examples of what the configurations could look like.

```yml
# config.yml
course_id: "CS141"
title: "Intermediate Systems Programming"
term: "Spring 2026"

# Global defaults
default_visibility: ["admin", "instructor", "ta"]

# Web UI theme
theme: dracula
```

```yml
# chronology.yml

anchors:
  semester_start: 2026-01-05
  final: 2026-03-14

# Named policies used by activities
policies:
  standard_lab:
    release:
        for: 
            section1: "Monday 08:00"
            section2: "Wednesday 08:00"
    deadline:
        duration: +7d
        time: 23:59 # After seven twenty-four hour days, end before midnight

  reading_quiz:
    release: "Monday 08:00"
    deadline: "Friday 17:00"   # Simple fixed-offset within the week
```


```yml
# groups.yml
# This could be offloaded to an external service,
# but for simplicity here is fine.
admin: ["prof_smith@univ.edu"]
ta: ["grad_ta_1@univ.edu", "grad_ta_2@univ.edu"]
student: ["tom@univ.edu"]
```

```yml
# modules/01-concurrency/mod.yml

id: "concurrency-1"
title: "Module 1: Threads and Synchronization"

# High-level scheduling for the whole module
schedule:
  anchor: "semester_start"
  week: 1

readings:
  - id: "thread-basics"
    title: "Introduction to Pthreads"
    source: "readings/intro.md"

slides:
  - id: "lec1-slides"
    title: "Lecture 1: Race Conditions"
    source: "slides/lec1.md"

activities:
  - "labs/hashmap/activity.yml"

resources:
  - title: "Pthread Cheat Sheet"
    file: "docs/pthreads.pdf"
```

```yml
# modules/01-concurrency/labs/hashmap/activity.yml
title: "Lab 1: Thread-Safe Hash Map"
points: 100
exact-grading: false

# Uses the policy from chronology.yml
# Based on Module 1 (Week 1), this releases Jan 5 and is due Jan 11.
schedule:
  policy: "standard_lab"

initialization:
  copy:
    - src: "https://assets.univ.edu/lib/libperformance.a"
      dest: "harness/libperformance.a"
  exec: "gcc -shared -o harness/check.so harness/src/check.c"
  creates: ["harness/check.so"]

uploads:
  - file: "hashmap.c"
    required: true
  - file: "hashmap.h"
    required: true

reading:
  source: "index.md"
  points: 10   # Quiz questions in index.md normalized to 10/100

autograding:
  # Staging: harness/ and submission/ are prepared automatically
  exec: "make -C harness run_tests"
  points: 70

human-grading:
  exact: true # Do not normalize score
  overflow: true # Allow extra credit
  setup: "make debug"
  rubric:
    items:
      "Clean Mutex Handover": 10
      "Proper Error Handling": 10
```

# Proposed Use

```bash
dasko init # create a new course in the current directory with useful defaults
dasko serve api # Serve dasko as a REST api
dasko serve webui # serve dasko as an LMS web interface
...

# Connet to a remote-broadcast dasko REST api.
# Of course requires authentication.
dasko remote set https://remote.url.edu/path 

# Print information like the current connect course
# Changes printout by context
dasko status

# Move into a new workspace that has the students submission
# along with any specified harness.
# --lock to prevent other graders from trying to grade the same assignment
dasko submission checkout student1@ucr.edu lab1 --lock

# While having checkout out a submission,
# see the rubric items
dasko status

# grade the first rubric item
dasko grade \
    --item 1 \
    --score 3 \
    -m "Multiple memory leaks in the source code" 

# Commit the graded rubric items to the dasko remote
dasko grade --commit

# checkout next ungraded student,
# moving the lock
dasko grade --continue
```

# Technical

Dasko will store internal data in a folder called `.dasko/`.
This will contain student submissions, grades, and other persistent data.

`dasko remote set` will download a `.dasko/` folder with metadata.
