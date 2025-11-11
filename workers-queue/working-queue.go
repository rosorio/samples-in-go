package main

import (
    "fmt"
    "sync"
    "time"
    "math/rand"
)

type Task struct {
    Name string
    Done bool
}


func executeTask(worker int, task chan *Task, wg *sync.WaitGroup) {
    /* On exit, decrease the WaitGroup counter */
    defer wg.Done()

    /* While we can grab tasks from the channel */
    for t :=  range task {
            /* mark the task as done */
            (*t).Done = true
            n := rand.Intn(1000) + 1

            /* Log */
            fmt.Printf("[worker %d] Running task %s and sleep %d seconds\n",
                worker, (*t).Name, n)

            /* Sleep */
            time.Sleep (time.Duration(n)*time.Millisecond )
    }

    /* Exit */
    fmt.Printf("[worker %d] exit.\n", worker)
}


func main() {

    /* Initialize the random source */
    rand.Seed(time.Now().UnixNano())

    /* Create a list of tasks */
    taskList := []Task{ {"Un", false},
                        {"Deux", false},
                        {"Trois", false},
                        {"Quatre", false},
                        {"Cinq", false},
                        {"Six", false},
                        {"Sept", false},
                        {"Huit", false},
                        {"Neux", false},
                        {"Dix", false}}
    
    /* Create the workig queue */
    in  := make(chan *Task, 3)

    /* Create the sync WaitGroup */
    var wg sync.WaitGroup


    /* Create the workers */
    for i := 0; i<3 ; i++ {
        /* increase the WaitGroup count */
        wg.Add(1)
        /* run the go routine */
        go executeTask(i, in, &wg)
    }

    /* fil the working queue */
    for id, _ := range taskList {
        fmt.Printf("[dispatcher] Send task %s\n", taskList[id].Name)
        /* pass a pointer to the taskList object, not a copy */
        in <- &taskList[id]
    }

    /* Close the queue */
    close(in)

    /* Wait for threads completion */
    wg.Wait()

    /* print a status of the tasks */
    for _, task := range taskList {
        fmt.Println(task.Name, task.Done)
    }
}
