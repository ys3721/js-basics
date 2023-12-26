console.log("come on ys!")
const app = Vue.createApp({
    data() {
        return {
            showBooks:false,
            title:"The Final Empire",
            author:"Barndon Sandersong",
            age:45
        }
    }, 
    methods: {
        changeTitle(title) {
            console.log("clieck app.js changeTitle method");
            this.title = title;
        },
        toggleShow() {
            if (this.showBooks) {
                this.showBooks = false
            } else{
                this.showBooks = true
            }
        }
    }
})

app.mount('#app')

 