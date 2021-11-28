<template>
  <div class="resRoomWrap">
        <v-header></v-header>
    <div class="main">
      <room-res-plug></room-res-plug>
    </div>

<!--    <v-footer></v-footer>-->
  </div>
</template>

<script>
import vHeader from '@/components/meeting/vHeader'
import roomResPlug from '@/components/meeting/roomRerservation'
export default {
  name: 'homePage',
  components:{
    roomResPlug,
    vHeader,
  },
  data () {
    return {

    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.resRoomWrap{
  height:100%;
  display: flex;
  flex-direction: column;
}
.main{
  flex:1;
  overflow: auto;
}
</style>
