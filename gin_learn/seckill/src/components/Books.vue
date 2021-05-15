<template>
  <div>
    <table>
      <thead>
      <th>ID</th>
      <th>书籍</th>
      <th>作者</th>
      <th>作者</th>
      </thead>
      <tbody>
      <tr v-for="(a,b) in books" :key="b" style="border: 2px solid red">
        <td>{{ a.id }}</td>
        <td>{{ a.name }}</td>
        <td>{{ a.age }}</td>
        <td>
          <router-link :to="{name:'books_detail',params:{id:a.id}}">
            {{ a.name }}
          </router-link>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "Books",
  data() {
    return {
      books: []
    }
  },
  created() {
    this.getBooks()
  },
  methods: {
    getBooks() {
      this.$axios.get('/api11/axios').then(
          res => {
            console.log(res)
            this.books = res.data.books
          }
      ).catch(
          res => {
            console.log(res)
          }
      )
    },
  },

}
</script>

<style scoped>

</style>
