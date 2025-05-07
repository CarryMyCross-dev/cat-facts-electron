<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <button @click="handleClick">Click me</button>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'HelloWorld',
  methods: {
    async handleClick() {
      try {
        const response = await axios.get('http://localhost:8080/catfacts');

        console.log(response.data);
      } catch (ex) {
        console.log(ex);
      }
    }
  },
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
