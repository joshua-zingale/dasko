```bash
dasko init # Create a new dasko course with standard files pre-populated with default

dasko generate # generate scaffolding for activities, readings, etc

dasko remote set remote.url.edu


# Pulling activity from remote
dasko activity checkout hw1

dasko submission commit --activity hw1 --user bob123@univ.edu --data '{"q1": 23.5}' **/*.c **/*.h Makefile 

dasko submission checkout --activity hw1 --user bob123@univ.edu --lock

dasko submission grade --activity hw1 --user bob123@univ.edu --item 1 --score 5 --message "Missing doc comments in header files"

# If the submission is currently checkout out, dasko can autopopulate the activity and user fields
dasko submission grade --item 2 --score 10

dasko submission grade commit --message "The code is well written, but lacks comments"

dasko submission grade publish --activity hw1


# Used at any point to print out a helpful message on what the current state of dasko is
dasko status # Grading hw1 for bob123@ucr.edu, awaiting scores for rubric items 3,4,5
```