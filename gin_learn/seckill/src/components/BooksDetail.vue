<template>
  <div>
    detail
  </div>
</template>

<script>
export default {
  name: "BooksDetail",
  mounted() {
    this.getDeail()
  },
  methods: {
    getDeail() {
      console.log(this.$route.params.id)
      this.$axios.get("/api11/detail", {
        params: {id: this.$route.params.id}
      }).then(res => {
        console.log("detail")
        console.log(res.data.books)
      }).catch(res => {
        console.log("data")
      })
    },
  }
}
</script>

<style scoped>

</style>
