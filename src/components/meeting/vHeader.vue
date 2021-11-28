<template>
  <b-container fluid>
      <Navbar :dashboard="true" :logout_prop="false" />
  </b-container>
</template>

<script>
import 'font-awesome/css/font-awesome.css'
import Navbar from '@/components/navbar/Navbar.vue'
export default {
  name:'vHeader',
  data(){
   return {

   }
  },
  components: {
      Navbar,
  },
  methods:{

  },
  mounted(){

  },
  created(){}
}
</script>

<style lang="less">

</style>
<style lang="less" scoped>

.header{
    background:#fff;
    box-shadow:0 1px 3px rgba(26, 26, 26, 0.1);
    height:52px;
    display:flex;
    justify-content: center;
    align-items:center;
    .headerMain{
        text-align:left;
        font-size:16px;
        color:#66b1ff;
    }
}
</style>
