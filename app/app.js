var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue !',
      clients: []
    },
    mounted () {
        this.getClients();
      },
  methods: {
    getClients() {
      axios.get("/clients/").then((response) => {
        this.clients = response.data
      }).catch( 
          error => { console.log(error) }
      )
    },
    block(id) {
        axios.put("/clients/"+id+"/block").then((response) => {
           this.getClients()
          }).catch( 
              error => { console.log(error) }
          )
    },
    unblock(id) {
        axios.put("/clients/"+id+"/unblock").then((response) => {
            this.getClients()
          }).catch( 
              error => { console.log(error) }
          )
    }
  }
  })