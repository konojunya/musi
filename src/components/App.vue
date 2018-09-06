<template>
  <div id="app">
    <!-- loading -->
    <div class="loading" v-if="loading">
      <h1>loading</h1>
    </div>
    <!-- success -->
    <div class="content" v-else>
      <h1>hello</h1>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app',
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
        this.city = res.data.city_name;
        this.playlists = res.data.playlist.playlists.items
        this.loading = false
      },
      err => console.error(err)
    )
  },
  data: {
    loading: true,
    city: "読み込み中...",
    playlists: []
  },
  watch: {
    playlists() {
      console.log(this.playlists[0])
    }
  },
  methods: {
    logout() {
      document.cookie = "musi-token=; max-age=0";
      location.reload()
    }
  }
}
</script>
