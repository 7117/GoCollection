<template>
  <div>
    <a-button @click="datarequest">测试</a-button>
    <span>{{ code }}</span>
    <span>{{ msg }}</span>
    <span>{{ Age }}</span>

    <span v-for="(a,b) in arrs" :key="a">
      <p>{{ a }}</p>
    </span>

    <router-link to="books">书籍</router-link>
  </div>
</template>

<script>
import ant from 'ant-design-vue'

export default {
  name: "Index",
  data() {
    return {
      code: "",
      msg: "",
      Age: null,
      arrs: [],
    }
  },
  created() {
    this.datarequest()
  },
  methods: {
    datarequest() {
      this.$axios.get('/api11/axios')
          .then(response => {   // 成功的回调
            console.log(response.data);
            this.code = response.data.code
            this.msg = response.data.msg
            this.Age = response.data.user.age
            this.arrs = response.data.arrs
          })
          .catch(error => {  // 失败的回调
            console.log(error);
          });
    }
  },
  components: {
    [ant.name]: ant
  }
}
</script>

<style scoped>

</style>
