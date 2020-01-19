var app = new Vue({
    el: '#app',
    data: {
      clients: [],
      blockeds: []
    },
    mounted () {
        this.getClients();
        this.getBlockeds();
      },
  methods: {
    getClients() {
      axios.get("/clients/").then((response) => {
        this.clients = response.data
      }).catch( 
          error => { console.log(error) }
      )
    },
    getBlockeds() {
      axios.get("/blocked/").then((response) => {
        this.blockeds = response.data
      }).catch( 
          error => { console.log(error) }
      )
    },
    block(id) {
        axios.put("/clients/"+id+"/block").then((response) => {
           this.getClients()
           this.getBlockeds()
          }).catch( 
              error => { console.log(error) }
          )
    },
    unblock(id) {
        axios.put("/clients/"+id+"/unblock").then((response) => {
            this.getClients()
            this.getBlockeds()
          }).catch( 
              error => { console.log(error) }
          )
    }
  }
  })