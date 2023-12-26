<template>
    <div class="container">
        <div class="task">
            <!--title-->
            <div class="title">
                <h1>To Do List</h1>
            </div>
            <!-- from -->
            <div class="form">
                <input type="text" placeholder="New Task" v-model="newTask" @keyup.enter="addTask" />
                <button @click="addTask"><i class="fas fa-plus"></i></button>
            </div>
            <!-- task list-->
            <div class="taskItems">
                <ul>
                    <tetm v-bind:task="task" v-for="(task, index) in tasks" :key="task.id" @remove="removeTask(index)"
                        @completeTask="completeTask(task)"></tetm>
                </ul>
            </div>
            <!-- buttons -->
            <div class="clearBtns">
                <button @click="clearCompleted">Clear completed</button>
                <button @click="clearAll">Clear all</button>
            </div>
            <!-- pending task -->
            <div class="pendingTasks">
                <span>Pending Tasks: {{ incomplete }}</span>
            </div>
        </div>
    </div>
</template>

<script src="./node_quickstart/index.js"></script>
<script>
import TaskItem from './Task-item.vue'
export default {
    name: "Task",
    props: ['tasks'],
    components:{
      'tetm':  TaskItem
    },
    date() {
        return {
            newTask:"value"
        }
    },
    computed: {
        incomplete(){
            return this.tasks.filter(this.inProgress).length;
        }

    },
    methods: {
        addTask() {
            if(this.newTask) {
                this.tasks.push({
                    title: this.newTask,
                    completed: false,
                });
                this.newTask = "";
            }
        },
        inProgress(task) {
            return !this.isCompleted(task)
        },
        isCompleted(task) {
            return task.completed
        },
        clearAll() {
            this.tasks = []
        }, 
        clearCompleted() {
            this.tasks = this.tasks.filter(this.inProgress)
        }, 
        removeTask(index) {
            this.tasks.splice(index, 1)
        },
        completeTask(task) {
            task.completed = !task.completed
        },
    },
};
</script>