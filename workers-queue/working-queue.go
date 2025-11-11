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
    defer wg.Done()
    for t :=  range task {
            (*t).Done = true
            n := rand.Intn(1000) + 1
            fmt.Printf("[worker %d] Running task %s and sleep %d seconds\n",
                worker, (*t).Name, n)
            (*t).Done = true
            time.Sleep (time.Duration(n)*time.Millisecond )
    }
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
    /* Create the sync mecanism */
    var wg sync.WaitGroup


    /* Create the workers */
    for i := 0; i<3 ; i++ {
        wg.Add(1)
        go executeTask(i, in, &wg)
    }

    /* fil the working queue */
    for id, _ := range taskList {
        fmt.Printf("[dispatcher] Send task %s\n", taskList[id].Name)
        in <- &taskList[id]
    }

    close(in)
    wg.Wait()

    for _, task := range taskList {
        fmt.Println(task.Name, task.Done)
    }
}
