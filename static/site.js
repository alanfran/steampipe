var queryUrl = "http://localhost:8080/query/"

var vue = new Vue({
  el: '#app',
  data: {
    item: {},
    errormsg: "",
    serverUrl: "",
    exampleServer: "zs.nekonet.xyz"
  },
  methods: {
    onEnter: function(e) {
      //console.log(e)
      e.stopPropagation()
      e.preventDefault()
      e.returnValue = false
      this.queryServer(this.serverUrl)
    },
    clickExample: function(e) {
      this.serverUrl = this.exampleServer
      this.queryServer(this.serverUrl)
    },
    queryServer: function(server) {
      console.log("Querying " + queryUrl + server)
      this.$http.get(queryUrl + server).then(response => {
        this.errormsg = ""
        this.item = response.data

        this.item.PlayerString = this.item.Players + " / " + this.item.MaxPlayers
        this.item.url = "steam://connect/"+ server

        console.log(response.data)
      }, response => {
        this.errormsg = "Error querying server."
        console.log(response)
      })
    }
  }
})
