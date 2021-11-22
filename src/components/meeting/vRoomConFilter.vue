<template>
  <div class="conditionBox">
    <li class="conditionItem">
      <h3 class="conditionTitle">Area：</h3>
      <ul class="regionLst">
        <li class="region" :key="key" v-for="(item,key) in regionLst" :class="{'active':item.status}" @click="selectRegionCondition(key)">{{item.name}}</li>
      </ul>
    </li>
    <li class="conditionItem">
      <h3 class="conditionTitle">Facility ：</h3>
      <ul class="facility">
        <el-checkbox-group v-model="facilityList" @change="changeRoom()">
          <el-checkbox label="TV"></el-checkbox>
          <el-checkbox label="Projiection"></el-checkbox>
          <el-checkbox label="Whiteboard"></el-checkbox>
          <el-checkbox label="Mic"></el-checkbox>
          <el-checkbox label="Computer"></el-checkbox>
        </el-checkbox-group>
      </ul>
    </li>
  </div>
</template>

<script>
import 'font-awesome/css/font-awesome.css'
  export default {
    name: 'vRoomConFilter',
    data() {
      return {
        facilityList: [],
        regionLst: [{
          name: 'All',
          status: true
        }, {
          name: 'AC1',
          status: false
        }, {
          name: 'AC2',
          status: false
        }, {
          name: 'AC3',
          status: false
        }, {
          name: 'Lab1',
          status: false
        }, {
          name: 'Lab2',
          status: false
        }, {
          name: 'Lab3',
          status: false
        }, {
          name: 'Lab4',
          status: false
        }, {
          name: 'Lab5',
          status: false
        }, {
          name: 'Centre1',
          status: false
        }, {
          name: 'Centre2',
          status: false
        }, {
          name: 'Centre3',
          status: false
        }, {
          name: 'Centre4',
          status: false
        }],
      }
    },
    components: {

    },
    methods: {
      changeRoom(){
          this.$emit('change')
      },
      //根据地域删选会议室信息
      selectRegionCondition(key) {
        this.regionLst.forEach((item) => {
          item.status = false
        })
        this.regionLst[key].status = true
        this.getRegionName()
        this.moveStep = 3
        this.changeRoom()
      },
      //根据地域的激活状态获插件的地域名称
      getRegionName() {
        this.regionLst.forEach((item) => {
          if (item.status) this.regionName = item.name
        })
      },
    },
    mounted() {

    },
    created() {}
  }

</script>

<!--<style lang="less">-->

<!--</style>-->
<style lang="less" scoped>
@themeColor:#66b1ff;
@fontColor:#606266;
//@disabledColor: #c8c9cc
  // 删选条件样式
  .conditionBox {
    margin-bottom: 40px;

    .conditionItem {
      display: flex;

      .conditionTitle {
        font-size: 14px;
        padding: 10px 0;
        width: 90px;
      }

      .facility {
        padding-top: 10px;
      }

      .regionLst {
        display: flex;
        flex-wrap: wrap;
        flex: 1;

        .region {
          font-size: 14px;
          padding: 10px 15px;
          cursor: pointer;
          color: @fontColor;

          &:hover {
            color: @themeColor;
          }
        }

        .region.active {
          background-color: fadeout(@themeColor, 80%);
          border-radius: 6px;
          font-weight: 700;
          color: @themeColor;
        }
      }
    }
  }

</style>
