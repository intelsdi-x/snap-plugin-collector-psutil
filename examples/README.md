# Example tasks

[This](task-psutil.json) example task will collect metrics from **psutil** and publish 
them to file.  

## Running the example

### Requirements 
 * `docker` and `docker-compose` are **installed** and **configured** 

Running the sample is as *easy* as running the script `./run-file-psutil.sh`. 

## Files

- [run-file-psutil.sh](run-file-psutil.sh) 
    - The example is launched with this script
- [file-psutil.sh](file-psutil.sh)
    - Downloads `snapteld`, `snaptel`, `snap-plugin-collector-psutil`,
    `snap-plugin-publisher-file-file` and starts the task
    [task-psutil.json](task-mpsutil.json).
- [tasks/task-psutil.json](tasks/task-psutil.json)
    - Snap task definition
- [.setup.sh](.setup.sh)
    - Verifies dependencies and starts the containers.  It's called 
    by [run-file-psutil.sh](run-file-psutil.sh).