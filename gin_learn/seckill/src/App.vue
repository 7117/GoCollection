<template>
  <div>
    <router-link to="/home">home</router-link>
    <p></p>
    <router-link to="/">Index</router-link>
    <p></p>
    <router-link to="/books_detail">books_detail</router-link>
    <p></p>
    <router-link to="/books_test">books_test</router-link>
    <p></p>
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'App',
}
</script>

<style>

</style>
