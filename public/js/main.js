new Vue({
  el: "#app",
  delimiters: ['${', '}'],
  mounted() {
    navigator.geolocation.getCurrentPosition(
      async data => {
        const { latitude, longitude } = data.coords;
        const res = await axios.get("/api/tracks", {
          params: {
            latitude,
            longitude
          }
        });
        console.log(res.data);
        this.city = res.data.city_name;
      },
      err => console.error(err)
    )
  },
  data: {
    city: "読み込み中...",
  },
  methods: {
    logout() {
      document.cookie = "musi-token=; max-age=0";
      location.reload()
    }
  }
})