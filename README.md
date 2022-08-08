
# data-engineer-challenge

This project is my solution for the following hiring challgenge https://github.com/tamediadigital/hiring-challenges/tree/master/data-engineer-challenge



## Run Locally
The assumption is that you run kafka on port 9092 on your local machine. If not, check the makefile and put in the necessary information.

Other requirements:

* Docker

Clone the project

```bash
  git clone https://github.com/cnpog/data-engineer-challenge-solution
```

Go to the project directory

```bash
  cd data-engineer-challenge-solution
```

How to run

```bash
  make dockerAll

  // kafka input, console log 1 minute interval
  make runDockerBasic

  // kafka input/output 1 minute interval
  make runDockerAdvanced

  // no kafka, loads payload.json 3 and 5 second interval
  make runDockerBonus
```

see [REPORT.md](https://github.com/cnpog/data-engineer-challenge-solution/DOCUMENTATION.md) for project description