<template>
  <div class="page">
      <Navbar :dashboard="true" :logout_prop="false" />
      <SprintBacklog v-if="sprintbacklogalive"></SprintBacklog>
      <Backlog v-if="backlogalive"></Backlog>
      <NewSprint v-if="newsprintalive"></NewSprint>
      <OldSprints v-if="oldsprintalive"></OldSprints>
      <ProjectPageOverview v-if="projectpage"></ProjectPageOverview>
<!--      <SprintHeader v-if="sprintbacklogalive"></SprintHeader>-->
<!--      <leftMenu />-->
      <el-row >
        <el-col :span="12">
          <el-menu
            class="el-menu-vertical-demo"
            @open="handleOpen"
            @close="handleClose"
            background-color="#444F61"
            text-color="#fff"
            active-text-color="#ffd04b">
<!--            <el-submenu index="1">-->
<!--              <template slot="title">-->
<!--                <i class="el-icon-location"></i>-->
<!--                <span>导航一</span>-->
<!--              </template>-->
<!--              <el-menu-item-group>-->
<!--                <template slot="title">分组一</template>-->
<!--                <el-menu-item index="1-1">选项1</el-menu-item>-->
<!--                <el-menu-item index="1-2">选项2</el-menu-item>-->
<!--              </el-menu-item-group>-->
<!--              <el-menu-item-group title="分组2">-->
<!--                <el-menu-item index="1-3">选项3</el-menu-item>-->
<!--              </el-menu-item-group>-->
<!--              <el-submenu index="1-4">-->
<!--                <template slot="title">选项4</template>-->
<!--                <el-menu-item index="1-4-1">选项1</el-menu-item>-->
<!--              </el-submenu>-->
<!--            </el-submenu>-->
            <el-menu-item index="1"  v-on:click="reload_Backlog">
              <i class="el-icon-location"></i>
              <span slot="title">Backlog</span>
            </el-menu-item>
            <el-menu-item index="2" v-on:click="reload_sprintbacklog">
              <i class="el-icon-menu"></i>
              <span slot="title">Sprint Backlog</span>
            </el-menu-item>
            <el-menu-item index="3">
              <i class="el-icon-document"></i>
              <span slot="title">导航三</span>
            </el-menu-item>
            <el-submenu index="1">
              <template slot="title">
                <i class="el-icon-location"></i>
                <span>Sprint</span>
              </template>
              <el-menu-item-group>
                <template slot="title">team1</template>
                <el-menu-item index="3-1" v-on:click="reload_newsprint">new Sprint</el-menu-item>
                <el-menu-item index="3-2" v-on:click="reload_oldsprint">old Sprint</el-menu-item>
              </el-menu-item-group>
              <el-menu-item-group title="分组2">
                <el-menu-item index="3-3">选项3</el-menu-item>
              </el-menu-item-group>
              <el-submenu index="3-4">
                <template slot="title" >选项4</template>
                <el-menu-item index="1-4-1">选项1</el-menu-item>
              </el-submenu>
            </el-submenu>
            <el-menu-item index="4" v-on:click="reload_projectpage">
              <i class="el-icon-setting"></i>
              <span slot="title">Project page</span>
            </el-menu-item>
          </el-menu>
        </el-col>
      </el-row>
    </div>
</template>

<script>
import Navbar from '@/components/navbar/Navbar.vue'
import SprintBacklog from '@/components/SprintBacklog.vue'
import Backlog from '@/components/Backlog.vue'
import NewSprint from '@/components/NewSprint.vue'
import OldSprints from '@/components/OldSprints.vue'
import ProjectPageOverview from '@/components/ProjectPageOverview.vue'
// import leftMenu from '@/components/leftMenu.vue'
// import SprintHeader from '@/components/sprintbacklog/SprintHeader.vue'
// import LoadStories from '@/components/sprintbacklog/stories/LoadStories.vue'
// import UserStoryCard from '@/components/sprintbacklog/stories/UserStoryCard.vue'
// import TaskGroup from '@/components/sprintbacklog/tasks/TaskGroup.vue'
// import AddStories from '@/components/sprintbacklog/stories/AddStories.vue'
export default {
  name: 'Home',
  components: {
    Navbar,
    SprintBacklog,
    Backlog,
    NewSprint,
    ProjectPageOverview,
    OldSprints
  },
  methods: {
    handleOpen (key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath)
    },
    reload_sprintbacklog () {
      this.sprintbacklogalive = false
      this.$nextTick(function () {
        this.sprintbacklogalive = true
        this.backlogalive = false
        this.newsprintalive = false
        this.projectpage = false
      })
    },
    reload_Backlog () {
      this.backlogalive = false
      this.$nextTick(function () {
        this.backlogalive = true
        this.sprintbacklogalive = false
        this.newsprintalive = false
        this.projectpage = false
        this.oldsprintalive = false
      })
    },
    reload_newsprint () {
      this.newsprintalive = false
      this.$nextTick(function () {
        this.newsprintalive = true
        this.sprintbacklogalive = false
        this.backlogalive = false
        this.projectpage = false
        this.oldsprintalive = false
      })
    },
    reload_oldsprint () {
      this.oldsprintalive = false
      this.$nextTick(function () {
        this.oldsprintalive = true
        this.newsprintalive = false
        this.sprintbacklogalive = false
        this.backlogalive = false
        this.projectpage = false
      })
    },
    reload_projectpage () {
      this.projectpage = false
      this.$nextTick(function () {
        this.projectpage = true
        this.newsprintalive = false
        this.sprintbacklogalive = false
        this.backlogalive = false
        this.oldsprintalive = false
      })
    }
  },
  data () {
    return {
      sprintbacklogalive: false,
      backlogalive: false,
      newsprintalive: false,
      projectpage: false,
      oldsprintalive: false
    }
  }
}
</script>

<style scoped>
.page {
  background-color: #fff;
  height: 100%;
  weight: 100%;
}
.el-menu-vertical-demo{
  position: fixed;
  top: 100px;
  bottom: 0;
  left: 0;
  width:150px;
  background-color: #444F61;
  overflow-y: auto;
}
</style>
