<template>
    <div id="app">
        <task v-bind:tasks="tasks"></task>
    </div>
</template>


<script>
import Task from "./components/Task";

export default {
    name:"App",
    components: {
        Task,
    },
    data() {
        return {
            tasks: [
                {
                    id:1,
                    title:"Learn Vue JS",
                    completed:true,
                },
                {
                    id:2,
                    title:"Watch nexflix",
                    completed:false,
                },
                {
                    id:3,
                    title:"Go Shopping",
                    completed:true,
                },
            ]
        }
    }
}
</script>